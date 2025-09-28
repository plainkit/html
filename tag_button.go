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
	ApplyButton(*ButtonAttrs, *[]Component)
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
		ar.ApplyButton(a, &kids)
	}
	return Node{Tag: "button", Attrs: a, Kids: kids}
}

func (g Global) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o CommandOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Command = o.v
}
func (o CommandforOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Commandfor = o.v
}
func (o DisabledOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o FormactionOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formaction = o.v
}
func (o FormenctypeOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formenctype = o.v
}
func (o FormmethodOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formmethod = o.v
}
func (o FormnovalidateOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formnovalidate = true
}
func (o FormtargetOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Formtarget = o.v
}
func (o NameOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o PopovertargetOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Popovertarget = o.v
}
func (o PopovertargetactionOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Popovertargetaction = o.v
}
func (o TypeOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o ValueOpt) ApplyButton(a *ButtonAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *ButtonAttrs) WriteAttrs(sb *strings.Builder) {
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
