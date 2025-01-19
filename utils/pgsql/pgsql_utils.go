package pgsql

import (
	"my-api/utils/errors"
	"strings"

	"github.com/lib/pq"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Code {
	case "23505":
		return errors.NewBadRequestError("unique_violation")
	}
	return errors.NewInternalServerError("error processing request")
}
