package reader

import (
	"fmt"
	"path/filepath"

	"github.com/BhaveshKaushal/base-lib/pkg/errors"
	"github.com/spf13/afero"
)

const (
	FILE_PATH_ERROR_CODE = 1001
)

type FileConfigReader struct {
	paths    []string
	name     string
	fileType string
	required bool
	fs       afero.Fs
	priority int
}

func NewFileConfigReader(paths []string, required bool, name, fileType string, priority int) (FileConfigReader, error) {
	fcr := FileConfigReader{
		paths:    paths,
		required: required,
		name:     name,
		fileType: fileType,
		priority: priority,
		fs:       afero.NewOsFs(),
	}

	er := fcr.validateAndAbsolutePaths()

	return fcr, er
}

func (fcr FileConfigReader) validateAndAbsolutePaths() error {
	for i, path := range fcr.paths {
		str, err := filepath.Abs(path)
		if err != nil {
			if fcr.required {
				return errors.NewErr(FILE_PATH_ERROR_CODE, err, fmt.Sprintf("File Config Path Error: %s", path), "config")
			} else {
				//TODO use logger to add warning message for non optional file error
			}
		}
		fcr.paths[i] = str

	}
	return nil
}

func (fcr FileConfigReader) GetPriority() int {
	return fcr.priority
}

/*func (fcr *FileConfigReader) ReadConfig() (map[string]interface{}, errors.Err) {
	configMap := make(map[string]interface{})

}*/
