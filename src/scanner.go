package main

import "strconv"

type Scanner struct {
	lox     *Lox
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string, lox *Lox) *Scanner {
	return &Scanner{
		lox:     lox,
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		// Start at the beginning of the next lexeme
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, *NewToken(EOF, "", nil, s.line))
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	currentByte := s.advance() // get next byte
	switch currentByte {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '*':
		s.addToken(STAR, nil)
	case '!':
		s.addToken(matchTernary(s.match('='), BANG_EQUAL, BANG), nil)
	case '=':
		s.addToken(matchTernary(s.match('='), EQUAL_EQUAL, EQUAL), nil)
	case '<':
		s.addToken(matchTernary(s.match('='), LESS_EQUAL, LESS), nil)
	case '>':
		s.addToken(matchTernary(s.match('='), GREATER_EQUAL, GREATER), nil)
	case '/':
		if s.match('/') { // deal with comments
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
	case ' ', '\r', '\t':
	// do nothing
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if isDigit(currentByte) {
			s.number()
		} else {
			s.lox.error(s.line, "Unexpected character")
			// TODO: not have each invalid character be its own error
		}
	}
}

func matchTernary(match bool, trueToken TokenType, falseToken TokenType) TokenType {
	if match {
		return trueToken
	} else {
		return falseToken
	}
}

func (s *Scanner) advance() byte {
	r := s.source[s.current]
	s.current += 1
	return r
}

func (s *Scanner) addToken(typ TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, *NewToken(typ, text, literal, s.line))
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	// found a match
	s.current += 1
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.source[s.current]
}

// lox does not support escape sequences
func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line += 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.lox.error(s.line, "Unterminated string.")
		return
	}

	// the closing "
	s.advance()

	// tring surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	// Look for fractional part
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// Consume the dot
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	// Convert the string to float64 (all numbers in lox are doubles)
	value, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		panic("floats not parsing as floats!")
	}

	s.addToken(NUMBER, value)
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return '\x00'
	}
	return s.source[s.current+1]
}
