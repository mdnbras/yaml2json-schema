package parser

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v3"
)

func encodeYAML(data interface{}) ([]byte, error) {
	var buf bytes.Buffer

	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)

	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("erro ao codificar YAML: %w", err)
	}

	_ = encoder.Close()
	return buf.Bytes(), nil
}
