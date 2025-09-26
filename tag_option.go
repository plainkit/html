package html

import "strings"

type OptionAttrs struct {
	Global   GlobalAttrs
	Disabled bool
	Label    string
	Selected bool
	Value    string
}

type OptionArg interface {
	applyOption(*OptionAttrs, *[]Component)
}

func defaultOptionAttrs() *OptionAttrs {
	return &OptionAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Option(args ...OptionArg) Node {
	a := defaultOptionAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyOption(a, &kids)
	}
	return Node{Tag: "option", Attrs: a, Kids: kids}
}

func (g Global) applyOption(a *OptionAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o DisabledOpt) applyOption(a *OptionAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o LabelOpt) applyOption(a *OptionAttrs, _ *[]Component) {
	a.Label = o.v
}
func (o SelectedOpt) applyOption(a *OptionAttrs, _ *[]Component) {
	a.Selected = true
}
func (o ValueOpt) applyOption(a *OptionAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *OptionAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Disabled {
		BoolAttr(sb, "disabled")
	}
	if a.Label != "" {
		Attr(sb, "label", a.Label)
	}
	if a.Selected {
		BoolAttr(sb, "selected")
	}
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
}
