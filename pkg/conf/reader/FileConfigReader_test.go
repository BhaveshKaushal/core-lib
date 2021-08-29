package reader

import (
	"os"
	"testing"

	"github.com/BhaveshKaushal/base-lib/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidateAndAbsolutePaths(t *testing.T) {
	tests := []struct {
		name   string
		input  *FileConfigReader
		output *errors.Err
	}{
		{
			name:   "Validation Success",
			input:  NewFileConfigReader([]string{os.Getenv("HOME")}, false, "base.yaml", "yaml"),
			output: nil,
		},
		
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			err := test.input.validateAndAbsolutePaths()
			assert.Equal(t, test.output, err)
		})
	}
}
