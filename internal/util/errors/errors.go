package errors

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func CheckErr(err error, msg string) {
	if err != nil {
		cobra.CheckErr(errors.Wrap(err, msg))
	}
}

func New(message string) error {
	return errors.New(message)
}

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}
