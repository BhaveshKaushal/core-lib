package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	tests := []struct {
		name string
		readers []ConfigReader
	}{
		{
			name: "Success Initialize",
		},
		{
			name: "Success Initialize with readers",
			readers: []ConfigReader{
				&FileConfigReader{
					
				},
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			err := Init(test.readers...)
			assert.Nil(t, err)
			assert.NotNil(t, appConfigurtion)
			assert.NotNil(t, appConfigurtion.configs)
			assert.NotNil(t, appConfigurtion.finalizedAppConfig)
		})
	}
}



