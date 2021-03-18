package main

import (
	"fmt"
	"math"
)

/**
// plays.json
{
   "hamlet":{"name":"Hamlet","type":"tragedy"},
   "as-like":{"name":"As You Like It","type":"comedy"},
   "othello":{"name":"Othello","type":"tragedy"}
}

// invoices.json
{
   "customer":"BigCo",
   "performances":[
      {
         "playID":"hamlet",
         "audience":55
      },
      {
         "playID":"as-like",
         "audience":35
      },
      {
         "playID":"othello",
         "audience":40
      }
   ]
}
*/

type Play map[string]map[string]string

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func statement(invoice Invoice, plays Play) string {
	totalAmount := 0.0
	volumeCredits := 0.0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	// formate = new Intl.NumberFormat("en-US", {style: "currency", currency: "USD", minimumFractionDigits: 2}).format

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := 0.0

		switch play["type"] {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (float64(perf.Audience - 30))
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(float64(perf.Audience-20))
			}
			thisAmount += 300 * float64(perf.Audience)
		default:
			panic(fmt.Sprintf("unknow type: %s", play["type"]))
		}

		// add volume credits
		volumeCredits += math.Max(float64(perf.Audience-30), 0)
		// add extra credit for every ten comedy attendees
		if "comedy" == play["type"] {
			volumeCredits += math.Floor(float64(perf.Audience / 5))
		}

		// print line for this order
		result += fmt.Sprintf("  %s: %f (%d seats)\n", play["name"], thisAmount/100, perf.Audience ) // TODO: format currency $580.00
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amoucnt owed is %f \n", totalAmount/100) // TODO: format currency $1,730.00
	result += fmt.Sprintf("you earned %d credits\n", int(volumeCredits))
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
	plays := map[string]map[string]string{
		"hamlet":  {"name": "Hamlet", "type": "tragedy"},
		"as-like": {"name": "As You Like It", "type": "comedy"},
		"othello": {"name": "Othello", "type": "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
