package merger

import (
	"fmt"
	"strings"

	"github.com/mdnbras/yaml2json-schema/internal/model"
	"github.com/mdnbras/yaml2json-schema/internal/validator"
)

func Merge(root *model.Field, metadata map[string]model.FieldMetadata) error {
	visited := make(map[string]bool)

	for path, meta := range metadata {
		if err := applyMetadata(root, path, meta); err != nil {
			return err
		}
		visited[path] = true
	}

	if err := validator.ValidateDynamicMaps(root, ""); err != nil {
		return err
	}

	return nil
}

func applyMetadata(root *model.Field, path string, meta model.FieldMetadata) error {
	tokens := strings.Split(path, ".")
	current := root

	for _, token := range tokens {
		switch {
		case token == "*":
			if current.AdditionalProperties == nil {
				current.AdditionalProperties = &model.Field{
					Name: "*",
				}
			}
			current = current.AdditionalProperties

		case strings.HasSuffix(token, "[]"):
			base := strings.TrimSuffix(token, "[]")

			if current.Properties == nil {
				current.Properties = make(map[string]*model.Field)
			}

			if _, ok := current.Properties[base]; !ok {
				current.Properties[base] = &model.Field{
					Name: base,
					Type: "array",
				}
			}

			arrayField := current.Properties[base]

			// Garante que é array
			if arrayField.Type != "" && arrayField.Type != "array" {
				return fmt.Errorf(
					"conflito de tipo no path '%s': esperado array mas é %s",
					path,
					arrayField.Type,
				)
			}

			if arrayField.Items == nil {
				arrayField.Items = &model.Field{
					Name: base + "_item",
				}
			}

			current = arrayField.Items

		default:
			if current.Properties == nil {
				current.Properties = make(map[string]*model.Field)
			}

			if _, ok := current.Properties[token]; !ok {
				current.Properties[token] = &model.Field{
					Name: token,
				}
			}

			current = current.Properties[token]
		}
	}

	current.Description = meta.Description
	current.Required = meta.Required

	if meta.Type != "" {

		// Não sobrescreve tipo estrutural diferente
		if current.Type != "" && current.Type != meta.Type {
			// Permite multi-type (string|number)
			if strings.Contains(meta.Type, "|") {
				current.Type = meta.Type
			} else {
				return fmt.Errorf(
					"conflito de tipo no path '%s': YAML=%s CSV=%s",
					path,
					current.Type,
					meta.Type,
				)
			}
		} else {
			current.Type = meta.Type
		}
	}

	return nil
}
