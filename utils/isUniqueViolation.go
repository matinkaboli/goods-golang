package utils

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return true
		}
	}

	return false
}
