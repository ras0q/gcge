package handler

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type GenOpts struct {
	Output    string
	IsPrivate bool
}

func (h *Handlers) ExecuteGen(in string, opts GenOpts) error {
	file, err := h.Srv.Analyzer.AnalyzeFile(in)
	if err != nil {
		return errors.Wrap(err, "Could not analyze file")
	}

	res, err := h.Srv.Generator.GenerateConstructors(file, opts.Output, opts.IsPrivate)
	if err != nil {
		return errors.Wrap(err, "Could not generate constructors")
	}

	if len(opts.Output) == 0 {
		fmt.Fprintln(os.Stdout, string(res))
	} else {
		if err := ioutil.WriteFile(opts.Output, res, fs.ModePerm); err != nil {
			return errors.Wrap(err, "Could not write to file")
		}
	}

	return nil
}
