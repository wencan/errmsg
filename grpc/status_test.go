package grpc_test

/*
 * wencan
 * 2019-07-02
 */

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wencan/errmsg"
	errmsg_grpc "github.com/wencan/errmsg/grpc"
	"google.golang.org/grpc/codes"
)

func TestStatus(t *testing.T) {
	errString := "this is a test"
	err := errmsg.WrapError(errmsg.ErrNotFound, errors.New(errString))
	status := errmsg_grpc.Status(err)
	assert.Equal(t, codes.NotFound, status.Code())
	assert.Equal(t, errString, status.Message())
}

func TestCustomStatus(t *testing.T) {
	errCustom := errmsg.ErrStatus(1004)

	err := errmsg.WrapError(errCustom, errors.New("test"))
	status := errmsg_grpc.Status(err)
	assert.Equal(t, codes.Unknown, status.Code())

	errmsg_grpc.BindCode(errCustom, codes.Internal)
	status = errmsg_grpc.Status(err)
	assert.Equal(t, codes.Internal, status.Code())
}
