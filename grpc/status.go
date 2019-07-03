package grpc

/*
 * Convert errmsg.ErrMsg to grpc/status.Status.
 *
 * wencan
 * 2019-07-02
 */

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Status Convert errmsg.ErrMsg to grpc/status.Status.
// Usage: Status(err).Err()
func Status(err error) *status.Status {
	if err == nil {
		return status.New(codes.OK, "")
	}

	code := Code(err)
	return status.New(code, err.Error())
}
