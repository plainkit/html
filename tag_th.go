package blox

import "strings"

// Th
type ThAttrs struct {
	Global  GlobalAttrs
	Colspan int
	Rowspan int
	Headers string
	Scope   string
}

type ThArg interface {
	applyTh(*ThAttrs, *[]Component)
}

func defaultThAttrs() *ThAttrs {
	return &ThAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Th(args ...ThArg) Node {
	a := defaultThAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTh(a, &kids)
	}
	return Node{Tag: "th", Attrs: a, Kids: kids}
}

// Th-specific options
type ScopeOpt struct{ v string }

func Scope(v string) ScopeOpt { return ScopeOpt{v} }

func (g Global) applyTh(a *ThAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyTh(_ *ThAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyTh(_ *ThAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ColspanOpt) applyTh(a *ThAttrs, _ *[]Component)  { a.Colspan = o.v }
func (o RowspanOpt) applyTh(a *ThAttrs, _ *[]Component)  { a.Rowspan = o.v }
func (o HeadersOpt) applyTh(a *ThAttrs, _ *[]Component)  { a.Headers = o.v }
func (o ScopeOpt) applyTh(a *ThAttrs, _ *[]Component)    { a.Scope = o.v }

func (a *ThAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Colspan > 1 {
		attr(sb, "colspan", itoa(a.Colspan))
	}
	if a.Rowspan > 1 {
		attr(sb, "rowspan", itoa(a.Rowspan))
	}
	if a.Headers != "" {
		attr(sb, "headers", a.Headers)
	}
	if a.Scope != "" {
		attr(sb, "scope", a.Scope)
	}
}
