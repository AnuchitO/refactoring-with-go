package main

import (
	"anuchito.com/testing/cmd/readfile/bill"
	"fmt"

	_ "embed"
)

//go:embed bill/big.json
var big string

func main() {
	fmt.Println("big", big)
	bill.Show()
}