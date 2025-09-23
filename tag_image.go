package html

import "strings"

type ImageAttrs struct {
	Global GlobalAttrs
}

type ImageArg interface {
	applyImage(*ImageAttrs, *[]Component)
}

func defaultImageAttrs() *ImageAttrs {
	return &ImageAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Image(args ...ImageArg) Node {
	a := defaultImageAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyImage(a, &kids)
	}
	return Node{Tag: "image", Attrs: a, Kids: kids}
}

func (g Global) applyImage(a *ImageAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyImage(_ *ImageAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyImage(_ *ImageAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *ImageAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
