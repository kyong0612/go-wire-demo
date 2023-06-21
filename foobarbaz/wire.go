//go:build wireinject

package foobarbaz

import (
	"context"

	"github.com/google/wire"
)

func initializeBaz(context.Context) (Baz, error) {
	wire.Build(MegaSet)
	return Baz{}, nil
}
