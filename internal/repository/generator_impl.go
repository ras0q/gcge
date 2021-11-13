package repository

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Ras96/gcg/internal/model"
	"github.com/Ras96/gcg/internal/util/errors"
	"golang.org/x/tools/imports"
)

type generatorRepository struct {
	Tmpl *template.Template
	Opts *imports.Options
}

func NewGeneratorRepository() GeneratorRepository {
	return &generatorRepository{
		Tmpl: nil,
		Opts: &imports.Options{
			AllErrors: true,
			Comments:  true,
		},
	}
}

var fmap = template.FuncMap{
	"title": strings.Title,
}

func (r *generatorRepository) GenerateConstructors(file *model.File, filename string) ([]byte, error) {
	r.Tmpl = template.New("constructor").Funcs(fmap)
	if _, err := r.Tmpl.Parse(string(model.GenTmpl)); err != nil {
		return nil, errors.Wrap(err, "Could not parse templates")
	}

	w := &bytes.Buffer{}
	if err := r.writeConstructors(w, file); err != nil {
		return nil, errors.Wrap(err, "Could not write constructors")
	}

	out, err := r.format(w, filename)
	if err != nil {
		return nil, errors.Wrap(err, "Could not format output")
	}

	return out, nil
}

func (r *generatorRepository) writeConstructors(w *bytes.Buffer, file *model.File) error {
	b := bufio.NewWriter(w)
	if err := r.Tmpl.Execute(b, file); err != nil {
		return errors.Wrap(err, "Could not execute template")
	}

	if err := b.Flush(); err != nil {
		return errors.Wrap(err, "Could not flush buffer")
	}

	return nil
}

func (r *generatorRepository) format(w *bytes.Buffer, filename string) ([]byte, error) {
	formatted, err := imports.Process(filename, w.Bytes(), r.Opts)
	if err != nil {
		if len(filename) == 0 {
			fmt.Fprintln(os.Stdout, w.String())
		} else {
			_ = ioutil.WriteFile(filename, w.Bytes(), fs.ModePerm)
		}

		fmt.Fprintln(os.Stderr, "Error occurred. Instead, gcg output the unformatted file")
		fmt.Fprintln(os.Stderr, "")

		return nil, errors.Wrap(err, "Could not format file")
	}

	return formatted, nil
}
