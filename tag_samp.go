package html

import "strings"

type SampAttrs struct {
	Global GlobalAttrs
}

type SampArg interface {
	applySamp(*SampAttrs, *[]Component)
}

func defaultSampAttrs() *SampAttrs {
	return &SampAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Samp(args ...SampArg) Node {
	a := defaultSampAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySamp(a, &kids)
	}
	return Node{Tag: "samp", Attrs: a, Kids: kids}
}

func (g Global) applySamp(a *SampAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *SampAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
