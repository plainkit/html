package html

import "strings"

type ProgressAttrs struct {
	Global GlobalAttrs
	Max    string
	Value  string
}

type ProgressArg interface {
	applyProgress(*ProgressAttrs, *[]Component)
}

func defaultProgressAttrs() *ProgressAttrs {
	return &ProgressAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Progress(args ...ProgressArg) Node {
	a := defaultProgressAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyProgress(a, &kids)
	}
	return Node{Tag: "progress", Attrs: a, Kids: kids}
}

func (g Global) applyProgress(a *ProgressAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyProgress(_ *ProgressAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyProgress(_ *ProgressAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o MaxOpt) applyProgress(a *ProgressAttrs, _ *[]Component) {
	a.Max = o.v
}
func (o ValueOpt) applyProgress(a *ProgressAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *ProgressAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Max != "" {
		Attr(sb, "max", a.Max)
	}
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
}
