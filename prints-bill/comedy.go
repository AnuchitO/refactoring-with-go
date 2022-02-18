package main

import "math"

type Comedy struct {
	Name string
}

func (c Comedy) playName() string {
	return c.Name
}

func (c Comedy) amountFor(audience int) float64 {
	amount := 30000.0
	if audience > 20 {
		amount += 10000 + 500*(float64(audience-20))
	}
	amount += 300 * float64(audience)
	return amount
}

func (c Comedy) volumeCreditsFor(audience int) float64 {
	credits := 0.0
	credits += math.Max(float64(audience-30), 0)
	// if "comedy" == play.Kind { // no need
	credits += math.Floor(float64(audience / 5))
	// }
	return credits
}
