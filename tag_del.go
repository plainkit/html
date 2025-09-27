package html

import "strings"

type DelAttrs struct {
	Global   GlobalAttrs
	Cite     string
	Datetime string
}

type DelArg interface {
	applyDel(*DelAttrs, *[]Component)
}

func defaultDelAttrs() *DelAttrs {
	return &DelAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Del(args ...DelArg) Node {
	a := defaultDelAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDel(a, &kids)
	}
	return Node{Tag: "del", Attrs: a, Kids: kids}
}

func (g Global) applyDel(a *DelAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o CiteOpt) applyDel(a *DelAttrs, _ *[]Component) {
	a.Cite = o.v
}
func (o DatetimeOpt) applyDel(a *DelAttrs, _ *[]Component) {
	a.Datetime = o.v
}

func (a *DelAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Cite != "" {
		Attr(sb, "cite", a.Cite)
	}
	if a.Datetime != "" {
		Attr(sb, "datetime", a.Datetime)
	}
}
