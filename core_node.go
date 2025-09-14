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

// Children exposes the node's children for traversals that need to walk
// the component tree (e.g., asset collection). This enables upstream
// code to discover nested components without callers having to pass
// them explicitly.
func (n Node) Children() []Component { return n.Kids }

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

// Allow passing any Component (Node) directly as a child argument to content tags,
// removing the need to wrap with Child(...)/C(...).
// Node implements many *Arg interfaces by appending itself to the children slice.

// Common block-level containers
func (n Node) applyHtml(_ *HtmlAttrs, kids *[]Component)   { *kids = append(*kids, n) }
func (n Node) applyHead(_ *HeadAttrs, kids *[]Component)   { *kids = append(*kids, n) }
func (n Node) applyBody(_ *BodyAttrs, kids *[]Component)   { *kids = append(*kids, n) }
func (n Node) applyDiv(_ *DivAttrs, kids *[]Component)   { *kids = append(*kids, n) }
func (n Node) applyMain(_ *MainAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyHeader(_ *HeaderAttrs, kids *[]Component) {
	*kids = append(*kids, n)
}
func (n Node) applyFooter(_ *FooterAttrs, kids *[]Component)   { *kids = append(*kids, n) }
func (n Node) applySection(_ *SectionAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyArticle(_ *ArticleAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyNav(_ *NavAttrs, kids *[]Component)         { *kids = append(*kids, n) }

// Inline and text containers
func (n Node) applySpan(_ *SpanAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyP(_ *PAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyA(_ *AAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyButton(_ *ButtonAttrs, kids *[]Component) { *kids = append(*kids, n) }

// Headings
func (n Node) applyH1(_ *H1Attrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyH2(_ *H2Attrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyH3(_ *H3Attrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyH4(_ *H4Attrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyH5(_ *H5Attrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyH6(_ *H6Attrs, kids *[]Component) { *kids = append(*kids, n) }

// Lists
func (n Node) applyUl(_ *UlAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyOl(_ *OlAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyLi(_ *LiAttrs, kids *[]Component) { *kids = append(*kids, n) }

// Forms and related
func (n Node) applyForm(_ *FormAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyLabel(_ *LabelAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) applyFieldset(_ *FieldsetAttrs, kids *[]Component) { *kids = append(*kids, n) }

// Tables
func (n Node) applyTable(_ *TableAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyThead(_ *TheadAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyTbody(_ *TbodyAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyTfoot(_ *TfootAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyTr(_ *TrAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) applyTh(_ *ThAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) applyTd(_ *TdAttrs, kids *[]Component)       { *kids = append(*kids, n) }
