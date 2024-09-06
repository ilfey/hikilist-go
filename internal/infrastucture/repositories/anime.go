package repositories

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
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
	logger loggerInterface.Logger
	db     postgres.RW
}

func NewAnime(logger loggerInterface.Logger, db postgres.RW) *Anime {
	return &Anime{
		logger: logger,
		db:     db,
	}
}

func (r *Anime) WithTx(tx postgres.RW) repositoryInterface.Anime {
	return &Anime{
		db: tx,
	}
}

func (r *Anime) Create(ctx context.Context, cm *dto.AnimeCreateRequestDTO) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return r.logger.CriticalPropagate(err)
	}

	err = r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		r.logger.Log(err)

		return ErrAnimeCreateFailed
	}

	return nil
}

func (r *Anime) Get(ctx context.Context, conds any) (*agg.AnimeDetail, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, r.logger.CriticalPropagate(err)
	}

	var dm agg.AnimeDetail

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		r.logger.Log(err)

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
		return nil, r.logger.CriticalPropagate(err)
	}

	var items []*agg.AnimeListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.logger.Log(err)

		return nil, ErrAnimesFindFailed
	}

	return items, nil
}

func (r *Anime) FindFromCollectionWithPaginator(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) ([]*agg.AnimeListItem, error) {
	sql, args, err := r.FindFromCollectionWithPaginatorSQL(dto)
	if err != nil {
		return nil, r.logger.CriticalPropagate(err)
	}

	var items []*agg.AnimeListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.logger.Log(err)

		return nil, ErrAnimesFindInCollectionFailed
	}

	return items, nil
}

func (r *Anime) CountInCollection(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) (uint64, error) {
	sql, args, err := r.CountInCollectionSQL(dto)
	if err != nil {
		return 0, r.logger.CriticalPropagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.logger.Log(err)

		return 0, ErrAnimesCountInCollectionFailed
	}

	return count, nil
}

func (r *Anime) Count(ctx context.Context, conds any) (uint64, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, r.logger.CriticalPropagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.logger.Log(err)

		return 0, ErrAnimesCountFailed
	}

	return count, nil
}
