package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var engine []string
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		engine = append(engine, text)
	}
	fmt.Printf("Engine:\n")

	for i, line := range engine {
		// fmt.Printf("%s", line)

		for j, c := range line {
			if !unicode.IsDigit(c) && c != '.' {
				// found a symbol, look around
				if (i > 0 && j > 0) {
					if unicode.IsDigit(engine[i][j]) {

					}
				}

			}
		}
	}


}
