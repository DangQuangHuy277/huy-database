package parser

import (
	"slices"
	"strings"
	"testing"
)

func TestSQLScanner_NextToken(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{"SELECT * FROM mytable;",
			[]Token{NewTestToken("SELECT"), NewTestToken("FROM"), NewTestToken("mytable"), NewTestToken(";")}},
		{"INSERT INTO thistable VALUES (1);", []Token{NewTestToken("INSERT"), NewTestToken("INTO"), NewTestToken("thistable"),
			NewTestToken("VALUES"), NewTestToken("("), NewTestToken("1"), NewTestToken(")"), NewTestToken(";")}},
	}

	equalComplexObjects := func(a, b Token) bool {
		return a.GetString() == b.GetString()
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			scanner := NewSQLScanner(strings.NewReader(tt.input))
			tokenList := make([]Token, 0)
			for token := scanner.NextToken(); token != nil; token = scanner.NextToken() {
				tokenList = append(tokenList, token)
			}

			slices.EqualFunc(tokenList, tt.expected, equalComplexObjects)
		})
	}
}

type TestToken struct {
	str string
}

func (t *TestToken) GetString() string {
	return t.str
}

func NewTestToken(str string) *TestToken {
	return &TestToken{
		str: str,
	}
}
