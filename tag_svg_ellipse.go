package blox

import "strings"

// Ellipse
type EllipseAttrs struct {
	Global           GlobalAttrs
	Cx               string
	Cy               string
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

type EllipseArg interface {
	applyEllipse(*EllipseAttrs, *[]Component)
}

func defaultEllipseAttrs() *EllipseAttrs {
	return &EllipseAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ellipse(args ...EllipseArg) Node {
	a := defaultEllipseAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyEllipse(a, &kids)
	}
	return Node{Tag: "ellipse", Attrs: a, Kids: kids}
}

// Ellipse-specific options
type EllipseCxOpt struct{ v string }
type EllipseCyOpt struct{ v string }
type EllipseRxOpt struct{ v string }
type EllipseRyOpt struct{ v string }

func EllipseCx(v string) EllipseCxOpt { return EllipseCxOpt{v} }
func EllipseCy(v string) EllipseCyOpt { return EllipseCyOpt{v} }
func EllipseRx(v string) EllipseRxOpt { return EllipseRxOpt{v} }
func EllipseRy(v string) EllipseRyOpt { return EllipseRyOpt{v} }

func (g Global) applyEllipse(a *EllipseAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyEllipse(_ *EllipseAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyEllipse(_ *EllipseAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o EllipseCxOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)        { a.Cx = o.v }
func (o EllipseCyOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)        { a.Cy = o.v }
func (o EllipseRxOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)        { a.Rx = o.v }
func (o EllipseRyOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)        { a.Ry = o.v }
func (o PathLengthOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)       { a.PathLength = o.v }
func (o FillOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyEllipse(a *EllipseAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyEllipse(a *EllipseAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applyEllipse(a *EllipseAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *EllipseAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Cx != "" {
		attr(sb, "cx", a.Cx)
	}
	if a.Cy != "" {
		attr(sb, "cy", a.Cy)
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
