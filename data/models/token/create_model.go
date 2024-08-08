package token

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/data/database"
	"github.com/rotisserie/eris"
)

type CreateModel struct {
	ID uint `json:"-"`

	Token     string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (cm *CreateModel) Insert(ctx context.Context) error {
	sql, args, err := cm.InsertSQL()
	if err != nil {
		return err
	}

	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		return eris.Wrap(err, "failed to insert token")
	}

	return nil
}

func (cm *CreateModel) InsertSQL() (string, []any, error) {
	sql, args, err := sq.Insert("tokens").
		Columns(
			"token",
			"created_at",
		).
		Values(
			cm.Token,
			time.Now(),
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", nil, eris.Wrap(err, "failed to build token insert query")
	}

	return sql, args, nil
}
