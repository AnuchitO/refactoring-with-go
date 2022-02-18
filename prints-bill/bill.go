package main

import (
	"fmt"
	"math"
)

type Play struct {
	Name string
	Kind string
}

type Plays map[string]Player

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Player interface {
	playName() string
	amountFor(audience int) float64
	volumeCreditsFor(audience int) float64
}

func playKind(play Play) string {
	return play.Kind
}

func (play Play) playName() string {
	return play.Name
}

func playFor(plays Plays, perf Performance) Player {
	return plays[perf.PlayID]
}

func (play Play) amountFor(audience int) float64 {
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

func (play Play) volumeCreditsFor(audience int) float64 {
	credits := 0.0
	credits += math.Max(float64(audience-30), 0)
	if "comedy" == play.Kind {
		credits += math.Floor(float64(audience / 5))
	}
	return credits
}

func totalAmount(rates []Rate) float64 {
	result := 0.0
	for _, r := range rates {
		result += r.Amount
	}
	return result
}

func totalVolumeCredits(rates []Rate) float64 {
	result := 0.0
	for _, r := range rates {
		result += r.Credits
	}
	return result
}

type Rate struct {
	Play     Player
	Amount   float64
	Audience int
	Credits  float64
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
		audience := perf.Audience
		rates = append(rates, Rate{
			Play:     play,
			Audience: audience,
			Amount:   play.amountFor(audience),
			Credits:  play.volumeCreditsFor(audience),
		})
	}

	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              rates,
		TotalAmount:        totalAmount(rates),
		TotalVolumeCredits: totalVolumeCredits(rates),
	}

	return renderPlainText(bill)
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, rate := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", rate.Play.playName(), rate.Amount/100, rate.Audience)
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
		"hamlet":  Play{Name: "Hamlet", Kind: "tragedy"},
		"as-like": Play{Name: "As You Like It", Kind: "comedy"},
		"othello": Play{Name: "Othello", Kind: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
