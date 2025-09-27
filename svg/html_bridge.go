package svg

import "github.com/plainkit/html"

// NodeAdapter wraps html.Node to implement SvgArg interface
type NodeAdapter struct {
	Node html.Node
}

func (a NodeAdapter) applySvg(_ *SvgAttrs, kids *[]html.Component) {
	*kids = append(*kids, a.Node)
}

// AdaptNode wraps an html.Node so it can be used as SvgArg
func AdaptNode(node html.Node) SvgArg {
	return NodeAdapter{Node: node}
}
