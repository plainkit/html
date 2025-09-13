package blox

import "strings"

type FormAttrs struct {
	Global        GlobalAttrs
	Action        string
	Method        string
	Enctype       string
	AcceptCharset string
	Autocomplete  string
	Novalidate    bool
	Target        string
}

type FormArg interface {
	applyForm(*FormAttrs, *[]Component)
}

func defaultFormAttrs() *FormAttrs {
	return &FormAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
		Method: "get",
	}
}

func Form(args ...FormArg) Component {
	a := defaultFormAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyForm(a, &kids)
	}
	return Node{Tag: "form", Attrs: a, Kids: kids}
}

// Form-specific options
type ActionOpt struct{ v string }
type MethodOpt struct{ v string }
type EnctypeOpt struct{ v string }
type AcceptCharsetOpt struct{ v string }
type AutocompleteOpt struct{ v string }
type NovalidateOpt struct{}

func Action(v string) ActionOpt               { return ActionOpt{v} }
func Method(v string) MethodOpt               { return MethodOpt{v} }
func Enctype(v string) EnctypeOpt             { return EnctypeOpt{v} }
func AcceptCharset(v string) AcceptCharsetOpt { return AcceptCharsetOpt{v} }
func Autocomplete(v string) AutocompleteOpt   { return AutocompleteOpt{v} }
func Novalidate() NovalidateOpt               { return NovalidateOpt{} }

func (g Global) applyForm(a *FormAttrs, _ *[]Component)           { g.do(&a.Global) }
func (o TxtOpt) applyForm(_ *FormAttrs, kids *[]Component)        { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyForm(_ *FormAttrs, kids *[]Component)      { *kids = append(*kids, o.c) }
func (o ActionOpt) applyForm(a *FormAttrs, _ *[]Component)        { a.Action = o.v }
func (o MethodOpt) applyForm(a *FormAttrs, _ *[]Component)        { a.Method = o.v }
func (o EnctypeOpt) applyForm(a *FormAttrs, _ *[]Component)       { a.Enctype = o.v }
func (o AcceptCharsetOpt) applyForm(a *FormAttrs, _ *[]Component) { a.AcceptCharset = o.v }
func (o AutocompleteOpt) applyForm(a *FormAttrs, _ *[]Component)  { a.Autocomplete = o.v }
func (o NovalidateOpt) applyForm(a *FormAttrs, _ *[]Component)    { a.Novalidate = true }

func (a *FormAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Action != "" {
		attr(sb, "action", a.Action)
	}
	if a.Method != "" {
		attr(sb, "method", a.Method)
	}
	if a.Enctype != "" {
		attr(sb, "enctype", a.Enctype)
	}
	if a.AcceptCharset != "" {
		attr(sb, "accept-charset", a.AcceptCharset)
	}
	if a.Autocomplete != "" {
		attr(sb, "autocomplete", a.Autocomplete)
	}
	if a.Novalidate {
		boolAttr(sb, "novalidate")
	}
	if a.Target != "" {
		attr(sb, "target", a.Target)
	}
}
