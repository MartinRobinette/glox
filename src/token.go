package main

import "fmt"

type Token struct {
	literal interface{}
	lexeme  string
	typ     TokenType
	line    int
}

func NewToken(typ TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		typ:     typ,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%s %s %v", t.typ.ToString(), t.lexeme, t.literal)
}
