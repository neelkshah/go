package main

import "fmt"

func newCard() string {
	return "Five of spades"
}

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)
	fmt.Println(hand)
	handPrime, _ := deal(cards, 5)
	fmt.Println(handPrime)
	cards.shuffle()
	handShuffled, _ := deal(cards, 5)
	fmt.Println(handShuffled)
}
