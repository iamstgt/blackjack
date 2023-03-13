package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/brandondoesdev/slice-swap"
)

// Create a new type of deck which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"H", "D", "S", "C"}
	cardValues := [13]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+value)
		}
	}
	return cards
}

// d   : The actual copy of the deck we're working with is available in the functi as a variable called "d". d -> "cards" in this case.
// deck: Every variable of type "deck" can call this function on itself.
func (d *deck) print() {
	for _, card := range *d {
		fmt.Println(card)
	}
}

func (d *deck) printLenCap() {
	fmt.Println("len=%d cap=%d %v\n", len(*d), cap(*d), *d)
}

// This is the insufficient random function because the go random number generator always uses the exact same seed.
// To fix this, generating new seed value and using it as the seed to generate any random numbers are necessary.
// func (d deck) shuffle() {*
// 	for index := range d {
// 		newIndex := rand.Intn(len(d)-1)
// 		swap.SwapS(d, index, newIndex)
// 	}
// }

func (d *deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // For generating random seed.
	r := rand.New(source)

	for index := range *d {
		newIndex := r.Intn(len(*d) - 1)
		swap.SwapS(*d, index, newIndex)
	}
}

// // The method below with/woithout pointer didn't work...
// func (d deck) deal (deck, deck, deck) {
// 	return d[:2], d[2:4], d[4:]
// }
func (d deck) deal() (deck, deck, deck) {
	var playerHand, dealerHand, remainingCards deck
	for i := 0; i < 2; i++ {
		playerHand = append(playerHand, d[0])
		d = d[1:]
		dealerHand = append(dealerHand, d[0])
		d = d[1:]
	}
	remainingCards = d
	return playerHand, dealerHand, remainingCards
}

func (d *deck) hitOrStand(remainingCards *deck) {
	for {
		fmt.Println("\n--1: hit, 2: stand --")
		var userInput int
		fmt.Scan(&userInput)
		switch userInput {
		case 1:
			// Append a card to hands
			*d = append(*d, (*remainingCards)[0])
			// Remove that card from remainingCards
			*remainingCards = (*remainingCards)[1:]
		case 2:
			return
		default:
			fmt.Printf("Expected input is 1 or 2, but got %v", userInput)
			continue
		}
		fmt.Println(*d)
	}
}

func conclude(playerHand *deck, dealerHand *deck) {
	var sumOfPlayerHand int
	var sumOfDealerHand int
	for _, card := range *playerHand {
		value, _ := strconv.Atoi(card[1:])
		sumOfPlayerHand += value
	}
	for _, card := range *dealerHand {
		value, _ := strconv.Atoi(card[1:])
		sumOfDealerHand += value
	}
	if sumOfPlayerHand > 21 && sumOfDealerHand <= 21 {
		fmt.Println("Player: BURST")
		fmt.Println("Player: LOSE")
	} else if sumOfPlayerHand <= 21 && sumOfDealerHand > 21 {
		fmt.Println("Dealer: BURST")
		fmt.Println("Player: WIN")
	} else if sumOfPlayerHand > 21 && sumOfDealerHand > 21 {
		fmt.Println("Both: BURST")
		fmt.Println("DRAW")
	} else if sumOfPlayerHand == sumOfDealerHand {
		fmt.Println("DRAW")
	} else if 21-sumOfPlayerHand < 21-sumOfDealerHand {
		fmt.Println("Player WIN")
	} else if 21-sumOfPlayerHand > 21-sumOfDealerHand {
		fmt.Println("Player LOSE")
	}
}
