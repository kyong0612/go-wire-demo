package model

import (
	"fmt"
	"time"
)

type Message string

type Greeter struct {
	message Message
	grumpy  bool
}

func (g Greeter) Greet() Message {
	if g.grumpy {
		return Message("Go away!")
	}
	return g.message
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{
		message: m,
		grumpy:  grumpy,
	}
}

type Event struct {
	greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.grumpy {
		return Event{}, fmt.Errorf("could not create event: event greeter is grumpy")
	}
	return Event{greeter: g}, nil
}

func (e Event) Start() {
	fmt.Println(e.greeter.Greet())
}
