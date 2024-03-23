package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("{Type:%s, Literal:%s}", t.Type, t.Literal)
}

const (
	ILLEGAL = "ILLEGAL" // tokens we don't classify
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	FSLASH   = "/"
	ASTERISK = "*"

	// Comparison
	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
