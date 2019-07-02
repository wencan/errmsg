package errmsg

/*
 * error status
 * error status code copy from https://cloud.google.com/apis/design/errors
 *
 * wencan
 * 2019-07-02
 */

import (
	"fmt"
)

// ErrStatus error status
type ErrStatus uint16

const (
	// ErrOK No error.
	ErrOK ErrStatus = iota

	// ErrInvalidArgument Client specified an invalid argument.
	// Check error message and error details for more information.
	ErrInvalidArgument

	// ErrFailedPrecondition Request can not be executed in the current system state,
	// such as deleting a non-empty directory.
	ErrFailedPrecondition

	// ErrOutOfRange Client specified an invalid range.
	ErrOutOfRange

	// ErrUnauthenticated Request not authenticated due to missing,
	// invalid, or expired OAuth token.
	ErrUnauthenticated

	// ErrPermissionDenied Client does not have sufficient permission.
	// This can happen because the OAuth token does not have the right scopes,
	// the client doesn't have permission,
	// or the API has not been enabled for the client project.
	ErrPermissionDenied

	// ErrNotFound A specified resource is not found,
	// or the request is rejected by undisclosed reasons,
	// such as whitelisting.
	ErrNotFound

	// ErrAborted Concurrency conflict,
	// such as read-modify-write conflict.
	ErrAborted

	// ErrAlreadyExists The resource that a client tried to create already exists.
	ErrAlreadyExists

	// ErrResourceExhausted Either out of resource quota or reaching rate limiting.
	// The client should look for google.rpc.QuotaFailure error detail for more information.
	ErrResourceExhausted

	// ErrCancelled Request cancelled by the client.
	ErrCancelled

	// ErrDataLoss Unrecoverable data loss or data corruption.
	// The client should report the error to the user.
	ErrDataLoss

	// ErrUnknown Unknown server error. Typically a server bug.
	ErrUnknown

	// ErrInternal Internal server error. Typically a server bug.
	ErrInternal

	// ErrNotImplemented API method not implemented by the server.
	ErrNotImplemented

	// ErrUnavailable Service unavailable. Typically the server is down.
	ErrUnavailable

	// ErrDeadlineExceeded  Request deadline exceeded.
	// This will happen only if the caller sets a deadline that is shorter than the method's default deadline
	// (i.e. requested deadline is not enough for the server to process the request) and the request did not finish within the deadline.
	ErrDeadlineExceeded
)

var status2Name = map[ErrStatus]string{
	ErrOK:                 "Ok",
	ErrInvalidArgument:    "InvalidArgument",
	ErrFailedPrecondition: "FailedPrecondition",
	ErrOutOfRange:         "OutOfRange",
	ErrUnauthenticated:    "Unauthenticated",
	ErrPermissionDenied:   "PermissionDenied",
	ErrNotFound:           "NotFound",
	ErrAborted:            "Aborted",
	ErrAlreadyExists:      "AlreadyExists",
	ErrResourceExhausted:  "ResourceExhausted",
	ErrCancelled:          "Cancelled",
	ErrDataLoss:           "DataLoss",
	ErrUnknown:            "Unknown",
	ErrInternal:           "Internal",
	ErrNotImplemented:     "NotImplemented",
	ErrUnavailable:        "Unavailable",
	ErrDeadlineExceeded:   "DeadlineExceeded",
}

func (status ErrStatus) String() string {
	name, exists := status2Name[status]
	if exists {
		return name
	}
	return fmt.Sprintf("Err#%d", status)
}

// BindErrStatus Bind custom error status with name.
func BindErrStatus(status ErrStatus, name string) {
	status2Name[status] = name
}
