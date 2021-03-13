package parser

import (
	"testing"

	ast "github.com/goCompiler/AST"
	"github.com/goCompiler/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 1000;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let' got=%q", s.TokenLiteral())
		return false
	}

	letSmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%q", s)
		return false
	}

	if letSmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letSmt.Name.Value)
		return false
	}

	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("LetSmt.Name.TokenLiteral() not '%s'. got='%s'", name, letSmt.Name.TokenLiteral())
		return false
	}
	return true
}
