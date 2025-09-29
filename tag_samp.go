package html

import "strings"

type SampAttrs struct {
	Global GlobalAttrs
}

type SampArg interface {
	ApplySamp(*SampAttrs, *[]Component)
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
		ar.ApplySamp(a, &kids)
	}

	return Node{Tag: "samp", Attrs: a, Kids: kids}
}

func (g Global) ApplySamp(a *SampAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *SampAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
