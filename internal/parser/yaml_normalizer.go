package parser

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func NormalizeYAML(input []byte) ([]byte, error) {
	var data interface{}

	if err := yaml.Unmarshal(input, &data); err != nil {
		return nil, fmt.Errorf("erro ao decodificar YAML: %w", err)
	}

	return encodeYAML(data)
}
