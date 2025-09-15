package html

import "strings"

// Circle
type CircleAttrs struct {
	Global           GlobalAttrs
	Cx               string
	Cy               string
	R                string
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

type CircleArg interface {
	applyCircle(*CircleAttrs, *[]Component)
}

func defaultCircleAttrs() *CircleAttrs {
	return &CircleAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Circle(args ...CircleArg) Node {
	a := defaultCircleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCircle(a, &kids)
	}
	return Node{Tag: "circle", Attrs: a, Kids: kids}
}

// Circle-specific options
type CxOpt struct{ v string }
type CyOpt struct{ v string }
type ROpt struct{ v string }
type PathLengthOpt struct{ v string }
type FillOpt struct{ v string }
type FillOpacityOpt struct{ v string }
type FillRuleOpt struct{ v string }
type StrokeOpt struct{ v string }
type StrokeWidthOpt struct{ v string }
type StrokeDasharrayOpt struct{ v string }
type StrokeDashoffsetOpt struct{ v string }
type StrokeLinecapOpt struct{ v string }
type StrokeLinejoinOpt struct{ v string }
type StrokeOpacityOpt struct{ v string }
type StrokeMiterlimitOpt struct{ v string }
type TransformOpt struct{ v string }
type OpacityOpt struct{ v string }

func Cx(v string) CxOpt                             { return CxOpt{v} }
func Cy(v string) CyOpt                             { return CyOpt{v} }
func R(v string) ROpt                               { return ROpt{v} }
func PathLength(v string) PathLengthOpt             { return PathLengthOpt{v} }
func Fill(v string) FillOpt                         { return FillOpt{v} }
func FillOpacity(v string) FillOpacityOpt           { return FillOpacityOpt{v} }
func FillRule(v string) FillRuleOpt                 { return FillRuleOpt{v} }
func Stroke(v string) StrokeOpt                     { return StrokeOpt{v} }
func StrokeWidth(v string) StrokeWidthOpt           { return StrokeWidthOpt{v} }
func StrokeDasharray(v string) StrokeDasharrayOpt   { return StrokeDasharrayOpt{v} }
func StrokeDashoffset(v string) StrokeDashoffsetOpt { return StrokeDashoffsetOpt{v} }
func StrokeLinecap(v string) StrokeLinecapOpt       { return StrokeLinecapOpt{v} }
func StrokeLinejoin(v string) StrokeLinejoinOpt     { return StrokeLinejoinOpt{v} }
func StrokeOpacity(v string) StrokeOpacityOpt       { return StrokeOpacityOpt{v} }
func StrokeMiterlimit(v string) StrokeMiterlimitOpt { return StrokeMiterlimitOpt{v} }
func Transform(v string) TransformOpt               { return TransformOpt{v} }
func Opacity(v string) OpacityOpt                   { return OpacityOpt{v} }

func (g Global) applyCircle(a *CircleAttrs, _ *[]Component)              { g.do(&a.Global) }
func (o TxtOpt) applyCircle(_ *CircleAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyCircle(_ *CircleAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o CxOpt) applyCircle(a *CircleAttrs, _ *[]Component)               { a.Cx = o.v }
func (o CyOpt) applyCircle(a *CircleAttrs, _ *[]Component)               { a.Cy = o.v }
func (o ROpt) applyCircle(a *CircleAttrs, _ *[]Component)                { a.R = o.v }
func (o PathLengthOpt) applyCircle(a *CircleAttrs, _ *[]Component)       { a.PathLength = o.v }
func (o FillOpt) applyCircle(a *CircleAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyCircle(a *CircleAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyCircle(a *CircleAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyCircle(a *CircleAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyCircle(a *CircleAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyCircle(a *CircleAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyCircle(a *CircleAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyCircle(a *CircleAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyCircle(a *CircleAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyCircle(a *CircleAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyCircle(a *CircleAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applyCircle(a *CircleAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applyCircle(a *CircleAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *CircleAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Cx != "" {
		attr(sb, "cx", a.Cx)
	}
	if a.Cy != "" {
		attr(sb, "cy", a.Cy)
	}
	if a.R != "" {
		attr(sb, "r", a.R)
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
