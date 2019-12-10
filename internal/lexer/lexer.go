package lexer

import (
	"strings"

	"github.com/KeisukeYamashita/go-vcl/internal/token"
)

// Lexer ...
type Lexer struct {
	input   string
	pos     int
	readPos int
	char    byte
}

// NewLexer ...
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.init()
	return l
}

func (l *Lexer) init() {
	l.readChar()
}

// readChar retrieves the byte from readPos
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

// readIndentifier ...
func (l *Lexer) readIndentifier() string {
	pos := l.pos
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readString() string {
	pos := l.pos + 1
	for l.char != 0 && l.char != ';' {
		l.readChar()
		if l.char == '"' {
			if l.peekChar() != '/' {
				break
			}
		}
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}

	return l.input[l.readPos]
}

func (l *Lexer) eatWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

// NextToken ...
func (l *Lexer) NextToken() token.Token {
	l.eatWhiteSpace()

	tok := token.Token{}
	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(char)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = token.NewToken(token.ASSIGN, l.char)
		}
	case '~':
		tok = token.NewToken(token.MATCH, l.char)
	case ',':
		tok = token.NewToken(token.COMMA, l.char)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.char)
	case '#':
		tok = token.NewToken(token.COMMENT, l.char)
	case '(':
		tok = token.NewToken(token.LPAREN, l.char)
	case ')':
		tok = token.NewToken(token.RPAREN, l.char)
	case '{':
		tok = token.NewToken(token.LBRACE, l.char)
	case '}':
		tok = token.NewToken(token.RBRACE, l.char)
	case '!':
		tok = token.NewToken(token.BANG, l.char)
	case '+':
		tok = token.NewToken(token.PLUS, l.char)
	case '"':
		s := l.readString()
		if strings.Contains(s, "/") {
			tok.Type = token.CIDR
			s = "\"" + s // CIDR format is "35.0.0.0"/24 which we have to wrap by ".
		} else {
			tok.Type = token.STRING
		}
		tok.Literal = s
	case '|':
		// it will be always &&
		if l.peekChar() == '|' {
			char := l.char
			l.readChar()
			literal := string(char) + string(char)
			tok = token.Token{Type: token.OR, Literal: literal}
		}
	case '&':
		// it will be always &&
		if l.peekChar() == '&' {
			char := l.char
			l.readChar()
			literal := string(char) + string(char)
			tok = token.Token{Type: token.AND, Literal: literal}
		}
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		// TODO(KeisukeYamashita): add cidr
		if isLetter(l.char) {
			tok.Literal = l.readIndentifier()
			tok.Type = token.LookupIndent(tok.Literal)
			return tok // early return not to walk step
		} else if isDigit(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok // early return not to walk step
		} else {
			tok = token.NewToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' || char == '.'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}