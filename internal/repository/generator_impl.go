package repository

import (
	"bufio"
	"bytes"
	"strings"
	"text/template"

	"github.com/Ras96/gcg/internal/model"
	"github.com/Ras96/gcg/internal/util/errors"
	"golang.org/x/tools/imports"
)

type generatorRepository struct {
	Tmpl   *template.Template
	Opts   *imports.Options
	Output model.Filename
}

func NewGeneratorRepository(opts *imports.Options) GeneratorRepository {
	return &generatorRepository{
		Opts: opts,
	}
}

var fmap = template.FuncMap{
	"title": strings.Title,
}

func (r *generatorRepository) GenerateConstructors(file *model.File) (string, error) {
	r.Tmpl = template.New("constructor.tmpl").Funcs(fmap)
	if _, err := r.Tmpl.ParseFiles("./template/constructor.tmpl"); err != nil {
		return "", errors.Wrap(err, "Could not parse template files")
	}

	w := &bytes.Buffer{}
	if err := r.writeConstructors(w, file); err != nil {
		return "", errors.Wrap(err, "Could not write constructors")
	}

	out, err := r.format(w)
	if err != nil {
		return "", errors.Wrap(err, "Could not format output")
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

func (r *generatorRepository) format(w *bytes.Buffer) (string, error) {
	formatted, err := imports.Process(string(r.Output), w.Bytes(), r.Opts)
	if err != nil {
		return "", errors.Wrap(err, "Could not format output")
	}

	return string(formatted), nil
}
