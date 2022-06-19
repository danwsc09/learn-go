package main

import (
	"math/rand"
	"time"
)

var CoinOptions []string = []string{"Heads", "Tails"}

type Coin struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c Coin) FlipCoin() string {
	randInt := rand.Intn(2)
	return CoinOptions[randInt]
}
