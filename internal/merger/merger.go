package merger

import (
	_ "fmt"
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
		switch token {

		case "*":
			if current.AdditionalProperties == nil {
				current.AdditionalProperties = &model.Field{
					Name: "*",
					Type: meta.Type,
				}
			}
			current = current.AdditionalProperties

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
		current.Type = meta.Type
	}

	return nil
}
