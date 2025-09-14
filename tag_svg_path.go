package blox

import "strings"

// Path
type PathAttrs struct {
	Global           GlobalAttrs
	D                string
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

type PathArg interface {
	applyPath(*PathAttrs, *[]Component)
}

func defaultPathAttrs() *PathAttrs {
	return &PathAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Path(args ...PathArg) Node {
	a := defaultPathAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPath(a, &kids)
	}
	return Node{Tag: "path", Attrs: a, Kids: kids}
}

// Path-specific options
type DOpt struct{ v string }

func D(v string) DOpt { return DOpt{v} }

func (g Global) applyPath(a *PathAttrs, _ *[]Component)              { g.do(&a.Global) }
func (o TxtOpt) applyPath(_ *PathAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyPath(_ *PathAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o DOpt) applyPath(a *PathAttrs, _ *[]Component)                { a.D = o.v }
func (o PathLengthOpt) applyPath(a *PathAttrs, _ *[]Component)       { a.PathLength = o.v }
func (o FillOpt) applyPath(a *PathAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyPath(a *PathAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyPath(a *PathAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyPath(a *PathAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyPath(a *PathAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyPath(a *PathAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyPath(a *PathAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyPath(a *PathAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyPath(a *PathAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyPath(a *PathAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyPath(a *PathAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o TransformOpt) applyPath(a *PathAttrs, _ *[]Component)        { a.Transform = o.v }
func (o OpacityOpt) applyPath(a *PathAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *PathAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.D != "" {
		attr(sb, "d", a.D)
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
