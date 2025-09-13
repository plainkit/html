package blox

import "strings"

// Dialog
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
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dialog(args ...DialogArg) Component {
	a := defaultDialogAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDialog(a, &kids)
	}
	return Node{Tag: "dialog", Attrs: a, Kids: kids}
}

func (g Global) applyDialog(a *DialogAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyDialog(_ *DialogAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyDialog(_ *DialogAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o OpenOpt) applyDialog(a *DialogAttrs, _ *[]Component)     { a.Open = true }

func (a *DialogAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Open {
		boolAttr(sb, "open")
	}
}
