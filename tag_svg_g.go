package html

import "strings"

// G (Group)
type GAttrs struct {
	Global           GlobalAttrs
	Fill             string
	FillOpacity      string
	FillRule         string
	Stroke           string
	StrokeWidth      string
	StrokeDasharray  string
	StrokeDashoffset string
	StrokeLinecap    string
	StrokeLinejoin   string
	StrokeOpacity    string
	StrokeMiterlimit string
	Transform        string
	Opacity          string
}

type GArg interface {
	applyG(*GAttrs, *[]Component)
}

func defaultGAttrs() *GAttrs {
	return &GAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func G(args ...GArg) Node {
	a := defaultGAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyG(a, &kids)
	}
	return Node{Tag: "g", Attrs: a, Kids: kids}
}

func (g Global) applyG(a *GAttrs, _ *[]Component)              { g.do(&a.Global) }
func (o TxtOpt) applyG(_ *GAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyG(_ *GAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o FillOpt) applyG(a *GAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyG(a *GAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyG(a *GAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyG(a *GAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyG(a *GAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyG(a *GAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyG(a *GAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyG(a *GAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyG(a *GAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyG(a *GAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyG(a *GAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applyG(a *GAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applyG(a *GAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *GAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Fill != "" {
		attr(sb, "fill", a.Fill)
	}
	if a.FillOpacity != "" {
		attr(sb, "fill-opacity", a.FillOpacity)
	}
	if a.FillRule != "" {
		attr(sb, "fill-rule", a.FillRule)
	}
	if a.Stroke != "" {
		attr(sb, "stroke", a.Stroke)
	}
	if a.StrokeWidth != "" {
		attr(sb, "stroke-width", a.StrokeWidth)
	}
	if a.StrokeDasharray != "" {
		attr(sb, "stroke-dasharray", a.StrokeDasharray)
	}
	if a.StrokeDashoffset != "" {
		attr(sb, "stroke-dashoffset", a.StrokeDashoffset)
	}
	if a.StrokeLinecap != "" {
		attr(sb, "stroke-linecap", a.StrokeLinecap)
	}
	if a.StrokeLinejoin != "" {
		attr(sb, "stroke-linejoin", a.StrokeLinejoin)
	}
	if a.StrokeOpacity != "" {
		attr(sb, "stroke-opacity", a.StrokeOpacity)
	}
	if a.StrokeMiterlimit != "" {
		attr(sb, "stroke-miterlimit", a.StrokeMiterlimit)
	}
	if a.Transform != "" {
		attr(sb, "transform", a.Transform)
	}
	if a.Opacity != "" {
		attr(sb, "opacity", a.Opacity)
	}
}
