package html

import "strings"

type StyleAttrs struct {
	Global GlobalAttrs
	Media  string
	Nonce  string
	Type   string
}

type StyleArg interface {
	applyStyle(*StyleAttrs, *[]Component)
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
		ar.applyStyle(a, &kids)
	}
	return Node{Tag: "style", Attrs: a, Kids: kids}
}

func (g Global) applyStyle(a *StyleAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o MediaOpt) applyStyle(a *StyleAttrs, _ *[]Component) {
	a.Media = o.v
}
func (o NonceOpt) applyStyle(a *StyleAttrs, _ *[]Component) {
	a.Nonce = o.v
}
func (o TypeOpt) applyStyle(a *StyleAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *StyleAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Media != "" {
		Attr(sb, "media", a.Media)
	}
	if a.Nonce != "" {
		Attr(sb, "nonce", a.Nonce)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
