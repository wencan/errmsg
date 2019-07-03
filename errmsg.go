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
	"runtime"
	"runtime/debug"
	"strings"
)

const (
	FstdFlag   = 1 << iota // only record error status and message
	Fshortfile             // attach file name and line number, overrides Flongfile
	Flongfile              // attach full file path and line number
)

var flags = FstdFlag

// SetFlags set global flags
func SetFlags(flag int) {
	flags = flag
}

// ErrMsg error detail
type ErrMsg struct {
	error

	Status  ErrStatus `json:"status"`
	Message string    `json:"message"`

	File string `json:"-"`
	Line int    `json:"-"`

	Stack string `json:"stack,omitempty"`
}

// WrapError Wrap a error object, attach error detail.
func WrapError(status ErrStatus, err error) *ErrMsg {
	errMsg := &ErrMsg{
		error:   err,
		Status:  status,
		Message: err.Error(),
	}
	errMsg.appendFileLineIfNeed()
	return errMsg
}

// WrapErrorWithStack Wrap a error object, attach error detail and call stack
func WrapErrorWithStack(status ErrStatus, err error) *ErrMsg {
	errMsg := &ErrMsg{
		error:   err,
		Status:  status,
		Message: err.Error(),
		Stack:   string(debug.Stack()),
	}
	errMsg.appendFileLineIfNeed()
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

// MarshalJSON Implement json.Marshaler
func (errMsg *ErrMsg) MarshalJSON() ([]byte, error) {
	type Alias ErrMsg
	err := Alias(*errMsg)
	return json.Marshal(err)
}
