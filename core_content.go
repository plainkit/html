package html

import "strings"

// Non-global content helpers (unified for all tags that accept children)

type TxtOpt struct {
	s string
}

func Text(s string) TxtOpt {
	return TxtOpt{s}
}

// T is an alias for Text to reduce verbosity
func T(s string) TxtOpt { return Text(s) }

type UnsafeTxtOpt struct {
	s string
}

func UnsafeText(s string) UnsafeTxtOpt {
	return UnsafeTxtOpt{s}
}

type ChildOpt struct {
	c Component
}

func Child(c Component) ChildOpt {
	return ChildOpt{c}
}

func C(c Component) ChildOpt { return Child(c) }

// FragmentNode represents a collection of components that render without a wrapper element,
// similar to React's <> fragment syntax. It implements Component and can be used
// anywhere a single Component is expected, but renders as multiple sibling elements.
type FragmentNode struct {
	children []Component
}

// Fragment creates a fragment containing the given components.
// Like React fragments, this renders the children directly without any wrapper element.
//
// Example:
//
//	Fragment(
//	  Div(T("First child")),
//	  P(T("Second child")),
//	  Span(T("Third child")),
//	)
//
// This renders as three sibling elements with no containing wrapper.
func Fragment(children ...Component) ChildOpt {
	return C(FragmentNode{children: children})
}

// F is an alias for Fragment to reduce verbosity
func F(children ...Component) ChildOpt { return Fragment(children...) }

// render implements Component interface by rendering each child component
// in sequence without any wrapper element.
func (f FragmentNode) render(sb *strings.Builder) {
	for _, child := range f.children {
		child.render(sb)
	}
}

// Children exposes the fragment's children for traversals that need to walk
// the component tree (e.g., asset collection).
func (f FragmentNode) Children() []Component {
	return f.children
}

// String returns the text content
func (t TxtOpt) String() string { return t.s }

// String returns the unsafe text content
func (t UnsafeTxtOpt) String() string { return t.s }
