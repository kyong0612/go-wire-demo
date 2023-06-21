package foobarbaz

import "github.com/google/wire"

var SuperSet = wire.NewSet(
	ProviderFoo,
	ProviderBar,
	ProviderBaz,
)

var MegaSet = wire.NewSet(SuperSet)
