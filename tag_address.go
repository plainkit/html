package html

import "strings"

type AddressAttrs struct {
	Global GlobalAttrs
}

type AddressArg interface {
	ApplyAddress(*AddressAttrs, *[]Component)
}

func defaultAddressAttrs() *AddressAttrs {
	return &AddressAttrs{
		Global: GlobalAttrs{
			Style:  "",
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
		ar.ApplyAddress(a, &kids)
	}

	return Node{Tag: "address", Attrs: a, Kids: kids}
}

func (g Global) ApplyAddress(a *AddressAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *AddressAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
