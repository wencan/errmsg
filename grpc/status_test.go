package grpc_test

/*
 * wencan
 * 2019-07-02
 */

import (
	"errors"
	"testing"

	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/assert"
	"github.com/wencan/errmsg"
	errmsg_grpc "github.com/wencan/errmsg/grpc"
	"google.golang.org/grpc/codes"
)

func TestStatus(t *testing.T) {
	errString := "this is a test"
	err := errmsg.WrapError(errmsg.ErrNotFound, errors.New(errString))
	st := errmsg_grpc.Status(err)
	assert.Equal(t, codes.NotFound, st.Code())
	assert.Equal(t, errString, st.Message())
	assert.Equal(t, codes.NotFound, status.Code(st.Err()))
}

func TestCustomStatus(t *testing.T) {
	errCustom := errmsg.ErrStatus(1004)

	err := errmsg.WrapError(errCustom, errors.New("test"))
	st := errmsg_grpc.Status(err)
	assert.Equal(t, codes.Unknown, st.Code())
	assert.Equal(t, codes.Unknown, status.Code(st.Err()))

	errmsg_grpc.BindCode(errCustom, codes.Internal)
	st = errmsg_grpc.Status(err)
	assert.Equal(t, codes.Internal, st.Code())
	assert.Equal(t, codes.Internal, status.Code(st.Err()))
}
