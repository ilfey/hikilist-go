package repositories

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/ilfey/hikilist-go/pkg/models/action"
	"github.com/ilfey/hikilist-go/pkg/models/user"
)

type User interface {
	WithTx(tx DBRW) User

	Create(ctx context.Context, cm *user.CreateModel) error
	Get(ctx context.Context, conds any) (*user.DetailModel, error)
	FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*user.ListItemModel, error)
	Count(ctx context.Context, conds any) (uint, error)
	UpdateUsername(ctx context.Context, userId uint, oldUsername, newUsername string) error
	UpdateLastOnline(ctx context.Context, userId uint) error
	UpdatePassword(ctx context.Context, userId uint, hash string) error
	Delete(ctx context.Context, conds any) error
}

type UserImpl struct {
	db     DBRW
	action Action
}

func NewUser(db DBRW, actionRepo Action) User {
	return &UserImpl{
		db:     db,
		action: actionRepo,
	}
}

func (r *UserImpl) WithTx(tx DBRW) User {
	return &UserImpl{
		action: r.action,
		db: tx,
	}
}

func (r *UserImpl) Create(ctx context.Context, cm *user.CreateModel) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		sql, args, err := r.CreateSQL(cm)
		if err != nil {
			return err
		}

		err = r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
		if err != nil {
			return err
		}

		// Create action
		actionCm := action.NewRegisterUserAction(cm.ID)

		return r.action.WithTx(tx).Create(ctx, actionCm)
	})
}

func (r *UserImpl) Get(ctx context.Context, conds any) (*user.DetailModel, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, err
	}

	var dm user.DetailModel

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		return nil, err
	}

	return &dm, nil
}

func (r *UserImpl) FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*user.ListItemModel, error) {
	sql, args, err := r.FindWithPaginatorSQL(p, conds)
	if err != nil {
		return nil, err
	}

	var items []*user.ListItemModel

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *UserImpl) Count(ctx context.Context, conds any) (uint, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, err
	}

	var count uint

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *UserImpl) UpdateUsername(ctx context.Context, userId uint, oldUsername, newUsername string) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		sql, args, err := r.UpdateUsernameSQL(userId, newUsername)
		if err != nil {
			return err
		}

		_, err = r.db.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}

		// Create action
		actionCm := action.NewUpdateUsernameAction(userId, oldUsername, newUsername)

		return r.action.WithTx(tx).Create(ctx, actionCm)
	})

}

func (r *UserImpl) UpdateLastOnline(ctx context.Context, userId uint) error {
	sql, args, err := r.UpdateLastOnlineSQL(userId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserImpl) UpdatePassword(ctx context.Context, userId uint, hash string) error {
	sql, args, err := r.UpdatePasswordSQL(userId, hash)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserImpl) Delete(ctx context.Context, conds any) error {
	sql, args, err := r.DeleteSQL(conds)
	if err != nil {
		return err
	}

	var id uint

	return r.db.QueryRow(ctx, sql, args...).Scan(&id)
}
