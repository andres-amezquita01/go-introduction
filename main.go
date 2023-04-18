package main

import "fmt"

func main() {
	card := getCard()
	fmt.Println(card)
}

func getCard() string {
	return "Five of Diamonds"
}
