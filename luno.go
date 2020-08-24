package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

// Coin is the model for the current pair state
type Coin struct {
	Symbol    string
	Timestamp uint64
	Bid       float64
	Ask       float64
	LastTrade float64
	Volume    float64
	IsActive  bool
}

// Monitor detects and alerts significant changes in a coin
func (coin Coin) Monitor(alert chan<- string) {
	clock := time.Tick(2 * time.Second)
	for {
		<-clock
		err := coin.getTick()
		if err == nil {
			// Dummy alert for testing
			alert <- fmt.Sprintf("%s: %v", coin.Symbol, coin.Bid)
		}
	}
}

// GetCoin retrieves the current price for the specified pair
func (coin *Coin) getTick() error {
	// Call Luno API & parse result
	url := "https://api.luno.com/api/1/ticker?pair=" + coin.Symbol
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var luno Luno
	err = json.Unmarshal(body, &luno)
	if err != nil {
		return err
	}

	// Map API response model to Coin
	coin.IsActive = (luno.Status == "ACTIVE")
	coin.Timestamp = luno.Timestamp
	coin.Bid, _ = strconv.ParseFloat(luno.Bid, 64)
	coin.Ask, _ = strconv.ParseFloat(luno.Ask, 64)
	coin.LastTrade, _ = strconv.ParseFloat(luno.LastTrade, 64)
	coin.Volume, _ = strconv.ParseFloat(luno.Volume, 64)

	return nil
}
