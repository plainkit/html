package html

import "strings"

type TextareaAttrs struct {
	Global       GlobalAttrs
	Autocomplete string
	Cols         string
	Dirname      string
	Disabled     bool
	Form         string
	Maxlength    string
	Minlength    string
	Name         string
	Placeholder  string
	Readonly     bool
	Required     bool
	Rows         string
	Wrap         string
}

type TextareaArg interface {
	applyTextarea(*TextareaAttrs, *[]Component)
}

func defaultTextareaAttrs() *TextareaAttrs {
	return &TextareaAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Textarea(args ...TextareaArg) Node {
	a := defaultTextareaAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTextarea(a, &kids)
	}
	return Node{Tag: "textarea", Attrs: a, Kids: kids}
}

func (g Global) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AutocompleteOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Autocomplete = o.v
}
func (o ColsOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Cols = o.v
}
func (o DirnameOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Dirname = o.v
}
func (o DisabledOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o MaxlengthOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Maxlength = o.v
}
func (o MinlengthOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Minlength = o.v
}
func (o NameOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o PlaceholderOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Placeholder = o.v
}
func (o ReadonlyOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Readonly = true
}
func (o RequiredOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Required = true
}
func (o RowsOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Rows = o.v
}
func (o WrapOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) {
	a.Wrap = o.v
}

func (a *TextareaAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Autocomplete != "" {
		Attr(sb, "autocomplete", a.Autocomplete)
	}
	if a.Cols != "" {
		Attr(sb, "cols", a.Cols)
	}
	if a.Dirname != "" {
		Attr(sb, "dirname", a.Dirname)
	}
	if a.Disabled {
		BoolAttr(sb, "disabled")
	}
	if a.Form != "" {
		Attr(sb, "form", a.Form)
	}
	if a.Maxlength != "" {
		Attr(sb, "maxlength", a.Maxlength)
	}
	if a.Minlength != "" {
		Attr(sb, "minlength", a.Minlength)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Placeholder != "" {
		Attr(sb, "placeholder", a.Placeholder)
	}
	if a.Readonly {
		BoolAttr(sb, "readonly")
	}
	if a.Required {
		BoolAttr(sb, "required")
	}
	if a.Rows != "" {
		Attr(sb, "rows", a.Rows)
	}
	if a.Wrap != "" {
		Attr(sb, "wrap", a.Wrap)
	}
}
