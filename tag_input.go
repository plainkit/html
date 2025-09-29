package html

import "strings"

type InputAttrs struct {
	Global              GlobalAttrs
	Accept              string
	Alpha               bool
	Alt                 string
	Autocomplete        string
	Checked             bool
	Colorspace          string
	Dirname             string
	Disabled            bool
	Form                string
	Formaction          string
	Formenctype         string
	Formmethod          string
	Formnovalidate      bool
	Formtarget          string
	Height              string
	List                string
	Max                 string
	Maxlength           string
	Min                 string
	Minlength           string
	Multiple            bool
	Name                string
	Pattern             string
	Placeholder         string
	Popovertarget       string
	Popovertargetaction string
	Readonly            bool
	Required            bool
	Size                string
	Src                 string
	Step                string
	Type                string
	Value               string
	Width               string
}

type InputArg interface {
	ApplyInput(*InputAttrs, *[]Component)
}

func defaultInputAttrs() *InputAttrs {
	return &InputAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Input(args ...InputArg) Node {
	a := defaultInputAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyInput(a, &kids)
	}

	return Node{Tag: "input", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyInput(a *InputAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AcceptOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Accept = o.v
}
func (o AlphaOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Alpha = true
}
func (o AltOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Alt = o.v
}
func (o AutocompleteOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Autocomplete = o.v
}
func (o CheckedOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Checked = true
}
func (o ColorspaceOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Colorspace = o.v
}
func (o DirnameOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Dirname = o.v
}
func (o DisabledOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o FormactionOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Formaction = o.v
}
func (o FormenctypeOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Formenctype = o.v
}
func (o FormmethodOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Formmethod = o.v
}
func (o FormnovalidateOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Formnovalidate = true
}
func (o FormtargetOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Formtarget = o.v
}
func (o HeightOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o ListOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.List = o.v
}
func (o MaxOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Max = o.v
}
func (o MaxlengthOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Maxlength = o.v
}
func (o MinOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Min = o.v
}
func (o MinlengthOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Minlength = o.v
}
func (o MultipleOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Multiple = true
}
func (o NameOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o PatternOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Pattern = o.v
}
func (o PlaceholderOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Placeholder = o.v
}
func (o PopovertargetOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Popovertarget = o.v
}
func (o PopovertargetactionOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Popovertargetaction = o.v
}
func (o ReadonlyOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Readonly = true
}
func (o RequiredOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Required = true
}
func (o SizeOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Size = o.v
}
func (o SrcOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o StepOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Step = o.v
}
func (o TypeOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o ValueOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Value = o.v
}
func (o WidthOpt) ApplyInput(a *InputAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *InputAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Accept != "" {
		Attr(sb, "accept", a.Accept)
	}

	if a.Alpha {
		BoolAttr(sb, "alpha")
	}

	if a.Alt != "" {
		Attr(sb, "alt", a.Alt)
	}

	if a.Autocomplete != "" {
		Attr(sb, "autocomplete", a.Autocomplete)
	}

	if a.Checked {
		BoolAttr(sb, "checked")
	}

	if a.Colorspace != "" {
		Attr(sb, "colorspace", a.Colorspace)
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

	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}

	if a.List != "" {
		Attr(sb, "list", a.List)
	}

	if a.Max != "" {
		Attr(sb, "max", a.Max)
	}

	if a.Maxlength != "" {
		Attr(sb, "maxlength", a.Maxlength)
	}

	if a.Min != "" {
		Attr(sb, "min", a.Min)
	}

	if a.Minlength != "" {
		Attr(sb, "minlength", a.Minlength)
	}

	if a.Multiple {
		BoolAttr(sb, "multiple")
	}

	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}

	if a.Pattern != "" {
		Attr(sb, "pattern", a.Pattern)
	}

	if a.Placeholder != "" {
		Attr(sb, "placeholder", a.Placeholder)
	}

	if a.Popovertarget != "" {
		Attr(sb, "popovertarget", a.Popovertarget)
	}

	if a.Popovertargetaction != "" {
		Attr(sb, "popovertargetaction", a.Popovertargetaction)
	}

	if a.Readonly {
		BoolAttr(sb, "readonly")
	}

	if a.Required {
		BoolAttr(sb, "required")
	}

	if a.Size != "" {
		Attr(sb, "size", a.Size)
	}

	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}

	if a.Step != "" {
		Attr(sb, "step", a.Step)
	}

	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}

	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}

	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
