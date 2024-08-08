package postgres

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rotisserie/eris"
)

func AsPgError(err error) *pgconn.PgError {
	var pgErr *pgconn.PgError
	if eris.As(err, &pgErr) {
		return pgErr
	}

	return nil
}

func PgErrCodeEquals(err error, code string) bool {
	pgErr := AsPgError(err)

	return pgErr != nil && pgErr.Code == code
}
