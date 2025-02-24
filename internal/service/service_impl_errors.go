package service

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type GenericResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e *GenericResponse) Error() string {
	return fmt.Sprintf("status %d: %s", e.StatusCode, e.Message)
}

func isMySQLDuplicateEntryError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		return true
	}

	return false
}
