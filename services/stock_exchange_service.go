package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Israel-Ferreira/stock-challenge-cli/config"
	"github.com/Israel-Ferreira/stock-challenge-cli/models"
)

func AsyncFindTicker(ticker string, tickerChan chan models.StockExchange, errorChan chan error) {

	var stockExchange models.StockExchange

	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s",
		ticker,
		config.ApiKey,
	)

	r, err := http.Get(url)

	if err != nil {
		errorChan <- err
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		err := errors.New("error: ação não encontrada")
		errorChan <- err
	}

	if err := json.NewDecoder(r.Body).Decode(&stockExchange); err != nil {
		errorChan <- err
	}

	tickerChan <- stockExchange
}
