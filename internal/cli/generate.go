package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/mdnbras/yaml2json-schema/internal/config"
	"github.com/mdnbras/yaml2json-schema/internal/generator"
	"github.com/mdnbras/yaml2json-schema/internal/merger"
	"github.com/mdnbras/yaml2json-schema/internal/parser"
	"github.com/mdnbras/yaml2json-schema/internal/validator"
	"github.com/spf13/cobra"
)

var generateOpts = &config.GenerateOptions{}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gera um JSON Schema a partir de YAML + CSV",
	Long: `Gera um JSON Schema Draft-07 combinando:
- Estrutura definida em YAML
- Metadados definidos em CSV (descrição, obrigatório, tipo)

O schema gerado é validado contra o meta-schema oficial Draft-07.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGenerate(cmd)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(
		&generateOpts.YamlPath,
		"yaml",
		"y",
		"",
		"Caminho do arquivo YAML de entrada",
	)

	generateCmd.Flags().StringVarP(
		&generateOpts.CsvPath,
		"csv",
		"c",
		"",
		"Caminho do arquivo CSV com metadados",
	)

	generateCmd.Flags().StringVarP(
		&generateOpts.OutputPath,
		"output",
		"o",
		"schema.json",
		"Caminho do arquivo JSON Schema de saída",
	)

	generateCmd.Flags().StringVar(
		&generateOpts.SchemaVersion,
		"schema-version",
		"https://json-schema.org/draft-07/schema#",
		"Versão do JSON Schema",
	)

	generateCmd.Flags().Bool(
		"yaml-diff",
		false,
		"Exibe o diff entre o YAML original e o YAML normalizado (anchors resolvidas)",
	)

	generateCmd.Flags().String(
		"dump-normalized-yaml-file",
		"",
		"Salva o YAML normalizado (anchors resolvidas) no arquivo informado",
	)

	generateCmd.Flags().String(
		"remove-anchor-bases",
		"",
		"Lista de chaves YAML (separadas por vírgula) a serem removidas após resolver anchors",
	)

	_ = generateCmd.MarkFlagRequired("yaml")
	_ = generateCmd.MarkFlagRequired("csv")
}

func runGenerate(cmd *cobra.Command) error {
	verbose, _ := cmd.Flags().GetBool("verbose")
	generateOpts.Verbose = verbose
	if verbose {
		fmt.Println("▶ Iniciando geração do JSON Schema (Draft-07)")
	}

	if verbose {
		fmt.Println("▶ Lendo YAML:", generateOpts.YamlPath)
	}

	basesArg, _ := cmd.Flags().GetString("remove-anchor-bases")
	var bases []string
	if basesArg != "" {
		bases = strings.Split(basesArg, ",")
	}

	showDiff, _ := cmd.Flags().GetBool("yaml-diff")
	dumpFile, _ := cmd.Flags().GetString("dump-normalized-yaml-file")

	rootField, normalizedYAML, err := parser.ParseYAML(
		generateOpts.YamlPath,
		showDiff,
		bases,
	)
	if err != nil {
		return fmt.Errorf("erro ao processar YAML: %w", err)
	}

	if dumpFile != "" {
		if err := os.WriteFile(dumpFile, normalizedYAML, 0644); err != nil {
			return fmt.Errorf(
				"erro ao salvar YAML normalizado em '%s': %w",
				dumpFile,
				err,
			)
		}
	}

	if verbose {
		fmt.Println("▶ Lendo CSV:", generateOpts.CsvPath)
	}

	metadata, err := parser.ParseCSV(generateOpts.CsvPath)
	if err != nil {
		return fmt.Errorf("erro ao processar CSV: %w", err)
	}

	if verbose {
		fmt.Println("▶ Aplicando metadados (merge)")
	}

	if err := merger.Merge(rootField, metadata); err != nil {
		return fmt.Errorf("erro ao aplicar metadados: %w", err)
	}

	if verbose {
		fmt.Println("▶ Gerando JSON Schema")
	}

	schemaBytes, err := generator.Generate(
		rootField,
		generateOpts.SchemaVersion,
	)
	if err != nil {
		return fmt.Errorf("erro ao gerar JSON Schema: %w", err)
	}

	if verbose {
		fmt.Println("▶ Validando JSON Schema contra Draft-07")
	}

	if err := validator.ValidateDraft07Schema(schemaBytes); err != nil {
		return err
	}

	if verbose {
		fmt.Println("▶ Salvando arquivo:", generateOpts.OutputPath)
	}

	if err := os.WriteFile(generateOpts.OutputPath, schemaBytes, 0644); err != nil {
		return fmt.Errorf("erro ao escrever arquivo de saída: %w", err)
	}

	if verbose {
		fmt.Println("✔ JSON Schema Draft-07 gerado e validado com sucesso")
	}

	return nil
}
