package blox

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

// attrWriter lets each tag write its own attributes (no central switch)
type attrWriter interface {
	writeAttrs(*strings.Builder)
}

type Node struct {
	Tag   string
	Attrs any         // must implement attrWriter
	Kids  []Component // empty for void tags
	Void  bool
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

func Render(c Component) string {
	var sb strings.Builder
	c.render(&sb)
	return sb.String()
}

// helper functions for attribute writing
func attr(sb *strings.Builder, k, v string) {
	sb.WriteString(" ")
	sb.WriteString(k)
	sb.WriteString(`="`)
	sb.WriteString(html.EscapeString(v))
	sb.WriteString(`"`)
}

func boolAttr(sb *strings.Builder, k string) {
	sb.WriteString(" ")
	sb.WriteString(k)
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
