package html

import "strings"

// Polygon
type PolygonAttrs struct {
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

type PolygonArg interface {
	applyPolygon(*PolygonAttrs, *[]Component)
}

func defaultPolygonAttrs() *PolygonAttrs {
	return &PolygonAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Polygon(args ...PolygonArg) Node {
	a := defaultPolygonAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPolygon(a, &kids)
	}
	return Node{Tag: "polygon", Attrs: a, Kids: kids}
}

// Polygon-specific options
type PointsOpt struct{ v string }

func Points(v string) PointsOpt { return PointsOpt{v} }

func (g Global) applyPolygon(a *PolygonAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyPolygon(_ *PolygonAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyPolygon(_ *PolygonAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o PointsOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)           { a.Points = o.v }
func (o PathLengthOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)       { a.PathLength = o.v }
func (o FillOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyPolygon(a *PolygonAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyPolygon(a *PolygonAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applyPolygon(a *PolygonAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *PolygonAttrs) writeAttrs(sb *strings.Builder) {
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
