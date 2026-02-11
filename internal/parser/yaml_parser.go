package parser

import (
	_ "bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mdnbras/yaml2json-schema/internal/model"
	"github.com/mdnbras/yaml2json-schema/internal/utils"
	"gopkg.in/yaml.v3"
)

func ParseYAML(
	path string,
	showDiff bool,
	anchorBases []string,
) (*model.Field, []byte, error) {

	original, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	normalized, err := NormalizeYAML(original)
	if err != nil {
		return nil, nil, err
	}

	cleaned, err := RemoveAnchorBases(normalized, anchorBases)
	if err != nil {
		return nil, nil, err
	}

	if showDiff {
		diff := utils.DiffYAML(original, cleaned)
		if diff != "" {
			fmt.Println("▶ Diff YAML (original × final):")
			fmt.Println(diff)
		}
	}

	var node yaml.Node
	if err := yaml.Unmarshal(cleaned, &node); err != nil {
		return nil, nil, err
	}

	field, err := parseNode("root", node.Content[0])
	if err != nil {
		return nil, nil, err
	}

	return field, cleaned, nil
}

func parseNode(name string, node *yaml.Node) (*model.Field, error) {

	if node.Kind == yaml.AliasNode {
		if node.Alias == nil {
			return nil, fmt.Errorf("alias YAML inválido em %s", name)
		}
		return parseNode(name, node.Alias)
	}

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

			child, err := parseNode(key, value)
			if err != nil {
				return nil, err
			}

			if strings.HasPrefix(key, "[") && strings.HasSuffix(key, "]") {
				field.AdditionalProperties = child
				continue
			}

			field.Properties[key] = child
		}

	case yaml.SequenceNode:
		field.Type = "array"

		if len(node.Content) == 0 {
			return nil, fmt.Errorf("array vazio não suportado: %s", name)
		}

		item, err := parseNode(name+"_item", node.Content[0])
		if err != nil {
			return nil, err
		}

		field.Items = item

	case yaml.ScalarNode:
		field.Type = inferScalarType(node)

	default:
		return nil, fmt.Errorf("tipo YAML não suportado: %d", node.Kind)
	}

	return field, nil
}

func inferScalarType(node *yaml.Node) string {
	// Se o YAML já informa o tipo explicitamente (ex: "string", "number")
	switch node.Value {
	case "string":
		return "string"
	case "number":
		return "number"
	case "integer":
		return "integer"
	case "boolean":
		return "boolean"
	case "object":
		return "object"
	case "array":
		return "array"
	case "null":
		return "null"
	}

	// Inferência básica baseada no tag do YAML
	switch node.Tag {
	case "!!str":
		return "string"
	case "!!int":
		return "integer"
	case "!!float":
		return "number"
	case "!!bool":
		return "boolean"
	case "!!null":
		return "null"
	default:
		// fallback seguro
		return "string"
	}
}
