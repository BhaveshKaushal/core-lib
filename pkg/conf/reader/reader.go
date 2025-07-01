package reader

import (
	"github.com/BhaveshKaushal/base-lib/pkg/errors"
	"github.com/BhaveshKaushal/base-lib/pkg/logger"
)

type (
	ConfigReader interface {
		priority
		ReadConfig() (map[string]interface{}, error)
	}
	priority interface {
		GetPriority() int
	}

	appConfig struct {
		//sorted order of config
		//configPriority []int

		//Configurations available for the app
		configs map[int]map[string]interface{}

		//Final configuration after merging all the available configs based on priority
		finalizedAppConfig map[string]interface{}
	}
)

var (
	appConfigurtion *appConfig
)

func Init(readers ...ConfigReader) error {
	var err error
	appConfigurtion, err = newAppConfig(readers...)
	return err
}

func newAppConfig(readers ...ConfigReader) (*appConfig, error) {
	appConfig := &appConfig{
		configs:            make(map[int]map[string]interface{}),
		finalizedAppConfig: make(map[string]interface{}),
	}

	appConfig.AddReader(readers...)

	return appConfig, nil
}

func (ac *appConfig) AddReader(readers ...ConfigReader) {
	for _, reader := range readers {
		config, err := reader.ReadConfig()
		if err != nil {
			customErr := errors.NewErr(errors.ErrCodeConfig, err, "Error reading config", "ConfigReader")
			logger.Fatal("Error reading config", customErr, nil)
		}
		ac.configs[reader.GetPriority()] = config
	}
}
