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
	playerHand, dealerHand, remainingCards := cards.deal()
	fmt.Println("--Player Hands--")
	playerHand.print()
	fmt.Println("--Dealer Hands--")
	dealerHand.print()
	// fmt.Println("--Remaining Cards--")
	// remainingCards.print()

	// cards.printLenCap() // 6 8 [H1 H2 D2 D1 D3 D3]
	// playerHand.printLenCap() // 2 8 [H1 H2]
	// dealerHand.printLenCap() // 2 6 [H3 D1]
	// remainingCards.printLenCap() // 2 4 [D2 D3]

	fmt.Println("--Player Turn--")
	playerHand.hitOrStand(&remainingCards)

	fmt.Println("--Dealer Turn--")
	dealerHand.hitOrStand(&remainingCards)

	fmt.Println("--Open--")
	fmt.Println("--Player Hands--")
	playerHand.print()
	fmt.Println("--Dealer Turn--")
	dealerHand.print()

	fmt.Println("--Result--")
	conclude(&playerHand, &dealerHand)
}
