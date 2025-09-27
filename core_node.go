package html

import (
	"html"
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

type AttrWriter interface {
	WriteAttrs(*strings.Builder)
}

type Node struct {
	Tag       string
	Attrs     any         // must implement AttrWriter
	Kids      []Component // empty for void tags
	Void      bool
	AssetCSS  string // CSS to be collected by asset system
	AssetJS   string // JavaScript to be collected by asset system
	AssetName string // Name for asset deduplication
}

func (n Node) render(sb *strings.Builder) {
	sb.WriteString("<")
	sb.WriteString(n.Tag)
	if aw, ok := n.Attrs.(AttrWriter); ok {
		aw.WriteAttrs(sb)
	}
	if n.Void {
		sb.WriteString("/>")
	} else {
		sb.WriteString(">")
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
	sb.WriteString(`="`)
	sb.WriteString(html.EscapeString(v))
	sb.WriteString(`"`)
}

func BoolAttr(sb *strings.Builder, k string) {
	sb.WriteString(" ")
	sb.WriteString(k)
}

func (n Node) applyA(_ *AAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) applyApplet(_ *AppletAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyArea(_ *AreaAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyAudio(_ *AudioAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyBase(_ *BaseAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyBasefont(_ *BasefontAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyBlockquote(_ *BlockquoteAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) applyBody(_ *BodyAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyBr(_ *BrAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyButton(_ *ButtonAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyCanvas(_ *CanvasAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyCaption(_ *CaptionAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) applyCol(_ *ColAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyColgroup(_ *ColgroupAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyData(_ *DataAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyDel(_ *DelAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyDetails(_ *DetailsAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) applyDialog(_ *DialogAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyDir(_ *DirAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyDiv(_ *DivAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyDl(_ *DlAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyEmbed(_ *EmbedAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyFieldset(_ *FieldsetAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyFont(_ *FontAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyForm(_ *FormAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyFrame(_ *FrameAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyFrameset(_ *FramesetAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyH1(_ *H1Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyH2(_ *H2Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyH3(_ *H3Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyH4(_ *H4Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyH5(_ *H5Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyH6(_ *H6Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyHead(_ *HeadAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyHr(_ *HrAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyHtml(_ *HtmlAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyIframe(_ *IframeAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyImg(_ *ImgAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyInput(_ *InputAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyIns(_ *InsAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyIsindex(_ *IsindexAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) applyLabel(_ *LabelAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyLegend(_ *LegendAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyLi(_ *LiAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyLink(_ *LinkAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyMap(_ *MapAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyMenu(_ *MenuAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyMeta(_ *MetaAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyMeter(_ *MeterAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyObject(_ *ObjectAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyOl(_ *OlAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyOptgroup(_ *OptgroupAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyOption(_ *OptionAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyOutput(_ *OutputAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyP(_ *PAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) applyParam(_ *ParamAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyPre(_ *PreAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) applyProgress(_ *ProgressAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyQ(_ *QAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) applyScript(_ *ScriptAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applySelect(_ *SelectAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applySlot(_ *SlotAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applySource(_ *SourceAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) applyStyle(_ *StyleAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyTable(_ *TableAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyTbody(_ *TbodyAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyTd(_ *TdAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyTemplate(_ *TemplateAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyTextarea(_ *TextareaAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) applyTfoot(_ *TfootAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyTh(_ *ThAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyThead(_ *TheadAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyTime(_ *TimeAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) applyTr(_ *TrAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyTrack(_ *TrackAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) applyUl(_ *UlAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) applyVideo(_ *VideoAttrs, kids *[]Component)           { *kids = append(*kids, n) }

func (o TxtOpt) applyA(_ *AAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyApplet(_ *AppletAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyArea(_ *AreaAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyAudio(_ *AudioAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyBase(_ *BaseAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyBasefont(_ *BasefontAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyBlockquote(_ *BlockquoteAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyBody(_ *BodyAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyBr(_ *BrAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyButton(_ *ButtonAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyCanvas(_ *CanvasAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyCaption(_ *CaptionAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyCol(_ *ColAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyColgroup(_ *ColgroupAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyData(_ *DataAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyDel(_ *DelAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyDetails(_ *DetailsAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyDialog(_ *DialogAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyDir(_ *DirAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyDiv(_ *DivAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyDl(_ *DlAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyEmbed(_ *EmbedAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyFieldset(_ *FieldsetAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyFont(_ *FontAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyForm(_ *FormAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyFrame(_ *FrameAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyFrameset(_ *FramesetAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyH1(_ *H1Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyH2(_ *H2Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyH3(_ *H3Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyH4(_ *H4Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyH5(_ *H5Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyH6(_ *H6Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyHead(_ *HeadAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyHr(_ *HrAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyHtml(_ *HtmlAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyIframe(_ *IframeAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyImg(_ *ImgAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyInput(_ *InputAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyIns(_ *InsAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyIsindex(_ *IsindexAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyLabel(_ *LabelAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyLegend(_ *LegendAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyLi(_ *LiAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyLink(_ *LinkAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyMap(_ *MapAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyMenu(_ *MenuAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyMeta(_ *MetaAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyMeter(_ *MeterAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyObject(_ *ObjectAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyOl(_ *OlAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyOptgroup(_ *OptgroupAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyOption(_ *OptionAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyOutput(_ *OutputAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyP(_ *PAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyParam(_ *ParamAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyPre(_ *PreAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyProgress(_ *ProgressAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyQ(_ *QAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyScript(_ *ScriptAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applySelect(_ *SelectAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applySlot(_ *SlotAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applySource(_ *SourceAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyStyle(_ *StyleAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTable(_ *TableAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTbody(_ *TbodyAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTd(_ *TdAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTemplate(_ *TemplateAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyTextarea(_ *TextareaAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) applyTfoot(_ *TfootAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTh(_ *ThAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyThead(_ *TheadAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTime(_ *TimeAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTr(_ *TrAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyTrack(_ *TrackAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyUl(_ *UlAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) applyVideo(_ *VideoAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }

func (o UnsafeTxtOpt) applyA(_ *AAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyApplet(_ *AppletAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyArea(_ *AreaAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyAudio(_ *AudioAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyBase(_ *BaseAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyBasefont(_ *BasefontAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyBlockquote(_ *BlockquoteAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyBody(_ *BodyAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyBr(_ *BrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyButton(_ *ButtonAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyCanvas(_ *CanvasAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyCaption(_ *CaptionAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyCol(_ *ColAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyColgroup(_ *ColgroupAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyData(_ *DataAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyDel(_ *DelAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyDetails(_ *DetailsAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyDialog(_ *DialogAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyDir(_ *DirAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyDiv(_ *DivAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyDl(_ *DlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyEmbed(_ *EmbedAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyFieldset(_ *FieldsetAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyFont(_ *FontAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyForm(_ *FormAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyFrame(_ *FrameAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyFrameset(_ *FramesetAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyH1(_ *H1Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyH2(_ *H2Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyH3(_ *H3Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyH4(_ *H4Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyH5(_ *H5Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyH6(_ *H6Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyHr(_ *HrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyHtml(_ *HtmlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyIframe(_ *IframeAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyImg(_ *ImgAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyInput(_ *InputAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyIns(_ *InsAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyIsindex(_ *IsindexAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyLabel(_ *LabelAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyLegend(_ *LegendAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyLi(_ *LiAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyLink(_ *LinkAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyMap(_ *MapAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyMenu(_ *MenuAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyMeta(_ *MetaAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyMeter(_ *MeterAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyObject(_ *ObjectAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyOl(_ *OlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyOptgroup(_ *OptgroupAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyOption(_ *OptionAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyOutput(_ *OutputAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyP(_ *PAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyParam(_ *ParamAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyPre(_ *PreAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyProgress(_ *ProgressAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyQ(_ *QAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyScript(_ *ScriptAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applySelect(_ *SelectAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applySlot(_ *SlotAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applySource(_ *SourceAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyStyle(_ *StyleAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTable(_ *TableAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTbody(_ *TbodyAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTd(_ *TdAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTemplate(_ *TemplateAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTextarea(_ *TextareaAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTfoot(_ *TfootAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTh(_ *ThAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyThead(_ *TheadAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTime(_ *TimeAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTr(_ *TrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyTrack(_ *TrackAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyUl(_ *UlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) applyVideo(_ *VideoAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}

func (o ChildOpt) applyA(_ *AAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) applyApplet(_ *AppletAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyArea(_ *AreaAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyAudio(_ *AudioAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyBase(_ *BaseAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyBasefont(_ *BasefontAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyBlockquote(_ *BlockquoteAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ChildOpt) applyBody(_ *BodyAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyBr(_ *BrAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyButton(_ *ButtonAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyCanvas(_ *CanvasAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyCaption(_ *CaptionAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) applyCol(_ *ColAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyColgroup(_ *ColgroupAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyData(_ *DataAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyDel(_ *DelAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyDetails(_ *DetailsAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) applyDialog(_ *DialogAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyDir(_ *DirAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyDiv(_ *DivAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyDl(_ *DlAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyEmbed(_ *EmbedAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyFieldset(_ *FieldsetAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyFont(_ *FontAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyForm(_ *FormAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyFrame(_ *FrameAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyFrameset(_ *FramesetAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyH1(_ *H1Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyH2(_ *H2Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyH3(_ *H3Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyH4(_ *H4Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyH5(_ *H5Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyH6(_ *H6Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyHead(_ *HeadAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyHr(_ *HrAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyHtml(_ *HtmlAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyIframe(_ *IframeAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyImg(_ *ImgAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyInput(_ *InputAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyIns(_ *InsAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyIsindex(_ *IsindexAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) applyLabel(_ *LabelAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyLegend(_ *LegendAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyLi(_ *LiAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyLink(_ *LinkAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyMap(_ *MapAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyMenu(_ *MenuAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyMeta(_ *MetaAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyMeter(_ *MeterAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyObject(_ *ObjectAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyOl(_ *OlAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyOptgroup(_ *OptgroupAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyOption(_ *OptionAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyOutput(_ *OutputAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyP(_ *PAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) applyParam(_ *ParamAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyPre(_ *PreAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) applyProgress(_ *ProgressAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyQ(_ *QAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) applyScript(_ *ScriptAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applySelect(_ *SelectAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applySlot(_ *SlotAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applySource(_ *SourceAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) applyStyle(_ *StyleAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTable(_ *TableAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTbody(_ *TbodyAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTd(_ *TdAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTemplate(_ *TemplateAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTextarea(_ *TextareaAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTfoot(_ *TfootAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTh(_ *ThAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyThead(_ *TheadAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTime(_ *TimeAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTr(_ *TrAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyTrack(_ *TrackAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) applyUl(_ *UlAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) applyVideo(_ *VideoAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
