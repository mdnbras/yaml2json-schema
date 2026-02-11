package validator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

func ValidateDataAgainstSchema(
	schemaJSON []byte,
	input []byte,
) error {

	compiler := jsonschema.NewCompiler()

	if err := compiler.AddResource(
		"schema.json",
		bytes.NewReader(schemaJSON),
	); err != nil {
		return fmt.Errorf("erro ao carregar schema: %w", err)
	}

	schema, err := compiler.Compile("schema.json")
	if err != nil {
		return fmt.Errorf("schema inválido: %w", err)
	}

	var data interface{}

	trimmed := strings.TrimSpace(string(input))

	if len(trimmed) > 0 && (trimmed[0] == '{' || trimmed[0] == '[') {
		if err := json.Unmarshal(input, &data); err != nil {
			return fmt.Errorf("erro ao ler JSON: %w", err)
		}
	} else {
		if err := yaml.Unmarshal(input, &data); err != nil {
			return fmt.Errorf("erro ao ler YAML: %w", err)
		}
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erro ao converter dados para JSON: %w", err)
	}

	var normalized interface{}
	if err := json.Unmarshal(jsonBytes, &normalized); err != nil {
		return fmt.Errorf("erro ao normalizar dados: %w", err)
	}

	if err := schema.Validate(normalized); err != nil {
		return fmt.Errorf("input inválido segundo o schema: %w", err)
	}

	return nil
}
