package repository

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/Ras96/gcg/internal/model"
	"github.com/pkg/errors"
)

const MAXCAP = 1000

type parserRepository struct{}

func NewParserRepository() ParserRepository {
	return &parserRepository{}
}

func (r *parserRepository) ParseFile(filename string) (*model.File, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse file")
	}

	packageName := f.Name.Name
	imports := r.parseImportSpecs(f.Imports)

	structs := make([]model.Struct, 0, MAXCAP)

	for name, obj := range f.Scope.Objects {
		ts, ok := obj.Decl.(*ast.TypeSpec)
		if !ok {
			continue
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}

		flds := r.parseFields(st.Fields.List)

		structs = append(structs, *model.NewStruct(name, flds))
	}

	return model.NewFile(packageName, imports, structs), nil
}

func (r *parserRepository) parseImportSpecs(is []*ast.ImportSpec) []model.Import {
	imports := make([]model.Import, len(is))

	for i, imp := range is {
		imports[i] = model.Import{
			Name: imp.Name.Name,
			Path: imp.Path.Value,
		}
	}

	return imports
}

func (r *parserRepository) parseFields(f []*ast.Field) []model.Field {
	fields := make([]model.Field, len(f))

	for i, fld := range f {
		org := fld.Names[0].Name
		fname := model.NewName(org, lower(org)) // TODO: lowerを柔軟にする

		ftype := model.Type{}
		if starExpr, ok := fld.Type.(*ast.StarExpr); ok {
			ftype = *r.parseExpr(starExpr.X, "", true)
		} else {
			ftype = *r.parseExpr(fld.Type, "", false)
		}

		fields[i] = *model.NewField(fname, &ftype)
	}

	return fields
}

func (r *parserRepository) parseExpr(f ast.Expr, prefix model.Prefix, isStar bool) *model.Type {
	switch t := f.(type) {
	case *ast.Ident:
		return model.NewType(isStar, prefix, "", t.Name)
	case *ast.SelectorExpr:
		return model.NewType(isStar, prefix, t.X.(*ast.Ident).Name, t.Sel.Name)
	case *ast.ArrayType:
		return r.parseExpr(t.Elt, prefix.Add("[]"), isStar)
	case *ast.StarExpr:
		return r.parseExpr(t.X, prefix.Add("*"), isStar)
	default:
		return nil
	}
}

func lower(s string) string {
	return strings.ToLower(s)
}
