package html

import "strings"

// Use
type UseAttrs struct {
	Global           GlobalAttrs
	Href             string
	X                string
	Y                string
	Width            string
	Height           string
	Transform        string
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
	Opacity          string
}

type UseArg interface {
	applyUse(*UseAttrs, *[]Component)
}

func defaultUseAttrs() *UseAttrs {
	return &UseAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Use(args ...UseArg) Node {
	a := defaultUseAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyUse(a, &kids)
	}
	return Node{Tag: "use", Attrs: a, Kids: kids}
}

// Use-specific options
type UseXOpt struct{ v string }
type UseYOpt struct{ v string }
type UseWidthOpt struct{ v string }
type UseHeightOpt struct{ v string }

func UseX(v string) UseXOpt           { return UseXOpt{v} }
func UseY(v string) UseYOpt           { return UseYOpt{v} }
func UseWidth(v string) UseWidthOpt   { return UseWidthOpt{v} }
func UseHeight(v string) UseHeightOpt { return UseHeightOpt{v} }

func (g Global) applyUse(a *UseAttrs, _ *[]Component)              { g.do(&a.Global) }
func (o TxtOpt) applyUse(_ *UseAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyUse(_ *UseAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o HrefOpt) applyUse(a *UseAttrs, _ *[]Component)             { a.Href = o.v }
func (o UseXOpt) applyUse(a *UseAttrs, _ *[]Component)             { a.X = o.v }
func (o UseYOpt) applyUse(a *UseAttrs, _ *[]Component)             { a.Y = o.v }
func (o UseWidthOpt) applyUse(a *UseAttrs, _ *[]Component)         { a.Width = o.v }
func (o UseHeightOpt) applyUse(a *UseAttrs, _ *[]Component)        { a.Height = o.v }
func (o TransformOpt) applyUse(a *UseAttrs, _ *[]Component)        { a.Transform = o.v }
func (o FillOpt) applyUse(a *UseAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyUse(a *UseAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyUse(a *UseAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyUse(a *UseAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyUse(a *UseAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyUse(a *UseAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyUse(a *UseAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyUse(a *UseAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyUse(a *UseAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyUse(a *UseAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyUse(a *UseAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o OpacityOpt) applyUse(a *UseAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *UseAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Href != "" {
		attr(sb, "href", a.Href)
	}
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
	if a.Transform != "" {
		attr(sb, "transform", a.Transform)
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
	if a.Opacity != "" {
		attr(sb, "opacity", a.Opacity)
	}
}
