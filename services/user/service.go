package userService

import (
	"context"

	"github.com/ilfey/hikilist-go/data/entities"
	authModels "github.com/ilfey/hikilist-go/data/models/auth"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/rotisserie/eris"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, rm *authModels.RegisterModel) (*userModels.DetailModel, error)

	Get(ctx context.Context, conds ...any) (*userModels.DetailModel, error)

	Paginate(ctx context.Context, p *userModels.Paginate, whereArgs ...any) (*userModels.ListModel, error)
}

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) Create(ctx context.Context, rm *authModels.RegisterModel) (*userModels.DetailModel, error) {
	hashedPassword := errorsx.Must(
		bcrypt.GenerateFromPassword(
			[]byte(rm.Password),
			bcrypt.DefaultCost,
		),
	)

	// Create entity
	userEntity := &entities.User{
		Username: rm.Username,
		Password: string(hashedPassword),
	}

	var dm userModels.DetailModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			// Create user
			result := tx.Create(userEntity)
			if result.Error != nil {
				if eris.Is(result.Error, gorm.ErrRecordNotFound) {
					eris.Wrap(result.Error, "user already exist")
				}

				return eris.Wrap(result.Error, "failed create user")
			}

			// Create register userAction
			result = tx.Create(&entities.UserAction{
				UserID:      userEntity.ID,
				Title:       "Регистрация аккаунта",
				Description: "Начало вашего пути на сайте Hikilist.",
			})
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed create register userAction")
			}

			result = tx.Model(&entities.User{}).First(&dm, userEntity.ID)
			if result.Error != nil {
				return eris.Wrap(result.Error, "failed get detail model after create user")
			}

			return nil
		})

	return &dm, err
}

func (s *service) Get(ctx context.Context, conds ...any) (*userModels.DetailModel, error) {
	var dm userModels.DetailModel

	result := s.db.WithContext(ctx).
		Model(&entities.User{}).
		First(&dm, conds...)
	if result.Error != nil {
		return nil, eris.Wrapf(result.Error, "failed get detail model with conds %+v", conds)
	}

	return &dm, nil
}

func (s *service) Paginate(ctx context.Context, p *userModels.Paginate, whereArgs ...any) (*userModels.ListModel, error) {
	var lm *userModels.ListModel

	err := s.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var (
				items []*userModels.ListItemModel
				count int64
			)

			// Build query
			query := s.db.Model(&entities.User{}).Scopes(p.Scope)

			if len(whereArgs) != 0 {
				query = query.Where(whereArgs[0], whereArgs[1:]...)
			}

			// Get list of users like whereArgs
			result := query.Find(&items)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get list of users with whereArgs: %+v", whereArgs)
			}

			lm = userModels.NewListModel(items)

			// Get count of users like whereArgs
			result = query.Count(&count)
			if result.Error != nil {
				return eris.Wrapf(result.Error, "failed get count with whereArgs: %+v", whereArgs)
			}

			lm.WithCount(count)

			return nil
		})

	return lm, err
}
