package blox

import "strings"

// Style
type StyleAttrs struct {
	Global GlobalAttrs
	Type   string
	Media  string
}

type StyleArg interface {
	applyStyle(*StyleAttrs, *[]Component)
}

func defaultStyleAttrs() *StyleAttrs {
	return &StyleAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Style(args ...StyleArg) Node {
	a := defaultStyleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyStyle(a, &kids)
	}
	return Node{Tag: "style", Attrs: a, Kids: kids}
}

// Style-specific options
type MediaOpt struct{ v string }

func Media(v string) MediaOpt { return MediaOpt{v} }

func (g Global) applyStyle(a *StyleAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyStyle(_ *StyleAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyStyle(_ *StyleAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o MediaOpt) applyStyle(a *StyleAttrs, _ *[]Component)    { a.Media = o.v }

func (a *StyleAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Media != "" {
		attr(sb, "media", a.Media)
	}
}
