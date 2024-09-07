package repositories

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/georgysavva/scany/v2/pgxscan"
)

var (
	ErrUserCreateFailed = errtype.NewInternalRepositoryError("unable to store user")
	ErrUserAlreadyExist = errtype.NewEntityAlreadyExistsError("user", "username")
	ErrUserUpdateFailed = errtype.NewInternalRepositoryError("unable to update user")
	ErrUserNotFoundById = errtype.NewEntityNotFoundError("database", "user", "id")
	ErrUsersFindFailed  = errtype.NewInternalRepositoryError("unable to find users")
	ErrUsersCountFailed = errtype.NewInternalRepositoryError("unable to count users")
	ErrUserGetFailed    = errtype.NewInternalRepositoryError("unable to get user")
	ErrUserDeleteFailed = errtype.NewInternalRepositoryError("unable to delete user")
)

type User struct {
	log    loggerInterface.Logger
	db     postgres.RW
	action repositoryInterface.Action
}

func NewUser(container diInterface.ServiceContainer) (*User, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	db, err := container.GetPostgresDatabase()
	if err != nil {
		return nil, log.Propagate(err)
	}

	action, err := container.GetActionRepository()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &User{
		log:    log,
		action: action,
		db:     db,
	}, nil
}

func (r *User) WithTx(tx postgres.RW) repositoryInterface.User {
	return &User{
		log:    r.log,
		action: r.action,
		db:     tx,
	}
}

func (r *User) Create(ctx context.Context, cm *dto.UserCreateRequestDTO) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		// Create user.
		sql, args, err := r.CreateSQL(cm)
		if err != nil {
			return r.log.Propagate(err)
		}

		err = r.db.QueryRow(ctx, sql, args...).Scan(&cm.UserID)
		if err != nil {
			r.log.Error(err)

			if postgres.PgErrCodeEquals(err, postgres.UniqueViolation) {
				return ErrUserAlreadyExist
			}

			return ErrUserCreateFailed
		}

		// Create action.
		actionCm := dto.NewRegisterUserAction(cm.UserID)

		return r.action.WithTx(tx).Create(ctx, actionCm)
	})
}

func (r *User) Get(ctx context.Context, conds any) (*agg.UserDetail, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var dm agg.UserDetail

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		r.log.Error(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFoundById
		}

		return nil, ErrUserGetFailed
	}

	return &dm, nil
}

func (r *User) Find(ctx context.Context, dto *dto.UserListRequestDTO, conds any) ([]*agg.UserListItem, error) {
	sql, args, err := r.FindWithPaginatorSQL(dto, conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var items []*agg.UserListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Error(err)

		return nil, ErrUsersFindFailed
	}

	return items, nil
}

func (r *User) Count(ctx context.Context, conds any) (uint64, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, r.log.Propagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Error(err)

		return 0, ErrUsersCountFailed
	}

	return count, nil
}

func (r *User) UpdateUsername(ctx context.Context, userId uint64, oldUsername, newUsername string) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		// Update user username.
		sql, args, err := r.UpdateUsernameSQL(userId, newUsername)
		if err != nil {
			return r.log.Propagate(err)
		}

		_, err = r.db.Exec(ctx, sql, args...)
		if err != nil {
			r.log.Error(err)

			if errors.Is(err, pgx.ErrNoRows) {
				return ErrUserNotFoundById
			}

			return ErrUserUpdateFailed
		}

		// Create action
		actionCm := dto.NewUpdateUsernameAction(userId, oldUsername, newUsername)

		return r.action.WithTx(tx).Create(ctx, actionCm)
	})

}

func (r *User) UpdateLastOnline(ctx context.Context, userId uint64) error {
	sql, args, err := r.UpdateLastOnlineSQL(userId)
	if err != nil {
		return r.log.Propagate(err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Error(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFoundById
		}

		return ErrUserUpdateFailed
	}

	return nil
}

func (r *User) UpdatePassword(ctx context.Context, userId uint64, hash string) error {
	sql, args, err := r.UpdatePasswordSQL(userId, hash)
	if err != nil {
		return r.log.Propagate(err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Error(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFoundById
		}

		return ErrUserUpdateFailed
	}

	return nil
}

func (r *User) Delete(ctx context.Context, conds any) error {
	sql, args, err := r.DeleteSQL(conds)
	if err != nil {
		return r.log.Propagate(err)
	}

	var id uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFoundById
		}

		return ErrUserDeleteFailed
	}

	return nil
}
