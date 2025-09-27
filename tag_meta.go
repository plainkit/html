package html

import "strings"

type MetaAttrs struct {
	Global    GlobalAttrs
	Charset   string
	Content   string
	Httpequiv string
}

type MetaArg interface {
	applyMeta(*MetaAttrs, *[]Component)
}

func defaultMetaAttrs() *MetaAttrs {
	return &MetaAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Meta(args ...MetaArg) Node {
	a := defaultMetaAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMeta(a, &kids)
	}
	return Node{Tag: "meta", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyMeta(a *MetaAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o CharsetOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.Charset = o.v
}
func (o ContentOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.Content = o.v
}
func (o HttpequivOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.Httpequiv = o.v
}

func (a *MetaAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Charset != "" {
		Attr(sb, "charset", a.Charset)
	}
	if a.Content != "" {
		Attr(sb, "content", a.Content)
	}
	if a.Httpequiv != "" {
		Attr(sb, "http-equiv", a.Httpequiv)
	}
}
