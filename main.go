package main

import (
	"fmt"

	"github.com/kyong0612/go-wire-demo/model"
)

func main() {
	fmt.Println("###### Not Use Wire #####")
	message := model.NewMessage()
	greeter := model.NewGreeter(message)
	event := model.NewEvent(greeter)

	event.Start()

	fmt.Println("##### Use Wire #####")
	e := InitializeEvent()

	e.Start()
}
