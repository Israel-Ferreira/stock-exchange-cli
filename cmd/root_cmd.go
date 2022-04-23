package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func RootCmd() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(CryptoCmd())
	rootCmd.AddCommand(StockCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Erro ao Executar o Comando")
	}
}
