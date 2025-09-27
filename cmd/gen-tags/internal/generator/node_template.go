package generator

const nodeTemplate = `package html


import (
	"html"
	"strconv"
	"strings"
)

type Component interface {
	render(*strings.Builder)
}

type TextNode string

func (t TextNode) render(sb *strings.Builder) {
	sb.WriteString(html.EscapeString(string(t)))
}

type UnsafeTextNode string

func (t UnsafeTextNode) render(sb *strings.Builder) {
	sb.WriteString(string(t))
}

type attrWriter interface {
	writeAttrs(*strings.Builder)
}

type Node struct {
	Tag       string
	Attrs     any         // must implement attrWriter
	Kids      []Component // empty for void tags
	Void      bool
	AssetCSS  string // CSS to be collected by asset system
	AssetJS   string // JavaScript to be collected by asset system
	AssetName string // Name for asset deduplication
}

func (n Node) render(sb *strings.Builder) {
	sb.WriteString("<")
	sb.WriteString(n.Tag)
	if aw, ok := n.Attrs.(attrWriter); ok {
		aw.writeAttrs(sb)
	}
	sb.WriteString(">")
	if !n.Void {
		for _, k := range n.Kids {
			k.render(sb)
		}
		sb.WriteString("</")
		sb.WriteString(n.Tag)
		sb.WriteString(">")
	}
}

func (n Node) Children() []Component { return n.Kids }

func (n Node) CSS() string  { return n.AssetCSS }
func (n Node) JS() string   { return n.AssetJS }
func (n Node) Name() string { return n.AssetName }

func (n Node) WithAssets(css, js, name string) Node {
	return Node{
		Tag:       n.Tag,
		Attrs:     n.Attrs,
		Kids:      n.Kids,
		Void:      n.Void,
		AssetCSS:  css,
		AssetJS:   js,
		AssetName: name,
	}
}

func Render(c Component) string {
	var sb strings.Builder
	c.render(&sb)
	return sb.String()
}

func Attr(sb *strings.Builder, k, v string) {
	sb.WriteString(" ")
	sb.WriteString(k)
	sb.WriteString(` + "`" + `="` + "`" + `)
	sb.WriteString(html.EscapeString(v))
	sb.WriteString(` + "`" + `"` + "`" + `)
}

func BoolAttr(sb *strings.Builder, k string) {
	sb.WriteString(" ")
	sb.WriteString(k)
}

{{range .TagNames}}func (n Node) apply{{.}}(_ *{{.}}Attrs, kids *[]Component) { *kids = append(*kids, n) }
{{end}}

{{range .TagNames}}func (o TxtOpt) apply{{.}}(_ *{{.}}Attrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
{{end}}

{{range .TagNames}}func (o UnsafeTxtOpt) apply{{.}}(_ *{{.}}Attrs, kids *[]Component) { *kids = append(*kids, UnsafeTextNode(o.s)) }
{{end}}

{{range .TagNames}}func (o ChildOpt) apply{{.}}(_ *{{.}}Attrs, kids *[]Component) { *kids = append(*kids, o.c) }
{{end}}`
