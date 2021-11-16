package handler

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/Ras96/gcg/internal/util/errors"
)

func (h *Handlers) ExecuteGen(in string, out string) error {
	file, err := h.Repo.Analyzer.AnalyzeFile(in)
	if err != nil {
		return errors.Wrap(err, "Could not analyze file")
	}

	res, err := h.Repo.Generator.GenerateConstructors(file, out)
	if err != nil {
		return errors.Wrap(err, "Could not generate constructors")
	}

	if len(out) == 0 {
		fmt.Fprintln(os.Stdout, string(res))
	} else {
		if err := ioutil.WriteFile(out, res, fs.ModePerm); err != nil {
			return errors.Wrap(err, "Could not write to file")
		}
	}

	return nil
}
