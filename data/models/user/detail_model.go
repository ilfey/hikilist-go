package userModels

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

// var _ baseModels.DetailModel[DetailModel] = &DetailModel{}

// Модель пользователя
type DetailModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`
	Password string `json:"-"`

	LastOnline *time.Time `json:"last_online"`

	CreatedAt time.Time `json:"created_at"`
}

func (m *DetailModel) TableName() string {
	return "users"
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

func (m *DetailModel) Update(ctx context.Context) error {
	_, err := orm.Update(m).
		Where(map[string]any{
			"ID": m.ID,
		}).
		Exec(ctx, database.Instance())

	return err
}
