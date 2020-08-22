package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Luno struct {
	Pair      string `json:"pair"`
	Timestamp uint64 `json:"timestamp"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastTrade string `json:"last_trade"`
	Volume    string `json:"rolling_24_hour_volume"`
	Status    string `json:"status"`
}

type Tick struct {
	Pair      string
	Timestamp uint64
	Bid       float64
	Ask       float64
	LastTrade float64
	Volume    float64
	Status    string
}

func getTick(url string) (Tick, error) {
	var tick Tick

	resp, err := http.Get(url)
	if err != nil {
		return tick, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var luno Luno
	err = json.Unmarshal(body, &luno)
	if err != nil {
		return tick, err
	}

	tick.Pair = luno.Pair
	tick.Status = luno.Status
	tick.Timestamp = luno.Timestamp
	tick.Bid, _ = strconv.ParseFloat(luno.Bid, 64)
	tick.Ask, _ = strconv.ParseFloat(luno.Ask, 64)
	tick.LastTrade, _ = strconv.ParseFloat(luno.LastTrade, 64)
	tick.Volume, _ = strconv.ParseFloat(luno.Volume, 64)

	return tick, nil
}

func main() {
	tick, err := getTick("https://api.mybitx.com/api/1/ticker?pair=XBTZAR")
	if err != nil {
		fmt.Println(err)
	}
	if tick.Status == "ACTIVE" {
		fmt.Printf("%+v\n", tick)
	}
}
