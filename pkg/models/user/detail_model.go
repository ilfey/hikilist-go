package user

import (
	"time"
)

// Модель пользователя
type DetailModel struct {
	ID uint `json:"id"`

	Username string `json:"username"`
	Password string `json:"-"`

	LastOnline *time.Time `json:"last_online"`

	CreatedAt time.Time `json:"created_at"`
}

// func (dm *DetailModel) Get(ctx context.Context, conds map[string]any) error {
// 	sql, args, err := dm.GetSQL(conds)
// 	if err != nil {
// 		return err
// 	}

// 	err = pgxscan.Get(ctx, database.Instance(), dm, sql, args...)
// 	if err != nil {
// 		return eris.Wrap(err, "failed to get user")
// 	}
// 	return nil
// }

// func (DetailModel) GetSQL(conds map[string]any) (string, []any, error) {
// 	sql, args, err := sq.Select(
// 		"id",
// 		"username",
// 		"password",
// 		"last_online",
// 		"created_at",
// 	).
// 		From("users").
// 		Where(conds).
// 		ToSql()
// 	if err != nil {
// 		return "", nil, eris.Wrap(err, "failed to build user select query")
// 	}

// 	return sql, args, nil
// }

// func (m *DetailModel) Delete(ctx context.Context) error {
// 	sql, args, err := m.DeleteSQL()
// 	if err != nil {
// 		return err
// 	}

// 	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&m.ID)
// 	if err != nil {
// 		return eris.Wrap(err, "failed to delete user")
// 	}

// 	return nil
// }

// func (m *DetailModel) DeleteSQL() (string, []any, error) {
// 	sql, args, err := sq.Delete("users").
// 		Where(sq.Eq{"id": m.ID}).
// 		Suffix("RETURNING id").
// 		ToSql()
// 	if err != nil {
// 		return "", nil, eris.Wrap(err, "failed to build user delete query")
// 	}

// 	return sql, args, nil
// }

// func (m *DetailModel) UpdateLastOnline(ctx context.Context) error {
// 	sql, args, err := m.UpdateLastOnlineSQL()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = database.Instance().Exec(ctx, sql, args...)
// 	if err != nil {
// 		return eris.Wrap(err, "failed to update user")
// 	}

// 	return nil
// }

// func (m *DetailModel) UpdateLastOnlineSQL() (string, []any, error) {
// 	sql, args, err := sq.Update("users").
// 		SetMap(map[string]any{
// 			"last_online": time.Now(),
// 		}).
// 		Where(sq.Eq{"id": m.ID}).
// 		ToSql()
// 	if err != nil {
// 		return "", nil, eris.Wrap(err, "failed to build user last online update query")
// 	}

// 	return sql, args, nil
// }

// func (m *DetailModel) UpdatePasswordSQL() (string, []any, error) {
// 	sql, args, err := sq.Update("users").
// 		SetMap(map[string]any{
// 			"password": m.Password,
// 		}).
// 		Where(sq.Eq{"id": m.ID}).
// 		ToSql()
// 	if err != nil {
// 		return "", nil, eris.Wrap(err, "failed to build user password update query")
// 	}

// 	return sql, args, nil
// }

// func (m *DetailModel) UpdateUsernameSQL() (string, []any, error) {
// 	sql, args, err := sq.Update("users").
// 		SetMap(map[string]any{
// 			"username": m.Username,
// 		}).
// 		Where(sq.Eq{"id": m.ID}).
// 		ToSql()
// 	if err != nil {
// 		return "", nil, eris.Wrap(err, "failed to build user username update query")
// 	}

// 	return sql, args, nil
// }
