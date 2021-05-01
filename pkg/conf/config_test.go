package conf

import(
  "testing"
  "github.com/BhaveshKaushal/base-lib/pkg/mocks"
  "github.com/stretchr/testify/assert" 
)

var (
		mockApp = &mocks.MockApp{}
)
func TestInitialize(t *testing.T) {
	
	tests := []struct {
		name           string
		expectedOutput interface{}
	}{
		{
			name: "Success Initilaize",
			expectedOutput: (*error)(nil),
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			err := Initialize(mockApp)
			assert.Equal(t,test.expectedOutput, err)
		})
	}
}