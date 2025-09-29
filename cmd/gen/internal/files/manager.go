package files

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
)

type Manager struct {
	Dir string
}

func NewManager(dir string) Manager {
	return Manager{Dir: dir}
}

func (m Manager) EnsureDir() error {
	return os.MkdirAll(m.Dir, 0o755)
}

func (m Manager) Clean(shouldRemove func(name string) bool) error {
	entries, err := os.ReadDir(m.Dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("read output directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if shouldRemove != nil && !shouldRemove(name) {
			continue
		}

		if err := os.Remove(filepath.Join(m.Dir, name)); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove %s: %w", name, err)
		}
	}

	return nil
}

func (m Manager) WriteGoFile(name, source string) error {
	formatted, err := format.Source([]byte(source))
	if err != nil {
		debugPath := filepath.Join(m.Dir, name+".unformatted")
		_ = os.WriteFile(debugPath, []byte(source), 0o644)

		return fmt.Errorf("format %s source: %w", name, err)
	}

	path := filepath.Join(m.Dir, name)
	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", name, err)
	}

	return nil
}

func (m Manager) ReadFile(name string) (string, error) {
	path := filepath.Join(m.Dir, name)

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", name, err)
	}

	return string(data), nil
}

func (m Manager) FileExists(name string) bool {
	_, err := os.Stat(filepath.Join(m.Dir, name))
	return err == nil
}
