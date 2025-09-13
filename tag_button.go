package blox

import "strings"

type ButtonAttrs struct {
	Global         GlobalAttrs
	Type           string
	Name           string
	Value          string
	Disabled       bool
	Form           string
	Formaction     string
	Formenctype    string
	Formmethod     string
	Formnovalidate bool
	Formtarget     string
}

type ButtonArg interface {
	applyButton(*ButtonAttrs, *[]Component)
}

func defaultButtonAttrs() *ButtonAttrs {
	return &ButtonAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
		Type: "button",
	}
}

func Button(args ...ButtonArg) Component {
	a := defaultButtonAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyButton(a, &kids)
	}
	return Node{Tag: "button", Attrs: a, Kids: kids}
}

// Button-specific options
type ButtonTypeOpt struct{ v string }
type ButtonNameOpt struct{ v string }
type ButtonValueOpt struct{ v string }

func ButtonType(v string) ButtonTypeOpt   { return ButtonTypeOpt{v} }
func ButtonName(v string) ButtonNameOpt   { return ButtonNameOpt{v} }
func ButtonValue(v string) ButtonValueOpt { return ButtonValueOpt{v} }

func (g Global) applyButton(a *ButtonAttrs, _ *[]Component)         { g.do(&a.Global) }
func (o TxtOpt) applyButton(_ *ButtonAttrs, kids *[]Component)      { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyButton(_ *ButtonAttrs, kids *[]Component)    { *kids = append(*kids, o.c) }
func (o ButtonTypeOpt) applyButton(a *ButtonAttrs, _ *[]Component)  { a.Type = o.v }
func (o ButtonNameOpt) applyButton(a *ButtonAttrs, _ *[]Component)  { a.Name = o.v }
func (o ButtonValueOpt) applyButton(a *ButtonAttrs, _ *[]Component) { a.Value = o.v }
func (o DisabledOpt) applyButton(a *ButtonAttrs, _ *[]Component)    { a.Disabled = true }

func (a *ButtonAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Name != "" {
		attr(sb, "name", a.Name)
	}
	if a.Value != "" {
		attr(sb, "value", a.Value)
	}
	if a.Disabled {
		boolAttr(sb, "disabled")
	}
	if a.Form != "" {
		attr(sb, "form", a.Form)
	}
}
