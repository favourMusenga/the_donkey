package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/favourMusenga/the_donkey/lexer"
	"github.com/favourMusenga/the_donkey/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading input: %v", err)
			continue
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
