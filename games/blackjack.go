package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var SUITS []string = []string{"C", "S", "H", "D"}
var NUMS []string = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

type strSlice []string

func (arr *strSlice) pop() string {
	if len(*arr) == 0 {
		panic("pop from an empty array")
	}
	n := len(*arr)
	returnVal := (*arr)[n-1]
	*arr = (*arr)[:n-1]
	return returnVal
}

func main() {
	fmt.Println("Game start!")
	var deck strSlice = generateShuffledDeck()
	fmt.Println(deck)

	var hand strSlice = []string{}
	var userInput string

	for i := 0; i < 2; i++ {
		hand = append(hand, deck.pop())
	}
	// TODO: game loop
	/*
		if min-handvalue < 21,
			ask user input for Hit / Stay
			if hit: deal card to user
			if stay: exit user turn
		if min-handvalue == 21
			exit loop and user wins
		if min-handvalue > 21
			exit loop and user loses
	*/
	for calculateHand(hand) < 21 {
		fmt.Println("Your turn: hit (h) or stay (s)?")
		fmt.Scanln(&userInput)
		if userInput[0] == 'h' {
			hand = append(hand, deck.pop())

			continue
		}
		break

	}
}

func generateShuffledDeck() []string {
	var deck = []string{}

	for i := 0; i < len(SUITS); i++ {
		for j := 0; j < len(NUMS); j++ {
			deck = append(deck, SUITS[i]+NUMS[j])
		}
	}

	deck = shuffleDeck(deck)

	return deck
}

func shuffleDeck(deck []string) []string {
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

func getNumberValue(card string) []int {
	if val := card[:len(card)-1]; val == "A" {
		return []int{1, 11}
	} else if val == "10" || val == "J" || val == "Q" || val == "K" {
		return []int{10}
	} else {
		i, _ := strconv.Atoi(val)
		return []int{i}
	}
}

/*
1. Generate a deck and shuffle
2. Deal 2 cards to player
3. Hit / Stay depending on player choice
4. Evaluate
5. Repeat 3
*/
