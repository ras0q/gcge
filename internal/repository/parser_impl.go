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

	packageName := r.parsePkgName(f.Name)
	imports := r.parseImportSpecs(f.Imports)
	structs := r.parseObjectsToStructs(f.Scope.Objects)

	return model.NewFile(packageName, imports, structs), nil
}

func (r *parserRepository) parsePkgName(name *ast.Ident) string {
	return name.Name
}

func (r *parserRepository) parseImportSpecs(is []*ast.ImportSpec) []model.Import {
	imports := make([]model.Import, len(is))

	for i, imp := range is {
		if imp.Name != nil {
			imports[i].Name = imp.Name.Name
		}

		if imp.Path != nil {
			imports[i].Path = imp.Path.Value
		}
	}

	return imports
}

func (r *parserRepository) parseObjectsToStructs(obj map[string]*ast.Object) []model.Struct {
	structs := make([]model.Struct, 0, MAXCAP)

	for name, obj := range obj {
		ts, ok := obj.Decl.(*ast.TypeSpec)
		if !ok {
			continue
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}

		flds := r.parseFields(st.Fields.List)

		structs = append(structs, *model.NewStruct(name, flds, isPrivate(name)))
	}

	return structs
}

func (r *parserRepository) parseFields(f []*ast.Field) []model.Field {
	fields := make([]model.Field, len(f))

	for i, fld := range f {
		if len(fld.Names) == 0 {
			continue
		}

		org := fld.Names[0].Name
		fname := model.NewName(org, toArgName(org))

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

var goIdentifiers = map[string]struct{}{
	"break":       {},
	"case":        {},
	"chan":        {},
	"const":       {},
	"continue":    {},
	"default":     {},
	"defer":       {},
	"else":        {},
	"fallthrough": {},
	"for":         {},
	"func":        {},
	"go":          {},
	"goto":        {},
	"if":          {},
	"import":      {},
	"interface":   {},
	"map":         {},
	"package":     {},
	"range":       {},
	"return":      {},
	"select":      {},
	"struct":      {},
	"switch":      {},
	"type":        {},
	"var":         {},
}

func toArgName(s string) string {
	l := strings.ToLower(s)
	if _, ok := goIdentifiers[l]; ok {
		return l + "_"
	}

	return l
}

func isPrivate(s string) bool {
	return s[:1] == strings.ToLower(s[:1])
}
