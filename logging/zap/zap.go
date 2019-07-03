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
// Usage: field = zap.Object(key, Marshaler(err))
func Marshaler(err error) zapcore.ObjectMarshalerFunc {
	return func(enc zapcore.ObjectEncoder) error {
		if err == nil {
			return nil
		}

		errMsg, ok := err.(*errmsg.ErrMsg)
		if ok {
			enc.AddString("status", errMsg.Status.String())
			enc.AddString("message", errMsg.Message)
			if errMsg.File != "" {
				enc.AddString("file", errMsg.File)
			}
			if errMsg.Line != 0 {
				enc.AddInt("line", errMsg.Line)
			}
		} else {
			enc.AddString("errorString", err.Error())
		}
		return nil
	}
}

// Fields Convert errmsg.ErrMsg as []zap.Field
func Fields(err error) []zap.Field {
	if err == nil {
		return []zap.Field{}
	}

	fields := []zap.Field{}
	errMsg, ok := err.(*errmsg.ErrMsg)
	if ok {
		fields = append(fields, zap.String("status", errMsg.Status.String()))
		fields = append(fields, zap.String("message", errMsg.Message))
		if errMsg.File != "" {
			fields = append(fields, zap.String("file", errMsg.File))
		}
		if errMsg.Line != 0 {
			fields = append(fields, zap.Int("line", errMsg.Line))
		}
	} else {
		fields = append(fields, zap.Error(err))
	}
	return fields
}
