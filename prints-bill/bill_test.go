package main

import (
	"testing"
)

func TestPrintBill(t *testing.T) {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		},
	}
	plays := map[string]map[string]string{
		"hamlet":  {"name": "Hamlet", "type": "tragedy"},
		"as-like": {"name": "As You Like It", "type": "comedy"},
		"othello": {"name": "Othello", "type": "tragedy"},
	}

	bill := statement(inv, plays)

	exp := `Statement for Bigco
  Hamlet: $650.00 (55 seats)
  As You Like It: $580.00 (35 seats)
  Othello: $500.00 (40 seats)
Amount owed is $1730.00
you earned 47 credits
`

	if exp != bill {
		t.Errorf("expect %q but got %q", exp, bill)
	}
}

// func TestPrintBillHtml(t *testing.T) {
// 	inv := Invoice{
// 		Customer: "Bigco",
// 		Performances: []Performance{
// 			{PlayID: "hamlet", Audience: 55},
// 			{PlayID: "as-like", Audience: 35},
// 			{PlayID: "othello", Audience: 40},
// 		},
// 	}
//
// 	plays := map[string]map[string]string{
// 		"hamlet":  {"name": "Hamlet", "type": "tragedy"},
// 		"as-like": {"name": "As You Like It", "type": "comedy"},
// 		"othello": {"name": "Othello", "type": "tragedy"},
// 	}
//
// 	bill := statementHtml(inv, plays)
//
// 	exp := `<h1>Statement for Bigco</h1>
// <table>
// <tr><th>Play</th><th>Seats</th><th>Cost</th></tr>
// <tr><td>Hamlet</td><td>55</td><td>$650.00</td></tr>
// <tr><td>As You Like It</td><td>35</td><td>$580.00</td></tr>
// <tr><td>Othello</td><td>40</td><td>$500.00</td></tr>
// </table>
// <p>Amount owed is <em>$1730.00</em></p>
// <p>you earned <em>47</em> credits</p>
// `
// 	if exp != bill {
// 		t.Errorf("expect %q but got %q", exp, bill)
// 	}
// }
