package animeService

import (
	"context"
	"errors"

	"github.com/ilfey/hikilist-go/data/entities"
	animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, cm *animeModels.CreateModel) (*animeModels.DetailModel, error)

	Get(ctx context.Context, conds ...any) (*animeModels.DetailModel, error)

	PaginateFromCollection(ctx context.Context, p *animeModels.Paginate, id uint) (*animeModels.ListModel, error)

	Paginate(ctx context.Context, p *animeModels.Paginate, whereArgs ...any) (*animeModels.ListModel, error)

	ResolveShiki(*animeModels.ShikiDetailModel) error
}

// Имплементация.

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) Create(ctx context.Context, cm *animeModels.CreateModel) (*animeModels.DetailModel, error) {
	var dm animeModels.DetailModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			// Create anime
			entity := cm.ToEntity()

			result := tx.Create(entity)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed create entity")
			}

			// Get created detail anime
			result = tx.Model(&entities.Anime{}).
				Preload("Related", func(tx *gorm.DB) *gorm.DB { // TODO: Test this
					return tx.Select("*")
				}).
				First(&dm, entity.ID)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed get detail model after create anime")
			}

			return result.Error
		})

	return &dm, err
}

func (s *service) ResolveShiki(model *animeModels.ShikiDetailModel) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var searchEntity entities.Anime

		if err := tx.First(&searchEntity, map[string]any{
			"shiki_id": model.ID,
		}).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			createEntity := model.ToEntity()

			return tx.Create(createEntity).Error
		}

		// Если время обновления существующей записи больше времени создания новой записи
		if searchEntity.UpdatedAt.After(*model.UpdatedAt) {
			return nil
		}

		updateEntity := model.ToEntity()

		updateEntity.ID = searchEntity.ID

		return tx.Model(&updateEntity).Updates(updateEntity).Error
	})
}

func (s *service) Get(ctx context.Context, conds ...any) (*animeModels.DetailModel, error) {
	var dm animeModels.DetailModel

	result := s.db.WithContext(ctx).
		Model(&entities.Anime{}).
		Preload("Related", func(tx *gorm.DB) *gorm.DB { // TODO: Test this
			return tx.Select("*")
		}).
		First(&dm, conds...)
	if result.Error != nil {
		return nil, eris.Wrapf(result.Error, "failed get detail model with conds: %+v", conds)
	}

	return &dm, nil
}

func (s *service) Paginate(ctx context.Context, p *animeModels.Paginate, whereArgs ...any) (*animeModels.ListModel, error) {
	var lm *animeModels.ListModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var (
				items []*animeModels.ListItemModel
				count int64
			)

			// Build query
			query := s.db.Model(&entities.Anime{}).Scopes(p.Scope)

			// Set whereArgs if exists
			if len(whereArgs) != 0 {
				query = query.Where(whereArgs[0], whereArgs[1:]...)
			}

			// Get list of animes
			result := query.Find(&items)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get list of animes with whereArgs: %+v", whereArgs)
			}

			lm = animeModels.NewListModel(items)

			// Get count of animes
			result = query.Count(&count)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get count with whereArgs: %+v", whereArgs)
			}

			lm.WithCount(count)

			return nil
		})

	return lm, err
}

func (s *service) PaginateFromCollection(ctx context.Context, p *animeModels.Paginate, id uint) (*animeModels.ListModel, error) {
	var lm *animeModels.ListModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var (
				collection entities.Collection
				items      []*animeModels.ListItemModel
				count      int64
			)

			// Get collection
			result := tx.First(&collection, id)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get collection with id: %d", id)
			}

			count = int64(len(collection.Animes))

			if count == 0 {
				lm = animeModels.NewListModel(items)

				return nil
			}

			// Get animes
			result = tx.Model(&entities.Anime{}).Scopes(p.Scope).Find(&items, collection.Animes)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get animes from collection with id: %d", id)
			}

			lm = animeModels.NewListModel(items)
			lm.WithCount(count)

			return nil
		})

	return lm, err
}
