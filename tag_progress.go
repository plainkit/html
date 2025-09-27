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
	g.Do(&a.Global)
}

func (o MaxOpt) applyProgress(a *ProgressAttrs, _ *[]Component) {
	a.Max = o.v
}
func (o ValueOpt) applyProgress(a *ProgressAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *ProgressAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Max != "" {
		Attr(sb, "max", a.Max)
	}
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
}
