package html

import "strings"

// HeadStyleAttrs attributes for the <style> tag
type HeadStyleAttrs struct {
	Global GlobalAttrs
	Type   string
	Media  string
}

type HeadStyleArg interface {
	applyHeadStyle(*HeadStyleAttrs, *[]Component)
}

func defaultHeadStyleAttrs() *HeadStyleAttrs {
	return &HeadStyleAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

type StyleComponent Node

func (style StyleComponent) render(sb *strings.Builder) {
	Node(style).render(sb)
}

func HeadStyle(args ...HeadStyleArg) StyleComponent {
	a := defaultHeadStyleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHeadStyle(a, &kids)
	}
	return StyleComponent{Tag: "style", Attrs: a, Kids: kids}
}

// Style-specific options
type MediaOpt struct{ v string }
type StyleTypeOpt struct{ v string }

func Media(v string) MediaOpt         { return MediaOpt{v} }
func StyleType(v string) StyleTypeOpt { return StyleTypeOpt{v} }

func (g Global) applyHeadStyle(a *HeadStyleAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyHeadStyle(_ *HeadStyleAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o UnsafeTxtOpt) applyHeadStyle(_ *HeadStyleAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o ChildOpt) applyHeadStyle(_ *HeadStyleAttrs, kids *[]Component)  { *kids = append(*kids, o.c) }
func (o MediaOpt) applyHeadStyle(a *HeadStyleAttrs, _ *[]Component)     { a.Media = o.v }
func (o StyleTypeOpt) applyHeadStyle(a *HeadStyleAttrs, _ *[]Component) { a.Type = o.v }

// Compile-time type safety: Style can be added to Head
func (style StyleComponent) applyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, style)
}

func (a *HeadStyleAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Media != "" {
		attr(sb, "media", a.Media)
	}
}
