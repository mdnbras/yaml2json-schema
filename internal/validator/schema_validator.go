package validator

import (
	"bytes"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

// ValidateDraft07Schema valida se o JSON informado
// é um JSON Schema válido segundo o meta-schema oficial Draft-07.
//
// IMPORTANTE:
// - Essa validação NÃO valida dados
// - Ela valida o PRÓPRIO schema como documento
// - Qualquer erro aqui significa que o schema gerado é inválido
func ValidateDraft07Schema(schemaJSON []byte) error {
	compiler := jsonschema.NewCompiler()

	// Por padrão o compiler já conhece o meta-schema Draft-07
	// pois ele é referenciado pelo campo "$schema"
	if err := compiler.AddResource(
		"schema.json",
		bytes.NewReader(schemaJSON),
	); err != nil {
		return fmt.Errorf(
			"erro ao carregar JSON Schema para validação: %w",
			err,
		)
	}

	// A validação REAL acontece aqui.
	// Se o schema não obedecer ao Draft-07,
	// o erro será lançado neste ponto.
	if _, err := compiler.Compile("schema.json"); err != nil {
		return fmt.Errorf(
			"JSON Schema inválido segundo Draft-07: %w",
			err,
		)
	}

	return nil
}
