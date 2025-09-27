package html

import "strings"

type SourceAttrs struct {
	Global GlobalAttrs
	Media  string
	Sizes  string
	Src    string
	Srcset string
	Type   string
}

type SourceArg interface {
	applySource(*SourceAttrs, *[]Component)
}

func defaultSourceAttrs() *SourceAttrs {
	return &SourceAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Source(args ...SourceArg) Node {
	a := defaultSourceAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySource(a, &kids)
	}
	return Node{Tag: "source", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applySource(a *SourceAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o MediaOpt) applySource(a *SourceAttrs, _ *[]Component) {
	a.Media = o.v
}
func (o SizesOpt) applySource(a *SourceAttrs, _ *[]Component) {
	a.Sizes = o.v
}
func (o SrcOpt) applySource(a *SourceAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o SrcsetOpt) applySource(a *SourceAttrs, _ *[]Component) {
	a.Srcset = o.v
}
func (o TypeOpt) applySource(a *SourceAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *SourceAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Media != "" {
		Attr(sb, "media", a.Media)
	}
	if a.Sizes != "" {
		Attr(sb, "sizes", a.Sizes)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Srcset != "" {
		Attr(sb, "srcset", a.Srcset)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
