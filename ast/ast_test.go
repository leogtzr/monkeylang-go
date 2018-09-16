package ast

import (
	"chango/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	pString := program.String()
	const expectedProgramString = `let myVar = anotherVar;`
	if pString != expectedProgramString {
		t.Errorf("program.String() wrong got=%q, expected: %q", pString, expectedProgramString)
	}
}
