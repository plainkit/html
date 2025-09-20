package html

import "strings"

// Line
type LineAttrs struct {
	Global           GlobalAttrs
	X1               string
	Y1               string
	X2               string
	Y2               string
	PathLength       string
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

type LineArg interface {
	applyLine(*LineAttrs, *[]Component)
}

func defaultLineAttrs() *LineAttrs {
	return &LineAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Line(args ...LineArg) Node {
	a := defaultLineAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyLine(a, &kids)
	}
	return Node{Tag: "line", Attrs: a, Kids: kids}
}

// Line-specific options
type X1Opt struct{ v string }
type Y1Opt struct{ v string }
type X2Opt struct{ v string }
type Y2Opt struct{ v string }

func X1(v string) X1Opt { return X1Opt{v} }
func Y1(v string) Y1Opt { return Y1Opt{v} }
func X2(v string) X2Opt { return X2Opt{v} }
func Y2(v string) Y2Opt { return Y2Opt{v} }

func (g Global) applyLine(a *LineAttrs, _ *[]Component)              { g.do(&a.Global) }
func (o TxtOpt) applyLine(_ *LineAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyLine(_ *LineAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o X1Opt) applyLine(a *LineAttrs, _ *[]Component)               { a.X1 = o.v }
func (o Y1Opt) applyLine(a *LineAttrs, _ *[]Component)               { a.Y1 = o.v }
func (o X2Opt) applyLine(a *LineAttrs, _ *[]Component)               { a.X2 = o.v }
func (o Y2Opt) applyLine(a *LineAttrs, _ *[]Component)               { a.Y2 = o.v }
func (o PathLengthOpt) applyLine(a *LineAttrs, _ *[]Component)       { a.PathLength = o.v }
func (o FillOpt) applyLine(a *LineAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyLine(a *LineAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyLine(a *LineAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyLine(a *LineAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyLine(a *LineAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyLine(a *LineAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyLine(a *LineAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyLine(a *LineAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyLine(a *LineAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyLine(a *LineAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyLine(a *LineAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applyLine(a *LineAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applyLine(a *LineAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *LineAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.X1 != "" {
		attr(sb, "x1", a.X1)
	}
	if a.Y1 != "" {
		attr(sb, "y1", a.Y1)
	}
	if a.X2 != "" {
		attr(sb, "x2", a.X2)
	}
	if a.Y2 != "" {
		attr(sb, "y2", a.Y2)
	}
	if a.PathLength != "" {
		attr(sb, "pathLength", a.PathLength)
	}
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
