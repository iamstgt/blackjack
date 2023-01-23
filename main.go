package main

import (
	"fmt"
)

func main() {
	// Call the function called "newDeck" returning cards whose type is "deck".
	cards := newDeck()

	// Shuffle the cards.
	cards.shuffle()

	// // Call the receiver function called "print".
	// cards.print()

	// Call the function to deal cards for player, dealer and remains.
	playerHand, dealerHand, remainingCards := deal(cards, 2)
	fmt.Println("--Player Hands--")
	playerHand.print()
	fmt.Println("--Dealer Hands--")
	dealerHand.print()
	fmt.Println("--Remaining Cards--")
	remainingCards.print()
}