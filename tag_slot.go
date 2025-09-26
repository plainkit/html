package html

import "strings"

type SlotAttrs struct {
	Global GlobalAttrs
	Name   string
}

type SlotArg interface {
	applySlot(*SlotAttrs, *[]Component)
}

func defaultSlotAttrs() *SlotAttrs {
	return &SlotAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Slot(args ...SlotArg) Node {
	a := defaultSlotAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySlot(a, &kids)
	}
	return Node{Tag: "slot", Attrs: a, Kids: kids}
}

func (g Global) applySlot(a *SlotAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o NameOpt) applySlot(a *SlotAttrs, _ *[]Component) {
	a.Name = o.v
}

func (a *SlotAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
}
