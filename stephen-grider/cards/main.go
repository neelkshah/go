package main

import "fmt"

func newCard() string {
	return "Five of spades"
}

func main() {
	card := newCard()
	fmt.Println(card)
}
