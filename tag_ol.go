package blox

import "strings"

// OL (Ordered List)
type OlAttrs struct {
	Global   GlobalAttrs
	Start    int
	Type     string
	Reversed bool
}

type OlArg interface {
	applyOl(*OlAttrs, *[]Component)
}

func defaultOlAttrs() *OlAttrs {
	return &OlAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ol(args ...OlArg) Component {
	a := defaultOlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyOl(a, &kids)
	}
	return Node{Tag: "ol", Attrs: a, Kids: kids}
}

// OL-specific options
type StartOpt struct{ v int }
type TypeOpt struct{ v string }
type ReversedOpt struct{}

func Start(v int) StartOpt  { return StartOpt{v} }
func Type(v string) TypeOpt { return TypeOpt{v} }
func Reversed() ReversedOpt { return ReversedOpt{} }

func (g Global) applyOl(a *OlAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyOl(_ *OlAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyOl(_ *OlAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o StartOpt) applyOl(a *OlAttrs, _ *[]Component)    { a.Start = o.v }
func (o TypeOpt) applyOl(a *OlAttrs, _ *[]Component)     { a.Type = o.v }
func (o ReversedOpt) applyOl(a *OlAttrs, _ *[]Component) { a.Reversed = true }

func (a *OlAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Start > 0 {
		attr(sb, "start", itoa(a.Start))
	}
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Reversed {
		boolAttr(sb, "reversed")
	}
}
