package html

import "strings"

type LabelAttrs struct {
	Global GlobalAttrs
	For    string
}

type LabelArg interface {
	applyLabel(*LabelAttrs, *[]Component)
}

func defaultLabelAttrs() *LabelAttrs {
	return &LabelAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Label(args ...LabelArg) Node {
	a := defaultLabelAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyLabel(a, &kids)
	}
	return Node{Tag: "label", Attrs: a, Kids: kids}
}

func (g Global) applyLabel(a *LabelAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o ForOpt) applyLabel(a *LabelAttrs, _ *[]Component) {
	a.For = o.v
}

func (a *LabelAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.For != "" {
		Attr(sb, "for", a.For)
	}
}
