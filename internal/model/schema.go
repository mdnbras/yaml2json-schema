package model

// Schema representa um JSON Schema Draft-07.
// Esta estrutura é compatível com o meta-schema oficial
// e suporta objects, arrays, maps dinâmicos e validação forte.
type Schema struct {
	// Meta
	Schema string `json:"$schema,omitempty"`

	// Core
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`

	// Object
	Properties map[string]*Schema `json:"properties,omitempty"`
	Required   []string           `json:"required,omitempty"`

	// Array
	Items *Schema `json:"items,omitempty"`

	// additionalProperties pode ser:
	// - false  → objeto fechado
	// - Schema → map dinâmico (Draft-07)
	AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
}
