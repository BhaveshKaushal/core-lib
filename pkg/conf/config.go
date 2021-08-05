package conf

import (
	"github.com/BhaveshKaushal/base-lib/pkg/base"
	errors"github.com/BhaveshKaushal/base-lib/pkg/errors"
)


func Initialize(app base.App) *errors.Err {
	if app.Name() == "" {
		return errors.MissingAppName
	}

	return nil
}