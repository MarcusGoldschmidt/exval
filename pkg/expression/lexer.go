package expression

import (
	"fmt"
)

//go:generate stringer -type=tokenType -output=token_type_string.go
type tokenType int

const (
	symbol tokenType = iota
	variable
	logicalOperation
	endOfExpression
)

type token struct {
	value     string
	tokenType tokenType
}

func (t *token) String() string {
	return fmt.Sprintf("Token: %s, Type: %s", t.value, t.tokenType.String())
}

func (t *token) Equal(compare *token) bool {
	return t.value == compare.value && t.tokenType == compare.tokenType
}

type lexer struct {
	input        string
	readPosition int
}

func newLexer(input string) *lexer {
	return &lexer{input: input, readPosition: 0}
}

func (l *lexer) getAllTokens() ([]*token, error) {
	result := make([]*token, 0)

	for {
		token, err := l.nextToken()

		if err != nil {
			return nil, err
		}

		result = append(result, token)

		if token.tokenType == endOfExpression {
			break
		}
	}

	return result, nil
}

func (l *lexer) nextToken() (*token, error) {
	var t *token

	if l.readPosition >= len(l.input) {
		return &token{value: "", tokenType: endOfExpression}, nil
	}

	value := l.input[l.readPosition]

	if value == ' ' || value == '\t' || value == '\r' {
		l.nextPosition()
		return l.nextToken()
	}

	if value == '(' || value == ')' {
		t = &token{value: string(value), tokenType: symbol}
		l.nextPosition()
		return t, nil
	}

	word := l.getNextWord()

	if word == "OR" || word == "AND" {
		t = &token{value: word, tokenType: logicalOperation}
		l.nextPosition(word)
		return t, nil
	}

	t = &token{value: word, tokenType: variable}
	l.nextPosition(word)
	return t, nil
}

func (l *lexer) getNextWord() string {
	partition := l.input[l.readPosition:]

	for i, v := range partition {

		_, ok := allSymbols[string(v)]

		if ok || v == ' ' {
			return partition[:i]
		}
	}

	return partition
}

func (l *lexer) nextPosition(word ...string) {
	if len(word) > 0 {
		l.readPosition += len(word[0])
		return
	}

	l.readPosition++
}
