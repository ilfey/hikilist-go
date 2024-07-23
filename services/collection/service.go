package collectionService

import (
	"context"
	"fmt"

	"github.com/ilfey/hikilist-go/data/entities"
	collectionModels "github.com/ilfey/hikilist-go/data/models/collection"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, cm *collectionModels.CreateModel) (*collectionModels.DetailModel, error)
	Get(ctx context.Context, conds ...any) (*collectionModels.DetailModel, error)
	Paginate(ctx context.Context, p *collectionModels.Paginate, whereArgs ...any) (*collectionModels.ListModel, error)
}

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) Create(ctx context.Context, cm *collectionModels.CreateModel) (*collectionModels.DetailModel, error) {
	var dm collectionModels.DetailModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			// Create collection
			entity := cm.ToEntity()

			result := tx.Create(entity)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed create entity")
			}

			// Create userAction
			result = tx.Create(&entities.UserAction{
				UserID:      entity.UserID,
				Title:       "Создание коллекции",
				Description: fmt.Sprintf("Вы создали коллекцию \"%s\".", entity.Name),
			})
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed create userAction")
			}

			// Get created detail collection
			result = tx.Model(&entities.Collection{}).
				Preload("User", func(tx *gorm.DB) *gorm.DB { // TODO: Test this
					return tx.Select("*")
				}).
				First(&dm, entity.ID)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed get detail model after create collection")
			}

			return result.Error
		})

	return &dm, err
}

func (s *service) Get(ctx context.Context, conds ...any) (*collectionModels.DetailModel, error) {
	var dm collectionModels.DetailModel

	result := s.db.WithContext(ctx).
		Model(&entities.Collection{}).
		Preload("User", func(tx *gorm.DB) *gorm.DB { // TODO: Test this
			return tx.Select("*")
		}).
		Preload("Animes", func(tx *gorm.DB) *gorm.DB { // TODO: Test this
			return tx.Select("*")
		}).
		First(&dm, conds...)
	if result.Error != nil {
		return nil, eris.Wrapf(result.Error, "failed get collection with conds: %+v", conds)
	}

	return &dm, nil
}
func (s *service) Paginate(ctx context.Context, p *collectionModels.Paginate, whereArgs ...any) (*collectionModels.ListModel, error) {
	var lm *collectionModels.ListModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var (
				items []*collectionModels.ListItemModel
				count int64
			)

			// Build query
			query := s.db.Model(&entities.Collection{}).Scopes(p.Scope)

			if len(whereArgs) != 0 {
				query = query.Where(whereArgs[0], whereArgs[1:]...)
			}

			// Get list of animes like whereArgs
			result := query.Find(&items)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get list of animes with whereArgs: %+v", whereArgs)
			}

			lm = collectionModels.NewListModel(items)

			// Get count of animes like whereArgs
			result = query.Count(&count)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get count with whereArgs: %+v", whereArgs)
			}

			lm.WithCount(count)

			return nil
		})

	return lm, err
}
