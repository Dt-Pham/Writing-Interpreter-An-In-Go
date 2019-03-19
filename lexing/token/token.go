package token

// TokenType represent the type of token in Monkey Programming Language
type TokenType string

// Token struct int Monkey Programming Language
type Token struct {
	Type    TokenType
	Literal string
}
