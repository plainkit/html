package html

import "strings"

type LiAttrs struct {
	Global GlobalAttrs
	Type   string
	Value  string
}

type LiArg interface {
	applyLi(*LiAttrs, *[]Component)
}

func defaultLiAttrs() *LiAttrs {
	return &LiAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Li(args ...LiArg) Node {
	a := defaultLiAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyLi(a, &kids)
	}
	return Node{Tag: "li", Attrs: a, Kids: kids}
}

func (g Global) applyLi(a *LiAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o TypeOpt) applyLi(a *LiAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o ValueOpt) applyLi(a *LiAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *LiAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
}
