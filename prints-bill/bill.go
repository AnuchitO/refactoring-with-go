package main

import (
	"fmt"
	"math"
)

type Play struct {
	Name string
	Type string
}

type Plays map[string]Play

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func statement(invoice Invoice, plays Plays) string {
	totalAmount := 0.0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	for _, perf := range invoice.Performances {
		// print line for this order
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", playName(playFor(plays, perf)), amountFor(perf, playFor(plays, perf))/100, perf.Audience)
	}
	for _, perf := range invoice.Performances {
		totalAmount += amountFor(perf, playFor(plays, perf))
	}
	result += fmt.Sprintf("Amoucnt owed is $%.2f\n", totalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", totalVolumeCredits(invoice, plays))
	return result
}

func totalVolumeCredits(invoice Invoice, plays Plays) float64 {
	credits := 0.0
	for _, perf := range invoice.Performances {
		credits += volumeCreditsFor(perf, plays)
	}
	return credits
}

func volumeCreditsFor(perf Performance, plays Plays) float64 {
	credits := 0.0
	credits += math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == playType(playFor(plays, perf)) {
		credits += math.Floor(float64(perf.Audience / 5))
	}
	return credits
}
func playName(play Play) string {
	return play.Name
}

func playType(play Play) string {
	return play.Type
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func amountFor(perf Performance, play Play) float64 {
	amount := 0.0
	kind := playType(play)
	switch kind {
	case "tragedy":
		amount = 40000
		if perf.Audience > 30 {
			amount += 1000 * (float64(perf.Audience - 30))
		}
	case "comedy":
		amount = 30000
		if perf.Audience > 20 {
			amount += 10000 + 500*(float64(perf.Audience-20))
		}
		amount += 300 * float64(perf.Audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", kind))
	}
	return amount
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}

	_plays := Plays{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}

	bill := statement(inv, _plays)
	fmt.Println(bill)
}
