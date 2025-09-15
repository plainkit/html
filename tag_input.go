package html

import "strings"

type InputAttrs struct {
	Global         GlobalAttrs
	Type           string
	Name           string
	Value          string
	Placeholder    string
	Required       bool
	Disabled       bool
	Readonly       bool
	Multiple       bool
	Checked        bool
	Min            string
	Max            string
	Step           string
	Pattern        string
	Size           int
	Maxlength      int
	Minlength      int
	Accept         string
	Form           string
	Formaction     string
	Formenctype    string
	Formmethod     string
	Formnovalidate bool
	Formtarget     string
	List           string
}

type InputArg interface {
	applyInput(*InputAttrs)
}

func defaultInputAttrs() *InputAttrs {
	return &InputAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
		Type: "text",
	}
}

func Input(args ...InputArg) Node {
	a := defaultInputAttrs()
	for _, ar := range args {
		ar.applyInput(a)
	}
	return Node{Tag: "input", Attrs: a, Void: true}
}

// Input-specific options
type InputTypeOpt struct{ v string }
type InputNameOpt struct{ v string }
type InputValueOpt struct{ v string }
type PlaceholderOpt struct{ v string }
type RequiredOpt struct{}
type DisabledOpt struct{}
type ReadonlyOpt struct{}
type MultipleOpt struct{}
type CheckedOpt struct{}
type MinOpt struct{ v string }
type MaxOpt struct{ v string }
type StepOpt struct{ v string }
type PatternOpt struct{ v string }
type SizeOpt struct{ v int }
type MaxlengthOpt struct{ v int }
type MinlengthOpt struct{ v int }
type AcceptOpt struct{ v string }
type FormOpt struct{ v string }
type FormactionOpt struct{ v string }
type FormenctypeOpt struct{ v string }
type FormmethodOpt struct{ v string }
type FormnovalidateOpt struct{}
type FormtargetOpt struct{ v string }
type ListOpt struct{ v string }

func InputType(v string) InputTypeOpt     { return InputTypeOpt{v} }
func InputName(v string) InputNameOpt     { return InputNameOpt{v} }
func InputValue(v string) InputValueOpt   { return InputValueOpt{v} }
func Placeholder(v string) PlaceholderOpt { return PlaceholderOpt{v} }
func Required() RequiredOpt               { return RequiredOpt{} }
func Disabled() DisabledOpt               { return DisabledOpt{} }
func Readonly() ReadonlyOpt               { return ReadonlyOpt{} }
func Multiple() MultipleOpt               { return MultipleOpt{} }
func Checked() CheckedOpt                 { return CheckedOpt{} }
func Min(v string) MinOpt                 { return MinOpt{v} }
func Max(v string) MaxOpt                 { return MaxOpt{v} }
func Step(v string) StepOpt               { return StepOpt{v} }
func Pattern(v string) PatternOpt         { return PatternOpt{v} }
func Size(v int) SizeOpt                  { return SizeOpt{v} }
func Maxlength(v int) MaxlengthOpt        { return MaxlengthOpt{v} }
func Minlength(v int) MinlengthOpt        { return MinlengthOpt{v} }
func Accept(v string) AcceptOpt           { return AcceptOpt{v} }
func FormAttr(v string) FormOpt           { return FormOpt{v} }
func Formaction(v string) FormactionOpt   { return FormactionOpt{v} }
func Formenctype(v string) FormenctypeOpt { return FormenctypeOpt{v} }
func Formmethod(v string) FormmethodOpt   { return FormmethodOpt{v} }
func Formnovalidate() FormnovalidateOpt   { return FormnovalidateOpt{} }
func Formtarget(v string) FormtargetOpt   { return FormtargetOpt{v} }
func List(v string) ListOpt               { return ListOpt{v} }

func (g Global) applyInput(a *InputAttrs)            { g.do(&a.Global) }
func (o InputTypeOpt) applyInput(a *InputAttrs)      { a.Type = o.v }
func (o InputNameOpt) applyInput(a *InputAttrs)      { a.Name = o.v }
func (o InputValueOpt) applyInput(a *InputAttrs)     { a.Value = o.v }
func (o PlaceholderOpt) applyInput(a *InputAttrs)    { a.Placeholder = o.v }
func (o RequiredOpt) applyInput(a *InputAttrs)       { a.Required = true }
func (o DisabledOpt) applyInput(a *InputAttrs)       { a.Disabled = true }
func (o ReadonlyOpt) applyInput(a *InputAttrs)       { a.Readonly = true }
func (o MultipleOpt) applyInput(a *InputAttrs)       { a.Multiple = true }
func (o CheckedOpt) applyInput(a *InputAttrs)        { a.Checked = true }
func (o MinOpt) applyInput(a *InputAttrs)            { a.Min = o.v }
func (o MaxOpt) applyInput(a *InputAttrs)            { a.Max = o.v }
func (o StepOpt) applyInput(a *InputAttrs)           { a.Step = o.v }
func (o PatternOpt) applyInput(a *InputAttrs)        { a.Pattern = o.v }
func (o SizeOpt) applyInput(a *InputAttrs)           { a.Size = o.v }
func (o MaxlengthOpt) applyInput(a *InputAttrs)      { a.Maxlength = o.v }
func (o MinlengthOpt) applyInput(a *InputAttrs)      { a.Minlength = o.v }
func (o AcceptOpt) applyInput(a *InputAttrs)         { a.Accept = o.v }
func (o FormOpt) applyInput(a *InputAttrs)           { a.Form = o.v }
func (o FormactionOpt) applyInput(a *InputAttrs)     { a.Formaction = o.v }
func (o FormenctypeOpt) applyInput(a *InputAttrs)    { a.Formenctype = o.v }
func (o FormmethodOpt) applyInput(a *InputAttrs)     { a.Formmethod = o.v }
func (o FormnovalidateOpt) applyInput(a *InputAttrs) { a.Formnovalidate = true }
func (o FormtargetOpt) applyInput(a *InputAttrs)     { a.Formtarget = o.v }
func (o ListOpt) applyInput(a *InputAttrs)           { a.List = o.v }

func (a *InputAttrs) writeAttrs(sb *strings.Builder) {
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
	if a.Multiple {
		boolAttr(sb, "multiple")
	}
	if a.Checked {
		boolAttr(sb, "checked")
	}
	if a.Min != "" {
		attr(sb, "min", a.Min)
	}
	if a.Max != "" {
		attr(sb, "max", a.Max)
	}
	if a.Step != "" {
		attr(sb, "step", a.Step)
	}
	if a.Pattern != "" {
		attr(sb, "pattern", a.Pattern)
	}
	if a.Size > 0 {
		attr(sb, "size", itoa(a.Size))
	}
	if a.Maxlength > 0 {
		attr(sb, "maxlength", itoa(a.Maxlength))
	}
	if a.Minlength > 0 {
		attr(sb, "minlength", itoa(a.Minlength))
	}
	if a.Accept != "" {
		attr(sb, "accept", a.Accept)
	}
	if a.Form != "" {
		attr(sb, "form", a.Form)
	}
	if a.Formaction != "" {
		attr(sb, "formaction", a.Formaction)
	}
	if a.Formenctype != "" {
		attr(sb, "formenctype", a.Formenctype)
	}
	if a.Formmethod != "" {
		attr(sb, "formmethod", a.Formmethod)
	}
	if a.Formnovalidate {
		boolAttr(sb, "formnovalidate")
	}
	if a.Formtarget != "" {
		attr(sb, "formtarget", a.Formtarget)
	}
	if a.List != "" {
		attr(sb, "list", a.List)
	}
}
