//go:build wireinject
// +build wireinject

package interfacebind

import "github.com/google/wire"

func initializeBar() (Bar, error) {
	wire.Build(Set)
	return "", nil
}
