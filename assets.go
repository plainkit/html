package blox

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
		} else {
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

// CSS returns a style component with all collected CSS
func (a *Assets) CSS() Component {
	if len(a.css) == 0 {
		return TextNode("")
	}
	return Style(Text(strings.Join(a.css, "\n\n")))
}

// JS returns a script component with all collected JavaScript
func (a *Assets) JS() Component {
	if len(a.js) == 0 {
		return TextNode("")
	}
	return Script(UnsafeText(strings.Join(a.js, "\n\n")))
}

// HasAssets returns true if any CSS or JS was collected
func (a *Assets) HasAssets() bool {
	return len(a.css) > 0 || len(a.js) > 0
}

// Reset clears all collected assets
func (a *Assets) Reset() {
	a.css = a.css[:0]
	a.js = a.js[:0]
	a.registry = make(map[string]bool)
}
