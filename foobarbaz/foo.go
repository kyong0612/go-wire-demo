package foobarbaz

type Foo struct {
	X int
}

func ProviderFoo() Foo {
	return Foo{X: 42}
}
