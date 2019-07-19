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

func TestCode(t *testing.T) {
	var err error = errmsg.WrapError(errmsg.ErrNotFound, errors.New("test"))
	code := errmsg_grpc.Code(err)
	assert.Equal(t, codes.NotFound, code)
}

func TestGRPCCode(t *testing.T) {
	var err error = status.Errorf(codes.Internal, "test")
	code := errmsg_grpc.Code(err)
	assert.Equal(t, codes.Internal, code)
}

func TestOverwriteCode(t *testing.T) {
	errmsg_grpc.BindCode(errmsg.ErrNotFound, codes.InvalidArgument)
	defer errmsg_grpc.BindCode(errmsg.ErrNotFound, codes.NotFound)

	err := errmsg.WrapError(errmsg.ErrNotFound, errors.New("test"))
	code := errmsg_grpc.Code(err)
	assert.Equal(t, codes.InvalidArgument, code)
}

func TestCustomCode(t *testing.T) {
	errCustom := errmsg.ErrStatus(1003)

	err := errmsg.WrapError(errCustom, errors.New("test"))
	code := errmsg_grpc.Code(err)
	assert.Equal(t, codes.Unknown, code)

	errmsg_grpc.BindCode(errCustom, codes.Internal)
	code = errmsg_grpc.Code(err)
	assert.Equal(t, codes.Internal, code)
}
