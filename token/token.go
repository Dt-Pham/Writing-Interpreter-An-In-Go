package token

// TokenType represent the type of token in Monkey Programming Language
type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literal

	IDENT = "IDENT"
	INT   = "INT"

	// Operators

	ASSIGN = "="
	PLUS   = "+"

	// Delimiters

	COMMAS    = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "{"

	// Keywords

	LET      = "LET"
	FUNCTION = "FUNCTION"
)

// Token struct int Monkey Programming Language
type Token struct {
	Type    TokenType
	Literal string
}
