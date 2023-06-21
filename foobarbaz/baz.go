package foobarbaz

import (
	"context"
	"errors"
)

type Baz struct {
	X int
}

func ProviderBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("bar.X is zero")
	}
	return Baz(bar), nil
}
