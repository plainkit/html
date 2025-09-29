package html

import "strings"

type ThAttrs struct {
	Global  GlobalAttrs
	Abbr    string
	Colspan string
	Headers string
	Rowspan string
	Scope   string
}

type ThArg interface {
	ApplyTh(*ThAttrs, *[]Component)
}

func defaultThAttrs() *ThAttrs {
	return &ThAttrs{
		Global: GlobalAttrs{
			Style:  "",
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
		ar.ApplyTh(a, &kids)
	}

	return Node{Tag: "th", Attrs: a, Kids: kids}
}

func (g Global) ApplyTh(a *ThAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AbbrOpt) ApplyTh(a *ThAttrs, _ *[]Component) {
	a.Abbr = o.v
}
func (o ColspanOpt) ApplyTh(a *ThAttrs, _ *[]Component) {
	a.Colspan = o.v
}
func (o HeadersOpt) ApplyTh(a *ThAttrs, _ *[]Component) {
	a.Headers = o.v
}
func (o RowspanOpt) ApplyTh(a *ThAttrs, _ *[]Component) {
	a.Rowspan = o.v
}
func (o ScopeOpt) ApplyTh(a *ThAttrs, _ *[]Component) {
	a.Scope = o.v
}

func (a *ThAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Abbr != "" {
		Attr(sb, "abbr", a.Abbr)
	}

	if a.Colspan != "" {
		Attr(sb, "colspan", a.Colspan)
	}

	if a.Headers != "" {
		Attr(sb, "headers", a.Headers)
	}

	if a.Rowspan != "" {
		Attr(sb, "rowspan", a.Rowspan)
	}

	if a.Scope != "" {
		Attr(sb, "scope", a.Scope)
	}
}
