package main

import (
	"fmt"
	"os"

	"github.com/kyong0612/go-wire-demo/model"
)

func main() {
	fmt.Println("###### Not Use Wire #####")
	message := model.NewMessage()
	greeter := model.NewGreeter(message)
	e, err := model.NewEvent(greeter)
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()

	fmt.Println("##### Use Wire #####")
	e, err = InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()
}
