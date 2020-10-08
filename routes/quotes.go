package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gabhendm/gostockbot/models"
	"github.com/prometheus/common/log"
)

func SymbolSearch(stock string) (models.AVSearchResponse, error) {
	quoteRequest := models.AVRequest{}
	quoteResponse := models.AVSearchResponse{}
	fmt.Println(stock)
	if stock != "" {
		endpoint := quoteRequest.FormatRequestSearch("SYMBOL_SEARCH", stock)
		fmt.Println(endpoint)
		resp, err := http.Get(endpoint)
		if err != nil {
			return quoteResponse, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err.Error())
			return quoteResponse, err
		}
		if err := json.Unmarshal(body, &quoteResponse); err != nil {
			log.Error(err.Error())
			return quoteResponse, err
		}

		return quoteResponse, nil
	} else {
		log.Error("Empty Stock!")
		return quoteResponse, errors.New("Empty Stock")
	}
}

func GetStockPrice(stock string) (models.AVQuoteResponse, error) {
	var quoteRequest models.AVRequest
	quoteResponse := models.AVQuoteResponse{}
	if stock != "" {
		endpoint := quoteRequest.FormatRequest("GLOBAL_QUOTE", stock)
		resp, err := http.Get(endpoint)
		if err != nil {
			return quoteResponse, err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err.Error())
			return quoteResponse, err
		}
		if err := json.Unmarshal(body, &quoteResponse); err != nil {
			log.Error(err.Error())
			return quoteResponse, err
		}

	}
	return quoteResponse, nil
}

func GetStockOverview(stock string) (models.AVOverviewResponse, error) {
	var overviewRequest models.AVRequest
	var overviewResponse models.AVOverviewResponse
	if stock != "" {
		endpoint := overviewRequest.FormatRequest("OVERVIEW", stock)
		fmt.Println(endpoint)
		resp, err := http.Get(endpoint)
		if err != nil {
			return overviewResponse, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err.Error())
			return overviewResponse, err
		}
		if err := json.Unmarshal(body, &overviewResponse); err != nil {
			log.Error(err.Error())
			return overviewResponse, err
		}

	}
	return overviewResponse, nil
}
