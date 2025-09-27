package html

import "strings"

type FormAttrs struct {
	Global        GlobalAttrs
	AcceptCharset string
	Action        string
	Autocomplete  string
	Enctype       string
	Method        string
	Name          string
	Novalidate    bool
	Target        string
}

type FormArg interface {
	applyForm(*FormAttrs, *[]Component)
}

func defaultFormAttrs() *FormAttrs {
	return &FormAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Form(args ...FormArg) Node {
	a := defaultFormAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyForm(a, &kids)
	}
	return Node{Tag: "form", Attrs: a, Kids: kids}
}

func (g Global) applyForm(a *FormAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AcceptCharsetOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.AcceptCharset = o.v
}
func (o ActionOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.Action = o.v
}
func (o AutocompleteOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.Autocomplete = o.v
}
func (o EnctypeOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.Enctype = o.v
}
func (o MethodOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.Method = o.v
}
func (o NameOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o NovalidateOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.Novalidate = true
}
func (o TargetOpt) applyForm(a *FormAttrs, _ *[]Component) {
	a.Target = o.v
}

func (a *FormAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.AcceptCharset != "" {
		Attr(sb, "accept-charset", a.AcceptCharset)
	}
	if a.Action != "" {
		Attr(sb, "action", a.Action)
	}
	if a.Autocomplete != "" {
		Attr(sb, "autocomplete", a.Autocomplete)
	}
	if a.Enctype != "" {
		Attr(sb, "enctype", a.Enctype)
	}
	if a.Method != "" {
		Attr(sb, "method", a.Method)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Novalidate {
		BoolAttr(sb, "novalidate")
	}
	if a.Target != "" {
		Attr(sb, "target", a.Target)
	}
}
