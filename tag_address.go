package html

import "strings"

// Address
type AddressAttrs struct {
	Global GlobalAttrs
}

type AddressArg interface {
	applyAddress(*AddressAttrs, *[]Component)
}

func defaultAddressAttrs() *AddressAttrs {
	return &AddressAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Address(args ...AddressArg) Node {
	a := defaultAddressAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyAddress(a, &kids)
	}
	return Node{Tag: "address", Attrs: a, Kids: kids}
}

func (g Global) applyAddress(a *AddressAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyAddress(_ *AddressAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyAddress(_ *AddressAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *AddressAttrs) writeAttrs(sb *strings.Builder)             { writeGlobal(sb, &a.Global) }
