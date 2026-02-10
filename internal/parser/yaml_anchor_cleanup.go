package parser

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func RemoveAnchorBases(
	normalized []byte,
	bases []string,
) ([]byte, error) {

	if len(bases) == 0 {
		return normalized, nil
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(normalized, &data); err != nil {
		return nil, fmt.Errorf("erro ao decodificar YAML normalizado: %w", err)
	}

	for _, base := range bases {
		delete(data, base)
	}

	return encodeYAML(data)
}
