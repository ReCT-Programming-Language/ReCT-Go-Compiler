package symbols

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type LocalVariableSymbol struct {
	VariableSymbol

	Name     string
	ReadOnly bool
	Type     TypeSymbol
	UniqueID int
}

// implement the symbol interface
func (LocalVariableSymbol) SymbolType() SymbolType { return LocalVariable }
func (s LocalVariableSymbol) SymbolName() string   { return s.Name }

func (sym LocalVariableSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ LocalVariableSymbol ["+sym.Name+"]")
}

// implement the var interface
func (LocalVariableSymbol) IsGlobal() bool        { return false }
func (s LocalVariableSymbol) IsReadOnly() bool    { return s.ReadOnly }
func (s LocalVariableSymbol) VarType() TypeSymbol { return s.Type }
func (s LocalVariableSymbol) Fingerprint() string { return fmt.Sprintf("VL_%d", s.UniqueID) }

// constructor
func CreateLocalVariableSymbol(name string, readonly bool, typeSymbol TypeSymbol) LocalVariableSymbol {
	variableCounter++
	return LocalVariableSymbol{
		Name:     name,
		ReadOnly: readonly,
		Type:     typeSymbol,
		UniqueID: variableCounter,
	}
}
