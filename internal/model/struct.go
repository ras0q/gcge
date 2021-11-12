package model

type File struct {
	Package string
	Imports []Import
	Structs []Struct
}

type Import struct {
	Name string
	Path string
}

type Struct struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name Name
	Type Type
}

type Name struct {
	// フィールド名
	Original string
	// 変数として使う名称(e.g. ID->id, Name->name, LongName->longName)
	Argument string
}

type Type struct {
	IsStar  bool
	Prefix  Prefix
	Package string
	Name    string
}

type Prefix string

func (p Prefix) Add(following string) Prefix {
	return p + Prefix(following)
}

func NewFile(pkg string, imports []Import, structs []Struct) *File {
	return &File{
		Package: pkg,
		Imports: imports,
		Structs: structs,
	}
}

func NewStruct(name string, fields []Field) *Struct {
	return &Struct{
		Name:   name,
		Fields: fields,
	}
}

func NewField(name *Name, typ *Type) *Field {
	return &Field{
		Name: *name,
		Type: *typ,
	}
}

func NewName(original string, argument string) *Name {
	return &Name{
		Original: original,
		Argument: argument,
	}
}

func NewType(isStar bool, prefix Prefix, pkg string, name string) *Type {
	return &Type{
		IsStar:  isStar,
		Prefix:  prefix,
		Package: pkg,
		Name:    name,
	}
}
