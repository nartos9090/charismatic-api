package helpers_errors

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func FromSql(err error) Error {
	log.Print(err)

	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		switch err {
		case sql.ErrNoRows:
			return *NotFoundError
		}
	} else {
		switch sqlErr.Number {
		case 1062:
			return *DuplicateEntryError
		}
	}

	return *InternalServerError
}
