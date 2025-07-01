package reader

import (
	"fmt"
	"path/filepath"

	"github.com/BhaveshKaushal/base-lib/pkg/errors"
	"github.com/BhaveshKaushal/base-lib/pkg/logger"
	"github.com/spf13/afero"
)

const (
	FILE_PATH_ERROR_CODE = errors.ErrCodeConfigFile
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

	er := fcr.addAbsoultePaths()

	return fcr, er
}

func (fcr FileConfigReader) addAbsoultePaths() error {
	for i, path := range fcr.paths {
		str, err := filepath.Abs(path)
		if err != nil {
			if fcr.required {
				return errors.NewErr(FILE_PATH_ERROR_CODE, err, fmt.Sprintf("File Config Path Error: %s", path), "config")
			} else {
				// Log warning message for non-optional file error
				logger.Warn(
					"Failed to resolve absolute path for optional config file", 
					map[string]interface{}{
						"path":      path,
						"error":     err.Error(),
						"file_name": fcr.name,
						"file_type": fcr.fileType,
						"code":      errors.ErrCodeConfigFile,
					},
				)
				continue // Skip this path and continue with others
			}
		}
		fcr.paths[i] = str
	}
	return nil
}

func (fcr FileConfigReader) GetPriority() int {
	return fcr.priority
}

func (fcr *FileConfigReader) ReadConfig() (map[string]interface{}, error) {
	for _, basePath := range fcr.paths {
		filePath := filepath.Join(basePath, fmt.Sprintf("%s.%s", fcr.name, fcr.fileType))

		if exists, _ := afero.Exists(fcr.fs, filePath); exists {
			// Read and parse file based on fcr.fileType
			// Return parsed configuration
		}
	}

	if fcr.required {
		return nil, errors.NewErr(FILE_PATH_ERROR_CODE, nil,
			fmt.Sprintf("Required config file not found: %s.%s", fcr.name, fcr.fileType), "config")
	}

	return make(map[string]interface{}), nil
}
