package html

import "strings"

type MetaAttrs struct {
	Global    GlobalAttrs
	Charset   string
	Content   string
	HttpEquiv string
	Media     string
	Name      string
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
	g.Do(&a.Global)
}

func (o CharsetOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.Charset = o.v
}
func (o ContentOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.Content = o.v
}
func (o HttpEquivOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.HttpEquiv = o.v
}
func (o MediaOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.Media = o.v
}
func (o NameOpt) applyMeta(a *MetaAttrs, _ *[]Component) {
	a.Name = o.v
}

func (a *MetaAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Charset != "" {
		Attr(sb, "charset", a.Charset)
	}
	if a.Content != "" {
		Attr(sb, "content", a.Content)
	}
	if a.HttpEquiv != "" {
		Attr(sb, "http-equiv", a.HttpEquiv)
	}
	if a.Media != "" {
		Attr(sb, "media", a.Media)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
}
