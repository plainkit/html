package html

import "strings"

// Option
type OptionAttrs struct {
	Global   GlobalAttrs
	Value    string
	Selected bool
	Disabled bool
	Label    string
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

type OptionComponent Node

func (opt OptionComponent) render(sb *strings.Builder) {
	Node(opt).render(sb)
}

func Option(args ...OptionArg) OptionComponent {
	a := defaultOptionAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyOption(a, &kids)
	}
	return OptionComponent{Tag: "option", Attrs: a, Kids: kids}
}

// Option-specific options
type OptionValueOpt struct{ v string }
type SelectedOpt struct{}

func OptionValue(v string) OptionValueOpt { return OptionValueOpt{v} }
func Selected() SelectedOpt               { return SelectedOpt{} }

func (g Global) applyOption(a *OptionAttrs, _ *[]Component)         { g.do(&a.Global) }
func (o TxtOpt) applyOption(_ *OptionAttrs, kids *[]Component)      { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyOption(_ *OptionAttrs, kids *[]Component)    { *kids = append(*kids, o.c) }
func (o OptionValueOpt) applyOption(a *OptionAttrs, _ *[]Component) { a.Value = o.v }
func (o SelectedOpt) applyOption(a *OptionAttrs, _ *[]Component)    { a.Selected = true }
func (o DisabledOpt) applyOption(a *OptionAttrs, _ *[]Component)    { a.Disabled = true }
func (o LabelOpt) applyOption(a *OptionAttrs, _ *[]Component)       { a.Label = o.v }

// Compile-time type safety: Option can be added to Select and Optgroup
func (opt OptionComponent) applySelect(_ *SelectAttrs, kids *[]Component) {
	*kids = append(*kids, opt)
}

func (opt OptionComponent) applyOptgroup(_ *OptgroupAttrs, kids *[]Component) {
	*kids = append(*kids, opt)
}

func (a *OptionAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Value != "" {
		attr(sb, "value", a.Value)
	}
	if a.Selected {
		boolAttr(sb, "selected")
	}
	if a.Disabled {
		boolAttr(sb, "disabled")
	}
	if a.Label != "" {
		attr(sb, "label", a.Label)
	}
}
