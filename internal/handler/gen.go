package handler

import (
	"fmt"
	"os"

	"github.com/Ras96/gcg/internal/util/errors"
	"github.com/spf13/cobra"
)

func (h *handlers) Gen(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cobra.CheckErr(errors.New("Please provide an argument"))
	}

	file, err := h.repo.Parser.ParseFile(args[0])
	errors.CheckErr(err, "Could not parse file")

	res, err := h.repo.Generator.GenerateConstructors(file)
	errors.CheckErr(err, "Could not generate constructors")

	fmt.Fprintln(os.Stdout, res)
}
