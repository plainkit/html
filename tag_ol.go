package html

import "strings"

type OlAttrs struct {
	Global   GlobalAttrs
	Reversed bool
	Start    string
	Type     string
}

type OlArg interface {
	ApplyOl(*OlAttrs, *[]Component)
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
		ar.ApplyOl(a, &kids)
	}

	return Node{Tag: "ol", Attrs: a, Kids: kids}
}

func (g Global) ApplyOl(a *OlAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o ReversedOpt) ApplyOl(a *OlAttrs, _ *[]Component) {
	a.Reversed = true
}
func (o StartOpt) ApplyOl(a *OlAttrs, _ *[]Component) {
	a.Start = o.v
}
func (o TypeOpt) ApplyOl(a *OlAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *OlAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

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
