//go:generate go run github.com/Ras96/gcg@main gen $GOFILE -o gcg_gen.go

package example

type Hoge struct {
	A int
	B Fuga
	C Foo
}

type Fuga struct {
	C *int
	D []string
	E *[]*[]*[]*[]string
	G interface{}
}

type Foo string
