package html

import "strings"

type XmpAttrs struct {
	Global GlobalAttrs
}

type XmpArg interface {
	applyXmp(*XmpAttrs, *[]Component)
}

func defaultXmpAttrs() *XmpAttrs {
	return &XmpAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Xmp(args ...XmpArg) Node {
	a := defaultXmpAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyXmp(a, &kids)
	}
	return Node{Tag: "xmp", Attrs: a, Kids: kids}
}

func (g Global) applyXmp(a *XmpAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyXmp(_ *XmpAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyXmp(_ *XmpAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *XmpAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
