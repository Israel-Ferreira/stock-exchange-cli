package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CryptoCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "crypto [ticker]",
		Aliases: []string{"cr", "cry"},
		Args:    cobra.ExactArgs(1),
		Example: "crypto BTC_BRL",
		Short:   "Comando de consulta de cotação de criptoativos",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	return cmd
}
