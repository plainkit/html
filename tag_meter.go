package html

import "strings"

type MeterAttrs struct {
	Global  GlobalAttrs
	High    string
	Low     string
	Max     string
	Min     string
	Optimum string
	Value   string
}

type MeterArg interface {
	applyMeter(*MeterAttrs, *[]Component)
}

func defaultMeterAttrs() *MeterAttrs {
	return &MeterAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Meter(args ...MeterArg) Node {
	a := defaultMeterAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMeter(a, &kids)
	}
	return Node{Tag: "meter", Attrs: a, Kids: kids}
}

func (g Global) applyMeter(a *MeterAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o HighOpt) applyMeter(a *MeterAttrs, _ *[]Component) {
	a.High = o.v
}
func (o LowOpt) applyMeter(a *MeterAttrs, _ *[]Component) {
	a.Low = o.v
}
func (o MaxOpt) applyMeter(a *MeterAttrs, _ *[]Component) {
	a.Max = o.v
}
func (o MinOpt) applyMeter(a *MeterAttrs, _ *[]Component) {
	a.Min = o.v
}
func (o OptimumOpt) applyMeter(a *MeterAttrs, _ *[]Component) {
	a.Optimum = o.v
}
func (o ValueOpt) applyMeter(a *MeterAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *MeterAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.High != "" {
		Attr(sb, "high", a.High)
	}
	if a.Low != "" {
		Attr(sb, "low", a.Low)
	}
	if a.Max != "" {
		Attr(sb, "max", a.Max)
	}
	if a.Min != "" {
		Attr(sb, "min", a.Min)
	}
	if a.Optimum != "" {
		Attr(sb, "optimum", a.Optimum)
	}
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
}
