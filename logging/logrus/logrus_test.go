package logrus_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wencan/errmsg"
	errmsg_logrus "github.com/wencan/errmsg/logging/logrus"
)

func TestFields(t *testing.T) {
	fields := errmsg_logrus.Fields(nil)
	assert.Equal(t, errmsg.ErrOK.String(), fields["status"])

	err := errors.New("test")
	fields = errmsg_logrus.Fields(err)
	assert.Equal(t, errmsg.ErrUnknown.String(), fields["status"])
	assert.Equal(t, err.Error(), fields["message"])

	errMsg := errmsg.WrapError(errmsg.ErrCancelled, err)
	fields = errmsg_logrus.Fields(errMsg)
	assert.Equal(t, errmsg.ErrCancelled.String(), fields["status"])
	assert.Equal(t, err.Error(), fields["message"])
}
