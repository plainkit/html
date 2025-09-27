package spec

import "encoding/json"

// BrowserCompatData represents the structure of MDN browser compatibility data files
type BrowserCompatData struct {
	HTML struct {
		Elements map[string]map[string]json.RawMessage `json:"elements"`
	} `json:"html"`
}

// GlobalBrowserCompatData represents the structure of global attributes spec
type GlobalBrowserCompatData struct {
	HTML struct {
		GlobalAttributes map[string]json.RawMessage `json:"global_attributes"`
	} `json:"html"`
}

// Attribute represents an HTML attribute with its metadata
type Attribute struct {
	Field string
	Type  string
	Attr  string
}

// TagSpec contains all information needed to generate a tag file
type TagSpec struct {
	Name          string
	Void          bool
	Attributes    []Attribute
	ParentTargets []string
}

// GlobalAttributesSpec represents the structure of global_attributes.json
type GlobalAttributesSpec struct {
	Html struct {
		GlobalAttributes map[string]interface{} `json:"global_attributes"`
	} `json:"html"`
}
