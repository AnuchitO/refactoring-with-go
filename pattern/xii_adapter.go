package main

import "fmt"

type Publisher struct {
	subs []Subscriber
}

func (pb *Publisher) Notifies(message interface{}) {
	for _, sub := range pb.subs {
		sub.Notify(message)
	}
}

type Subscriber interface {
	Notify(message interface{})
}

// this is a function can have method
func sentMessage(s string) {
	fmt.Printf("let me sent some msg : %#v\n", s)
}

type Notify func(s string)

func (fn Notify) Notify(message interface{}) {
	msg, _ := message.(string)

	fn(msg)
}

func main() {
	pb := Publisher{
		subs: []Subscriber{Notify(sentMessage)},
	}

	pb.Notifies("hi, there!")
}
