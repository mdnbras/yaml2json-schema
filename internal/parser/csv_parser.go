package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/mdnbras/yaml2json-schema/internal/model"
)

func ParseCSV(path string) (map[string]model.FieldMetadata, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler CSV: %w", err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV deve conter cabeçalho e ao menos uma linha")
	}

	header := normalizeHeader(records[0])
	index := indexHeader(header)

	metadata := make(map[string]model.FieldMetadata)

	for i, row := range records[1:] {
		if len(row) < len(header) {
			return nil, fmt.Errorf("linha %d inválida no CSV", i+2)
		}

		path := row[index["path"]]

		metadata[path] = model.FieldMetadata{
			Path:        path,
			Description: row[index["description"]],
			Required:    parseBool(row[index["required"]]),
			Type:        row[index["type"]],
		}
	}

	return metadata, nil
}

func normalizeHeader(header []string) []string {
	for i := range header {
		header[i] = strings.ToLower(strings.TrimSpace(header[i]))
	}
	return header
}

func indexHeader(header []string) map[string]int {
	required := []string{"path", "description", "required", "type"}
	index := make(map[string]int)

	for i, h := range header {
		index[h] = i
	}

	for _, r := range required {
		if _, ok := index[r]; !ok {
			panic(fmt.Sprintf("coluna obrigatória ausente no CSV: %s", r))
		}
	}

	return index
}

func parseBool(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "true", "yes", "1", "y":
		return true
	default:
		return false
	}
}
