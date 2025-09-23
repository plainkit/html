package html

import "strings"

type OlAttrs struct {
	Global   GlobalAttrs
	Compact  string
	Reversed bool
	Start    string
	Type     string
}

type OlArg interface {
	applyOl(*OlAttrs, *[]Component)
}

func defaultOlAttrs() *OlAttrs {
	return &OlAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ol(args ...OlArg) Node {
	a := defaultOlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyOl(a, &kids)
	}
	return Node{Tag: "ol", Attrs: a, Kids: kids}
}

func (g Global) applyOl(a *OlAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyOl(_ *OlAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyOl(_ *OlAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o CompactOpt) applyOl(a *OlAttrs, _ *[]Component) {
	a.Compact = o.v
}
func (o ReversedOpt) applyOl(a *OlAttrs, _ *[]Component) {
	a.Reversed = true
}
func (o StartOpt) applyOl(a *OlAttrs, _ *[]Component) {
	a.Start = o.v
}
func (o TypeOpt) applyOl(a *OlAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *OlAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Compact != "" {
		Attr(sb, "compact", a.Compact)
	}
	if a.Reversed {
		BoolAttr(sb, "reversed")
	}
	if a.Start != "" {
		Attr(sb, "start", a.Start)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
