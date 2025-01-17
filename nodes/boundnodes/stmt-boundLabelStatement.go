package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type BoundLabelStatementNode struct {
	BoundStatementNode

	Label         BoundLabel
	UnboundSource nodes.SyntaxNode
}

// implement the interface
func (BoundLabelStatementNode) NodeType() BoundType { return BoundLabelStatement }
func (node BoundLabelStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundLabelStatementNode")
	fmt.Printf("%s  └ Label: %s\n", indent, node.Label)
}

func (node BoundLabelStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// constructor
func CreateBoundLabelStatementNode(label BoundLabel, src nodes.SyntaxNode) BoundLabelStatementNode {
	return BoundLabelStatementNode{
		Label:         label,
		UnboundSource: src,
	}
}
