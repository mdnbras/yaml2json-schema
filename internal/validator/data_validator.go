package validator

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

// ValidateDataAgainstSchema valida um YAML de entrada
// contra um JSON Schema Draft-07.
func ValidateDataAgainstSchema(
	schemaJSON []byte,
	inputYAML []byte,
) error {

	// 1️⃣ Carrega o schema
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

	// 2️⃣ YAML → interface{}
	var data interface{}
	if err := yaml.Unmarshal(inputYAML, &data); err != nil {
		return fmt.Errorf("erro ao ler input YAML: %w", err)
	}

	// 3️⃣ Normaliza YAML → JSON types
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erro ao converter YAML para JSON: %w", err)
	}

	var jsonData interface{}
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		return fmt.Errorf("erro ao normalizar dados: %w", err)
	}

	// 4️⃣ Validação real
	if err := schema.Validate(jsonData); err != nil {
		return fmt.Errorf("input inválido segundo o schema: %w", err)
	}

	return nil
}
