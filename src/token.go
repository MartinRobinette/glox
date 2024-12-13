package main

import "fmt"

type Token struct {
	literal interface{}
	lexeme  string
	typ     TokenType
	line    int
}

func (t *Token) NewToken(typ TokenType, lexeme string, literal interface{}, line int) {
	t.typ = typ
	t.lexeme = lexeme
	t.literal = literal
	t.line = line
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%s %s %v", t.typ.ToString(), t.lexeme, t.literal)
}
