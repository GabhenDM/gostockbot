package routes

import (
	"encoding/json"
	"fmt"
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
		/* if (models.GlobalQuote{} == quoteResponse.Quote) {
			http.Error(w, "Stock Not Found", 404)
			return
		} */
	}
	fmt.Println(quoteResponse.Quote)
	// response is for now in JSON, will change as integration is developed
	return quoteResponse, nil
}
