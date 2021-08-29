package errors

import (
	"github.com/pkg/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)
const(
	msg = "error message"
	
)
var (
	er = errors.New(msg)
)
func TestNewErr(t *testing.T) {
	
	tests := []struct {
		name           string
		code           int
		msg            string
		expectedOutput *Err
	}{
		{
			name:           "TestNewErr",
			code:           100,
			msg:            msg,
			expectedOutput: &Err{code: 100, message: msg, er: er},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			er := NewErrDefault(test.code, test.msg,"")
			assert.Equal(t, test.expectedOutput.Code(), er.Code())
			assert.Equal(t, test.expectedOutput.Message(), er.Message())
			assert.Equal(t, test.expectedOutput.Er().Error(), er.Er().Error())
		})
	}
}

func TestCause(t *testing.T) {
	er := errors.New(msg)
	tests := []struct {
		name           string
		code           int
		msg            string
		expectedOutput *Err
	}{
		{
			name:           "TestCause",
			code:           100,
			msg:            msg,
			expectedOutput: &Err{code: 100, message: msg, er: er},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			er := NewErrDefault(test.code, test.msg,"")
			assert.Equal(t, test.expectedOutput.Cause().Error(), er.Cause().Error())
			assert.Equal(t, test.expectedOutput.Message(), er.Message())
			assert.Equal(t, test.expectedOutput.Er().Error(), er.Er().Error())
		})
	}
}

func TestWrap(t *testing.T) {
	er := errors.New(msg)
	tests := []struct {
		name           string
		code           int
		msg            string
		expectedOutput *Err
	}{
		{
			name:           "TestWrap",
			code:           100,
			msg:            msg,
			expectedOutput: &Err{code: 100, message: msg, er: er},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			er := NewErr(test.code,  errors.New(test.msg), test.msg,"error")
			assert.Equal(t, test.expectedOutput.Wrap(msg).Error(), er.Wrap(msg).Error())
			assert.Equal(t, test.expectedOutput.Message(), er.Message())
			assert.Equal(t, test.expectedOutput.Wrap(msg).Error(), er.Wrap(msg).Error())
		})
	}
}

