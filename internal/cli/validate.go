package cli

import (
	"fmt"
	"os"

	"github.com/mdnbras/yaml2json-schema/internal/validator"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Valida um YAML de entrada contra um JSON Schema",
	Long: `Valida um arquivo YAML de entrada contra um JSON Schema Draft-07.
Falha caso o input não esteja em conformidade com o schema.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runValidate(cmd)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().String(
		"schema",
		"",
		"Caminho do arquivo schema.json",
	)

	validateCmd.Flags().String(
		"input",
		"",
		"Caminho do arquivo input.yaml",
	)

	_ = validateCmd.MarkFlagRequired("schema")
	_ = validateCmd.MarkFlagRequired("input")
}

func runValidate(cmd *cobra.Command) error {
	schemaPath, _ := cmd.Flags().GetString("schema")
	inputPath, _ := cmd.Flags().GetString("input")

	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("erro ao ler schema: %w", err)
	}

	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("erro ao ler input YAML: %w", err)
	}

	if err := validator.ValidateDataAgainstSchema(
		schemaBytes,
		inputBytes,
	); err != nil {
		return err
	}

	fmt.Println("✔ Input YAML/JSON é válido segundo o schema")
	return nil
}
