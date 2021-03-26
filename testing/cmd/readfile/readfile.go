package main

import (
	"anuchito.com/testing/cmd/readfile/bill"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	path := filepath.Join("bill","big.json")
	fmt.Println("path:", path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("invoices:", string(b))
	bill.Show()
}
