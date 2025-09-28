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
	ApplyMeter(*MeterAttrs, *[]Component)
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
		ar.ApplyMeter(a, &kids)
	}
	return Node{Tag: "meter", Attrs: a, Kids: kids}
}

func (g Global) ApplyMeter(a *MeterAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o HighOpt) ApplyMeter(a *MeterAttrs, _ *[]Component) {
	a.High = o.v
}
func (o LowOpt) ApplyMeter(a *MeterAttrs, _ *[]Component) {
	a.Low = o.v
}
func (o MaxOpt) ApplyMeter(a *MeterAttrs, _ *[]Component) {
	a.Max = o.v
}
func (o MinOpt) ApplyMeter(a *MeterAttrs, _ *[]Component) {
	a.Min = o.v
}
func (o OptimumOpt) ApplyMeter(a *MeterAttrs, _ *[]Component) {
	a.Optimum = o.v
}
func (o ValueOpt) ApplyMeter(a *MeterAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *MeterAttrs) WriteAttrs(sb *strings.Builder) {
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
