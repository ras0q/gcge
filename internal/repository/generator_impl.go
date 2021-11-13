package repository

import (
	"bufio"
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/Ras96/gcg/internal/model"
	"github.com/Ras96/gcg/internal/util/errors"
	"golang.org/x/tools/imports"
)

//go:embed template/constructor.tmpl
var genTmpl []byte

type generatorRepository struct {
	Tmpl *template.Template
	Opts *imports.Options
}

func NewGeneratorRepository() GeneratorRepository {
	return &generatorRepository{}
}

var fmap = template.FuncMap{
	"title": strings.Title,
}

func (r *generatorRepository) GenerateConstructors(file *model.File, filename string) (string, error) {
	r.Tmpl = template.New("constructor").Funcs(fmap)
	if _, err := r.Tmpl.Parse(string(genTmpl)); err != nil {
		return "", errors.Wrap(err, "Could not parse templates")
	}

	w := &bytes.Buffer{}
	if err := r.writeConstructors(w, file); err != nil {
		return "", errors.Wrap(err, "Could not write constructors")
	}

	out, err := r.format(w, filename)
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

func (r *generatorRepository) format(w *bytes.Buffer, filename string) (string, error) {
	formatted, err := imports.Process(filename, w.Bytes(), r.Opts)
	if err != nil {
		return "", errors.Wrap(err, "Could not format output")
	}

	return string(formatted), nil
}
