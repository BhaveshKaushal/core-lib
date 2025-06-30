package conf

import (
	"io"

	"github.com/BhaveshKaushal/base-lib/pkg/base"
	errors "github.com/BhaveshKaushal/base-lib/pkg/errors"
)

var (
	BASE_FILE_PRIORITY = 500
)
type (
	Config struct {
		reader io.Reader
	}
)

func Initialize(app base.App) *errors.Err {
	if app.Name() == "" {
		return errors.NewErrDefault(errors.ErrCodeConfigMissing, "Missing app name", "conf")
	}

	return nil
}

func InitializeWithConfig(name string, config *Config) {
	//TODO Use different config readers to load application specific configurations
}
