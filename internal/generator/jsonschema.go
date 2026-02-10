package generator

import (
	"encoding/json"
	"fmt"

	"github.com/mdnbras/yaml2json-schema/internal/model"
)

// Generate gera o JSON Schema Draft-07 a partir do modelo Field.
// O schema gerado é estruturalmente compatível com o meta-schema oficial.
func Generate(
	root *model.Field,
	schemaVersion string,
) ([]byte, error) {

	schema := buildSchema(root)

	// Draft-07 exige $schema explícito no root
	schema.Schema = schemaVersion

	output, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar JSON Schema: %w", err)
	}

	return output, nil
}

// buildSchema converte recursivamente model.Field -> model.Schema
// seguindo estritamente as regras do JSON Schema Draft-07.
func buildSchema(field *model.Field) *model.Schema {
	schema := &model.Schema{
		Type:        field.Type,
		Description: field.Description,
	}

	switch field.Type {

	case "object":
		// propriedades fixas
		if len(field.Properties) > 0 {
			schema.Properties = make(map[string]*model.Schema)

			for name, child := range field.Properties {
				schema.Properties[name] = buildSchema(child)

				if child.Required {
					schema.Required = append(schema.Required, name)
				}
			}

			if len(schema.Required) == 0 {
				schema.Required = nil
			}
		}

		// MAP dinâmico (additionalProperties como schema)
		if field.AdditionalProperties != nil {
			schema.AdditionalProperties = buildSchema(field.AdditionalProperties)
		} else {
			// objeto fechado
			closed := false
			schema.AdditionalProperties = &closed
		}

	case "array":
		if field.Items != nil {
			schema.Items = buildSchema(field.Items)
		}

	default:
		// string, number, boolean, integer, null
	}

	return schema
}
