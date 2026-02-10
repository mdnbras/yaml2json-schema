package merger

import "github.com/mdnbras/yaml2json-schema/internal/model"

var (
	metaMap map[string]model.FieldMetadata
	visited map[string]bool
)

func getMetadata(path string) (model.FieldMetadata, bool) {
	meta, ok := metaMap[path]
	return meta, ok
}

func markVisited(path string) {
	visited[path] = true
}
