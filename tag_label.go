package html

import "strings"

// Label
type LabelAttrs struct {
	Global GlobalAttrs
	For    string
	Form   string
}

type LabelArg interface {
	applyLabel(*LabelAttrs, *[]Component)
}

func defaultLabelAttrs() *LabelAttrs {
	return &LabelAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func FormLabel(args ...LabelArg) Node {
	a := defaultLabelAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyLabel(a, &kids)
	}
	return Node{Tag: "label", Attrs: a, Kids: kids}
}

// Label-specific options
type ForOpt struct{ v string }

func For(v string) ForOpt { return ForOpt{v} }

func (g Global) applyLabel(a *LabelAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyLabel(_ *LabelAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyLabel(_ *LabelAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ForOpt) applyLabel(a *LabelAttrs, _ *[]Component)      { a.For = o.v }
func (o FormOpt) applyLabel(a *LabelAttrs, _ *[]Component)     { a.Form = o.v }

func (a *LabelAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.For != "" {
		attr(sb, "for", a.For)
	}
	if a.Form != "" {
		attr(sb, "form", a.Form)
	}
}
