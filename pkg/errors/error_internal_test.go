package errors

import (
	"github.com/pkg/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErr(t *testing.T) {
	msg := "error message"
	er := errors.New(msg)
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
			er := NewErr(test.code, test.msg)
			assert.Equal(t, test.expectedOutput.Code(), er.Code())
			assert.Equal(t, test.expectedOutput.Message(), er.Message())
			assert.Equal(t, test.expectedOutput.Er().Error(), er.Er().Error())
		})
	}
}
