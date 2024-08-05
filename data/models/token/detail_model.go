package tokenModels

import (
	"context"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type DetailModel struct {
	ID    uint   `json:"-"`
	Token string `json:"-"`
}

func (DetailModel) TableName() string {
	return "tokens"
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

func (dm *DetailModel) Delete(ctx context.Context) error {
	_, err := orm.Delete(dm).
		Where(map[string]any{
			"ID": dm.ID,
		}).
		Exec(ctx, database.Instance())

	return err
}
