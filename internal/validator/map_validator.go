package validator

import (
	"fmt"

	"github.com/mdnbras/yaml2json-schema/internal/model"
)

func ValidateDynamicMaps(field *model.Field, path string) error {
	currentPath := path
	if field.Name != "root" {
		if path == "" {
			currentPath = field.Name
		} else {
			currentPath = path + "." + field.Name
		}
	}

	if field.Type == "object" {
		if field.AdditionalProperties == nil && len(field.Properties) == 0 {
			return fmt.Errorf(
				"map dinâmico sem definição de '*': %s",
				currentPath,
			)
		}

		for _, child := range field.Properties {
			if err := ValidateDynamicMaps(child, currentPath); err != nil {
				return err
			}
		}

		if field.AdditionalProperties != nil {
			return ValidateDynamicMaps(field.AdditionalProperties, currentPath+".*")
		}
	}

	return nil
}
