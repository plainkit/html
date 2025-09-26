package html

import "strings"

type ButtonAttrs struct {
	Global              GlobalAttrs
	Command             string
	Commandfor          string
	Disabled            bool
	Form                string
	Formaction          string
	Formenctype         string
	Formmethod          string
	Formnovalidate      bool
	Formtarget          string
	Name                string
	Popovertarget       string
	Popovertargetaction string
	Type                string
	Value               string
}

type ButtonArg interface {
	applyButton(*ButtonAttrs, *[]Component)
}

func defaultButtonAttrs() *ButtonAttrs {
	return &ButtonAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Button(args ...ButtonArg) Node {
	a := defaultButtonAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyButton(a, &kids)
	}
	return Node{Tag: "button", Attrs: a, Kids: kids}
}

func (g Global) applyButton(a *ButtonAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o CommandOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Command = o.v
}
func (o CommandforOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Commandfor = o.v
}
func (o DisabledOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o FormactionOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formaction = o.v
}
func (o FormenctypeOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formenctype = o.v
}
func (o FormmethodOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formmethod = o.v
}
func (o FormnovalidateOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formnovalidate = true
}
func (o FormtargetOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formtarget = o.v
}
func (o NameOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o PopovertargetOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Popovertarget = o.v
}
func (o PopovertargetactionOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Popovertargetaction = o.v
}
func (o TypeOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o ValueOpt) applyButton(a *ButtonAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *ButtonAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Command != "" {
		Attr(sb, "command", a.Command)
	}
	if a.Commandfor != "" {
		Attr(sb, "commandfor", a.Commandfor)
	}
	if a.Disabled {
		BoolAttr(sb, "disabled")
	}
	if a.Form != "" {
		Attr(sb, "form", a.Form)
	}
	if a.Formaction != "" {
		Attr(sb, "formaction", a.Formaction)
	}
	if a.Formenctype != "" {
		Attr(sb, "formenctype", a.Formenctype)
	}
	if a.Formmethod != "" {
		Attr(sb, "formmethod", a.Formmethod)
	}
	if a.Formnovalidate {
		BoolAttr(sb, "formnovalidate")
	}
	if a.Formtarget != "" {
		Attr(sb, "formtarget", a.Formtarget)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Popovertarget != "" {
		Attr(sb, "popovertarget", a.Popovertarget)
	}
	if a.Popovertargetaction != "" {
		Attr(sb, "popovertargetaction", a.Popovertargetaction)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
}
