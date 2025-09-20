package html

import "strings"

// Rect
type RectAttrs struct {
	Global           GlobalAttrs
	X                string
	Y                string
	Width            string
	Height           string
	Rx               string
	Ry               string
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

type RectArg interface {
	applyRect(*RectAttrs, *[]Component)
}

func defaultRectAttrs() *RectAttrs {
	return &RectAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Rect(args ...RectArg) Node {
	a := defaultRectAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyRect(a, &kids)
	}
	return Node{Tag: "rect", Attrs: a, Kids: kids}
}

// Rect-specific options
type RectXOpt struct{ v string }
type RectYOpt struct{ v string }
type RectWidthOpt struct{ v string }
type RectHeightOpt struct{ v string }
type RxOpt struct{ v string }
type RyOpt struct{ v string }

func X(v string) RectXOpt               { return RectXOpt{v} }
func Y(v string) RectYOpt               { return RectYOpt{v} }
func RectWidth(v string) RectWidthOpt   { return RectWidthOpt{v} }
func RectHeight(v string) RectHeightOpt { return RectHeightOpt{v} }
func Rx(v string) RxOpt                 { return RxOpt{v} }
func Ry(v string) RyOpt                 { return RyOpt{v} }

func (g Global) applyRect(a *RectAttrs, _ *[]Component)              { g.do(&a.Global) }
func (o TxtOpt) applyRect(_ *RectAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyRect(_ *RectAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o RectXOpt) applyRect(a *RectAttrs, _ *[]Component)            { a.X = o.v }
func (o RectYOpt) applyRect(a *RectAttrs, _ *[]Component)            { a.Y = o.v }
func (o RectWidthOpt) applyRect(a *RectAttrs, _ *[]Component)        { a.Width = o.v }
func (o RectHeightOpt) applyRect(a *RectAttrs, _ *[]Component)       { a.Height = o.v }
func (o RxOpt) applyRect(a *RectAttrs, _ *[]Component)               { a.Rx = o.v }
func (o RyOpt) applyRect(a *RectAttrs, _ *[]Component)               { a.Ry = o.v }
func (o PathLengthOpt) applyRect(a *RectAttrs, _ *[]Component)       { a.PathLength = o.v }
func (o FillOpt) applyRect(a *RectAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyRect(a *RectAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyRect(a *RectAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyRect(a *RectAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyRect(a *RectAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyRect(a *RectAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyRect(a *RectAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyRect(a *RectAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyRect(a *RectAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyRect(a *RectAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyRect(a *RectAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applyRect(a *RectAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applyRect(a *RectAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *RectAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.X != "" {
		attr(sb, "x", a.X)
	}
	if a.Y != "" {
		attr(sb, "y", a.Y)
	}
	if a.Width != "" {
		attr(sb, "width", a.Width)
	}
	if a.Height != "" {
		attr(sb, "height", a.Height)
	}
	if a.Rx != "" {
		attr(sb, "rx", a.Rx)
	}
	if a.Ry != "" {
		attr(sb, "ry", a.Ry)
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
