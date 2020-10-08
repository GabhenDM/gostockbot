package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabhendm/gostockbot/models"
	"github.com/prometheus/common/log"
)

func GetStockPrice(stock string) (models.AVQuoteResponse, error) {
	var quoteRequest models.AVQuoteRequest
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
