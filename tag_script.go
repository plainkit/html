package blox

import "strings"

// Script
type ScriptAttrs struct {
	Global GlobalAttrs
	Src    string
	Type   string
	Async  bool
	Defer  bool
}

type ScriptArg interface {
	applyScript(*ScriptAttrs, *[]Component)
}

func defaultScriptAttrs() *ScriptAttrs {
	return &ScriptAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Script(args ...ScriptArg) Node {
	a := defaultScriptAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyScript(a, &kids)
	}
	return Node{Tag: "script", Attrs: a, Kids: kids}
}

// Script-specific options
type AsyncOpt struct{}
type DeferOpt struct{}
type ScriptSrcOpt struct{ v string }

func Async() AsyncOpt            { return AsyncOpt{} }
func Defer() DeferOpt            { return DeferOpt{} }
func ScriptSrc(v string) ScriptSrcOpt { return ScriptSrcOpt{v} }

func (g Global) applyScript(a *ScriptAttrs, _ *[]Component)    { g.do(&a.Global) }
func (o TxtOpt) applyScript(_ *ScriptAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o UnsafeTxtOpt) applyScript(_ *ScriptAttrs, kids *[]Component) {
    *kids = append(*kids, UnsafeTextNode(o.s))
}
func (o ChildOpt) applyScript(_ *ScriptAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o AsyncOpt) applyScript(a *ScriptAttrs, _ *[]Component)    { a.Async = true }
func (o DeferOpt) applyScript(a *ScriptAttrs, _ *[]Component)    { a.Defer = true }
func (o ScriptSrcOpt) applyScript(a *ScriptAttrs, _ *[]Component) { a.Src = o.v }

func (a *ScriptAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Src != "" {
		attr(sb, "src", a.Src)
	}
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Async {
		boolAttr(sb, "async")
	}
	if a.Defer {
		boolAttr(sb, "defer")
	}
}
