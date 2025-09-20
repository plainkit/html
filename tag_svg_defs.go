package html

import "strings"

// Defs
type DefsAttrs struct {
	Global           GlobalAttrs
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

type DefsArg interface {
	applyDefs(*DefsAttrs, *[]Component)
}

func defaultDefsAttrs() *DefsAttrs {
	return &DefsAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Defs(args ...DefsArg) Node {
	a := defaultDefsAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDefs(a, &kids)
	}
	return Node{Tag: "defs", Attrs: a, Kids: kids}
}

func (g Global) applyDefs(a *DefsAttrs, _ *[]Component)              { g.do(&a.Global) }
func (o TxtOpt) applyDefs(_ *DefsAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyDefs(_ *DefsAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o TransformOpt) applyDefs(a *DefsAttrs, _ *[]Component)        { a.Transform = o.v }
func (o FillOpt) applyDefs(a *DefsAttrs, _ *[]Component)             { a.Fill = o.v }
func (o FillOpacityOpt) applyDefs(a *DefsAttrs, _ *[]Component)      { a.FillOpacity = o.v }
func (o FillRuleOpt) applyDefs(a *DefsAttrs, _ *[]Component)         { a.FillRule = o.v }
func (o StrokeOpt) applyDefs(a *DefsAttrs, _ *[]Component)           { a.Stroke = o.v }
func (o StrokeWidthOpt) applyDefs(a *DefsAttrs, _ *[]Component)      { a.StrokeWidth = o.v }
func (o StrokeDasharrayOpt) applyDefs(a *DefsAttrs, _ *[]Component)  { a.StrokeDasharray = o.v }
func (o StrokeDashoffsetOpt) applyDefs(a *DefsAttrs, _ *[]Component) { a.StrokeDashoffset = o.v }
func (o StrokeLinecapOpt) applyDefs(a *DefsAttrs, _ *[]Component)    { a.StrokeLinecap = o.v }
func (o StrokeLinejoinOpt) applyDefs(a *DefsAttrs, _ *[]Component)   { a.StrokeLinejoin = o.v }
func (o StrokeOpacityOpt) applyDefs(a *DefsAttrs, _ *[]Component)    { a.StrokeOpacity = o.v }
func (o StrokeMiterlimitOpt) applyDefs(a *DefsAttrs, _ *[]Component) { a.StrokeMiterlimit = o.v }
func (o OpacityOpt) applyDefs(a *DefsAttrs, _ *[]Component)          { a.Opacity = o.v }

func (a *DefsAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
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
