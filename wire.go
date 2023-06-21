//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kyong0612/go-wire-demo/model"
)

func InitializeEvent() model.Event {
	wire.Build(model.NewEvent, model.NewGreeter, model.NewMessage)
	return model.Event{}
}
