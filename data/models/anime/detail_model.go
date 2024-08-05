package animeModels

import (
	"context"
	"time"

	"github.com/ilfey/hikilist-go/data/database"
	"github.com/ilfey/hikilist-go/internal/orm"
)

type DetailModel struct {
	ID uint `json:"id"`

	Title            string  `json:"title"`
	Description      *string `json:"description"`
	Poster           *string `json:"poster"`
	Episodes         *uint   `json:"episodes"`
	EpisodesReleased uint    `json:"episodes_released"`

	MalID   *uint `json:"mal_id"`
	ShikiID *uint `json:"shiki_id"`

	Related []*ListItemModel `json:"related"`

	CreatedAt time.Time `json:"created_at"`
}

func (DetailModel) TableName() string {
	return "animes"
}

func (dm *DetailModel) Get(ctx context.Context, conds any) error {
	m, err := orm.Select(&DetailModel{}).
		Ignore("Related"). // TODO: fix this
		Where(conds).
		QueryRow(ctx, database.Instance())
	if err != nil {
		return err
	}

	*dm = *m

	return nil
}
