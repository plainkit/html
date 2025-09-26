package files

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

// Manager handles file operations for generated code
type Manager struct {
	outDir string
}

// NewManager creates a new file manager for the specified output directory
func NewManager(outDir string) *Manager {
	return &Manager{outDir: outDir}
}

// EnsureOutputDir ensures the output directory exists
func (m *Manager) EnsureOutputDir() error {
	if err := os.MkdirAll(m.outDir, 0755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	return nil
}

// CleanGeneratedFiles removes all generated tag files from the output directory
func (m *Manager) CleanGeneratedFiles() error {
	entries, err := os.ReadDir(m.outDir)
	if err != nil {
		// If directory doesn't exist, that's fine
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("read output directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()
		// Remove tag_*.go files (these are generated)
		if strings.HasPrefix(name, "tag_") && strings.HasSuffix(name, ".go") {
			path := filepath.Join(m.outDir, name)
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("remove %s: %w", name, err)
			}
		}
	}

	return nil
}

// WriteFormattedFile writes formatted Go source code to a file
func (m *Manager) WriteFormattedFile(filename, source string) error {
	// Format the source code
	formatted, err := format.Source([]byte(source))
	if err != nil {
		// If formatting fails, write the unformatted source for debugging
		unformattedPath := filepath.Join(m.outDir, filename+".unformatted")
		if writeErr := os.WriteFile(unformattedPath, []byte(source), 0644); writeErr != nil {
			return fmt.Errorf("write unformatted file: %w", writeErr)
		}
		return fmt.Errorf("format source (unformatted saved to %s): %w", unformattedPath, err)
	}

	// Write the formatted file
	path := filepath.Join(m.outDir, filename)
	if err := os.WriteFile(path, formatted, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

// GetOutputPath returns the full path for a file in the output directory
func (m *Manager) GetOutputPath(filename string) string {
	return filepath.Join(m.outDir, filename)
}
