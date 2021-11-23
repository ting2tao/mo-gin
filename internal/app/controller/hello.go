package controller

import "fmt"

type Message struct {
	msg string
}

// Greeter 招待员
type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

// NewMessage 构造函数
func NewMessage(msg string) Message {
	return Message{msg: msg}
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Message
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg.msg)
}

func NormalStart() {
	m := NewMessage("hello wire")
	g := NewGreeter(m)
	event := NewEvent(g)
	event.Start()
}

func WireStart() {
	e := InitializeEvent("Hello MyWire")
	e.Start()
}
