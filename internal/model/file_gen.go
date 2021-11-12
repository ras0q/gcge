package model

func NewFile(pkg string, imports []Import, structs []Struct) *File {
	return &File{
		Package: pkg,
		Imports: imports,
		Structs: structs,
	}
}

func NewImport(name string, path string) *Import {
	return &Import{
		Name: name,
		Path: path,
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
