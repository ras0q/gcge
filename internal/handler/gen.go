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
	opts *GenOpts
}

type GenOpts struct {
	Output string
}

func NewGenHandler(repo *repository.Repositories, opts *GenOpts) *genHandler {
	return &genHandler{
		repo: repo,
		opts: opts,
	}
}

func (h *genHandler) Run(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Please provide an argument")
	}

	file, err := h.repo.Analyzer.AnalyzeFile(args[0])
	if err != nil {
		return errors.Wrap(err, "Could not analyze file")
	}

	res, err := h.repo.Generator.GenerateConstructors(file, h.opts.Output)
	if err != nil {
		return errors.Wrap(err, "Could not generate constructors")
	}

	if len(h.opts.Output) == 0 {
		fmt.Fprintln(os.Stdout, string(res))
	} else {
		_ = ioutil.WriteFile(h.opts.Output, res, fs.ModePerm)
	}

	return nil
}
