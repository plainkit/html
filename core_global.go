package html

import "strings"

type GlobalAttrs struct {
	// Common core
	Id, Class, Title, Role, Lang, Dir, Slot, Part, Popover, Nonce, Is string
	AccessKey, ContentEditable, InputMode, EnterKeyHint, ExportParts  string
	ItemType, ItemId, ItemProp, ItemRef                               string
	XMLLang, XMLBase, VirtualKeyboardPolicy                           string

	// Style attribute as a single string
	Style string

	// Map attributes
	Aria   map[string]string // aria-*
	Data   map[string]string // data-*
	Events map[string]string // "onclick" -> "handler()"
	Custom map[string]string // custom attributes like hx-*, x-*, etc.

	// Pointers for tri-state values
	TabIndex                                             *int
	Spellcheck, Translate, Draggable, WritingSuggestions *string

	// Booleans
	Hidden, Inert, Autofocus, ItemScope bool
}

// Helper methods for setting attributes
func (g *GlobalAttrs) addClass(v string) {
	if v == "" {
		return
	}
	if g.Class == "" {
		g.Class = v
	} else {
		g.Class += " " + v
	}
}

func (g *GlobalAttrs) setStyle(style string) {
	g.Style = style
}

func (g *GlobalAttrs) setAria(k, v string) {
	if g.Aria == nil {
		g.Aria = map[string]string{}
	}
	g.Aria[k] = v
}

func (g *GlobalAttrs) setData(k, v string) {
	if g.Data == nil {
		g.Data = map[string]string{}
	}
	g.Data[k] = v
}

func (g *GlobalAttrs) setEvent(ev, handler string) {
	if g.Events == nil {
		g.Events = map[string]string{}
	}
	g.Events["on"+ev] = handler
}

func (g *GlobalAttrs) setCustom(k, v string) {
	if g.Custom == nil {
		g.Custom = map[string]string{}
	}
	g.Custom[k] = v
}

// Common writer reused by each tag's writeAttrs
func writeGlobal(sb *strings.Builder, g *GlobalAttrs) {
	// Simple scalars
	if g.Id != "" {
		attr(sb, "id", g.Id)
	}
	if g.Class != "" {
		attr(sb, "class", g.Class)
	}
	if g.Title != "" {
		attr(sb, "title", g.Title)
	}
	if g.Role != "" {
		attr(sb, "role", g.Role)
	}
	if g.Lang != "" {
		attr(sb, "lang", g.Lang)
	}
	if g.Dir != "" {
		attr(sb, "dir", g.Dir)
	}
	if g.Slot != "" {
		attr(sb, "slot", g.Slot)
	}
	if g.Part != "" {
		attr(sb, "part", g.Part)
	}
	if g.Popover != "" {
		attr(sb, "popover", g.Popover)
	}
	if g.Nonce != "" {
		attr(sb, "nonce", g.Nonce)
	}
	if g.Is != "" {
		attr(sb, "is", g.Is)
	}
	if g.AccessKey != "" {
		attr(sb, "accesskey", g.AccessKey)
	}
	if g.ContentEditable != "" {
		attr(sb, "contenteditable", g.ContentEditable)
	}
	if g.InputMode != "" {
		attr(sb, "inputmode", g.InputMode)
	}
	if g.EnterKeyHint != "" {
		attr(sb, "enterkeyhint", g.EnterKeyHint)
	}
	if g.ExportParts != "" {
		attr(sb, "exportparts", g.ExportParts)
	}
	if g.ItemScope {
		boolAttr(sb, "itemscope")
	}
	if g.ItemType != "" {
		attr(sb, "itemtype", g.ItemType)
	}
	if g.ItemId != "" {
		attr(sb, "itemid", g.ItemId)
	}
	if g.ItemProp != "" {
		attr(sb, "itemprop", g.ItemProp)
	}
	if g.ItemRef != "" {
		attr(sb, "itemref", g.ItemRef)
	}
	if g.XMLLang != "" {
		attr(sb, "xml:lang", g.XMLLang)
	}
	if g.XMLBase != "" {
		attr(sb, "xml:base", g.XMLBase)
	}
	if g.VirtualKeyboardPolicy != "" {
		attr(sb, "virtualkeyboardpolicy", g.VirtualKeyboardPolicy)
	}

	// Boolean attrs
	if g.Hidden {
		boolAttr(sb, "hidden")
	}
	if g.Inert {
		boolAttr(sb, "inert")
	}
	if g.Autofocus {
		boolAttr(sb, "autofocus")
	}

	// Pointer values
	if g.TabIndex != nil {
		attr(sb, "tabindex", itoa(*g.TabIndex))
	}
	if g.Spellcheck != nil {
		attr(sb, "spellcheck", *g.Spellcheck)
	}
	if g.Translate != nil {
		attr(sb, "translate", *g.Translate)
	}
	if g.Draggable != nil {
		attr(sb, "draggable", *g.Draggable)
	}
	if g.WritingSuggestions != nil {
		attr(sb, "writingsuggestions", *g.WritingSuggestions)
	}

	// Style attribute
	if g.Style != "" {
		attr(sb, "style", g.Style)
	}

	// Aria attributes
	for k, v := range g.Aria {
		if k != "" {
			attr(sb, "aria-"+k, v)
		}
	}

	// Data attributes
	for k, v := range g.Data {
		if k != "" {
			attr(sb, "data-"+k, v)
		}
	}

	// Event handlers
	for evAttr, handler := range g.Events {
		if evAttr != "" && handler != "" {
			attr(sb, evAttr, handler)
		}
	}

	// Custom attributes
	for k, v := range g.Custom {
		if k != "" && v != "" {
			attr(sb, k, v)
		}
	}
}
