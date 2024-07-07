package httphandlers

import (
	"fmt"
	"net/http"

	"github.com/johnfercher/go-hexagonal/internal/adapters/drivens/memdb"
)

type Error struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Inner  error  `json:"inner"`
}

func Err(status int, code string) *Error {
	e := &Error{
		Status: status,
		Code:   code,
	}

	return e
}

func (e Error) WithInner(err error) *Error {
	return &Error{
		Status: e.Status,
		Code:   e.Code,
		Inner:  err,
	}
}

func (e Error) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf("%d %s", e.Status, e.Code)
	}

	return fmt.Sprintf("%d %s %s", e.Status, e.Code, e.Inner.Error())
}

var (
	ErrUnmarshallRequest = Err(http.StatusInternalServerError, "cannot_unmarshall")
	ErrUnknown           = Err(http.StatusInternalServerError, "unknown_error")
)

var mapErr = map[error]*Error{
	memdb.ErrUserAlreadyCreated: Err(http.StatusBadRequest, memdb.ErrUserAlreadyCreated.Error()),
}
