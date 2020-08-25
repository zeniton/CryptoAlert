package main

import (
	"fmt"
)

func main() {
	alert := make(chan string)

	bitcoin := Coin{Symbol: "XBTZAR"}
	coins := []Coin{bitcoin}
	for _, coin := range coins {
		go coin.Monitor(alert)
	}

	for {
		fmt.Println(<-alert)
	}
}
