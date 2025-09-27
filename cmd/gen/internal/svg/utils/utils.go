package utils

import "strings"

// CamelCase converts kebab-case or snake_case to CamelCase
func CamelCase(name string) string {
	// Handle special cases for SVG attributes with dashes and mixed case
	delimiters := func(r rune) bool { return r == '-' || r == '_' }
	parts := strings.FieldsFunc(name, delimiters)

	if len(parts) == 0 {
		return name
	}

	// First part: capitalize first letter
	result := ""
	if len(parts[0]) > 0 {
		result = strings.ToUpper(parts[0][:1]) + parts[0][1:]
	}

	// Capitalize subsequent parts
	for i := 1; i < len(parts); i++ {
		p := parts[i]
		if len(p) > 0 {
			result += strings.ToUpper(p[:1]) + p[1:]
		}
	}

	return result
}

// GoType converts attribute type to Go type
func GoType(attrType string) string {
	if attrType == "bool" {
		return "bool"
	}
	return "string"
}
