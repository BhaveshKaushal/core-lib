package reader

import (
	errors "github.com/BhaveshKaushal/base-lib/pkg/errors"
)

type (
	ConfigReader interface {
		priority
		ReadConfig() errors.Error
	}
	priority interface {
		GetPriority() int
	}

	appConfigDetails struct {
		//sorted order of config
		configOrder []int

		//Configurations available for the app
		configs map[int]map[string]interface{}

		//Final configuration after merging all the available configs based on priority
		appConfiguration map[string]interface{}
	}
)

var appConfig *appConfigDetails

func Initialize(){
	appConfig = &appConfigDetails{
		configs:          make(map[int]map[string]interface{}),
		appConfiguration: make(map[string]interface{}),
	}
}

/*func (acd *appConfigDetails) AddReader(){
	
}*/