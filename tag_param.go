package html

import "strings"

type ParamAttrs struct {
	Global    GlobalAttrs
	Name      string
	Type      string
	Value     string
	Valuetype string
}

type ParamArg interface {
	applyParam(*ParamAttrs, *[]Component)
}

func defaultParamAttrs() *ParamAttrs {
	return &ParamAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Param(args ...ParamArg) Node {
	a := defaultParamAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyParam(a, &kids)
	}
	return Node{Tag: "param", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyParam(a *ParamAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o NameOpt) applyParam(a *ParamAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o TypeOpt) applyParam(a *ParamAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o ValueOpt) applyParam(a *ParamAttrs, _ *[]Component) {
	a.Value = o.v
}
func (o ValuetypeOpt) applyParam(a *ParamAttrs, _ *[]Component) {
	a.Valuetype = o.v
}

func (a *ParamAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
	if a.Valuetype != "" {
		Attr(sb, "valuetype", a.Valuetype)
	}
}
