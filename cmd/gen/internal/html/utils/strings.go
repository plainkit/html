package utils

import (
	"github.com/plainkit/html/cmd/gen/internal/html/spec"
)

// CamelCase converts kebab-case to CamelCase
func CamelCase(name string) string {
	return spec.CamelCase(name)
}

// GoType converts spec types to Go types
func GoType(t string) string {
	if t == "bool" {
		return "bool"
	}

	return "string"
}
