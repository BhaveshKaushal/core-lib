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
			expectedOutput: (*errors.Err)(nil),
			app:            mocks.NewMock("testApp"),
		},
		{
			name:           "Missing app name",
			expectedOutput: errors.MissingAppName,
			app:            mocks.NewMock(""),
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			err := Initialize(test.app)
			assert.Equal(t, test.expectedOutput, err)
		})
	}
}
