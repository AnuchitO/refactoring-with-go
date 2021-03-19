package main

import (
	"fmt"
	"math"
)

type Play interface {
	amountFor(audience int) float64
	volumeCreditsFor(audience int) float64
	Name() string
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

type rate struct {
	play     Play
	amount   float64
	audience int
	credit   float64
}

type bill struct {
	customer           string
	rates              []rate
	totalAmount        float64
	totalVolumeCredits float64
}

func statement(invoice Invoice, plays Plays) string {
	var rates []rate
	for _, perf := range invoice.Performances {
		play := playFor(plays, perf)
		audience := perf.Audience
		rates = append(rates, rate{
			play:     play,
			audience: audience,
			amount:   play.amountFor(audience),
			credit:   play.volumeCreditsFor(audience),
		})
	}
	bill := bill{
		customer:           invoice.Customer,
		rates:              rates,
		totalAmount:        totalAmounts(rates),
		totalVolumeCredits: totalVolumeCredits(rates),
	}

	return renderPlainText(bill)
}

type comedy struct {
	name string
	kind string
}

func totalVolumeCredits(rates []rate) float64 {
	result := 0.0
	for _, r := range rates {
		result += r.credit
	}
	return result
}

func totalAmounts(rates []rate) float64 {
	result := 0.0
	for _, r := range rates {
		result += r.amount
	}
	return result
}

func renderPlainText(bill bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.customer)
	for _, r := range bill.rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", r.play.Name(), r.amount/100, r.audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.totalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.totalVolumeCredits)
	return result
}

func (play comedy) volumeCreditsFor(audience int) float64 {
	credits := math.Max(float64(audience-30), 0)
	credits += math.Floor(float64(audience / 5))
	return credits
}

func playFor(plays Plays, perf Performance) Play {
	play := plays[perf.PlayID]
	return play
}

type tragedy struct {
	name string
	kind string
}

func (play tragedy) amountFor(audience int) float64 {
	amount := 40000.0
	if audience > 30 {
		amount += 1000 * (float64(audience - 30))
	}
	return amount
}

func (play tragedy) volumeCreditsFor(audience int) float64 {
	credits := math.Max(float64(audience-30), 0)
	return credits
}

func (play tragedy) Name() string {
	return play.name
}

func (play comedy) Name() string {
	return play.name
}

func (play comedy) amountFor(audience int) float64 {
	amount := 30000.0
	if audience > 20 {
		amount += 10000 + 500*(float64(audience-20))
	}
	amount += 300 * float64(audience)
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

	plays := Plays{
		"hamlet":  tragedy{name: "Hamlet", kind: "tragedy"},
		"as-like": comedy{name: "As You Like It", kind: "comedy"},
		"othello": tragedy{name: "Othello", kind: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
