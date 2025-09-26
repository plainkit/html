package html

import "strings"

type HrAttrs struct {
	Global  GlobalAttrs
	Align   string
	Noshade string
	Size    string
	Width   string
}

type HrArg interface {
	applyHr(*HrAttrs, *[]Component)
}

func defaultHrAttrs() *HrAttrs {
	return &HrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Hr(args ...HrArg) Node {
	a := defaultHrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHr(a, &kids)
	}
	return Node{Tag: "hr", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyHr(a *HrAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyHr(a *HrAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o NoshadeOpt) applyHr(a *HrAttrs, _ *[]Component) {
	a.Noshade = o.v
}
func (o SizeOpt) applyHr(a *HrAttrs, _ *[]Component) {
	a.Size = o.v
}
func (o WidthOpt) applyHr(a *HrAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *HrAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Noshade != "" {
		Attr(sb, "noshade", a.Noshade)
	}
	if a.Size != "" {
		Attr(sb, "size", a.Size)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
