package consultant

import "fmt"

type Consultant struct {
	Name         string
	Role         string
	AnnualIncome int
}

type Role int

const (
	Trainer Role = iota
	Coach
)

// New version 1
func New(name string, role string, annualIncome int) *Consultant {
	return &Consultant{Name: name, Role: role, AnnualIncome: annualIncome}
}

// New version 2
func NewRole(role Role) *Consultant {
	switch role {
	case Trainer:
		return &Consultant{"", "Trainer", 60000}
	case Coach:
		return &Consultant{"", "Coach", 80000}
	default:
		return &Consultant{}
	}
}

// New version 3
func NewTrainer() *Consultant {
	return &Consultant{"", "Trainer", 60000}
}

func NewCoach() *Consultant {
	return &Consultant{"", "Coach", 80000}
}

/**
package main
func main() {
  m := consultant.New(Coach)
  m.Name = "Gamer"
  fmt.Println(m)
}
*/
