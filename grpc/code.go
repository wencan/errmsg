package grpc

/*
 * Convert errmsg.ErrStatus to grpc/codes.Code
 * wencan
 * 2019-07-02
 */

import (
	"github.com/wencan/errmsg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var status2code = map[errmsg.ErrStatus]codes.Code{
	errmsg.ErrOK:                 codes.OK,
	errmsg.ErrInvalidArgument:    codes.InvalidArgument,
	errmsg.ErrFailedPrecondition: codes.FailedPrecondition,
	errmsg.ErrOutOfRange:         codes.OutOfRange,
	errmsg.ErrUnauthenticated:    codes.Unauthenticated,
	errmsg.ErrPermissionDenied:   codes.PermissionDenied,
	errmsg.ErrNotFound:           codes.NotFound,
	errmsg.ErrAborted:            codes.Aborted,
	errmsg.ErrAlreadyExists:      codes.AlreadyExists,
	errmsg.ErrResourceExhausted:  codes.ResourceExhausted,
	errmsg.ErrCancelled:          codes.Canceled,
	errmsg.ErrDataLoss:           codes.DataLoss,
	errmsg.ErrUnknown:            codes.Unknown,
	errmsg.ErrInternal:           codes.Internal,
	errmsg.ErrNotImplemented:     codes.Unimplemented,
	errmsg.ErrUnavailable:        codes.Unavailable,
	errmsg.ErrDeadlineExceeded:   codes.DeadlineExceeded,
}

// Code get GRPC code by error status
func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	code := status.Code(err)
	if code != codes.Unknown {
		return code
	}

	errMsg, ok := err.(*errmsg.ErrMsg)
	if ok {
		code, exists := status2code[errMsg.Status]
		if exists {
			return code
		}
	}

	return codes.Unknown
}

// BindCode Bind GRPC code with error status
func BindCode(status errmsg.ErrStatus, code codes.Code) {
	status2code[status] = code
}
