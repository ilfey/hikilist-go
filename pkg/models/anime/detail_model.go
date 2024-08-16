package anime

import (
	"time"
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
	UpdatedAt time.Time `json:"updated_at"`
}

// func (dm *DetailModel) Get(ctx context.Context, conds map[string]any) error {
// 	sql, args, err := dm.GetSQL(conds)
// 	if err != nil {
// 		return err
// 	}

// 	return pgxscan.Select(ctx, database.Instance(), dm, sql, args...)
// }

// func (DetailModel) GetSQL(conds map[string]any) (string, []any, error) {
// 	sql, args, err := sq.Select(
// 		"id",
// 		"title",
// 		"description",
// 		"poster",
// 		"episodes",
// 		"episodes_released",
// 		"mal_id",
// 		"shiki_id",
// 		"created_at",
// 	).
// 		From("animes").
// 		Where(conds).
// 		ToSql()
// 	if err != nil {
// 		return "", nil, eris.Wrap(err, "failed to build anime select query")
// 	}

// 	return sql, args, nil
// }
