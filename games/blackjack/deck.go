package main

import (
	"math/rand"
	"time"
)

type StrSlice []string

func (arr *StrSlice) pop() string {
	if len(*arr) == 0 {
		panic("pop from an empty array")
	}
	n := len(*arr)
	returnVal := (*arr)[n-1]
	*arr = (*arr)[:n-1]
	return returnVal
}

var SUITS []string = []string{"C", "S", "H", "D"}
var NUMS []string = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func GenerateShuffledDeck() []string {
	var deck []string

	for i := 0; i < len(SUITS); i++ {
		for j := 0; j < len(NUMS); j++ {
			deck = append(deck, SUITS[i]+NUMS[j])
		}
	}

	deck = ShuffleDeck(deck)

	return deck
}

func ShuffleDeck(deck []string) []string {
	var swap string
	var randInt int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(deck); i++ {
		randInt = rand.Intn(len(deck))
		swap = deck[i]
		deck[i] = deck[randInt]
		deck[randInt] = swap
	}

	return deck
}
