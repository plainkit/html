package blox

import "strings"

// Td
type TdAttrs struct {
	Global  GlobalAttrs
	Colspan int
	Rowspan int
	Headers string
}

type TdArg interface {
	applyTd(*TdAttrs, *[]Component)
}

func defaultTdAttrs() *TdAttrs {
	return &TdAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Td(args ...TdArg) Node {
	a := defaultTdAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTd(a, &kids)
	}
	return Node{Tag: "td", Attrs: a, Kids: kids}
}

// Td-specific options
type ColspanOpt struct{ v int }
type RowspanOpt struct{ v int }
type HeadersOpt struct{ v string }

func Colspan(v int) ColspanOpt    { return ColspanOpt{v} }
func Rowspan(v int) RowspanOpt    { return RowspanOpt{v} }
func Headers(v string) HeadersOpt { return HeadersOpt{v} }

func (g Global) applyTd(a *TdAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyTd(_ *TdAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyTd(_ *TdAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ColspanOpt) applyTd(a *TdAttrs, _ *[]Component)  { a.Colspan = o.v }
func (o RowspanOpt) applyTd(a *TdAttrs, _ *[]Component)  { a.Rowspan = o.v }
func (o HeadersOpt) applyTd(a *TdAttrs, _ *[]Component)  { a.Headers = o.v }

func (a *TdAttrs) writeAttrs(sb *strings.Builder) {
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
}
