package errmsg_test

/*
 * wencan
 * 2019-07-02
 */

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wencan/errmsg"
)

func TestStatusName(t *testing.T) {
	errMsg := errmsg.WrapError(errmsg.ErrInvalidArgument, errors.New("test"))
	assert.Equal(t, "InvalidArgument", errMsg.Status.String())
}

func TestStatusNameValueMapping(t *testing.T) {
	for status := errmsg.ErrOK; status <= errmsg.ErrDeadlineExceeded; status++ {
		assert.Equal(t, status, errmsg.FromStatusName(status.String()))
	}
}

func TestCustomStatus(t *testing.T) {
	ErrCustom := errmsg.ErrStatus(1001)
	name := "Custom"
	errmsg.BindErrStatus(ErrCustom, name)

	assert.Equal(t, ErrCustom, errmsg.FromStatusName(name))
	assert.Equal(t, name, ErrCustom.String())

	errMsg := errmsg.WrapError(ErrCustom, errors.New("test"))
	assert.Equal(t, ErrCustom, errMsg.Status)
	assert.Equal(t, name, errMsg.Status.String())
}
