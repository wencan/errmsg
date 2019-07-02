package errmsg_test

/*
 * wencan
 * 2019-07-02
 */

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wencan/errmsg"
)

func TestHTTPStatus(t *testing.T) {
	errMsg := errmsg.WrapError(errmsg.ErrNotFound, errors.New("test"))
	httpStatus := errmsg.HTTPStatus(errMsg)
	assert.Equal(t, http.StatusNotFound, httpStatus)
}

func TestOverwiteHTTPStatus(t *testing.T) {
	errmsg.BindHTTPStatus(errmsg.ErrAborted, http.StatusForbidden)
	defer errmsg.BindHTTPStatus(errmsg.ErrAborted, http.StatusConflict)

	errMsg := errmsg.WrapError(errmsg.ErrAborted, errors.New("test"))
	httpStatus := errmsg.HTTPStatus(errMsg)
	assert.Equal(t, http.StatusForbidden, httpStatus)
}

func TestCustomHTTPStatus(t *testing.T) {
	ErrCustom := errmsg.ErrStatus(1002)
	errMsg := errmsg.WrapError(ErrCustom, errors.New("test"))
	httpStatus := errmsg.HTTPStatus(errMsg)
	assert.Equal(t, http.StatusInternalServerError, httpStatus)

	errmsg.BindHTTPStatus(ErrCustom, http.StatusForbidden)
	httpStatus = errmsg.HTTPStatus(errMsg)
	assert.Equal(t, http.StatusForbidden, httpStatus)
}
