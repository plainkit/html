package blox

import "strings"

// Svg
type SvgAttrs struct {
	Global              GlobalAttrs
	Width               string
	Height              string
	ViewBox             string
	PreserveAspectRatio string
	Xmlns               string
	Version             string
	BaseProfile         string
	ContentScriptType   string
	ContentStyleType    string
	ZoomAndPan          string
	X                   string
	Y                   string
	Fill                string
	FillOpacity         string
	FillRule            string
	Stroke              string
	StrokeWidth         string
	StrokeDasharray     string
	StrokeDashoffset    string
	StrokeLinecap       string
	StrokeLinejoin      string
	StrokeOpacity       string
	StrokeMiterlimit    string
	Transform           string
	Opacity             string
}

type SvgArg interface {
	applySvg(*SvgAttrs, *[]Component)
}

func defaultSvgAttrs() *SvgAttrs {
	return &SvgAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Svg(args ...SvgArg) Node {
	a := defaultSvgAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySvg(a, &kids)
	}
	return Node{Tag: "svg", Attrs: a, Kids: kids}
}

// SVG-specific options
type ViewBoxOpt struct{ v string }
type BaseProfileOpt struct{ v string }
type SvgWidthOpt struct{ v string }
type SvgHeightOpt struct{ v string }
type PreserveAspectRatioOpt struct{ v string }
type ContentScriptTypeOpt struct{ v string }
type ContentStyleTypeOpt struct{ v string }
type ZoomAndPanOpt struct{ v string }
type SvgXOpt struct{ v string }
type SvgYOpt struct{ v string }

func ViewBox(v string) ViewBoxOpt                         { return ViewBoxOpt{v} }
func BaseProfile(v string) BaseProfileOpt                 { return BaseProfileOpt{v} }
func SvgWidth(v string) SvgWidthOpt                       { return SvgWidthOpt{v} }
func SvgHeight(v string) SvgHeightOpt                     { return SvgHeightOpt{v} }
func PreserveAspectRatio(v string) PreserveAspectRatioOpt { return PreserveAspectRatioOpt{v} }
func ContentScriptType(v string) ContentScriptTypeOpt     { return ContentScriptTypeOpt{v} }
func ContentStyleType(v string) ContentStyleTypeOpt       { return ContentStyleTypeOpt{v} }
func ZoomAndPan(v string) ZoomAndPanOpt                   { return ZoomAndPanOpt{v} }
func SvgX(v string) SvgXOpt                               { return SvgXOpt{v} }
func SvgY(v string) SvgYOpt                               { return SvgYOpt{v} }

func (g Global) applySvg(a *SvgAttrs, _ *[]Component)                 { g.do(&a.Global) }
func (o TxtOpt) applySvg(_ *SvgAttrs, kids *[]Component)              { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applySvg(_ *SvgAttrs, kids *[]Component)            { *kids = append(*kids, o.c) }
func (o ViewBoxOpt) applySvg(a *SvgAttrs, _ *[]Component)             { a.ViewBox = o.v }
func (o XmlnsOpt) applySvg(a *SvgAttrs, _ *[]Component)               { a.Xmlns = o.v }
func (o VersionOpt) applySvg(a *SvgAttrs, _ *[]Component)             { a.Version = o.v }
func (o BaseProfileOpt) applySvg(a *SvgAttrs, _ *[]Component)         { a.BaseProfile = o.v }
func (o SvgWidthOpt) applySvg(a *SvgAttrs, _ *[]Component)            { a.Width = o.v }
func (o SvgHeightOpt) applySvg(a *SvgAttrs, _ *[]Component)           { a.Height = o.v }
func (o PreserveAspectRatioOpt) applySvg(a *SvgAttrs, _ *[]Component) { a.PreserveAspectRatio = o.v }
func (o ContentScriptTypeOpt) applySvg(a *SvgAttrs, _ *[]Component)   { a.ContentScriptType = o.v }
func (o ContentStyleTypeOpt) applySvg(a *SvgAttrs, _ *[]Component)    { a.ContentStyleType = o.v }
func (o ZoomAndPanOpt) applySvg(a *SvgAttrs, _ *[]Component)          { a.ZoomAndPan = o.v }
func (o SvgXOpt) applySvg(a *SvgAttrs, _ *[]Component)                { a.X = o.v }
func (o SvgYOpt) applySvg(a *SvgAttrs, _ *[]Component)                { a.Y = o.v }
func (o FillOpt) applySvg(a *SvgAttrs, _ *[]Component)                { a.Fill = o.v }
func (o FillOpacityOpt) applySvg(a *SvgAttrs, _ *[]Component)         { a.FillOpacity = o.v }
func (o FillRuleOpt) applySvg(a *SvgAttrs, _ *[]Component)            { a.FillRule = o.v }
func (o StrokeOpt) applySvg(a *SvgAttrs, _ *[]Component)              { a.Stroke = o.v }
func (o StrokeWidthOpt) applySvg(a *SvgAttrs, _ *[]Component)         { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applySvg(a *SvgAttrs, _ *[]Component)     { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applySvg(a *SvgAttrs, _ *[]Component)    { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applySvg(a *SvgAttrs, _ *[]Component)       { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applySvg(a *SvgAttrs, _ *[]Component)      { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applySvg(a *SvgAttrs, _ *[]Component)       { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applySvg(a *SvgAttrs, _ *[]Component)    { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applySvg(a *SvgAttrs, _ *[]Component)           { a.Transform = o.v }
func (o OpacityOpt) applySvg(a *SvgAttrs, _ *[]Component)             { a.Opacity = o.v }

func (a *SvgAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Width != "" {
		attr(sb, "width", a.Width)
	}
	if a.Height != "" {
		attr(sb, "height", a.Height)
	}
	if a.ViewBox != "" {
		attr(sb, "viewBox", a.ViewBox)
	}
	if a.PreserveAspectRatio != "" {
		attr(sb, "preserveAspectRatio", a.PreserveAspectRatio)
	}
	if a.Xmlns != "" {
		attr(sb, "xmlns", a.Xmlns)
	}
	if a.Version != "" {
		attr(sb, "version", a.Version)
	}
	if a.BaseProfile != "" {
		attr(sb, "baseProfile", a.BaseProfile)
	}
	if a.ContentScriptType != "" {
		attr(sb, "contentScriptType", a.ContentScriptType)
	}
	if a.ContentStyleType != "" {
		attr(sb, "contentStyleType", a.ContentStyleType)
	}
	if a.ZoomAndPan != "" {
		attr(sb, "zoomAndPan", a.ZoomAndPan)
	}
	if a.X != "" {
		attr(sb, "x", a.X)
	}
	if a.Y != "" {
		attr(sb, "y", a.Y)
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
