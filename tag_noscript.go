package html

import "strings"

type NoscriptAttrs struct {
	Global GlobalAttrs
}

type NoscriptArg interface {
	ApplyNoscript(*NoscriptAttrs, *[]Component)
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
		ar.ApplyNoscript(a, &kids)
	}

	return Node{Tag: "noscript", Attrs: a, Kids: kids}
}

func (g Global) ApplyNoscript(a *NoscriptAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *NoscriptAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
