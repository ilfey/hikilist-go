package postgres

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func AsPgError(err error) *pgconn.PgError {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr
	}

	return nil
}

func PgErrCodeEquals(err error, code Code) bool {
	pgErr := AsPgError(err)

	return pgErr != nil && pgErr.Code == code
}

func IsNotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
