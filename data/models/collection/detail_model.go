package collectionModels

import (
	"context"
	"fmt"
	"time"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"

	// animeModels "github.com/ilfey/hikilist-go/data/models/anime"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
)

type DetailModel struct {
	ID uint `json:"id"`

	UserID uint `json:"-"`

	User *userModels.ListItemModel `json:"user"`

	Title string `json:"title"`

	Description *string `json:"description"`

	IsPublic bool `json:"is_public"`

	// Animes []*animeModels.ListItemModel `json:"animes"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DetailModel) TableName() string {
	return "collections"
}

func (dm *DetailModel) Get(ctx context.Context, conds any) error {
	m, err := orm.Select(dm).
		Resolve("User", func(ctx context.Context, dm *DetailModel) error {
			var user userModels.ListItemModel

			err := user.Get(ctx, fmt.Sprintf("%s.id = %d", user.TableName(), dm.UserID))
			if err != nil {
				return err
			}

			dm.User = &user

			return nil
		}).
		Where(conds).
		QueryRow(ctx, database.Instance())
	if err != nil {
		return err
	}

	*dm = *m

	return nil
}
