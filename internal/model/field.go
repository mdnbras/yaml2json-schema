package model

type Field struct {
	Name        string
	Type        string
	Description string
	Required    bool

	Properties           map[string]*Field
	Items                *Field
	AdditionalProperties *Field
}
