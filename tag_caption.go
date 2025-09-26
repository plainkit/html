package html

import "strings"

type CaptionAttrs struct {
	Global GlobalAttrs
	Align  string
}

type CaptionArg interface {
	applyCaption(*CaptionAttrs, *[]Component)
}

func defaultCaptionAttrs() *CaptionAttrs {
	return &CaptionAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Caption(args ...CaptionArg) Node {
	a := defaultCaptionAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCaption(a, &kids)
	}
	return Node{Tag: "caption", Attrs: a, Kids: kids}
}

func (g Global) applyCaption(a *CaptionAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyCaption(a *CaptionAttrs, _ *[]Component) {
	a.Align = o.v
}

func (a *CaptionAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
}
