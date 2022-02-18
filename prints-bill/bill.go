package main

import (
	"fmt"
	"math"
)

type Play struct {
	Name string
	Kind string
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

func playKind(play Play) string {
	return play.Kind
}

func playName(play Play) string {
	return play.Name
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func amountFor(plays Plays, perf Performance) float64 {
	amount := 0.0

	switch playKind(playFor(plays, perf)) {
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
		panic(fmt.Sprintf("unknow type: %s", playKind(playFor(plays, perf))))
	}

	return amount
}

func volumeCreditsFor(plays Plays, perf Performance) float64 {
	credits := 0.0
	credits += math.Max(float64(perf.Audience-30), 0)
	kind := playKind(playFor(plays, perf))
	if "comedy" == kind {
		credits += math.Floor(float64(perf.Audience / 5))
	}
	return credits
}

func totalAmount(plays Plays, performances []Performance) float64 {
	amounts := 0.0
	for _, perf := range performances {
		amounts += amountFor(plays, perf)
	}
	return amounts
}

func totalVolumeCredits(plays Plays, performances []Performance) float64 {
	credits := 0.0
	for _, perf := range performances {
		credits += volumeCreditsFor(plays, perf)
	}
	return credits
}

type Rate struct {
	Play     Play
	Amount   float64
	Audience int
}

type Bill struct {
	Customer           string
	Rates              []Rate
	TotalAmount        float64
	TotalVolumeCredits float64
}

func statement(invoice Invoice, plays Plays) string {
	var rates []Rate
	for _, perf := range invoice.Performances {
		play := playFor(plays, perf)
		amount := amountFor(plays, perf)
		audience := perf.Audience
		rates = append(rates, Rate{Play: play, Amount: amount, Audience: audience})
	}

	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              rates,
		TotalAmount:        totalAmount(plays, invoice.Performances),
		TotalVolumeCredits: totalVolumeCredits(plays, invoice.Performances),
	}

	return renderPlainText(bill)
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, rate := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", rate.Play.Name, rate.Amount/100, rate.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := Plays{
		"hamlet":  {Name: "Hamlet", Kind: "tragedy"},
		"as-like": {Name: "As You Like It", Kind: "comedy"},
		"othello": {Name: "Othello", Kind: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
