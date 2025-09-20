package html

import "strings"

// Source (void)
type SourceAttrs struct {
	Global GlobalAttrs
	Src    string
	Type   string
	Media  string
	Sizes  string
	Srcset string
}

type SourceArg interface {
	applySource(*SourceAttrs)
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
	for _, ar := range args {
		ar.applySource(a)
	}
	return Node{Tag: "source", Attrs: a, Void: true}
}

// Source-specific options
type SrcsetOpt struct{ v string }

func Srcset(v string) SrcsetOpt { return SrcsetOpt{v} }

func (g Global) applySource(a *SourceAttrs)    { g.do(&a.Global) }
func (o SrcOpt) applySource(a *SourceAttrs)    { a.Src = o.v }
func (o TypeOpt) applySource(a *SourceAttrs)   { a.Type = o.v }
func (o MediaOpt) applySource(a *SourceAttrs)  { a.Media = o.v }
func (o SizesOpt) applySource(a *SourceAttrs)  { a.Sizes = o.v }
func (o SrcsetOpt) applySource(a *SourceAttrs) { a.Srcset = o.v }

func (a *SourceAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Src != "" {
		attr(sb, "src", a.Src)
	}
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Media != "" {
		attr(sb, "media", a.Media)
	}
	if a.Sizes != "" {
		attr(sb, "sizes", a.Sizes)
	}
	if a.Srcset != "" {
		attr(sb, "srcset", a.Srcset)
	}
}
