package tokenModels

import (
	"context"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type CreateModel struct {
	ID uint `json:"-"`

	Token string `json:"-"`
}

func (CreateModel) TableName() string {
	return "tokens"
}

func (cm *CreateModel) Insert(ctx context.Context) error {
	id, err := orm.Insert(cm).
		Ignore("ID").
		Exec(ctx, database.Instance())

	if err != nil {
		return err
	}

	cm.ID = id

	return nil
}
