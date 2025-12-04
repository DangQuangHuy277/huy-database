package parser

import (
	"io"
)

type TokenStream interface {
	NextToken() Token
}

type SQLScanner struct {
	scanner io.RuneScanner
}

func NewSQLScanner(r io.RuneScanner) *SQLScanner {
	return &SQLScanner{scanner: r}
}

func (s *SQLScanner) NextToken() Token {
	char, _, err := s.scanner.ReadRune()
	if err != nil {
		return NewSQLToken(string(char), UNEXPECTED_CHAR)
	}
	switch char {
	case ';', '.', '(', ')', ',',
		'*', '+', '~', '%', '&':
		lex := string(char)
		return NewSQLToken(lex, operators[lex])

	case '<':
		next, _, err := s.scanner.ReadRune()
		if err != nil {
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}
		switch next {
		case '>':
			return NewSQLToken("<>", NOT_EQ2)
		case '<':
			return NewSQLToken("<<", LT2)
		case '=':
			return NewSQLToken("<=", LT_EQ)
		case ' ':
			return NewSQLToken("<", LT)
		}

	case '>':
		next, _, err := s.scanner.ReadRune()
		if err != nil {
			return NewSQLToken(string(char), UNEXPECTED_CHAR)
		}

		switch next {
		case '>':
			return NewSQLToken("<>", LT2)
		case '=':
			return NewSQLToken("<=", LT_EQ)
		case ' ':
			return NewSQLToken("<", LT)
		default:
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}

	case '=':
		next, _, err := s.scanner.ReadRune()
		if err != nil {
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}
		switch next {
		case '=':
			return NewSQLToken("==", EQ)
		case ' ':
			return NewSQLToken("=", ASSIGN)
		default:
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}
	case '|':
		next, _, err := s.scanner.ReadRune()
		if err != nil {
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}
		switch next {
		case '|':
			return NewSQLToken("||", PIPE2)
		case ' ':
			return NewSQLToken("|", PIPE)
		}
	case '!':
		next, _, err := s.scanner.ReadRune()
		if err != nil {
			return NewSQLToken("!", UNEXPECTED_CHAR)
		}
		switch next {
		case '=':
			return NewSQLToken("!=", NOT_EQ1)
		default:
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}
	case '-':
		next, _, err := s.scanner.ReadRune()
		if err != nil {
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}
		switch next {
		case '>':
			next2, _, err := s.scanner.ReadRune()
			if err != nil {
				return NewSQLToken(string(next2), UNEXPECTED_CHAR)
			}
			if next2 == ' ' {
				return NewSQLToken("->", JPTR)
			} else if next2 == '>' {
				return NewSQLToken("->", JPTR2)
			}
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		case ' ':
			return NewSQLToken("-", MINUS)
			// TODO: Handle case single line comment
		}
	case '/':
		next, _, err := s.scanner.ReadRune()
		if err != nil {
			return NewSQLToken(string(next), UNEXPECTED_CHAR)
		}
		if next == ' ' {
			return NewSQLToken("/", DIV)
		}
		// TODO : Handle case multiline comment
	}

	return nil
}
