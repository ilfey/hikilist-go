package user

type CreateModel struct {
	ID uint `json:"-"`

	Username string `json:"username"`
	Password string `json:"password"`
}

// func (cm *CreateModel) Insert(ctx context.Context) error {
// 	sql, args, err := cm.InsertSQL()
// 	if err != nil {
// 		return err
// 	}

// 	err = database.Instance().QueryRow(ctx, sql, args...).Scan(&cm.ID)
// 	if err != nil {
// 		return eris.Wrap(err, "failed to insert user")
// 	}

// 	return nil
// }

// func (cm *CreateModel) InsertSQL() (string, []any, error) {
// 	sql, args, err := sq.Insert("users").
// 		Columns(
// 			"username",
// 			"password",
// 			"created_at",
// 		).
// 		Values(
// 			cm.Username,
// 			cm.Password,
// 			time.Now(),
// 		).
// 		Suffix("RETURNING id").
// 		ToSql()
// 	if err != nil {
// 		return "", nil, eris.Wrap(err, "failed to build user insert query")
// 	}

// 	return sql, args, nil
// }
