package parser

import "io"

type Token interface {
	GetString() string
}

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

}
