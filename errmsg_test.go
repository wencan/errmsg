package errmsg_test

/*
 * test ErrMsg basic function
 *
 * wencan
 * 2019-07-02
 */

import (
	"errors"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wencan/errmsg"
)

func TestErrMsg(t *testing.T) {
	errString := "This is a test"
	err := errors.New(errString)
	errMsg := errmsg.WrapError(errmsg.ErrUnavailable, err)
	assert.Implements(t, (*error)(nil), errMsg)
	assert.Equal(t, errmsg.ErrUnavailable, errMsg.Status)
	assert.Equal(t, errString, errMsg.Message)
	assert.Equal(t, "", errMsg.Stack)
}

func TestErrMsgWithFileLine(t *testing.T) {
	errString := "This is a test"
	err := errors.New(errString)

	errmsg.SetFlags(errmsg.FstdFlag)
	errMsg := errmsg.WrapError(errmsg.ErrUnavailable, err)
	assert.Equal(t, "", errMsg.File)
	assert.Equal(t, 0, errMsg.Line)
	assert.Equal(t, "", errMsg.Stack)

	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Flongfile)
	errMsg = errmsg.WrapError(errmsg.ErrUnavailable, err)
	_, file, _, ok := runtime.Caller(0)
	if ok {
		assert.Equal(t, file, errMsg.File)
	}
	assert.NotEqual(t, 0, errMsg.Line)
	assert.Equal(t, "", errMsg.Stack)

	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Fshortfile)
	errMsg = errmsg.WrapError(errmsg.ErrUnavailable, err)
	assert.Equal(t, "errmsg_test.go", errMsg.File)
	assert.NotEqual(t, 0, errMsg.Line)
	assert.Equal(t, "", errMsg.Stack)
}

func TestErrMsgWithStack(t *testing.T) {
	errString := "This is a test"
	err := errors.New(errString)
	errMsg := errmsg.WrapErrorWithStack(errmsg.ErrUnavailable, err)
	assert.Implements(t, (*error)(nil), errMsg)
	assert.Equal(t, errmsg.ErrUnavailable, errMsg.Status)
	assert.Equal(t, errString, errMsg.Message)
	assert.NotEqual(t, "", errMsg.Stack)
}

func TestUnwrap(t *testing.T) {
	errMsg := errmsg.Unwrap(nil)
	assert.Equal(t, errmsg.ErrOK, errMsg.Status)

	err := errors.New("this is a test")
	errMsg = errmsg.Unwrap(err)
	assert.Equal(t, errmsg.ErrUnknown, errMsg.Status)
	assert.Equal(t, err.Error(), errMsg.Message)

	errMsg = errmsg.Unwrap(errmsg.WrapError(errmsg.ErrCancelled, err))
	assert.Equal(t, errmsg.ErrCancelled, errMsg.Status)
	assert.Equal(t, err.Error(), errMsg.Message)
}
