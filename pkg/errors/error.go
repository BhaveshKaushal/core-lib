package errors

import (
	"github.com/pkg/errors"
)

type Error interface {
	error
	Code() int
	Cause() error
	Message() string
	Er() error
	Wrap(string) error
}
type Err struct {
	code    int
	message string
	er      error
}

//TODO: Need to integrate logger
func NewErr(code int, msg string, err error) *Err {
	return &Err{code: code, message: msg, er: err}
}

func NewErrDefault(code int, msg string) *Err {
	return NewErr(code,msg,errors.New(msg))
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

func (er *Err) Wrap(msg string) error {
	return errors.Wrap(er.er, msg)
}
