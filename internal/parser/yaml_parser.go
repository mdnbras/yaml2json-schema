package parser

import (
	"fmt"
	"os"

	"github.com/mdnbras/yaml2json-schema/internal/model"
	"gopkg.in/yaml.v3"
)

func ParseYAML(path string) (*model.Field, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler YAML: %w", err)
	}

	var root yaml.Node
	if err := yaml.Unmarshal(content, &root); err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do YAML: %w", err)
	}

	if len(root.Content) == 0 {
		return nil, fmt.Errorf("YAML vazio ou inválido")
	}

	return parseYAMLNode("root", root.Content[0])
}

func parseYAMLNode(name string, node *yaml.Node) (*model.Field, error) {
	field := &model.Field{
		Name: name,
	}

	switch node.Kind {

	case yaml.MappingNode:
		field.Type = "object"
		field.Properties = make(map[string]*model.Field)

		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i].Value
			value := node.Content[i+1]

			child, err := parseYAMLNode(key, value)
			if err != nil {
				return nil, err
			}

			field.Properties[key] = child
		}

	case yaml.SequenceNode:
		field.Type = "array"

		if len(node.Content) == 0 {
			return nil, fmt.Errorf("array vazio não suportado (%s)", name)
		}

		item, err := parseYAMLNode(name+"_item", node.Content[0])
		if err != nil {
			return nil, err
		}

		field.Items = item

	case yaml.ScalarNode:
		field.Type = inferScalarType(node.Value)

	default:
		return nil, fmt.Errorf("tipo YAML não suportado: %v", node.Kind)
	}

	return field, nil
}

func inferScalarType(value string) string {
	switch value {
	case "true", "false":
		return "boolean"
	default:
		return "string"
	}
}
