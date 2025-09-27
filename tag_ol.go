package html

import "strings"

type OlAttrs struct {
	Global   GlobalAttrs
	Reversed bool
	Start    string
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

func (o ReversedOpt) applyOl(a *OlAttrs, _ *[]Component) {
	a.Reversed = true
}
func (o StartOpt) applyOl(a *OlAttrs, _ *[]Component) {
	a.Start = o.v
}

func (a *OlAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Reversed {
		BoolAttr(sb, "reversed")
	}
	if a.Start != "" {
		Attr(sb, "start", a.Start)
	}
}
