package expression

import (
	"fmt"
	"testing"
)

func CompareTokens(result []*token, expected []*token) error {
	if len(result) != len(expected) {
		return fmt.Errorf("expected %d tokens, got %d", len(expected), len(result))
	}

	for i, r := range result {
		if r.Equal(expected[i]) == false {
			return fmt.Errorf("expected %s, got %s at index %d", expected[i].String(), r.String(), i)
		}
	}

	return nil
}

func TestSimpleLexer(t *testing.T) {
	scenarios := []struct {
		input    string
		expected []*token
	}{
		{
			"(A) AND (B OR C)",
			[]*token{
				{value: "(", tokenType: symbol},
				{value: "A", tokenType: variable},
				{value: ")", tokenType: symbol},
				{value: "AND", tokenType: logicalOperation},
				{value: "(", tokenType: symbol},
				{value: "B", tokenType: variable},
				{value: "OR", tokenType: logicalOperation},
				{value: "C", tokenType: variable},
				{value: ")", tokenType: symbol},
				{value: "", tokenType: endOfExpression},
			},
		},
		{
			"(A)",
			[]*token{
				{value: "(", tokenType: symbol},
				{value: "A", tokenType: variable},
				{value: ")", tokenType: symbol},
				{value: "", tokenType: endOfExpression},
			},
		},
		{
			"A)",
			[]*token{
				{value: "A", tokenType: variable},
				{value: ")", tokenType: symbol},
				{value: "", tokenType: endOfExpression},
			},
		},

		{
			") AND (LONGNAMEAAAAA OR C)",
			[]*token{
				{value: ")", tokenType: symbol},
				{value: "AND", tokenType: logicalOperation},
				{value: "(", tokenType: symbol},
				{value: "LONGNAMEAAAAA", tokenType: variable},
				{value: "OR", tokenType: logicalOperation},
				{value: "C", tokenType: variable},
				{value: ")", tokenType: symbol},
				{value: "", tokenType: endOfExpression},
			},
		},
	}

	for _, scenario := range scenarios {
		lexer := newLexer(scenario.input)
		result, err := lexer.getAllTokens()

		if err != nil {
			t.Errorf("Error: %s", err)
		}

		err = CompareTokens(result, scenario.expected)

		if err != nil {
			t.Errorf("error: for scnario %s err: %s", scenario.input, err)
		}
	}

}

func TestGetNextWord(t *testing.T) {
	lexer := newLexer("A)")

	word := lexer.getNextWord()

	if word != "A" {
		t.Errorf("Expected A, got %s", word)
	}
}
