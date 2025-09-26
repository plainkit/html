package html

import "strings"

type SvgAttrs struct {
	Global GlobalAttrs
}

type SvgArg interface {
	applySvg(*SvgAttrs, *[]Component)
}

func defaultSvgAttrs() *SvgAttrs {
	return &SvgAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Svg(args ...SvgArg) Node {
	a := defaultSvgAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySvg(a, &kids)
	}
	return Node{Tag: "svg", Attrs: a, Kids: kids}
}

func (g Global) applySvg(a *SvgAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *SvgAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
