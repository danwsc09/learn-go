package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Start() {
	fmt.Println("Game start!")
	var deck StrSlice = GenerateShuffledDeck()
	fmt.Println(deck)

	var hand StrSlice = []string{}
	var userInput string

	for i := 0; i < 2; i++ {
		hand = append(hand, deck.pop())
	}

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
	fmt.Println("Your hand:", hand)
	for {
		fmt.Println("Your turn: hit (h) or stay (s)?")
		fmt.Scanln(&userInput)
		if userInput[0] == 'h' {
			hand = append(hand, deck.pop())
		} else {
			fmt.Println("Your hand:", hand)
			handValues := CalculateHand(hand)
			fmt.Println("Value:", handValues[len(handValues)-1])
			break
		}

		fmt.Println("Your hand:", hand)

		handValues := CalculateHand(hand)

		fmt.Println("Possible values:", handValues)

		if len(handValues) == 0 {
			fmt.Println("you're bust!")
			break
		} else if WinningHand(handValues) {
			fmt.Println("You win!")
			break
		}
	}
}

func CalculateHand(hand []string) []int {
	values := make(map[int]bool)
	if val := GetCardValue(hand[0]); val == 1 {
		values[1] = true
		values[11] = true
	} else {
		values[val] = true
	}

	for i := 1; i < len(hand); i++ {
		val := GetCardValue(hand[i])
		newValues := make(map[int]bool)
		for key, _ := range values {
			if key+val <= 21 {

				newValues[key+val] = true
			}
			if val == 1 && key+11 <= 21 {
				newValues[key+11] = true
			}
		}
		values = newValues
	}

	var filtered []int
	for key, _ := range values {
		filtered = append(filtered, key)
	}
	sort.Ints(filtered)
	return filtered
}

func GetCardValue(card string) int {
	if val := card[1:]; val == "A" {
		return 1
	} else if val == "10" || val == "J" || val == "Q" || val == "K" {
		return 10
	} else {
		i, _ := strconv.Atoi(val)
		return i
	}
}

func WinningHand(handValues []int) bool {
	for _, value := range handValues {
		if value == 21 {
			return true
		}
	}
	return false
}
