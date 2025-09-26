package html

import "strings"

type NoscriptAttrs struct {
	Global GlobalAttrs
}

type NoscriptArg interface {
	applyNoscript(*NoscriptAttrs, *[]Component)
}

func defaultNoscriptAttrs() *NoscriptAttrs {
	return &NoscriptAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Noscript(args ...NoscriptArg) Node {
	a := defaultNoscriptAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyNoscript(a, &kids)
	}
	return Node{Tag: "noscript", Attrs: a, Kids: kids}
}

func (g Global) applyNoscript(a *NoscriptAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *NoscriptAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
