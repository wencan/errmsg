package zap

/*
 * Convert ErrMsg to zap.Field
 *
 * wencan
 * 2019-07-03
 */

import (
	"github.com/wencan/errmsg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Marshaler Marshale errmsg.ErrMsg as zap.Field
//
//    field = zap.Object(key, Marshaler(err))
func Marshaler(err error) zapcore.ObjectMarshalerFunc {
	return func(enc zapcore.ObjectEncoder) error {
		errMsg := errmsg.Unwrap(err)

		enc.AddString("status", errMsg.Status.String())
		enc.AddString("message", errMsg.Message)
		if errMsg.File != "" {
			enc.AddString("file", errMsg.File)
		}
		if errMsg.Line != 0 {
			enc.AddInt("line", errMsg.Line)
		}
		if errMsg.Stack != "" {
			enc.AddString("stack", errMsg.Stack)
		}

		return nil
	}
}

// Fields Convert errmsg.ErrMsg as []zap.Field
func Fields(err error) []zap.Field {
	errMsg := errmsg.Unwrap(err)

	fields := []zap.Field{}
	fields = append(fields, zap.String("status", errMsg.Status.String()))
	fields = append(fields, zap.String("message", errMsg.Message))
	if errMsg.File != "" {
		fields = append(fields, zap.String("file", errMsg.File))
	}
	if errMsg.Line != 0 {
		fields = append(fields, zap.Int("line", errMsg.Line))
	}
	if errMsg.Stack != "" {
		fields = append(fields, zap.String("stack", errMsg.Stack))
	}

	return fields
}
