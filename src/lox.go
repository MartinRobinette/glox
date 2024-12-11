package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Lox struct {
	hadError bool
}

func NewLox() Lox {
	return Lox{
		hadError: false,
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	}

	lox := NewLox()
	if len(os.Args) == 2 {
		lox.runFile(os.Args[1])
	} else {
		lox.runPrompt()
	}
}

func (l *Lox) runFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("No such file: " + path)
		os.Exit(1)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	l.run(string(data))

	if l.hadError {
		os.Exit(65)
	}
}

func (l *Lox) runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // end of input
		}

		line := scanner.Text()
		l.run(line)
		l.hadError = false // allow for user error in prompt
	}
}

func (l *Lox) run(source string) {
	scanner := &Scanner{source: source}
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

//
// Error Reporting
//

func (l *Lox) error(line int, message string) {
	l.report(line, "", message)
}

func (l *Lox) report(line int, where string, message string) {
	fmt.Fprintln(os.Stderr,
		"[line "+strconv.Itoa(line)+"] Error"+where+": "+message)
	l.hadError = true
}
