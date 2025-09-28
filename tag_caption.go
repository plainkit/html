package html

import "strings"

type CaptionAttrs struct {
	Global GlobalAttrs
}

type CaptionArg interface {
	ApplyCaption(*CaptionAttrs, *[]Component)
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
		ar.ApplyCaption(a, &kids)
	}
	return Node{Tag: "caption", Attrs: a, Kids: kids}
}

func (g Global) ApplyCaption(a *CaptionAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *CaptionAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
