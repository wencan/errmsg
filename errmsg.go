package errmsg

/*
 * errmsg core
 * support kit/transport.DefaultErrorEncoder
 *
 * wencan
 * 2019-07-02
 */

import (
	"encoding/json"
	"errors"
	"runtime"
	"runtime/debug"
	"strings"
)

const (
	// FstdFlag only record error status and message
	FstdFlag = 1 << iota

	// Fshortfile attach file name and line number, overrides Flongfile
	Fshortfile

	// Flongfile attach full file path and line number
	Flongfile
)

var flags = FstdFlag

var (
	// NoError No error
	NoError = errors.New("OK")

	// NoErrMsg No error
	NoErrMsg = &ErrMsg{error: NoError, Status: ErrOK}
)

// SetFlags set global flags
func SetFlags(flag int) {
	flags = flag
}

// ErrMsg error detail
type ErrMsg struct {
	error

	Status  ErrStatus `json:"-"`
	Message string    `json:"message"`

	File string `json:"-"`
	Line int    `json:"-"`

	Stack string `json:"stack,omitempty"`
}

func wrapError(status ErrStatus, err error) *ErrMsg {
	if err == nil {
		panic("err must non-empty")
	}

	return &ErrMsg{
		error:   err,
		Status:  status,
		Message: err.Error(),
	}
}

// WrapError Wrap a error object, attach error detail.
func WrapError(status ErrStatus, err error) *ErrMsg {
	errMsg := wrapError(status, err)
	errMsg.appendFileLineIfNeed()
	return errMsg
}

// WrapErrorWithStack Wrap a error object, attach error detail and call stack
func WrapErrorWithStack(status ErrStatus, err error) *ErrMsg {
	errMsg := wrapError(status, err)
	errMsg.appendFileLineIfNeed()
	errMsg.Stack = string(debug.Stack())
	return errMsg
}

func (errMsg *ErrMsg) appendFileLineIfNeed() {
	if flags&(Flongfile|Fshortfile) != 0 {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		} else if flags&Fshortfile != 0 {
			parts := strings.Split(file, "/")
			file = parts[len(parts)-1]
		}
		errMsg.File = file
		errMsg.Line = line
	}
}

// MarshalJSON Implement json.Marshaler.
// Provide to the kit/transport.DefaultErrorEncoder.
func (errMsg *ErrMsg) MarshalJSON() ([]byte, error) {
	type Alias struct {
		Status string `json:"status"`
		ErrMsg
	}
	err := Alias{
		ErrMsg: *errMsg,
		Status: errMsg.Status.String(),
	}
	return json.Marshal(err)
}

// Unwrap Unwrap a error to *ErrMsg.
// return a NoErrMsg if err is nil;
// return a ErrMsg with unknown status if err not is instance of ErrMsg.
func Unwrap(err error) *ErrMsg {
	if err == nil {
		return NoErrMsg
	}

	errMsg, ok := err.(*ErrMsg)
	if !ok {
		errMsg = wrapError(ErrUnknown, err)
	}

	return errMsg
}
