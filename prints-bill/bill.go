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
type NPlays map[string]XYZ

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func statement(invoice Invoice, plays Plays) string {
	var pays []Rate
	for _, perf := range invoice.Performances {
		play := playFor(plays, perf)
		audience := perf.Audience
		pays = append(pays, Rate{
			Play:     play,
			Audience: audience,
			Amount:   play.amountFor(audience),
			Credit:   play.volumeCreditsFor(audience),
		})
	}
	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              pays,
		TotalAmount:        totalAmounts(pays),
		TotalVolumeCredits: totalVolumeCredits(pays),
	}

	return renderPlainText(bill)
}

func totalVolumeCredits(pays []Rate) float64 {
	result := 0.0
	for _, pay := range pays {
		result += pay.Credit
	}
	return result
}

func totalAmounts(pays []Rate) float64 {
	result := 0.0
	for _, pay := range pays {
		result += pay.Amount
	}
	return result
}

type Rate struct {
	Play     XYZ
	Amount   float64
	Audience int
	Credit   float64
}

type Bill struct {
	Customer           string
	Rates              []Rate
	TotalAmount        float64
	TotalVolumeCredits float64
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, pay := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", pay.Play.name(), pay.Amount/100, pay.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

func (play Play) volumeCreditsFor(audience int) float64 {
	credits := 0.0
	credits += math.Max(float64(audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == play.Type {
		credits += math.Floor(float64(audience / 5))
	}
	return credits
}

func playFor(plays Plays, perf Performance) XYZ {
	play := plays[perf.PlayID]
	switch play.Type {
	case "tragedy":
		return tragedy{Name: play.Name, Type: play.Type}
	}
	return play
}

type tragedy struct {
	Name string
	Type string
}

func (play tragedy) amountFor(audience int) float64 {
	amount := 0.0
	kind := play.Type
	switch kind {
	case "tragedy":
		amount = 40000
		if audience > 30 {
			amount += 1000 * (float64(audience - 30))
		}
	default:
		panic(fmt.Sprintf("not a type tragedy: %s", kind))
	}
	return amount
}

func (play tragedy) volumeCreditsFor(audience int) float64 {
	credits := 0.0
	credits += math.Max(float64(audience-30), 0)
	return credits
}

func (play tragedy) name() string {
	return play.Name
}

type XYZ interface {
	amountFor(audience int) float64
	volumeCreditsFor(audience int) float64
	name() string
}

func (play Play) name() string {
	return play.Name
}

func (play Play) amountFor(audience int) float64 {
	amount := 0.0
	kind := play.Type
	switch kind {
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

	plays := Plays{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
