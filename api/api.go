package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance api request format
type CoinBalanceParam struct {
	Username string
}

// Coin balance api response
type CoinBalanceResponse struct {
	HTTPResponseCode int
	CoinBalance      int64
}

//Error response
type Error struct {
	HTTPResponseCode int
	ErrorMessage 	 string
}

