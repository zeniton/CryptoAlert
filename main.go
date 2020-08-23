package main

import (
	"fmt"
)

func main() {
	btc, err := GetCoin("XBTZAR")
	if err != nil {
		fmt.Println(err)
	}
	if btc.IsActive {
		fmt.Printf("%+v\n", btc)
	}
}
