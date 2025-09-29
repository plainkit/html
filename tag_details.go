package html

import "strings"

type DetailsAttrs struct {
	Global GlobalAttrs
	Name   string
	Open   bool
}

type DetailsArg interface {
	ApplyDetails(*DetailsAttrs, *[]Component)
}

func defaultDetailsAttrs() *DetailsAttrs {
	return &DetailsAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Details(args ...DetailsArg) Node {
	a := defaultDetailsAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyDetails(a, &kids)
	}

	return Node{Tag: "details", Attrs: a, Kids: kids}
}

func (g Global) ApplyDetails(a *DetailsAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o NameOpt) ApplyDetails(a *DetailsAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o OpenOpt) ApplyDetails(a *DetailsAttrs, _ *[]Component) {
	a.Open = true
}

func (a *DetailsAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}

	if a.Open {
		BoolAttr(sb, "open")
	}
}
