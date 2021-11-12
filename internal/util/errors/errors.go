package errors

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var levelError = color.New(color.FgRed).Sprint("[Error]")

func Exit(cmd *cobra.Command, err error) {
	fmt.Println(levelError, err.Error())
	fmt.Println()

	_ = cmd.Help()

	os.Exit(1)
}

func New(message string) error {
	return errors.New(message)
}

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}
