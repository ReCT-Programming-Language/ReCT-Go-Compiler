package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type AssignmentExpressionNode struct {
	ExpressionNode

	InMain     bool
	Identifier lexer.Token
	Expression ExpressionNode
}

// implement node type from interface
func (AssignmentExpressionNode) NodeType() NodeType { return AssignmentExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node AssignmentExpressionNode) Span() print.TextSpan {
	return node.Identifier.Span.SpanBetween(node.Expression.Span())
}

// node print function
func (node AssignmentExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ AssignmentExpressionNode")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateAssignmentExpressionNode(id lexer.Token, expr ExpressionNode) AssignmentExpressionNode {
	return AssignmentExpressionNode{
		Identifier: id,
		Expression: expr,
	}
}

func CreateMainAssignmentExpressionNode(id lexer.Token, expr ExpressionNode) AssignmentExpressionNode {
	return AssignmentExpressionNode{
		Identifier: id,
		Expression: expr,
		InMain:     true,
	}
}
