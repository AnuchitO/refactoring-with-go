package main

import (
	"fmt"
	"math"
)

type Play struct {
	name string
	kind string
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

func totalAmount(rates []Rate) float64 {
	result := 0.0
	for _, r := range rates {
		result += r.amount
	}
	return result
}

func statement(invoice Invoice, plays Plays) string {
	var rates []Rate
	for _, perf := range invoice.Performances {
		audience := perf.Audience
		play := playFor(plays, perf)
		amount := play.amountFor(audience)
		credits := play.volumeCreditsFor(audience)
		rates = append(rates, Rate{play: play, amount: amount, audience: audience, credits: credits})
	}

	bill := Bill{
		customer:           invoice.Customer,
		rates:              rates,
		totalAmount:        totalAmount(rates),
		totalVolumeCredits: totalVolumeCredits(rates),
	}

	return renderPlainText(bill)
}

type Rate struct {
	play     Play
	amount   float64
	audience int
	credits  float64
}

type Bill struct {
	customer           string
	rates              []Rate
	totalAmount        float64
	totalVolumeCredits float64
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.customer)
	for _, r := range bill.rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", r.play.name, r.amount/100, r.audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.totalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.totalVolumeCredits)
	return result
}

func totalVolumeCredits(rates []Rate) (result float64) {
	for _, r := range rates {
		result += r.credits
	}
	return
}

type Player interface {
	volumeCreditsFor(audience int) float64
	amountFor(audience int) float64
}

// type comedy
// type tragedy

func (play Play) volumeCreditsFor(audience int) float64 {
	volumeCredits := 0.0
	// add volume credits
	volumeCredits += math.Max(float64(audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == play.kind {
		volumeCredits += math.Floor(float64(audience / 5))
	}
	return volumeCredits
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func (play Play) amountFor(audience int) float64 {
	result := 0.0
	switch play.kind {
	case "tragedy":
		result = 40000
		if audience > 30 {
			result += 1000 * (float64(audience - 30))
		}
	case "comedy":
		result = 30000
		if audience > 20 {
			result += 10000 + 500*(float64(audience-20))
		}
		result += 300 * float64(audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", play.kind))
	}

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
	_plays := map[string]Play{
		"hamlet":  {name: "Hamlet", kind: "tragedy"},
		"as-like": {name: "As You Like It", kind: "comedy"},
		"othello": {name: "Othello", kind: "tragedy"},
	}

	bill := statement(inv, _plays)
	fmt.Println(bill)
}
