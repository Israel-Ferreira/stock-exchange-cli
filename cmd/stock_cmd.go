package cmd

import (
	"fmt"

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
			fmt.Println(args[0])
		},
	}
}
