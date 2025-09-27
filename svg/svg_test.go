package svg

import (
	"strings"
	"testing"

	"github.com/plainkit/html"
)

func TestPathWithAttributes(t *testing.T) {
	// Test that a single path with d attribute renders correctly
	path := Path(AD("M10 10 L20 20"))
	rendered := html.Render(path)

	expected := `<path d="M10 10 L20 20"/>`
	if rendered != expected {
		t.Errorf("Expected %q, got %q", expected, rendered)
	}
}

func TestMultiplePathsInSvg(t *testing.T) {
	// Test that multiple paths render as siblings, not nested
	svgElement := Svg(
		AdaptNode(Path(AD("M10 10 L20 20"))),
		AdaptNode(Path(AD("M30 30 L40 40"))),
	)
	rendered := html.Render(svgElement)

	// Should have two separate path elements as siblings
	if !strings.Contains(rendered, `<path d="M10 10 L20 20"/>`) {
		t.Errorf("Missing first path in rendered output: %s", rendered)
	}
	if !strings.Contains(rendered, `<path d="M30 30 L40 40"/>`) {
		t.Errorf("Missing second path in rendered output: %s", rendered)
	}

	// Should not have nested paths
	if strings.Contains(rendered, "<path><path>") {
		t.Errorf("Paths should not be nested: %s", rendered)
	}

	t.Logf("Rendered SVG: %s", rendered)
}
