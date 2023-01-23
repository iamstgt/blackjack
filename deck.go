package main 

import (
	"fmt"
	"math/rand"
	"github.com/brandondoesdev/slice-swap"
	"time"
)

// Create a new type of deck which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"❤️", "♦️", "♠️", "♣️"}
	cardValues := [12]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+value)
		}
	}
	return cards
}

// d   : The actual copy of the deck we're working with is available in the functi as a variable called "d". d -> "cards" in this case.
// deck: Every variable of type "deck" can call this function on itself.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// This is not a receiver function. 
func deal(d deck, handSize int) (deck, deck, deck) {
	return d[:handSize], d[handSize:handSize+2], d[handSize+2:]
}

// This is the insufficient random function because the go random number generator always uses the exact same seed. 
// To fix this, generating new seed value and using it as the seed to generate any random numbers are necessary.
// func (d deck) shuffle() {
// 	for index := range d {
// 		newIndex := rand.Intn(len(d)-1)
// 		swap.SwapS(d, index, newIndex)
// 	}
// }

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // For generating random seed.
	r := rand.New(source)

	for index := range d {
		newIndex := r.Intn(len(d)-1)
		swap.SwapS(d, index, newIndex)
	}
}
