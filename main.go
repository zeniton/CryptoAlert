package main

import (
	"fmt"
)

func main() {
	alert := make(chan string)

	bitcoin := Coin{Symbol: "XBTZAR"}
	ether := Coin{Symbol: "ETHZAR"}
	coins := []Coin{bitcoin, ether}
	for _, coin := range coins {
		go coin.Monitor(alert)
	}

	for {
		fmt.Println(<-alert)
	}
}
