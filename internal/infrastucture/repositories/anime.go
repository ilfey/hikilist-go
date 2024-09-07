package repositories

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var (
	ErrAnimeCreateFailed             = errtype.NewInternalRepositoryError("unable to store anime")
	ErrAnimeGetFailed                = errtype.NewInternalRepositoryError("unable to get anime")
	ErrAnimeNotFoundById             = errtype.NewEntityNotFoundError("database", "anime", "id")
	ErrAnimesFindFailed              = errtype.NewInternalRepositoryError("unable to find animes")
	ErrAnimesCountFailed             = errtype.NewInternalRepositoryError("unable to count animes")
	ErrAnimesFindInCollectionFailed  = errtype.NewInternalRepositoryError("unable to find animes in collection")
	ErrAnimesCountInCollectionFailed = errtype.NewInternalRepositoryError("unable to count animes in collection")
)

type Anime struct {
	log loggerInterface.Logger
	db  postgres.RW
}

func NewAnime(container diInterface.ServiceContainer) (*Anime, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	db, err := container.GetPostgresDatabase()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Anime{
		log: log,
		db:  db,
	}, nil
}

func (r *Anime) WithTx(tx postgres.RW) repositoryInterface.Anime {
	return &Anime{
		db: tx,
	}
}

func (r *Anime) Create(ctx context.Context, cm *dto.AnimeCreateRequestDTO) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return r.log.Propagate(err)
	}

	err = r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		r.log.Error(err)

		return ErrAnimeCreateFailed
	}

	return nil
}

func (r *Anime) Get(ctx context.Context, conds any) (*agg.AnimeDetail, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var dm agg.AnimeDetail

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		r.log.Error(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAnimeNotFoundById
		}

		return nil, ErrAnimeGetFailed
	}

	return &dm, nil
}

func (r *Anime) FindWithPaginator(ctx context.Context, dto *dto.AnimeListRequestDTO, conds any) ([]*agg.AnimeListItem, error) {
	sql, args, err := r.FindWithPaginatorSQL(dto, conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var items []*agg.AnimeListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Error(err)

		return nil, ErrAnimesFindFailed
	}

	return items, nil
}

func (r *Anime) FindFromCollectionWithPaginator(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) ([]*agg.AnimeListItem, error) {
	sql, args, err := r.FindFromCollectionWithPaginatorSQL(dto)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var items []*agg.AnimeListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Error(err)

		return nil, ErrAnimesFindInCollectionFailed
	}

	return items, nil
}

func (r *Anime) CountInCollection(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) (uint64, error) {
	sql, args, err := r.CountInCollectionSQL(dto)
	if err != nil {
		return 0, r.log.Propagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Error(err)

		return 0, ErrAnimesCountInCollectionFailed
	}

	return count, nil
}

func (r *Anime) Count(ctx context.Context, conds any) (uint64, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, r.log.Propagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Error(err)

		return 0, ErrAnimesCountFailed
	}

	return count, nil
}
