package model

import "fmt"

type Message string

type Greeter struct {
	message Message
}

func (g Greeter) Greet() Message {
	return g.message
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{message: m}
}

type Event struct {
	greeter Greeter
}

func NewEvent(g Greeter) Event {
	return Event{greeter: g}
}

func (e Event) Start() {
	fmt.Println(e.greeter.Greet())
}
