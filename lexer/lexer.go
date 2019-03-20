package lexer

// Lexer ...
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New create a lexer from given string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
