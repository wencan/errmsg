package logrus

/*
 * Convert ErrMsg to logrus.Fields
 *
 * wencan
 * 2019-07-05
 */

import (
	"github.com/sirupsen/logrus"
	"github.com/wencan/errmsg"
)

// Fields Convert ErrMsg to logrus.Fields
func Fields(err error) logrus.Fields {
	errMsg := errmsg.Unwrap(err)

	fields := make(logrus.Fields)
	fields["status"] = errMsg.Status.String()
	fields["message"] = errMsg.Message
	if errMsg.File != "" {
		fields["file"] = errMsg.File
	}
	if errMsg.Line != 0 {
		fields["line"] = errMsg.Line
	}
	if errMsg.Stack != "" {
		fields["stack"] = errMsg.Stack
	}

	return fields
}
