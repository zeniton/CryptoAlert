package main

import (
	"fmt"
)

func main() {
	tick, err := GetTick("XBTZAR")
	if err != nil {
		fmt.Println(err)
	}
	if tick.IsActive {
		fmt.Printf("%+v\n", tick)
	}
}
