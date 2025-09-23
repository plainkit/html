package files

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

// Manager handles file operations for code generation
type Manager struct {
	outputDir string
}

// NewManager creates a new file manager
func NewManager(outputDir string) *Manager {
	return &Manager{outputDir: outputDir}
}

// EnsureOutputDir creates the output directory if it doesn't exist
func (m *Manager) EnsureOutputDir() error {
	return os.MkdirAll(m.outputDir, 0o755)
}

// WriteFormattedFile writes Go source code to a file with formatting
func (m *Manager) WriteFormattedFile(fileName, source string) error {
	filePath := filepath.Join(m.outputDir, fileName)

	// Format the source code
	formatted, err := format.Source([]byte(source))
	if err != nil {
		// Write unformatted version for debugging
		if debugErr := m.writeUnformattedDebugFile(fileName, source); debugErr != nil {
			fmt.Printf("Error writing debug file: %v\n", debugErr)
		}
		return fmt.Errorf("format source for %s: %w", fileName, err)
	}

	// Write the formatted file
	if err := os.WriteFile(filePath, formatted, 0o644); err != nil {
		return fmt.Errorf("write file %s: %w", fileName, err)
	}

	fmt.Printf("Successfully wrote %s (%d bytes)\n", fileName, len(formatted))
	return nil
}

// writeUnformattedDebugFile writes unformatted source for debugging
func (m *Manager) writeUnformattedDebugFile(fileName, source string) error {
	debugPath := filepath.Join(m.outputDir, fileName+".unformatted")
	return os.WriteFile(debugPath, []byte(source), 0o644)
}

// ReadFile reads a file from the filesystem
func (m *Manager) ReadFile(fileName string) (string, error) {
	filePath := filepath.Join(m.outputDir, fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("read file %s: %w", fileName, err)
	}
	return string(data), nil
}

// WriteFile writes content to a file (without formatting)
func (m *Manager) WriteFile(fileName, content string) error {
	filePath := filepath.Join(m.outputDir, fileName)
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write file %s: %w", fileName, err)
	}
	return nil
}

// CleanGeneratedFiles removes previously generated files
func (m *Manager) CleanGeneratedFiles() error {
	fmt.Println("Cleaning previously generated files...")

	entries, err := os.ReadDir(m.outputDir)
	if err != nil {
		return fmt.Errorf("read output directory: %w", err)
	}

	filesToRemove := []string{}
	for _, entry := range entries {
		name := entry.Name()
		if m.isGeneratedFile(name) {
			filesToRemove = append(filesToRemove, name)
		}
	}

	for _, fileName := range filesToRemove {
		filePath := filepath.Join(m.outputDir, fileName)
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			fmt.Printf("Warning: failed to remove %s: %v\n", fileName, err)
		} else {
			fmt.Printf("Removed %s\n", fileName)
		}
	}

	fmt.Printf("Cleaned %d generated files\n", len(filesToRemove))
	return nil
}

// isGeneratedFile determines if a file should be cleaned up
func (m *Manager) isGeneratedFile(fileName string) bool {
	return strings.HasPrefix(fileName, "tag_") ||
		fileName == "attrs.go" ||
		fileName == "core_options.go" ||
		fileName == "core_global.go"
}

// FileExists checks if a file exists in the output directory
func (m *Manager) FileExists(fileName string) bool {
	filePath := filepath.Join(m.outputDir, fileName)
	_, err := os.Stat(filePath)
	return err == nil
}
