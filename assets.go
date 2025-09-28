package html

import "strings"

// HasCSS interface for components that provide CSS
type HasCSS interface {
	CSS() string
}

// HasJS interface for components that provide JavaScript
type HasJS interface {
	JS() string
}

// HasName interface for components that provide explicit names for deduplication
type HasName interface {
	Name() string
}

// Assets collects CSS and JS from components at compile-time
type Assets struct {
	css      []string
	js       []string
	registry map[string]bool
}

// NewAssets creates a new asset collector
func NewAssets() *Assets {
	return &Assets{
		registry: make(map[string]bool),
	}
}

// Collect walks the component tree and gathers unique CSS/JS assets
func (a *Assets) Collect(components ...Component) {
	a.collectRecursive(components)
}

// collectRecursive walks components recursively, deduplicating by name
func (a *Assets) collectRecursive(components []Component) {
	for _, comp := range components {
		// Get component name for deduplication
		var compName string
		if namedComp, ok := comp.(HasName); ok {
			compName = namedComp.Name()
		}
		if compName == "" {
			// Fallback: use a simple string representation for unnamed components
			compName = "unnamed"
		}

		// Skip if we've already processed this component type
		if compName != "unnamed" && a.registry[compName] {
			continue // Already collected this component type
		}

		if compName != "unnamed" {
			a.registry[compName] = true
		}

		// Collect CSS if component provides it
		if cssComp, ok := comp.(HasCSS); ok {
			css := strings.TrimSpace(cssComp.CSS())
			if css != "" {
				a.css = append(a.css, css)
			}
		}

		// Collect JS if component provides it
		if jsComp, ok := comp.(HasJS); ok {
			js := strings.TrimSpace(jsComp.JS())
			if js != "" {
				a.js = append(a.js, js)
			}
		}

		// Try to get children if component has them
		if hasChildren, ok := comp.(interface{ Children() []Component }); ok {
			a.collectRecursive(hasChildren.Children())
		}
	}
}

// HasAssets returns true if any CSS or JS was collected
func (a *Assets) HasAssets() bool {
	return len(a.css) > 0 || len(a.js) > 0
}

// CSS returns the collected CSS snippets in insertion order.
func (a *Assets) CSS() []string {
	return append([]string(nil), a.css...)
}

// JS returns the collected JavaScript snippets in insertion order.
func (a *Assets) JS() []string {
	return append([]string(nil), a.js...)
}

// Reset clears all collected assets
func (a *Assets) Reset() {
	a.css = a.css[:0]
	a.js = a.js[:0]
	a.registry = make(map[string]bool)
}

// assetHook is a no-op component that contributes CSS/JS to the asset collector.
type assetHook struct {
	name string
	css  string
	js   string
}

func (h assetHook) render(_ *strings.Builder) {}
func (h assetHook) CSS() string               { return strings.TrimSpace(h.css) }
func (h assetHook) JS() string                { return strings.TrimSpace(h.js) }
func (h assetHook) Name() string              { return h.name }

// AssetHook creates a component that carries CSS/JS for collection without rendering output.
// Name is used for de-duplication across the page.
func AssetHook(name, css, js string) Component { return assetHook{name: name, css: css, js: js} }
