package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	coinOption string
	name       string
}

func (p *Player) GetRandCoinOption() string {
	randIndex := rand.Intn(2)
	p.coinOption = CoinOptions[randIndex]
	return p.coinOption
}

func (p *Player) SetCoinOption(opponentOption string) {
	if opponentOption == "Heads" {
		p.coinOption = "Tails"
	} else {
		p.coinOption = "Heads"
	}
}

func (p Player) DidPlayerWin(winningFlip string) {
	if winningFlip == p.coinOption {
		fmt.Printf("%s won with a flip of %s!\n", p.name, p.coinOption)
	} else {
		fmt.Printf("%s lost with a flip of %s!\n", p.name, p.coinOption)
	}
}
