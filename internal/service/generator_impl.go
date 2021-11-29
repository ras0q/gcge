package service

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
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

type generatorService struct {
	Tmpl *template.Template
	Opts *imports.Options
}

func NewGeneratorService() GeneratorService {
	return &generatorService{
		Tmpl: nil,
		Opts: &imports.Options{
			AllErrors: true,
			Comments:  true,
		},
	}
}

var (
	tmplPath  = "internal/template" // relative to project root
	tmplFiles = []string{
		tmplPath + "/main.tmpl",
		tmplPath + "/constructor.tmpl",
		tmplPath + "/util.tmpl",
	}
)

func (s *generatorService) GenerateConstructors(file *model.File, output string, isPrivate bool) ([]byte, error) {
	s.Tmpl = template.New("main.tmpl").Funcs(fmap(isPrivate))
	if _, err := s.Tmpl.ParseFiles(tmplFiles...); err != nil {
		return nil, errors.Wrap(err, "Could not parse templates")
	}

	w := &bytes.Buffer{}
	if err := s.writeConstructors(w, file); err != nil {
		return nil, errors.Wrap(err, "Could not write constructors")
	}

	out, err := s.format(w, output)
	if err != nil {
		return nil, errors.Wrap(err, "Could not format output")
	}

	return out, nil
}

func (s *generatorService) writeConstructors(w *bytes.Buffer, file *model.File) error {
	b := bufio.NewWriter(w)
	if err := s.Tmpl.Execute(b, file); err != nil {
		return errors.Wrap(err, "Could not execute template")
	}

	if err := b.Flush(); err != nil {
		return errors.Wrap(err, "Could not flush buffer")
	}

	return nil
}

func (s *generatorService) format(w *bytes.Buffer, filename string) ([]byte, error) {
	formatted, err := imports.Process(filename, w.Bytes(), s.Opts)
	if err != nil {
		if len(filename) == 0 {
			fmt.Fprintln(os.Stdout, w.String())
		} else {
			if err := ioutil.WriteFile(filename, w.Bytes(), fs.ModePerm); err != nil {
				return nil, errors.Wrap(err, "Could not write to file")
			}
		}

		fmt.Fprintln(os.Stderr, "Error occurred. Instead, gcg output the unformatted file")
		fmt.Fprintln(os.Stderr, "")

		return nil, errors.Wrap(err, "Could not format file")
	}

	return formatted, nil
}

func fmap(isPrivate bool) template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
		"funcName": func(funcName string) string {
			if isPrivate {
				return strings.ToLower(funcName[:1]) + funcName[1:]
			} else {
				return strings.Title(funcName)
			}
		},
	}
}
