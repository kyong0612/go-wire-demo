package foobar

import "github.com/google/wire"

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func providerMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

func providerBar(f Fooer) Bar {
	return Bar(f.Foo())
}

var InterfaceSet = wire.NewSet(
	providerMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	providerBar,
)
