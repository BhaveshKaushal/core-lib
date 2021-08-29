package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "Success Initialize",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			Initialize()
			assert.NotNil(t, appConfig)
			assert.NotNil(t, appConfig.appConfiguration)
			assert.NotNil(t, appConfig.configs)

		})
	}
}
