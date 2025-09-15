package blox

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
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func HeadStyle(args ...HeadStyleArg) Node {
	a := defaultHeadStyleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHeadStyle(a, &kids)
	}
	return Node{Tag: "style", Attrs: a, Kids: kids}
}

// MediaOpt sets the media attribute
type MediaOpt struct{ v string }

func Media(v string) MediaOpt { return MediaOpt{v} }

func (g Global) applyHeadStyle(a *HeadStyleAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyHeadStyle(_ *HeadStyleAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyHeadStyle(_ *HeadStyleAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o MediaOpt) applyHeadStyle(a *HeadStyleAttrs, _ *[]Component)    { a.Media = o.v }

func (a *HeadStyleAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Media != "" {
		attr(sb, "media", a.Media)
	}
}
