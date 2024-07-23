package userActionService

import (
	"context"

	"github.com/ilfey/hikilist-go/data/entities"
	userActionModels "github.com/ilfey/hikilist-go/data/models/user_action"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, cm *userActionModels.CreateModel) (*userActionModels.DetailModel, error)
	Get(ctx context.Context, conds ...any) (*userActionModels.DetailModel, error)
	Paginate(ctx context.Context, p *userActionModels.Paginate, whereArgs ...any) (*userActionModels.ListModel, error)
}

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) Create(ctx context.Context, cm *userActionModels.CreateModel) (*userActionModels.DetailModel, error) {
	var dm userActionModels.DetailModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			// Create anime
			entity := cm.ToEntity()

			result := tx.Create(entity)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed create entity")
			}

			// Get created detail anime
			result = tx.Model(&entities.UserAction{}).
				First(&dm, entity.ID)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed get detail model after create userAction")
			}

			return result.Error
		})

	return &dm, err
}

func (s *service) Get(ctx context.Context, conds ...any) (*userActionModels.DetailModel, error) {
	var dm userActionModels.DetailModel

	result := s.db.WithContext(ctx).
		Model(&entities.UserAction{}).
		First(&dm, conds...)
	if result.Error != nil {
		return nil, eris.Wrapf(result.Error, "failed get detail model with conds: %+v", conds)
	}

	return &dm, nil
}

func (s *service) Paginate(ctx context.Context, p *userActionModels.Paginate, whereArgs ...any) (*userActionModels.ListModel, error) {
	var lm *userActionModels.ListModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var (
				items []*userActionModels.ListItemModel
				count int64
			)

			// Build query
			query := s.db.Model(&entities.UserAction{}).Scopes(p.Scope)

			if len(whereArgs) != 0 {
				query = query.Where(whereArgs[0], whereArgs[1:]...)
			}

			// Get list of userActions like whereArgs
			result := query.Find(&items)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get list of userActions with whereArgs: %+v", whereArgs)
			}

			lm = userActionModels.NewListModel(items)

			// Get count of userActions like whereArgs
			result = query.Count(&count)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get count with whereArgs: %+v", whereArgs)
			}

			lm.WithCount(count)

			return nil
		})

	return lm, err
}
