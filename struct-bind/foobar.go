package structbind

import "github.com/google/wire"

type Foo int
type Bar int

func ProvideFoo() Foo { return Foo(10) }
func ProvideBar() Bar { return Bar(20) }

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"),
)
