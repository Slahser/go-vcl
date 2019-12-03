package token

// Token defineds a single VCL token
type Token struct {
	Type    Type
	Literal string
}

// Type is a set of lexical tokens of the VCL
type Type string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	COMMENT = "COMMENT"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	CIDR   = "CIDR"
	TRUE   = "TRUE"
	FALSE  = "FALSE"

	ASSIGN = "="
	MATCH  = "~"
	PLUS   = "+"
	BANG   = "!"
	EQUAL  = "=="
	AND    = "&&"
	OR     = "||"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	IF   = "IF"
	ELSE = "ELSE"

	RETURN     = "RETURN"
	IMPORT     = "IMPORT"
	ACL        = "ACL"
	BACKEND    = "BACKEND"
	SUBROUTINE = "SUBROUTINE"
	CALL       = "CALL"
)

// NewToken ...
func NewToken(tokenType Type, char byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

var keywords = map[string]Type{
	"sub":     SUBROUTINE,
	"call":    CALL,
	"true":    TRUE,
	"false":   FALSE,
	"if":      IF,
	"else":    ELSE,
	"return":  RETURN,
	"import":  IMPORT,
	"acl":     ACL,
	"backend": BACKEND,
}

// LookupIndent ...
func LookupIndent(indent string) Type {
	if tokenType, ok := keywords[indent]; ok {
		return tokenType
	}

	return IDENT
}
