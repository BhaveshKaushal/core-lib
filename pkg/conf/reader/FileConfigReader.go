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
}

func NewFileConfigReader(paths []string, required bool, name, fileType string) *FileConfigReader {
	fcr := &FileConfigReader{
		paths:    paths,
		required: required,
		name:     name,
		fileType: fileType,
	}

	fcr.validateAndAbsolutePaths()

	return fcr
}

func (fcr *FileConfigReader) validateAndAbsolutePaths() *errors.Err {
	for i, path := range fcr.paths {
		str, err := filepath.Abs(path)
		if err != nil {
			return errors.NewErr(FILE_PATH_ERROR_CODE, err, fmt.Sprintf("File Config Path Error: %s",path), "config")
		}
		fcr.paths[i] = str
		
	}
	return nil
}
