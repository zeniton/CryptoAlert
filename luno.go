package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Luno is tThe model for the JSON reponse from Luno API
type Luno struct {
	Pair      string `json:"pair"`
	Timestamp uint64 `json:"timestamp"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastTrade string `json:"last_trade"`
	Volume    string `json:"rolling_24_hour_volume"`
	Status    string `json:"status"`
}

// Tick is the model for the current pair state
type Tick struct {
	Pair      string
	Timestamp uint64
	Bid       float64
	Ask       float64
	LastTrade float64
	Volume    float64
	IsActive  bool
}

// GetTick retrieves the current price for the specified pair
func GetTick(pair string) (Tick, error) {
	tick := Tick{
		Pair: pair,
	}

	// Call Luno API & parse result
	url := "https://api.luno.com/api/1/ticker?pair=" + pair
	resp, err := http.Get(url)
	if err != nil {
		return tick, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tick, err
	}
	var luno Luno
	err = json.Unmarshal(body, &luno)
	if err != nil {
		return tick, err
	}

	// Map API response model to Tick
	tick.IsActive = (luno.Status == "ACTIVE")
	tick.Timestamp = luno.Timestamp
	tick.Bid, _ = strconv.ParseFloat(luno.Bid, 64)
	tick.Ask, _ = strconv.ParseFloat(luno.Ask, 64)
	tick.LastTrade, _ = strconv.ParseFloat(luno.LastTrade, 64)
	tick.Volume, _ = strconv.ParseFloat(luno.Volume, 64)

	return tick, nil
}
