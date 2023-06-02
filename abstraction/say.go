package main

import "fmt"

func main() {
	var word string
	fmt.Println("What do you want to say?")
	fmt.Scan(&word)
	fmt.Printf("Say, %s!\n", word)
}
