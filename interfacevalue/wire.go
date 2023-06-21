//go:build wireinject
// +build wireinject

package interfacevalue

import (
	"io"
	"os"

	"github.com/google/wire"
)

func injectReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
s
