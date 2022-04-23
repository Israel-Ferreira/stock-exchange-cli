package services

import (
	"errors"
	"fmt"
	"net/http"
)

func AsyncFindTicker(ticker string, tickerChan chan any, errorChan chan error) {

	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&outputsize=full&apikey=%s",
		ticker,
		"",
	)

	r, err := http.Get(url)

	if err != nil {
		errorChan <- err
	}

	if r.StatusCode != http.StatusOK {
		err := errors.New("")
		errorChan <- err
	}
}
