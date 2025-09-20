package html

import "strings"

// DT (Description Term)
type DtAttrs struct {
	Global GlobalAttrs
}

type DtArg interface {
	applyDt(*DtAttrs, *[]Component)
}

func defaultDtAttrs() *DtAttrs {
	return &DtAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dt(args ...DtArg) Node {
	a := defaultDtAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDt(a, &kids)
	}
	return Node{Tag: "dt", Attrs: a, Kids: kids}
}

func (g Global) applyDt(a *DtAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyDt(_ *DtAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyDt(_ *DtAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *DtAttrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
