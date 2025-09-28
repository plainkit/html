package html

import "strings"

type DialogAttrs struct {
	Global   GlobalAttrs
	Closedby string
	Open     bool
}

type DialogArg interface {
	ApplyDialog(*DialogAttrs, *[]Component)
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
		ar.ApplyDialog(a, &kids)
	}
	return Node{Tag: "dialog", Attrs: a, Kids: kids}
}

func (g Global) ApplyDialog(a *DialogAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o ClosedbyOpt) ApplyDialog(a *DialogAttrs, _ *[]Component) {
	a.Closedby = o.v
}
func (o OpenOpt) ApplyDialog(a *DialogAttrs, _ *[]Component) {
	a.Open = true
}

func (a *DialogAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Closedby != "" {
		Attr(sb, "closedby", a.Closedby)
	}
	if a.Open {
		BoolAttr(sb, "open")
	}
}
