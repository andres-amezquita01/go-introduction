package main

import "fmt"

func main() {
	cards := []string{getCard(), "Ace of Diamonds", getCard()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func getCard() string {
	return "Five of Diamonds"
}
