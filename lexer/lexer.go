package lexer

import (
	"github.com/goCompiler/token"
)

type Lexer struct {
	input        string
	position     int  // Current position in input(points to current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipeWhiteSpaces()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = NewToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = NewToken(token.LPAREN, l.ch)
	case ')':
		tok = NewToken(token.RPAREN, l.ch)
	case '{':
		tok = NewToken(token.LBRACE, l.ch)
	case '}':
		tok = NewToken(token.RBRACE, l.ch)
	case '+':
		tok = NewToken(token.PLUS, l.ch)
	case ',':
		tok = NewToken(token.COMMA, l.ch)
	case '-':
		tok = NewToken(token.MINUS, l.ch)
	case '*':
		tok = NewToken(token.ASTERISK, l.ch)
	case '>':
		tok = NewToken(token.GT, l.ch)
	case '<':
		tok = NewToken(token.LT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQUAL, Literal: literal}
		} else {
			tok = NewToken(token.BANG, l.ch)
		}
	case '/':
		tok = NewToken(token.SLACH, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipeWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return ch <= '9' && ch >= '0'
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
