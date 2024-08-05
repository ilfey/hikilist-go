package userActionModels

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type CreateModel struct {
	ID uint `json:"-"`

	UserID uint `json:"-"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"-"`
}

func (CreateModel) TableName() string {
	return "user_actions"
}

func (cm *CreateModel) Insert(ctx context.Context) error {
	cm.CreatedAt = time.Now()

	id, err := orm.Insert(cm).
		Ignore("ID").
		Exec(ctx, database.Instance())

	if err != nil {
		return err
	}

	cm.ID = id

	return nil
}
