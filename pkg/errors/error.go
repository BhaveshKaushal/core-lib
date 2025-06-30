package errors

import (
	"github.com/pkg/errors"
)

type Error interface {
	error
	Code() Code
	Cause() error
	Message() string
	Er() error
	Wrap(string) error
}
type Err struct {
	code    Code
	message string
	er      error
	app     string
}

//TODO: Need to integrate logger
func NewErr(code Code, err error, msg, app string) *Err {
	return &Err{code: code, message: msg, er: err, app: app}
}

func NewErrDefault(code Code, msg, app string) *Err {
	return NewErr(code,errors.New(msg), msg, app)
}


func (err *Err) Code() Code {
	return err.code
}

func (err *Err) Message() string {
	return err.message
}

func (err *Err) Er() error {
	return err.er
}

func (err *Err) Cause() error {
	return errors.Cause(err.er)
}

func (er *Err) Wrap(msg string) error {
	return errors.Wrap(er.er, msg)
}

func (er *Err) Error() string {
	if er.er == nil {
		return ""
	}
	return er.er.Error()
}