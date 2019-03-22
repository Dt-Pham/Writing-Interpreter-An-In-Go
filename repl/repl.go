package repl

import (
	"Writing-Interpreter-In-Go/lexer"
	"Writing-Interpreter-In-Go/token"
	"bufio"
	"fmt"
	"io"
)

// PROMPT ...
const PROMPT = ">> "

// Start ...
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		if ok := scanner.Scan(); !ok {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
