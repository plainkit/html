package html

import "strings"

type OutputAttrs struct {
	Global GlobalAttrs
	For    string
	Form   string
	Name   string
}

type OutputArg interface {
	applyOutput(*OutputAttrs, *[]Component)
}

func defaultOutputAttrs() *OutputAttrs {
	return &OutputAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Output(args ...OutputArg) Node {
	a := defaultOutputAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyOutput(a, &kids)
	}
	return Node{Tag: "output", Attrs: a, Kids: kids}
}

func (g Global) applyOutput(a *OutputAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o ForOpt) applyOutput(a *OutputAttrs, _ *[]Component) {
	a.For = o.v
}
func (o FormOpt) applyOutput(a *OutputAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o NameOpt) applyOutput(a *OutputAttrs, _ *[]Component) {
	a.Name = o.v
}

func (a *OutputAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.For != "" {
		Attr(sb, "for", a.For)
	}
	if a.Form != "" {
		Attr(sb, "form", a.Form)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
}
