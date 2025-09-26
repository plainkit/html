package html

import "strings"

type FontAttrs struct {
	Global GlobalAttrs
	Color  string
	Face   string
	Size   string
}

type FontArg interface {
	applyFont(*FontAttrs, *[]Component)
}

func defaultFontAttrs() *FontAttrs {
	return &FontAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Font(args ...FontArg) Node {
	a := defaultFontAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFont(a, &kids)
	}
	return Node{Tag: "font", Attrs: a, Kids: kids}
}

func (g Global) applyFont(a *FontAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o ColorOpt) applyFont(a *FontAttrs, _ *[]Component) {
	a.Color = o.v
}
func (o FaceOpt) applyFont(a *FontAttrs, _ *[]Component) {
	a.Face = o.v
}
func (o SizeOpt) applyFont(a *FontAttrs, _ *[]Component) {
	a.Size = o.v
}

func (a *FontAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Color != "" {
		Attr(sb, "color", a.Color)
	}
	if a.Face != "" {
		Attr(sb, "face", a.Face)
	}
	if a.Size != "" {
		Attr(sb, "size", a.Size)
	}
}
