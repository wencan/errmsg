package errmsg

/*
 * Convert ErrStatus to HTTP status
 *
 * wencan
 * 2019-07-02
 */

import "net/http"

var status2HTTPStatus = map[ErrStatus]int{
	ErrOK:                 http.StatusOK,
	ErrInvalidArgument:    http.StatusBadRequest,
	ErrFailedPrecondition: http.StatusBadRequest,
	ErrOutOfRange:         http.StatusBadRequest,
	ErrUnauthenticated:    http.StatusUnauthorized,
	ErrPermissionDenied:   http.StatusForbidden,
	ErrNotFound:           http.StatusNotFound,
	ErrAborted:            http.StatusConflict,
	ErrAlreadyExists:      http.StatusConflict,
	ErrResourceExhausted:  http.StatusTooManyRequests,
	ErrCancelled:          499, // A non-standard status code introduced by nginx
	ErrDataLoss:           http.StatusInternalServerError,
	ErrUnknown:            http.StatusInternalServerError,
	ErrInternal:           http.StatusInternalServerError,
	ErrNotImplemented:     http.StatusNotImplemented,
	ErrUnavailable:        http.StatusServiceUnavailable,
	ErrDeadlineExceeded:   http.StatusGatewayTimeout,
}

// HTTPStatus Get HTTP status by error status
func HTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	errMsg, ok := err.(*ErrMsg)
	if ok {
		httpStatus, exists := status2HTTPStatus[errMsg.Status]
		if exists {
			return httpStatus
		}
	}

	return http.StatusInternalServerError
}

// BindHTTPStatus Bind HTTP status with error status
func BindHTTPStatus(status ErrStatus, httpStatus int) {
	status2HTTPStatus[status] = httpStatus
}

// StatusCode Implement kit/transport/http.StatusCoder
func (errMsg *ErrMsg) StatusCode() int {
	return HTTPStatus(errMsg)
}
