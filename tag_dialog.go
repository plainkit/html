package html

import "strings"

type DialogAttrs struct {
	Global GlobalAttrs
	Open   bool
}

type DialogArg interface {
	applyDialog(*DialogAttrs, *[]Component)
}

func defaultDialogAttrs() *DialogAttrs {
	return &DialogAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dialog(args ...DialogArg) Node {
	a := defaultDialogAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDialog(a, &kids)
	}
	return Node{Tag: "dialog", Attrs: a, Kids: kids}
}

func (g Global) applyDialog(a *DialogAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o OpenOpt) applyDialog(a *DialogAttrs, _ *[]Component) {
	a.Open = true
}

func (a *DialogAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Open {
		BoolAttr(sb, "open")
	}
}
