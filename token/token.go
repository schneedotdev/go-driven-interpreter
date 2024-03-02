package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // tokens we don't classify
	EOF = "EOF"

	// Identifiers and literatls
	IDENT = "IDENT" // add, foobar, x, y
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	FSLASH = "/"
	ASTERISK = "*"

	// Comparison
	LT = "<"
	GT = ">"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}