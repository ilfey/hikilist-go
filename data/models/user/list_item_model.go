package userModels

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type ListItemModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`

	CreatedAt time.Time `json:"created_at"`
}

func (ListItemModel) TableName() string {
	return "users"
}

func (lim *ListItemModel) Get(ctx context.Context, conds any) error {
	m, err := orm.Select(lim).
		Where(conds).
		QueryRow(ctx, database.Instance())
	if err != nil {
		return err
	}

	*lim = *m

	return nil 
}
