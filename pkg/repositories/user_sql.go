package repositories

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/user"
)

func (r *UserImpl) CountSQL(conds any) (string, []any, error) {
	b := squirrel.Select("COUNT(*)").
		From(user.TableName)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.ToSql()
}

func (r *UserImpl) FindWithPaginatorSQL(p *paginate.Paginator, conds any) (string, []any, error) {
	b := squirrel.Select(
		"id",
		"username",
		"created_at",
	).
		From(user.TableName)

	if conds != nil {
		b = b.Where(conds)
	}

	return b.
		OrderBy(p.Order.ToQuery()).
		Offset(uint64(p.GetOffset(p.Page, p.Limit))).
		Limit(uint64(p.Limit)).
		ToSql()
}

func (r *UserImpl) CreateSQL(cm *user.CreateModel) (string, []any, error) {
	return squirrel.Insert(user.TableName).
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

func (r *UserImpl) GetSQL(conds any) (string, []any, error) {
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
		From(user.TableName).
		Where(conds).
		Limit(1).
		ToSql()
}

func (r *UserImpl) UpdateLastOnlineSQL(userId uint) (string, []any, error) {
	return squirrel.Update(user.TableName).
		SetMap(map[string]any{
			"last_online": time.Now(),
		}).
		Where(squirrel.Eq{"id": userId}).
		ToSql()
}

func (r *UserImpl) UpdatePasswordSQL(userId uint, password string) (string, []any, error) {
	return squirrel.Update(user.TableName).
		SetMap(map[string]any{
			"password": password,
		}).
		Where(squirrel.Eq{"id": userId}).
		ToSql()
}

func (r *UserImpl) UpdateUsernameSQL(userId uint, username string) (string, []any, error) {
	return squirrel.Update(user.TableName).
		SetMap(map[string]any{
			"username": username,
		}).
		Where(squirrel.Eq{"id": userId}).
		ToSql()
}

func (r *UserImpl) DeleteSQL(conds any) (string, []any, error) {
	if conds == nil {
		panic("conds is nil")
	}

	return squirrel.Delete(user.TableName).
		Where(conds).
		Suffix("RETURNING id").
		ToSql()
}
