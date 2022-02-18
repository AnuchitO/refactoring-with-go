package main

import "math"

type Tragedy struct {
	Name string
}

func (play Tragedy) playName() string {
	return play.Name
}

func (play Tragedy) amountFor(audience int) float64 {
	amount := 0.0
	amount = 40000
	if audience > 30 {
		amount += 1000 * (float64(audience - 30))
	}

	return amount
}

func (play Tragedy) volumeCreditsFor(audience int) float64 {
	credits := 0.0
	credits += math.Max(float64(audience-30), 0)
	return credits
}
