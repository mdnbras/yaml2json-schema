package utils

import (
	"github.com/sergi/go-diff/diffmatchpatch"
)

// DiffYAML retorna um diff unificado entre dois textos YAML
func DiffYAML(original, normalized []byte) string {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(
		string(original),
		string(normalized),
		false,
	)

	patches := dmp.PatchMake(diffs)
	return dmp.PatchToText(patches)
}
