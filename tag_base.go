package html

import "strings"

type BaseAttrs struct {
	Global GlobalAttrs
	Href   string
	Target string
}

type BaseArg interface {
	applyBase(*BaseAttrs, *[]Component)
}

func defaultBaseAttrs() *BaseAttrs {
	return &BaseAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Base(args ...BaseArg) Node {
	a := defaultBaseAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyBase(a, &kids)
	}
	return Node{Tag: "base", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyBase(a *BaseAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o HrefOpt) applyBase(a *BaseAttrs, _ *[]Component) {
	a.Href = o.v
}
func (o TargetOpt) applyBase(a *BaseAttrs, _ *[]Component) {
	a.Target = o.v
}

func (a *BaseAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Href != "" {
		Attr(sb, "href", a.Href)
	}
	if a.Target != "" {
		Attr(sb, "target", a.Target)
	}
}
