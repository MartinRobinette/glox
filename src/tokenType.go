package main

type TokenType int

const (
	// Single Cahracter tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

// Map to store TokenType-to-string mappings
var tokenTypeStrings = map[TokenType]string{
	LEFT_PAREN:    "LEFT_PAREN",
	RIGHT_PAREN:   "RIGHT_PAREN",
	LEFT_BRACE:    "LEFT_BRACE",
	RIGHT_BRACE:   "RIGHT_BRACE",
	COMMA:         "COMMA",
	DOT:           "DOT",
	MINUS:         "MINUS",
	PLUS:          "PLUS",
	SEMICOLON:     "SEMICOLON",
	SLASH:         "SLASH",
	STAR:          "STAR",
	BANG:          "BANG",
	BANG_EQUAL:    "BANG_EQUAL",
	EQUAL:         "EQUAL",
	EQUAL_EQUAL:   "EQUAL_EQUAL",
	GREATER:       "GREATER",
	GREATER_EQUAL: "GREATER_EQUAL",
	LESS:          "LESS",
	LESS_EQUAL:    "LESS_EQUAL",
	IDENTIFIER:    "IDENTIFIER",
	STRING:        "STRING",
	NUMBER:        "NUMBER",
	AND:           "AND",
	CLASS:         "CLASS",
	ELSE:          "ELSE",
	FALSE:         "FALSE",
	FOR:           "FOR",
	IF:            "IF",
	NIL:           "NIL",
	OR:            "OR",
	PRINT:         "PRINT",
	RETURN:        "RETURN",
	SUPER:         "SUPER",
	THIS:          "THIS",
	TRUE:          "TRUE",
	VAR:           "VAR",
	WHILE:         "WHILE",
	EOF:           "EOF",
}

func (t *TokenType) ToString() string {
	if name, ok := tokenTypeStrings[*t]; ok {
		return name
	}
	return "UNKNOWN" // should never occur
}

// possible change, removes the need for a seperate to string funstion, does increase the size of each token
// const (
// 	// Single Cahracter tokens
// 	LEFT_PAREN  TokenType = "("
// 	RIGHT_PAREN TokenType = ")"
// 	LEFT_BRACE  TokenType = "}"
// 	RIGHT_BRACE TokenType = ""
// 	COMMA       TokenType = ","
// 	DOT         TokenType = "."
// 	MINUS       TokenType = "-"
// 	PLUS        TokenType = "+"
// 	SEMICOLON   TokenType = ";"
// 	SLASH       TokenType = "/"
// 	STAR        TokenType = "*"
//
// 	// One or two character tokens
// 	BANG          TokenType = "!"
// 	BANG_EQUAL    TokenType = "!="
// 	EQUAL         TokenType = "="
// 	EQUAL_EQUAL   TokenType = "=="
// 	GREATER       TokenType = ">"
// 	GREATER_EQUAL TokenType = ">="
// 	LESS          TokenType = "<"
// 	LESS_EQUAL    TokenType = "<="
//
// 	// Literals
// 	IDENTIFIER TokenType = "IDENT"
// 	STRING     TokenType = "STRING"
// 	NUMBER     TokenType = "NUMBER"
//
// 	// Keywords
// 	AND    TokenType = "AND"
// 	CLASS  TokenType = "CLASS"
// 	ELSE   TokenType = "ELSE"
// 	FALSE  TokenType = "FALSE"
// 	FOR    TokenType = "FOR"
// 	IF     TokenType = "IF"
// 	NIL    TokenType = "NIL"
// 	OR     TokenType = "OR"
// 	PRINT  TokenType = "PRINT"
// 	RETURN TokenType = "RETURN"
// 	SUPER  TokenType = "SUPER"
// 	THIS   TokenType = "THIS"
// 	TRUE   TokenType = "TRUE"
// 	VAR    TokenType = "VAR"
// 	WHILE  TokenType = "WHILE"
//
// 	EOF TokenType = "EOF"
// )
