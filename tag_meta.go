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
	ApplyMeta(*MetaAttrs, *[]Component)
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
		ar.ApplyMeta(a, &kids)
	}
	return Node{Tag: "meta", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyMeta(a *MetaAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o CharsetOpt) ApplyMeta(a *MetaAttrs, _ *[]Component) {
	a.Charset = o.v
}
func (o ContentOpt) ApplyMeta(a *MetaAttrs, _ *[]Component) {
	a.Content = o.v
}
func (o HttpEquivOpt) ApplyMeta(a *MetaAttrs, _ *[]Component) {
	a.HttpEquiv = o.v
}
func (o MediaOpt) ApplyMeta(a *MetaAttrs, _ *[]Component) {
	a.Media = o.v
}
func (o NameOpt) ApplyMeta(a *MetaAttrs, _ *[]Component) {
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
