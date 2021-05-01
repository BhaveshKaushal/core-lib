package errors

import "github.com/pkg/errors"

type Err struct {
	code    int
	message string
	er      error
}

//TODO: Need to integrate logger
func NewErr(code int, msg string) *Err {
	return &Err{code: code, message: msg, er: errors.New(msg)}
}

func (err *Err) Code() int {
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

func (err *Err) WithMessage(msg string) error {
	return errors.WithMessage(err.Wrap(""), msg)
}

func (err *Err) WithMessagef(format string, args ...interface{}) error {
	return errors.WithMessagef(err.Wrap(""), format, args...)
}

func (er *Err) Wrap(msg string) error {
	return errors.Wrap(er.er, msg)
}

func (er *Err) Wrapf(format string, args ...interface{}) error {
	return errors.Wrapf(er.er, format, args...)
}


