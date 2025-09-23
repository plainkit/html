package html

import "strings"

// LI (List Item)
type LiAttrs struct {
	Global   GlobalAttrs
	Value    int
	valueSet bool
}

type LiArg interface {
	applyLi(*LiAttrs, *[]Component)
}

func defaultLiAttrs() *LiAttrs {
	return &LiAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Li(args ...LiArg) LiComponent {
	a := defaultLiAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyLi(a, &kids)
	}
	return LiComponent{Tag: "li", Attrs: a, Kids: kids}
}

// LI-specific options
type ValueOpt struct{ v int }

func Value(v int) ValueOpt { return ValueOpt{v} }

func (g Global) applyLi(a *LiAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyLi(_ *LiAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyLi(_ *LiAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ValueOpt) applyLi(a *LiAttrs, _ *[]Component) {
	a.Value = o.v
	a.valueSet = true
}

// Compile-time type safety: Li can be added to Ul and Ol
// This makes Li() return something that implements both UlArg and OlArg
type LiComponent Node

func (li LiComponent) render(sb *strings.Builder) {
	Node(li).render(sb)
}

func (li LiComponent) applyUl(_ *UlAttrs, kids *[]Component) {
	*kids = append(*kids, li)
}

func (li LiComponent) applyOl(_ *OlAttrs, kids *[]Component) {
	*kids = append(*kids, li)
}

func (a *LiAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.valueSet {
		attr(sb, "value", itoa(a.Value))
	}
}
