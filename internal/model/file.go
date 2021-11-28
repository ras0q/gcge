//go:generate go run github.com/Ras96/gcg@latest gen $GOFILE -o file_cst.go

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
	Name      string
	Fields    []Field
	IsPrivate bool
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
