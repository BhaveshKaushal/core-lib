package conf

import (
	"testing"

	"github.com/BhaveshKaushal/base-lib/pkg/base"
	errors "github.com/BhaveshKaushal/base-lib/pkg/errors"
	"github.com/BhaveshKaushal/base-lib/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {

	tests := []struct {
		name           string
		expectedOutput *errors.Err
		app            base.App
	}{
		{
			name:           "Success Initialize",
			expectedOutput: nil,
			app:            mocks.NewMock("testApp"),
		},
		{
			name:           "Missing app name",
			expectedOutput: errors.NewErrDefault(errors.ErrCodeConfigMissing, "Missing app name", "conf"),
			app:            mocks.NewMock(""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Initialize(test.app)
			
			if test.expectedOutput == nil {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, test.expectedOutput.Code(), err.Code())
				assert.Equal(t, test.expectedOutput.Message(), err.Message())
			}
		})
	}
}
