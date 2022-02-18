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

func amountFor(play Play, audience int) float64 {
	amount := 0.0
	switch play.Kind {
	case "tragedy":
		amount = 40000
		if audience > 30 {
			amount += 1000 * (float64(audience - 30))
		}
	case "comedy":
		amount = 30000
		if audience > 20 {
			amount += 10000 + 500*(float64(audience-20))
		}
		amount += 300 * float64(audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", play.Kind))
	}

	return amount
}

func volumeCreditsFor(play Play, audience int) float64 {
	credits := 0.0
	credits += math.Max(float64(audience-30), 0)
	if "comedy" == play.Kind {
		credits += math.Floor(float64(audience / 5))
	}
	return credits
}

func totalAmount(plays Plays, performances []Performance) float64 {
	amounts := 0.0
	for _, perf := range performances {
		play := playFor(plays, perf)
		amounts += amountFor(play, perf.Audience)
	}
	return amounts
}

func totalVolumeCredits(plays Plays, performances []Performance) float64 {
	credits := 0.0
	for _, perf := range performances {
		play := playFor(plays, perf)
		credits += volumeCreditsFor(play, perf.Audience)
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
		amount := amountFor(play, perf.Audience)
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
