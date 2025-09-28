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

func (n Node) ApplyA(_ *AAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) ApplyAbbr(_ *AbbrAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyAddress(_ *AddressAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyArea(_ *AreaAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyArticle(_ *ArticleAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyAside(_ *AsideAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyAudio(_ *AudioAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyB(_ *BAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) ApplyBase(_ *BaseAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyBdi(_ *BdiAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyBdo(_ *BdoAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyBlockquote(_ *BlockquoteAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) ApplyBody(_ *BodyAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyBr(_ *BrAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyButton(_ *ButtonAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyCanvas(_ *CanvasAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyCaption(_ *CaptionAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyCite(_ *CiteAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyCode(_ *CodeAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyCol(_ *ColAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyColgroup(_ *ColgroupAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplyData(_ *DataAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyDatalist(_ *DatalistAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplyDd(_ *DdAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyDel(_ *DelAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyDetails(_ *DetailsAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyDfn(_ *DfnAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyDialog(_ *DialogAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyDiv(_ *DivAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyDl(_ *DlAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyDt(_ *DtAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyEm(_ *EmAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyEmbed(_ *EmbedAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyFieldset(_ *FieldsetAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplyFigcaption(_ *FigcaptionAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) ApplyFigure(_ *FigureAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyFooter(_ *FooterAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyForm(_ *FormAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyH1(_ *H1Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyH2(_ *H2Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyH3(_ *H3Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyH4(_ *H4Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyH5(_ *H5Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyH6(_ *H6Attrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyHead(_ *HeadAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyHeader(_ *HeaderAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyHgroup(_ *HgroupAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyHr(_ *HrAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyHtml(_ *HtmlAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyI(_ *IAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) ApplyIframe(_ *IframeAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyImg(_ *ImgAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyInput(_ *InputAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyIns(_ *InsAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyKbd(_ *KbdAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyLabel(_ *LabelAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyLegend(_ *LegendAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyLi(_ *LiAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyLink(_ *LinkAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyMain(_ *MainAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyMap(_ *MapAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyMark(_ *MarkAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyMath(_ *MathAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyMenu(_ *MenuAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyMeta(_ *MetaAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyMeter(_ *MeterAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyNav(_ *NavAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyNoscript(_ *NoscriptAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplyObject(_ *ObjectAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyOl(_ *OlAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyOptgroup(_ *OptgroupAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplyOption(_ *OptionAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyOutput(_ *OutputAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyP(_ *PAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) ApplyPicture(_ *PictureAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyPre(_ *PreAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyProgress(_ *ProgressAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplyQ(_ *QAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) ApplyRp(_ *RpAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyRt(_ *RtAttrs, kids *[]Component)                 { *kids = append(*kids, n) }
func (n Node) ApplyRuby(_ *RubyAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyS(_ *SAttrs, kids *[]Component)                   { *kids = append(*kids, n) }
func (n Node) ApplySamp(_ *SampAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyScript(_ *ScriptAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplySearch(_ *SearchAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplySection(_ *SectionAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplySelect(_ *SelectAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplySelectedcontent(_ *SelectedcontentAttrs, kids *[]Component) {
	*kids = append(*kids, n)
}
func (n Node) ApplySlot(_ *SlotAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplySmall(_ *SmallAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplySource(_ *SourceAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplySpan(_ *SpanAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyStrong(_ *StrongAttrs, kids *[]Component)     { *kids = append(*kids, n) }
func (n Node) ApplyStyle(_ *StyleAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplySub(_ *SubAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplySummary(_ *SummaryAttrs, kids *[]Component)   { *kids = append(*kids, n) }
func (n Node) ApplySup(_ *SupAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyTable(_ *TableAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyTbody(_ *TbodyAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyTd(_ *TdAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyTemplate(_ *TemplateAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) ApplyTextarea(_ *TextareaAttrs, kids *[]Component) { *kids = append(*kids, n) }
func (n Node) ApplyTfoot(_ *TfootAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyTh(_ *ThAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyThead(_ *TheadAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyTime(_ *TimeAttrs, kids *[]Component)         { *kids = append(*kids, n) }
func (n Node) ApplyTitle(_ *TitleAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyTr(_ *TrAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyTrack(_ *TrackAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyU(_ *UAttrs, kids *[]Component)               { *kids = append(*kids, n) }
func (n Node) ApplyUl(_ *UlAttrs, kids *[]Component)             { *kids = append(*kids, n) }
func (n Node) ApplyVar(_ *VarAttrs, kids *[]Component)           { *kids = append(*kids, n) }
func (n Node) ApplyVideo(_ *VideoAttrs, kids *[]Component)       { *kids = append(*kids, n) }
func (n Node) ApplyWbr(_ *WbrAttrs, kids *[]Component)           { *kids = append(*kids, n) }

func (o TxtOpt) ApplyA(_ *AAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyAbbr(_ *AbbrAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyAddress(_ *AddressAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyArea(_ *AreaAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyArticle(_ *ArticleAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyAside(_ *AsideAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyAudio(_ *AudioAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyB(_ *BAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyBase(_ *BaseAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyBdi(_ *BdiAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyBdo(_ *BdoAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyBlockquote(_ *BlockquoteAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyBody(_ *BodyAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyBr(_ *BrAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyButton(_ *ButtonAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyCanvas(_ *CanvasAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyCaption(_ *CaptionAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyCite(_ *CiteAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyCode(_ *CodeAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyCol(_ *ColAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyColgroup(_ *ColgroupAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyData(_ *DataAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyDatalist(_ *DatalistAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyDd(_ *DdAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyDel(_ *DelAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyDetails(_ *DetailsAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyDfn(_ *DfnAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyDialog(_ *DialogAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyDiv(_ *DivAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyDl(_ *DlAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyDt(_ *DtAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyEm(_ *EmAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyEmbed(_ *EmbedAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyFieldset(_ *FieldsetAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyFigcaption(_ *FigcaptionAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyFigure(_ *FigureAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyFooter(_ *FooterAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyForm(_ *FormAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyH1(_ *H1Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyH2(_ *H2Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyH3(_ *H3Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyH4(_ *H4Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyH5(_ *H5Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyH6(_ *H6Attrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyHead(_ *HeadAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyHeader(_ *HeaderAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyHgroup(_ *HgroupAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyHr(_ *HrAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyHtml(_ *HtmlAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyI(_ *IAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyIframe(_ *IframeAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyImg(_ *ImgAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyInput(_ *InputAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyIns(_ *InsAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyKbd(_ *KbdAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyLabel(_ *LabelAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyLegend(_ *LegendAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyLi(_ *LiAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyLink(_ *LinkAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyMain(_ *MainAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyMap(_ *MapAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyMark(_ *MarkAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyMath(_ *MathAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyMenu(_ *MenuAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyMeta(_ *MetaAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyMeter(_ *MeterAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyNav(_ *NavAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyNoscript(_ *NoscriptAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyObject(_ *ObjectAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyOl(_ *OlAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyOptgroup(_ *OptgroupAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyOption(_ *OptionAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyOutput(_ *OutputAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyP(_ *PAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyPicture(_ *PictureAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyPre(_ *PreAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyProgress(_ *ProgressAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyQ(_ *QAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyRp(_ *RpAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyRt(_ *RtAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyRuby(_ *RubyAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyS(_ *SAttrs, kids *[]Component)           { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySamp(_ *SampAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyScript(_ *ScriptAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySearch(_ *SearchAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySection(_ *SectionAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplySelect(_ *SelectAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySelectedcontent(_ *SelectedcontentAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplySlot(_ *SlotAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySmall(_ *SmallAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySource(_ *SourceAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySpan(_ *SpanAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyStrong(_ *StrongAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyStyle(_ *StyleAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySub(_ *SubAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplySummary(_ *SummaryAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplySup(_ *SupAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTable(_ *TableAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTbody(_ *TbodyAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTd(_ *TdAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTemplate(_ *TemplateAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyTextarea(_ *TextareaAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o TxtOpt) ApplyTfoot(_ *TfootAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTh(_ *ThAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyThead(_ *TheadAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTime(_ *TimeAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTitle(_ *TitleAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTr(_ *TrAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyTrack(_ *TrackAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyU(_ *UAttrs, kids *[]Component)         { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyUl(_ *UlAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyVar(_ *VarAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyVideo(_ *VideoAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o TxtOpt) ApplyWbr(_ *WbrAttrs, kids *[]Component)     { *kids = append(*kids, TextNode(o.s)) }

func (o UnsafeTxtOpt) ApplyA(_ *AAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyAbbr(_ *AbbrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyAddress(_ *AddressAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyArea(_ *AreaAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyArticle(_ *ArticleAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyAside(_ *AsideAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyAudio(_ *AudioAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyB(_ *BAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyBase(_ *BaseAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyBdi(_ *BdiAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyBdo(_ *BdoAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyBlockquote(_ *BlockquoteAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyBody(_ *BodyAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyBr(_ *BrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyButton(_ *ButtonAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyCanvas(_ *CanvasAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyCaption(_ *CaptionAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyCite(_ *CiteAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyCode(_ *CodeAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyCol(_ *ColAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyColgroup(_ *ColgroupAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyData(_ *DataAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDatalist(_ *DatalistAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDd(_ *DdAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDel(_ *DelAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDetails(_ *DetailsAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDfn(_ *DfnAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDialog(_ *DialogAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDiv(_ *DivAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDl(_ *DlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyDt(_ *DtAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyEm(_ *EmAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyEmbed(_ *EmbedAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyFieldset(_ *FieldsetAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyFigcaption(_ *FigcaptionAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyFigure(_ *FigureAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyFooter(_ *FooterAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyForm(_ *FormAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyH1(_ *H1Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyH2(_ *H2Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyH3(_ *H3Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyH4(_ *H4Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyH5(_ *H5Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyH6(_ *H6Attrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyHeader(_ *HeaderAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyHgroup(_ *HgroupAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyHr(_ *HrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyHtml(_ *HtmlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyI(_ *IAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyIframe(_ *IframeAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyImg(_ *ImgAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyInput(_ *InputAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyIns(_ *InsAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyKbd(_ *KbdAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyLabel(_ *LabelAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyLegend(_ *LegendAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyLi(_ *LiAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyLink(_ *LinkAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyMain(_ *MainAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyMap(_ *MapAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyMark(_ *MarkAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyMath(_ *MathAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyMenu(_ *MenuAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyMeta(_ *MetaAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyMeter(_ *MeterAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyNav(_ *NavAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyNoscript(_ *NoscriptAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyObject(_ *ObjectAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyOl(_ *OlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyOptgroup(_ *OptgroupAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyOption(_ *OptionAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyOutput(_ *OutputAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyP(_ *PAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyPicture(_ *PictureAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyPre(_ *PreAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyProgress(_ *ProgressAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyQ(_ *QAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyRp(_ *RpAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyRt(_ *RtAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyRuby(_ *RubyAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyS(_ *SAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySamp(_ *SampAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyScript(_ *ScriptAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySearch(_ *SearchAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySection(_ *SectionAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySelect(_ *SelectAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySelectedcontent(_ *SelectedcontentAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySlot(_ *SlotAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySmall(_ *SmallAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySource(_ *SourceAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySpan(_ *SpanAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyStrong(_ *StrongAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyStyle(_ *StyleAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySub(_ *SubAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySummary(_ *SummaryAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplySup(_ *SupAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTable(_ *TableAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTbody(_ *TbodyAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTd(_ *TdAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTemplate(_ *TemplateAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTextarea(_ *TextareaAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTfoot(_ *TfootAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTh(_ *ThAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyThead(_ *TheadAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTime(_ *TimeAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTitle(_ *TitleAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTr(_ *TrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyTrack(_ *TrackAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyU(_ *UAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyUl(_ *UlAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyVar(_ *VarAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyVideo(_ *VideoAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o UnsafeTxtOpt) ApplyWbr(_ *WbrAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}

func (o ChildOpt) ApplyA(_ *AAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyAbbr(_ *AbbrAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyAddress(_ *AddressAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyArea(_ *AreaAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyArticle(_ *ArticleAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyAside(_ *AsideAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyAudio(_ *AudioAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyB(_ *BAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyBase(_ *BaseAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyBdi(_ *BdiAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyBdo(_ *BdoAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyBlockquote(_ *BlockquoteAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyBody(_ *BodyAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyBr(_ *BrAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyButton(_ *ButtonAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyCanvas(_ *CanvasAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyCaption(_ *CaptionAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyCite(_ *CiteAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyCode(_ *CodeAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyCol(_ *ColAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyColgroup(_ *ColgroupAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyData(_ *DataAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDatalist(_ *DatalistAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDd(_ *DdAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDel(_ *DelAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDetails(_ *DetailsAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDfn(_ *DfnAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDialog(_ *DialogAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDiv(_ *DivAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDl(_ *DlAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyDt(_ *DtAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyEm(_ *EmAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyEmbed(_ *EmbedAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyFieldset(_ *FieldsetAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyFigcaption(_ *FigcaptionAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyFigure(_ *FigureAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyFooter(_ *FooterAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyForm(_ *FormAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyH1(_ *H1Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyH2(_ *H2Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyH3(_ *H3Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyH4(_ *H4Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyH5(_ *H5Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyH6(_ *H6Attrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyHead(_ *HeadAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyHeader(_ *HeaderAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyHgroup(_ *HgroupAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyHr(_ *HrAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyHtml(_ *HtmlAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyI(_ *IAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyIframe(_ *IframeAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyImg(_ *ImgAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyInput(_ *InputAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyIns(_ *InsAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyKbd(_ *KbdAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyLabel(_ *LabelAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyLegend(_ *LegendAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyLi(_ *LiAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyLink(_ *LinkAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyMain(_ *MainAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyMap(_ *MapAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyMark(_ *MarkAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyMath(_ *MathAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyMenu(_ *MenuAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyMeta(_ *MetaAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyMeter(_ *MeterAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyNav(_ *NavAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyNoscript(_ *NoscriptAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyObject(_ *ObjectAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyOl(_ *OlAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyOptgroup(_ *OptgroupAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyOption(_ *OptionAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyOutput(_ *OutputAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyP(_ *PAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyPicture(_ *PictureAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyPre(_ *PreAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyProgress(_ *ProgressAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyQ(_ *QAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyRp(_ *RpAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyRt(_ *RtAttrs, kids *[]Component)                 { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyRuby(_ *RubyAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyS(_ *SAttrs, kids *[]Component)                   { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySamp(_ *SampAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyScript(_ *ScriptAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySearch(_ *SearchAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySection(_ *SectionAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySelect(_ *SelectAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySelectedcontent(_ *SelectedcontentAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}
func (o ChildOpt) ApplySlot(_ *SlotAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySmall(_ *SmallAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySource(_ *SourceAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySpan(_ *SpanAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyStrong(_ *StrongAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyStyle(_ *StyleAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySub(_ *SubAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySummary(_ *SummaryAttrs, kids *[]Component)   { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplySup(_ *SupAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTable(_ *TableAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTbody(_ *TbodyAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTd(_ *TdAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTemplate(_ *TemplateAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTextarea(_ *TextareaAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTfoot(_ *TfootAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTh(_ *ThAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyThead(_ *TheadAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTime(_ *TimeAttrs, kids *[]Component)         { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTitle(_ *TitleAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTr(_ *TrAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyTrack(_ *TrackAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyU(_ *UAttrs, kids *[]Component)               { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyUl(_ *UlAttrs, kids *[]Component)             { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyVar(_ *VarAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyVideo(_ *VideoAttrs, kids *[]Component)       { *kids = append(*kids, o.c) }
func (o ChildOpt) ApplyWbr(_ *WbrAttrs, kids *[]Component)           { *kids = append(*kids, o.c) }
