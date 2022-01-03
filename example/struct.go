//go:generate go run github.com/Ras96/gcg@latest gen $GOFILE -o gcg_gen.go

package example_test

type Hoge struct {
	A int
	B Fuga
	C Foo
}

type Fuga struct {
	A *int
	B []string
	C *[]*[]*[]*[]string
	D interface{}
	E map[*Hoge]map[*Fuga]*Hoge
}

type Foo string
