package html

import "strings"

type BasefontAttrs struct {
	Global GlobalAttrs
	Color  string
	Face   string
	Size   string
}

type BasefontArg interface {
	applyBasefont(*BasefontAttrs, *[]Component)
}

func defaultBasefontAttrs() *BasefontAttrs {
	return &BasefontAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Basefont(args ...BasefontArg) Node {
	a := defaultBasefontAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyBasefont(a, &kids)
	}
	return Node{Tag: "basefont", Attrs: a, Kids: kids}
}

func (g Global) applyBasefont(a *BasefontAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o ColorOpt) applyBasefont(a *BasefontAttrs, _ *[]Component) {
	a.Color = o.v
}
func (o FaceOpt) applyBasefont(a *BasefontAttrs, _ *[]Component) {
	a.Face = o.v
}
func (o SizeOpt) applyBasefont(a *BasefontAttrs, _ *[]Component) {
	a.Size = o.v
}

func (a *BasefontAttrs) writeAttrs(sb *strings.Builder) {
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
