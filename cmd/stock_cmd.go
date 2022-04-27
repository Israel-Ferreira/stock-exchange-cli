package cmd

import (
	"fmt"
	"os"

	"github.com/Israel-Ferreira/stock-challenge-cli/models"
	"github.com/Israel-Ferreira/stock-challenge-cli/services"
	"github.com/spf13/cobra"
)

func StockCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "stock [ticker]",
		Aliases: []string{"stk", "stx"},
		Short:   "Consulta ações listadas na B3",
		Long:    "Consulta ações listadas na B3 ou em outras bolsas do mundo (BDR)",
		Example: "stock PETR4",
		Run: func(cmd *cobra.Command, args []string) {
			tickerB3 := fmt.Sprintf("%s.SA", args[0])

			errChan := make(chan error)
			respChan := make(chan models.StockExchange)

			go services.AsyncFindTicker(tickerB3, respChan, errChan)

			select {
			case msgTicker := <-respChan:
				fmt.Printf("High: %s \n", msgTicker.GlobalQuote.High)
				fmt.Printf("Low: %s \n", msgTicker.GlobalQuote.Low)
				fmt.Printf("Close price: %s \n", msgTicker.GlobalQuote.Price)
			case errResp := <-errChan:
				fmt.Println(errResp)
				os.Exit(1)
			}
		},
	}
}
