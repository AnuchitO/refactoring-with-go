package main

import (
	"fmt"
)

type Publisher struct {
	subs []Subscriber
}

func (pb *Publisher) Subscribe(s Subscriber) {
	pb.subs = append(pb.subs, s)
}

func (pb *Publisher) Unsubscribe(s Subscriber) {
	// ...
}

func (pb *Publisher) Notifies(message interface{}) {
	for _, sub := range pb.subs {
		sub.Notify(message)
	}
}

type Subscriber interface {
	Notify(message interface{})
}

type Journalist struct {
	Publisher
	Name string
}

func NewJournalist(name string) Journalist {
	return Journalist{
		Name: name,
	}
}

func (jl Journalist) ProofReading() {
	jl.Notifies(jl.Name + "proof reading")
}

type NewspaperService struct{}

func (NewspaperService) Notify(message interface{}) {
	fmt.Printf("NewspaperService: has been called : %#v\n", message)
}

type MagazinesService struct{}

func (MagazinesService) Notify(message interface{}) {
	fmt.Printf("MagazinesService: has been called : %#v\n", message)
}

func main() {
	np := NewspaperService{}
	mg := MagazinesService{}

	p := NewJournalist("Nong")
	p.Subscribe(np)
	p.Subscribe(mg)

	p.ProofReading()
}
