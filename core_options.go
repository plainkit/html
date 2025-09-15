package blox

// Global option: one glue impl for all tags (methods are added in tag files)
type Global struct {
	f func(*GlobalAttrs)
}

func (g Global) do(ga *GlobalAttrs) {
	g.f(ga)
}

// Global attribute constructors
func Class(v string) Global {
	return Global{func(g *GlobalAttrs) { g.addClass(v) }}
}

func Id(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Id = v }}
}

func Title(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Title = v }}
}

func Lang(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Lang = v }}
}

func Dir(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Dir = v }}
}

func Role(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Role = v }}
}

func Hidden() Global {
	return Global{func(g *GlobalAttrs) { g.Hidden = true }}
}

func Inert() Global {
	return Global{func(g *GlobalAttrs) { g.Inert = true }}
}

func Autofocus() Global {
	return Global{func(g *GlobalAttrs) { g.Autofocus = true }}
}

func TabIndex(i int) Global {
	return Global{func(g *GlobalAttrs) { g.TabIndex = &i }}
}

func Draggable(b bool) Global {
	val := "false"
	if b {
		val = "true"
	}
	return Global{func(g *GlobalAttrs) { g.Draggable = &val }}
}

func Spellcheck(b bool) Global {
	val := "false"
	if b {
		val = "true"
	}
	return Global{func(g *GlobalAttrs) { g.Spellcheck = &val }}
}

func Translate(b bool) Global {
	val := "no"
	if b {
		val = "yes"
	}
	return Global{func(g *GlobalAttrs) { g.Translate = &val }}
}

func AccessKey(v string) Global {
	return Global{func(g *GlobalAttrs) { g.AccessKey = v }}
}

func Slot(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Slot = v }}
}

func Part(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Part = v }}
}

func Popover(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Popover = v }}
}

func Nonce(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Nonce = v }}
}

func IsAttr(v string) Global {
	return Global{func(g *GlobalAttrs) { g.Is = v }}
}

func ContentEditable(v string) Global {
	return Global{func(g *GlobalAttrs) { g.ContentEditable = v }}
}

func InputMode(v string) Global {
	return Global{func(g *GlobalAttrs) { g.InputMode = v }}
}

func EnterKeyHint(v string) Global {
	return Global{func(g *GlobalAttrs) { g.EnterKeyHint = v }}
}

func ExportParts(v string) Global {
	return Global{func(g *GlobalAttrs) { g.ExportParts = v }}
}

func ItemScope(b bool) Global {
	return Global{func(g *GlobalAttrs) { g.ItemScope = b }}
}

func ItemType(v string) Global {
	return Global{func(g *GlobalAttrs) { g.ItemType = v }}
}

func ItemId(v string) Global {
	return Global{func(g *GlobalAttrs) { g.ItemId = v }}
}

func ItemProp(v string) Global {
	return Global{func(g *GlobalAttrs) { g.ItemProp = v }}
}

func ItemRef(v string) Global {
	return Global{func(g *GlobalAttrs) { g.ItemRef = v }}
}

func XMLLang(v string) Global {
	return Global{func(g *GlobalAttrs) { g.XMLLang = v }}
}

func XMLBase(v string) Global {
	return Global{func(g *GlobalAttrs) { g.XMLBase = v }}
}

func VirtualKeyboardPolicy(v string) Global {
	return Global{func(g *GlobalAttrs) { g.VirtualKeyboardPolicy = v }}
}

func WritingSuggestions(b bool) Global {
	val := "false"
	if b {
		val = "true"
	}
	return Global{func(g *GlobalAttrs) { g.WritingSuggestions = &val }}
}

// Map-like convenience functions
func Data(k, v string) Global {
	return Global{func(g *GlobalAttrs) { g.setData(k, v) }}
}

func Aria(k, v string) Global {
	return Global{func(g *GlobalAttrs) { g.setAria(k, v) }}
}

func Style(k, v string) Global {
	return Global{func(g *GlobalAttrs) { g.setStyleKV(k, v) }}
}

func On(ev, handler string) Global {
	return Global{func(g *GlobalAttrs) { g.setEvent(ev, handler) }}
}

func Custom(k, v string) Global {
	return Global{func(g *GlobalAttrs) { g.setCustom(k, v) }}
}
