package repositories

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"time"

	"github.com/Masterminds/squirrel"
)

func (r *User) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From(UserTN)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *User) FindWithPaginatorSQL(dto *dto.UserListRequestDTO, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"username",
		"created_at",
	).
		From(UserTN)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.
		OrderBy("id ASC"). // OrderBy(dto.Order.ToQuery()).
		Offset((dto.Page - 1) * dto.Limit).
		Limit(dto.Limit).
		ToSql()
}

func (r *User) CreateSQL(cm *dto.UserCreateRequestDTO) (string, []any, error) {
	return squirrel.Insert(UserTN).
		Columns(
			"username",
			"password",
		).
		Values(
			cm.Username,
			cm.Password,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *User) GetSQL(conds any) (string, []any, error) {
	if conds == nil {
		panic("conds is nil")
	}

	return squirrel.Select(
		"id",
		"username",
		"password",
		"last_online",
		"created_at",
	).
		From(UserTN).
		Where(conds).
		Limit(1).
		ToSql()
}

func (r *User) UpdateLastOnlineSQL(userId uint64) (string, []any, error) {
	return squirrel.Update(UserTN).
		SetMap(map[string]any{
			"last_online": time.Now(),
		}).
		Where(squirrel.Eq{"id": userId}).
		ToSql()
}

func (r *User) UpdatePasswordSQL(userId uint64, password string) (string, []any, error) {
	return squirrel.Update(UserTN).
		SetMap(map[string]any{
			"password": password,
		}).
		Where(squirrel.Eq{"id": userId}).
		ToSql()
}

func (r *User) UpdateUsernameSQL(userId uint64, username string) (string, []any, error) {
	return squirrel.Update(UserTN).
		SetMap(map[string]any{
			"username": username,
		}).
		Where(squirrel.Eq{"id": userId}).
		ToSql()
}

func (r *User) DeleteSQL(conds any) (string, []any, error) {
	if conds == nil {
		panic("conds is nil")
	}

	return squirrel.Delete(UserTN).
		Where(conds).
		Suffix("RETURNING id").
		ToSql()
}
