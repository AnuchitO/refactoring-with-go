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
		}}
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
