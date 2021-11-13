package repository

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"

	"github.com/Ras96/gcg/internal/model"
	"github.com/pkg/errors"
)

const MAXCAP = 1000

type analyzerRepository struct{}

func NewAnalyzerRepository() AnalyzerRepository {
	return &analyzerRepository{}
}

func (r *analyzerRepository) AnalyzeFile(filename string) (*model.File, error) {
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

func (r *analyzerRepository) parsePkgName(name *ast.Ident) string {
	return name.Name
}

func (r *analyzerRepository) parseImportSpecs(is []*ast.ImportSpec) []model.Import {
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

func (r *analyzerRepository) parseObjectsToStructs(obj map[string]*ast.Object) []model.Struct {
	sa := convertSortedArr(obj)
	structs := make([]model.Struct, len(sa))

	for i, v := range sa {
		flds := r.parseFields(v.fields)
		structs[i] = *model.NewStruct(v.name, flds, isPrivate(v.name))
	}

	return structs
}

func (r *analyzerRepository) parseFields(f []*ast.Field) []model.Field {
	fields := make([]model.Field, len(f))

	for i, fld := range f {
		var org string

		if len(fld.Names) == 0 {
			org = fld.Type.(*ast.Ident).Name
		} else {
			org = fld.Names[0].Name
		}

		fname := model.NewName(org, toArgName(org))

		var ftype model.Type
		if starExpr, ok := fld.Type.(*ast.StarExpr); ok {
			ftype = *r.parseExpr(starExpr.X, "", true)
		} else {
			ftype = *r.parseExpr(fld.Type, "", false)
		}

		fields[i] = *model.NewField(*fname, ftype)
	}

	return fields
}

func (r *analyzerRepository) parseExpr(f ast.Expr, prefix model.Prefix, isStar bool) *model.Type {
	switch t := f.(type) {
	case *ast.Ident:
		return model.NewType(isStar, prefix, "", t.Name)
	case *ast.SelectorExpr:
		return model.NewType(isStar, prefix, t.X.(*ast.Ident).Name, t.Sel.Name)
	case *ast.InterfaceType:
		return model.NewType(isStar, prefix, "", "interface{}")
	case *ast.ArrayType:
		return r.parseExpr(t.Elt, prefix.Add("[]"), isStar)
	case *ast.StarExpr:
		return r.parseExpr(t.X, prefix.Add("*"), isStar)
	default:
		return model.NewType(isStar, prefix, "", "interface{}")
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
	var l string
	if !isPrivate(s) {
		l = strings.ToLower(s[:1]) + s[1:]
	}

	if _, ok := goIdentifiers[l]; ok {
		return l + "_"
	}

	return l
}

func isPrivate(s string) bool {
	return s[:1] == strings.ToLower(s[:1])
}

type structInfo struct {
	pos    token.Pos
	name   string
	fields []*ast.Field
}

func convertSortedArr(obj map[string]*ast.Object) []structInfo {
	arr := make([]structInfo, 0, MAXCAP)

	for name, v := range obj {
		ts, ok := v.Decl.(*ast.TypeSpec)
		if !ok {
			continue
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}

		arr = append(arr, structInfo{v.Pos(), name, st.Fields.List})
	}

	sort.Slice(arr, func(i int, j int) bool {
		return arr[i].pos < arr[j].pos
	})

	return arr
}
