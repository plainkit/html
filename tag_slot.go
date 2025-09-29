package html

import "strings"

type SlotAttrs struct {
	Global GlobalAttrs
	Name   string
}

type SlotArg interface {
	ApplySlot(*SlotAttrs, *[]Component)
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
		ar.ApplySlot(a, &kids)
	}

	return Node{Tag: "slot", Attrs: a, Kids: kids}
}

func (g Global) ApplySlot(a *SlotAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o NameOpt) ApplySlot(a *SlotAttrs, _ *[]Component) {
	a.Name = o.v
}

func (a *SlotAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
}
