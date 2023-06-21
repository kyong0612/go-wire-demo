//go:build wireinject
// +build wireinject

package structbind

import "github.com/google/wire"

func initializeFooBar() (FooBar, error) {
	panic(wire.Build(Set))
}
