package html

import "strings"

type CheckboxAttrs struct {
	Global GlobalAttrs
}

type CheckboxArg interface {
	applyCheckbox(*CheckboxAttrs, *[]Component)
}

func defaultCheckboxAttrs() *CheckboxAttrs {
	return &CheckboxAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Checkbox(args ...CheckboxArg) Node {
	a := defaultCheckboxAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCheckbox(a, &kids)
	}
	return Node{Tag: "checkbox", Attrs: a, Kids: kids}
}

func (g Global) applyCheckbox(a *CheckboxAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyCheckbox(_ *CheckboxAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyCheckbox(_ *CheckboxAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *CheckboxAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
