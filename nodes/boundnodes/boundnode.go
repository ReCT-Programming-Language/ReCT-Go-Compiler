package boundnodes

import "ReCT-Go-Compiler/symbols"

// incredibly cool interface for creating bound nodes
type BoundNode interface {
	NodeType() BoundType
	Print(indent string)
}

type BoundStatementNode interface {
	BoundNode
}

type BoundExpressionNode interface {
	BoundNode
	Type() symbols.TypeSymbol
}

// enum for all our node types
type BoundType int

const (
	// Statements
	BoundBlockStatement BoundType = iota
	BoundVariableDeclaration
	BoundIfStatement
	BoundWhileStatement
	BoundDoWhileStatement
	BoundForStatement
	BoundLabelStatement
	BoundGotoStatement
	BoundConditionalGotoStatement
	BoundReturnStatement
	BoundExpressionStatement

	// Expressions
	BoundErrorExpression
	BoundLiteralExpression
	BoundVariableExpression
	BoundAssignmentExpression
	BoundUnaryExpression
	BoundBinaryExpression
	BoundCallExpression
	BoundConversionExpression
	BoundFromToStatement
)