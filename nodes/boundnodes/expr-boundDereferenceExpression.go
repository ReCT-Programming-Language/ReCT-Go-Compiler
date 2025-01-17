package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundDereferenceExpressionNode struct {
	BoundExpressionNode

	Expression    BoundExpressionNode
	UnboundSource nodes.SyntaxNode
}

func (BoundDereferenceExpressionNode) NodeType() BoundType { return BoundDereferenceExpression }

func (node BoundDereferenceExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundDereferenceExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundReferenceExpressionNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func (node BoundDereferenceExpressionNode) IsPersistent() bool { return node.Expression.IsPersistent() }

// implement the expression node interface
func (node BoundDereferenceExpressionNode) Type() symbols.TypeSymbol {
	return node.Expression.Type().SubTypes[0]
}

func CreateBoundDereferenceExpressionNode(expression BoundExpressionNode, src nodes.SyntaxNode) BoundDereferenceExpressionNode {
	return BoundDereferenceExpressionNode{
		Expression:    expression,
		UnboundSource: src,
	}
}
