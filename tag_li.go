package blox

import "strings"

// LI (List Item)
type LiAttrs struct {
	Global GlobalAttrs
	Value  int
}

type LiArg interface {
	applyLi(*LiAttrs, *[]Component)
}

func defaultLiAttrs() *LiAttrs {
	return &LiAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Li(args ...LiArg) Component {
	a := defaultLiAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyLi(a, &kids)
	}
	return Node{Tag: "li", Attrs: a, Kids: kids}
}

// LI-specific options
type ValueOpt struct{ v int }

func Value(v int) ValueOpt { return ValueOpt{v} }

func (g Global) applyLi(a *LiAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyLi(_ *LiAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyLi(_ *LiAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ValueOpt) applyLi(a *LiAttrs, _ *[]Component)    { a.Value = o.v }

func (a *LiAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Value > 0 {
		attr(sb, "value", itoa(a.Value))
	}
}
