package html

import "strings"

// DL (Description List)
type DlAttrs struct {
	Global GlobalAttrs
}

type DlArg interface {
	applyDl(*DlAttrs, *[]Component)
}

func defaultDlAttrs() *DlAttrs {
	return &DlAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dl(args ...DlArg) Node {
	a := defaultDlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDl(a, &kids)
	}
	return Node{Tag: "dl", Attrs: a, Kids: kids}
}

func (g Global) applyDl(a *DlAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyDl(_ *DlAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyDl(_ *DlAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *DlAttrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
