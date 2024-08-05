package userActionModels

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/data/database"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type DetailModel struct {
	ID uint

	UserID uint
	User   *userModels.ListItemModel

	Title       string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (DetailModel) TableName() string {
	return "user_actions"
}

func (dm *DetailModel) Get(ctx context.Context, conds any) error {
	m, err := orm.Select(dm).
		Where(conds).
		QueryRow(ctx, database.Instance())
	if err != nil {
		return err
	}

	*dm = *m

	return nil
}
