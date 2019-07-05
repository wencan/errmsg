package zap_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wencan/errmsg"

	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	errmsg_zap "github.com/wencan/errmsg/logging/zap"
	"go.uber.org/zap"
)

func takeFieldContent(fields ...zap.Field) map[string]interface{} {
	core, hook := observer.New(zapcore.DebugLevel)
	logger := zap.New(core)

	logger.Info("this is a test", fields...)

	entries := hook.All()
	entry := entries[0]
	return entry.ContextMap()
}

func TestMarshaler(t *testing.T) {
	marshaler := errmsg_zap.Marshaler(nil)
	field := zap.Object("object", marshaler)
	kvs := takeFieldContent(field)
	kvs = kvs["object"].(map[string]interface{})
	assert.Equal(t, errmsg.ErrOK.String(), kvs["status"])

	err := errors.New("test")
	marshaler = errmsg_zap.Marshaler(err)
	field = zap.Object("object", marshaler)
	kvs = takeFieldContent(field)
	kvs = kvs["object"].(map[string]interface{})
	assert.Equal(t, errmsg.ErrUnknown.String(), kvs["status"])
	assert.Equal(t, err.Error(), kvs["message"])

	errMsg := errmsg.WrapError(errmsg.ErrCancelled, err)
	marshaler = errmsg_zap.Marshaler(errMsg)
	field = zap.Object("object", marshaler)
	kvs = takeFieldContent(field)
	kvs = kvs["object"].(map[string]interface{})
	assert.Equal(t, errmsg.ErrCancelled.String(), kvs["status"])
	assert.Equal(t, err.Error(), kvs["message"])
}

func TestFields(t *testing.T) {
	fields := errmsg_zap.Fields(nil)
	kvs := takeFieldContent(fields...)
	assert.Equal(t, errmsg.ErrOK.String(), kvs["status"])

	err := errors.New("test")
	fields = errmsg_zap.Fields(err)
	kvs = takeFieldContent(fields...)
	assert.Equal(t, errmsg.ErrUnknown.String(), kvs["status"])
	assert.Equal(t, err.Error(), kvs["message"])

	errMsg := errmsg.WrapError(errmsg.ErrCancelled, err)
	fields = errmsg_zap.Fields(errMsg)
	kvs = takeFieldContent(fields...)
	assert.Equal(t, errmsg.ErrCancelled.String(), kvs["status"])
	assert.Equal(t, err.Error(), kvs["message"])
}
