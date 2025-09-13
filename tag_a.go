package blox

import "strings"

type ATarget string

const (
	TargetSelf  ATarget = "_self"
	TargetBlank ATarget = "_blank"
)

type AAttrs struct {
	Global GlobalAttrs
	Href   string
	Target ATarget
	Rel    string
}

type AArg interface {
	applyA(*AAttrs, *[]Component)
}

func defaultAAttrs() *AAttrs {
	return &AAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
		Target: TargetSelf,
	}
}

func A(args ...AArg) Component {
	a := defaultAAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyA(a, &kids)
	}
	return Node{Tag: "a", Attrs: a, Kids: kids}
}

// Tag-specific options
type HrefOpt struct {
	v string
}

type TargetOpt struct {
	v ATarget
}

type RelOpt struct {
	v string
}

func Href(v string) HrefOpt {
	return HrefOpt{v}
}

func Target(v ATarget) TargetOpt {
	return TargetOpt{v}
}

func Rel(v string) RelOpt {
	return RelOpt{v}
}

// Global option glue
func (g Global) applyA(a *AAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyA(_ *AAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyA(_ *AAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Tag-specific option glue
func (o HrefOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Href = o.v
}

func (o TargetOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Target = o.v
}

func (o RelOpt) applyA(a *AAttrs, _ *[]Component) {
	if a.Rel == "" {
		a.Rel = o.v
	} else {
		a.Rel += " " + o.v
	}
}

// Attrs writer implementation
func (a *AAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Href != "" {
		attr(sb, "href", a.Href)
	}
	if a.Target != "" {
		attr(sb, "target", string(a.Target))
	}
	if a.Rel != "" {
		attr(sb, "rel", a.Rel)
	}
}
