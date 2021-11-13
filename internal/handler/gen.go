package handler

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/Ras96/gcg/internal/repository"
	"github.com/Ras96/gcg/internal/util/errors"
	"github.com/spf13/cobra"
)

type genHandler struct {
	repo *repository.Repositories
	opts *genOpts
}

type genOpts struct {
	output *string
}

func NewGenOpts(output *string) genOpts {
	return genOpts{
		output: output,
	}
}

func (h *genHandler) Run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cobra.CheckErr(errors.New("Please provide an argument"))
	}

	file, err := h.repo.Parser.ParseFile(args[0])
	errors.CheckErr(err, "Could not parse file")

	res, err := h.repo.Generator.GenerateConstructors(file, *h.opts.output)
	errors.CheckErr(err, "Could not generate constructors")

	if len(*h.opts.output) == 0 {
		fmt.Fprintln(os.Stdout, string(res))
	} else {
		ioutil.WriteFile(*h.opts.output, res, fs.ModePerm)
	}
}
