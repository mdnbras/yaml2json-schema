package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jsonschema",
	Short: "CLI para gerar JSON Schema a partir de YAML + CSV",
	Long:  "Ferramenta CLI escrita em Go para gerar JSON Schema baseado em estrutura YAML e metadados CSV.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Habilita logs detalhados")
}
