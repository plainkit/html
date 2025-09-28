package html

import "strings"

type TdAttrs struct {
	Global  GlobalAttrs
	Colspan string
	Headers string
	Rowspan string
}

type TdArg interface {
	ApplyTd(*TdAttrs, *[]Component)
}

func defaultTdAttrs() *TdAttrs {
	return &TdAttrs{
		Global: GlobalAttrs{
			Style:  "",
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
		ar.ApplyTd(a, &kids)
	}
	return Node{Tag: "td", Attrs: a, Kids: kids}
}

func (g Global) ApplyTd(a *TdAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o ColspanOpt) ApplyTd(a *TdAttrs, _ *[]Component) {
	a.Colspan = o.v
}
func (o HeadersOpt) ApplyTd(a *TdAttrs, _ *[]Component) {
	a.Headers = o.v
}
func (o RowspanOpt) ApplyTd(a *TdAttrs, _ *[]Component) {
	a.Rowspan = o.v
}

func (a *TdAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Colspan != "" {
		Attr(sb, "colspan", a.Colspan)
	}
	if a.Headers != "" {
		Attr(sb, "headers", a.Headers)
	}
	if a.Rowspan != "" {
		Attr(sb, "rowspan", a.Rowspan)
	}
}
