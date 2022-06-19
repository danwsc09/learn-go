package main

import "fmt"

func main() {
	fmt.Println("Starting...")
	theGame := CoinGame{}
	theGame.StartGame("Dan", "Phil")
}
