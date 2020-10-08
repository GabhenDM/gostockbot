package models

import (
	"fmt"

	"github.com/spf13/viper"
)

type AVQuoteRequest struct {
	url      string
	function string
	symbol   string
	apiKey   string
}

func (r AVQuoteRequest) FormatRequest(function string, symbol string) string {
	return fmt.Sprintf("%s?function=%s&symbol=%s&apikey=%s", viper.GetString("API_URL"), function, symbol+".SAO", viper.GetString("AV_API_KEY"))
}

type GlobalQuote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. changePercent"`
}

type AVQuoteResponse struct {
	Quote GlobalQuote
}
