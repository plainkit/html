package html

import "strings"

// Optgroup
type OptgroupAttrs struct {
	Global   GlobalAttrs
	Label    string
	Disabled bool
}

type OptgroupArg interface {
	applyOptgroup(*OptgroupAttrs, *[]Component)
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

type OptgroupComponent Node

func (optg OptgroupComponent) render(sb *strings.Builder) {
	Node(optg).render(sb)
}

func Optgroup(args ...OptgroupArg) OptgroupComponent {
	a := defaultOptgroupAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyOptgroup(a, &kids)
	}
	return OptgroupComponent{Tag: "optgroup", Attrs: a, Kids: kids}
}

func (g Global) applyOptgroup(a *OptgroupAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o LabelOpt) applyOptgroup(a *OptgroupAttrs, _ *[]Component)    { a.Label = o.v }
func (o DisabledOpt) applyOptgroup(a *OptgroupAttrs, _ *[]Component) { a.Disabled = true }

// Compile-time type safety: Optgroup can be added to Select
func (optg OptgroupComponent) applySelect(_ *SelectAttrs, kids *[]Component) {
	*kids = append(*kids, optg)
}

func (a *OptgroupAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Label != "" {
		attr(sb, "label", a.Label)
	}
	if a.Disabled {
		boolAttr(sb, "disabled")
	}
}
