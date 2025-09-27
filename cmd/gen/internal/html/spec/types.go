package spec

// Attribute represents an HTML attribute with its metadata.
type Attribute struct {
	Field string
	Type  string
	Attr  string
}

// TagSpec contains the information needed to generate a tag file.
type TagSpec struct {
	Name          string
	Void          bool
	Attributes    []Attribute
	ParentTargets []string
}
