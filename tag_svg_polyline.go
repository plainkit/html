package blox

import "strings"

// Polyline
type PolylineAttrs struct {
	Global           GlobalAttrs
	Points           string
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

type PolylineArg interface {
	applyPolyline(*PolylineAttrs, *[]Component)
}

func defaultPolylineAttrs() *PolylineAttrs {
	return &PolylineAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Polyline(args ...PolylineArg) Component {
	a := defaultPolylineAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPolyline(a, &kids)
	}
	return Node{Tag: "polyline", Attrs: a, Kids: kids}
}

func (g Global) applyPolyline(a *PolylineAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyPolyline(_ *PolylineAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyPolyline(_ *PolylineAttrs, kids *[]Component)        { *kids = append(*kids, o.c) }
func (o PointsOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)          { a.Points = o.v }
func (o PathLengthOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)      { a.PathLength = o.v }
func (o FillOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)            { a.Fill = o.v }
func (o FillOpacityOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)     { a.FillOpacity = o.v }
func (o FillRuleOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)        { a.FillRule = o.v }
func (o StrokeOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)          { a.Stroke = o.v }
func (o StrokeWidthOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)     { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyPolyline(a *PolylineAttrs, _ *[]Component) { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyPolyline(a *PolylineAttrs, _ *[]Component) {
	a.StrokeDashoffset = o.v
}
func (o StrokeLinecapOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)  { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyPolyline(a *PolylineAttrs, _ *[]Component) { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)  { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyPolyline(a *PolylineAttrs, _ *[]Component) {
	a.StrokeMiterlimit = o.v
}
func (o TransformOpt) applyPolyline(a *PolylineAttrs, _ *[]Component) { a.Transform = o.v }
func (o OpacityOpt) applyPolyline(a *PolylineAttrs, _ *[]Component)   { a.Opacity = o.v }

func (a *PolylineAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Points != "" {
		attr(sb, "points", a.Points)
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
