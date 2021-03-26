package structure

import "fmt"

type Stack interface {
	Push(element interface{})
	Pop() interface{}
}

type stack struct {
	elements []interface{}
}

func (s *stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *stack) Pop() interface{} {
	return s.elements[:len(s.elements)]
}

// NewStack ... expose an interface only
func NewStack() Stack {
	return &stack{}
}

// package main
func main() {
	s := NewStack() // structure.NewStack()
	s.Push("Nong")
	fmt.Println(s.Pop())
}
