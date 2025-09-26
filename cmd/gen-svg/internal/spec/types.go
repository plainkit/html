package spec

// Attribute represents an SVG element attribute
type Attribute struct {
	Field string // Go field name (CamelCase)
	Attr  string // HTML attribute name (kebab-case)
	Type  string // "string" or "bool"
}

// TagSpec represents the specification for a single SVG tag
type TagSpec struct {
	Name       string      // SVG element name
	Void       bool        // Whether the element is self-closing
	Attributes []Attribute // Element-specific attributes
}

// SvgElementAttributes represents the structure from wooorm/svg-element-attributes
type SvgElementAttributes map[string][]string
