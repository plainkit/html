package blox

import "strings"

// Textarea
type TextareaAttrs struct {
	Global      GlobalAttrs
	Name        string
	Rows        int
	Cols        int
	Placeholder string
	Required    bool
	Disabled    bool
	Readonly    bool
	Maxlength   int
	Minlength   int
	Wrap        string
	Form        string
}

type TextareaArg interface {
	applyTextarea(*TextareaAttrs, *[]Component)
}

func defaultTextareaAttrs() *TextareaAttrs {
	return &TextareaAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
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

// Textarea-specific options
type TextareaNameOpt struct{ v string }
type RowsOpt struct{ v int }
type ColsOpt struct{ v int }
type WrapOpt struct{ v string }

func TextareaName(v string) TextareaNameOpt { return TextareaNameOpt{v} }
func Rows(v int) RowsOpt                    { return RowsOpt{v} }
func Cols(v int) ColsOpt                    { return ColsOpt{v} }
func Wrap(v string) WrapOpt                 { return WrapOpt{v} }

func (g Global) applyTextarea(a *TextareaAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyTextarea(_ *TextareaAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyTextarea(_ *TextareaAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o TextareaNameOpt) applyTextarea(a *TextareaAttrs, _ *[]Component) { a.Name = o.v }
func (o RowsOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)         { a.Rows = o.v }
func (o ColsOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)         { a.Cols = o.v }
func (o PlaceholderOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)  { a.Placeholder = o.v }
func (o RequiredOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)     { a.Required = true }
func (o DisabledOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)     { a.Disabled = true }
func (o ReadonlyOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)     { a.Readonly = true }
func (o MaxlengthOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)    { a.Maxlength = o.v }
func (o MinlengthOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)    { a.Minlength = o.v }
func (o WrapOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)         { a.Wrap = o.v }
func (o FormOpt) applyTextarea(a *TextareaAttrs, _ *[]Component)         { a.Form = o.v }

func (a *TextareaAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Name != "" {
		attr(sb, "name", a.Name)
	}
	if a.Rows > 0 {
		attr(sb, "rows", itoa(a.Rows))
	}
	if a.Cols > 0 {
		attr(sb, "cols", itoa(a.Cols))
	}
	if a.Placeholder != "" {
		attr(sb, "placeholder", a.Placeholder)
	}
	if a.Required {
		boolAttr(sb, "required")
	}
	if a.Disabled {
		boolAttr(sb, "disabled")
	}
	if a.Readonly {
		boolAttr(sb, "readonly")
	}
	if a.Maxlength > 0 {
		attr(sb, "maxlength", itoa(a.Maxlength))
	}
	if a.Minlength > 0 {
		attr(sb, "minlength", itoa(a.Minlength))
	}
	if a.Wrap != "" {
		attr(sb, "wrap", a.Wrap)
	}
	if a.Form != "" {
		attr(sb, "form", a.Form)
	}
}
