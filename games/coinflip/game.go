package main

import "math/rand"

type CoinGame struct {
	players []Player
	coin    Coin
}

func (c *CoinGame) StartGame(player1Name, player2Name string) {
	c.players = append(c.players, Player{name: player1Name}, Player{name: player2Name})
	randIndex := rand.Intn(2)
	var otherIndex int
	if randIndex == 0 {
		otherIndex = 1
	} else {
		otherIndex = 0
	}

	playerPick := c.players[randIndex].GetRandCoinOption()
	c.players[otherIndex].SetCoinOption(playerPick)

	c.coin = Coin{}
	winningFlip := c.coin.FlipCoin()

	for _, player := range c.players {
		player.DidPlayerWin(winningFlip)
	}
}
