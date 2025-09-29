package html

import "strings"

type StyleAttrs struct {
	Global   GlobalAttrs
	Blocking string
	Media    string
}

type StyleArg interface {
	ApplyStyle(*StyleAttrs, *[]Component)
}

func defaultStyleAttrs() *StyleAttrs {
	return &StyleAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Style(args ...StyleArg) Node {
	a := defaultStyleAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyStyle(a, &kids)
	}

	return Node{Tag: "style", Attrs: a, Kids: kids}
}

func (g Global) ApplyStyle(a *StyleAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o BlockingOpt) ApplyStyle(a *StyleAttrs, _ *[]Component) {
	a.Blocking = o.v
}
func (o MediaOpt) ApplyStyle(a *StyleAttrs, _ *[]Component) {
	a.Media = o.v
}

func (a *StyleAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Blocking != "" {
		Attr(sb, "blocking", a.Blocking)
	}

	if a.Media != "" {
		Attr(sb, "media", a.Media)
	}
}
