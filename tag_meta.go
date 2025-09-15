package html

import "strings"

type MetaAttrs struct {
	Global    GlobalAttrs
	Name      string
	Content   string
	HttpEquiv string
	Charset   string
	Property  string
	Scheme    string
}

type MetaArg interface {
	applyMeta(*MetaAttrs)
}

func defaultMetaAttrs() *MetaAttrs {
	return &MetaAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Meta(args ...MetaArg) Node {
	a := defaultMetaAttrs()
	for _, ar := range args {
		ar.applyMeta(a)
	}
	return Node{Tag: "meta", Attrs: a, Void: true}
}

// Tag-specific options
type NameOpt struct{ v string }
type ContentOpt struct{ v string }
type HttpEquivOpt struct{ v string }
type CharsetOpt struct{ v string }
type PropertyOpt struct{ v string }
type SchemeOpt struct{ v string }

func Name(v string) NameOpt           { return NameOpt{v} }
func Content(v string) ContentOpt     { return ContentOpt{v} }
func HttpEquiv(v string) HttpEquivOpt { return HttpEquivOpt{v} }
func Charset(v string) CharsetOpt     { return CharsetOpt{v} }
func Property(v string) PropertyOpt   { return PropertyOpt{v} }
func Scheme(v string) SchemeOpt       { return SchemeOpt{v} }

// Global option glue
func (g Global) applyMeta(a *MetaAttrs) {
	g.do(&a.Global)
}

// Tag-specific option glue
func (o NameOpt) applyMeta(a *MetaAttrs)      { a.Name = o.v }
func (o ContentOpt) applyMeta(a *MetaAttrs)   { a.Content = o.v }
func (o HttpEquivOpt) applyMeta(a *MetaAttrs) { a.HttpEquiv = o.v }
func (o CharsetOpt) applyMeta(a *MetaAttrs)   { a.Charset = o.v }
func (o PropertyOpt) applyMeta(a *MetaAttrs)  { a.Property = o.v }
func (o SchemeOpt) applyMeta(a *MetaAttrs)    { a.Scheme = o.v }

// Attrs writer implementation
func (a *MetaAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Name != "" {
		attr(sb, "name", a.Name)
	}
	if a.Content != "" {
		attr(sb, "content", a.Content)
	}
	if a.HttpEquiv != "" {
		attr(sb, "http-equiv", a.HttpEquiv)
	}
	if a.Charset != "" {
		attr(sb, "charset", a.Charset)
	}
	if a.Property != "" {
		attr(sb, "property", a.Property)
	}
	if a.Scheme != "" {
		attr(sb, "scheme", a.Scheme)
	}
}
