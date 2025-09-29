package html

import "strings"

type OptgroupAttrs struct {
	Global   GlobalAttrs
	Disabled bool
	Label    string
}

type OptgroupArg interface {
	ApplyOptgroup(*OptgroupAttrs, *[]Component)
}

func defaultOptgroupAttrs() *OptgroupAttrs {
	return &OptgroupAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Optgroup(args ...OptgroupArg) Node {
	a := defaultOptgroupAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyOptgroup(a, &kids)
	}

	return Node{Tag: "optgroup", Attrs: a, Kids: kids}
}

func (g Global) ApplyOptgroup(a *OptgroupAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o DisabledOpt) ApplyOptgroup(a *OptgroupAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o LabelOpt) ApplyOptgroup(a *OptgroupAttrs, _ *[]Component) {
	a.Label = o.v
}

func (a *OptgroupAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Disabled {
		BoolAttr(sb, "disabled")
	}

	if a.Label != "" {
		Attr(sb, "label", a.Label)
	}
}
