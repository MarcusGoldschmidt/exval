package expression

import "testing"

func getSemantic(t *testing.T, input string) *semantic {
	tokens, err := newLexer(input).getAllTokens()
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return newSemantic(tokens)
}

func TestSimpleSemantic(t *testing.T) {
	scenarios := []struct {
		input     string
		variables SymbolTable
		expected  bool
	}{
		{
			"(B OR C)",
			SymbolTable{
				"A": 1,
				"B": 0,
				"C": 1,
			},
			true,
		},
	}

	for _, scenario := range scenarios {
		s := getSemantic(t, scenario.input)

		tree, err := s.getSemanticTree()
		if err != nil {
			t.Fatal(err)
		}

		if tree == nil {
			t.Fatal("tree is nil")
		}

		evaluate, err := tree.Evaluate(scenario.variables)
		if err != nil {
			t.Fatal(err)
		}

		if evaluate != scenario.expected {
			t.Error("expected: ", scenario.expected, " but got: ", evaluate)
		}
	}
}

func TestSemantic(t *testing.T) {
	s := getSemantic(t, "((A) OR (C))")

	tree, err := s.getSemanticTree()
	if err != nil {
		t.Fatal(err)
	}

	if tree == nil {
		t.Fatal("tree is nil")
	}
}
