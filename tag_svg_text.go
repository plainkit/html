package html

import "strings"

// SvgText
type SvgTextAttrs struct {
	Global           GlobalAttrs
	X                string
	Y                string
	Dx               string
	Dy               string
	Rotate           string
	TextLength       string
	LengthAdjust     string
	DominantBaseline string
	TextAnchor       string
	FontFamily       string
	FontSize         string
	FontWeight       string
	FontStyle        string
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

type SvgTextArg interface {
	applySvgText(*SvgTextAttrs, *[]Component)
}

func defaultSvgTextAttrs() *SvgTextAttrs {
	return &SvgTextAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func SvgText(args ...SvgTextArg) Node {
	a := defaultSvgTextAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySvgText(a, &kids)
	}
	return Node{Tag: "text", Attrs: a, Kids: kids}
}

// SvgText-specific options
type SvgTextXOpt struct{ v string }
type SvgTextYOpt struct{ v string }
type DxOpt struct{ v string }
type DyOpt struct{ v string }
type RotateOpt struct{ v string }
type TextLengthOpt struct{ v string }
type LengthAdjustOpt struct{ v string }
type DominantBaselineOpt struct{ v string }
type TextAnchorOpt struct{ v string }
type FontFamilyOpt struct{ v string }
type FontSizeOpt struct{ v string }
type FontWeightOpt struct{ v string }
type FontStyleOpt struct{ v string }

func SvgTextX(v string) SvgTextXOpt                 { return SvgTextXOpt{v} }
func SvgTextY(v string) SvgTextYOpt                 { return SvgTextYOpt{v} }
func Dx(v string) DxOpt                             { return DxOpt{v} }
func Dy(v string) DyOpt                             { return DyOpt{v} }
func Rotate(v string) RotateOpt                     { return RotateOpt{v} }
func TextLength(v string) TextLengthOpt             { return TextLengthOpt{v} }
func LengthAdjust(v string) LengthAdjustOpt         { return LengthAdjustOpt{v} }
func DominantBaseline(v string) DominantBaselineOpt { return DominantBaselineOpt{v} }
func TextAnchor(v string) TextAnchorOpt             { return TextAnchorOpt{v} }
func FontFamily(v string) FontFamilyOpt             { return FontFamilyOpt{v} }
func FontSize(v string) FontSizeOpt                 { return FontSizeOpt{v} }
func FontWeight(v string) FontWeightOpt             { return FontWeightOpt{v} }
func FontStyle(v string) FontStyleOpt               { return FontStyleOpt{v} }

func (g Global) applySvgText(a *SvgTextAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applySvgText(_ *SvgTextAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applySvgText(_ *SvgTextAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o SvgTextXOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)         { a.X = o.v }
func (o SvgTextYOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)         { a.Y = o.v }
func (o DxOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)               { a.Dx = o.v }
func (o DyOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)               { a.Dy = o.v }
func (o RotateOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)           { a.Rotate = o.v }
func (o TextLengthOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)       { a.TextLength = o.v }
func (o LengthAdjustOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)     { a.LengthAdjust = o.v }
func (o DominantBaselineOpt) applySvgText(a *SvgTextAttrs, _ *[]Component) { a.DominantBaseline = o.v }
func (o TextAnchorOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)       { a.TextAnchor = o.v }
func (o FontFamilyOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)       { a.FontFamily = o.v }
func (o FontSizeOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)         { a.FontSize = o.v }
func (o FontWeightOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)       { a.FontWeight = o.v }
func (o FontStyleOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)        { a.FontStyle = o.v }
func (o FillOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applySvgText(a *SvgTextAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applySvgText(a *SvgTextAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applySvgText(a *SvgTextAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *SvgTextAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.X != "" {
		attr(sb, "x", a.X)
	}
	if a.Y != "" {
		attr(sb, "y", a.Y)
	}
	if a.Dx != "" {
		attr(sb, "dx", a.Dx)
	}
	if a.Dy != "" {
		attr(sb, "dy", a.Dy)
	}
	if a.Rotate != "" {
		attr(sb, "rotate", a.Rotate)
	}
	if a.TextLength != "" {
		attr(sb, "textLength", a.TextLength)
	}
	if a.LengthAdjust != "" {
		attr(sb, "lengthAdjust", a.LengthAdjust)
	}
	if a.DominantBaseline != "" {
		attr(sb, "dominant-baseline", a.DominantBaseline)
	}
	if a.TextAnchor != "" {
		attr(sb, "text-anchor", a.TextAnchor)
	}
	if a.FontFamily != "" {
		attr(sb, "font-family", a.FontFamily)
	}
	if a.FontSize != "" {
		attr(sb, "font-size", a.FontSize)
	}
	if a.FontWeight != "" {
		attr(sb, "font-weight", a.FontWeight)
	}
	if a.FontStyle != "" {
		attr(sb, "font-style", a.FontStyle)
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
