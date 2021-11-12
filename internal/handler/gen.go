package handler

import (
	"fmt"

	"github.com/Ras96/gcg/internal/util/errors"
	"github.com/spf13/cobra"
)

func (h *handlers) Gen(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		errors.Exit(cmd, errors.New("Please provide an argument"))
	}

	file, err := h.repo.Parser.ParseFile(args[0])
	if err != nil {
		errors.Exit(cmd, errors.Wrap(err, "Could not parse file"))
	}

	fmt.Printf("%+v\n", file)
}
