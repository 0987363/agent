package models

import (
	"errors"
	"fmt"
)

func Error(v ...interface{}) error {
	return errors.New(fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) error {
	return errors.New(fmt.Sprintf(format, v...))
}
