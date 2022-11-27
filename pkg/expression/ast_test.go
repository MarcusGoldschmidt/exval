package expression

import "testing"

func TestEvaluate(t *testing.T) {
	scenarios := []struct {
		ast      *ast
		symbol   SymbolTable
		expected bool
	}{
		{
			&ast{
				value: "OR",
				left: &ast{
					value: 1,
				},
				right: &ast{
					value: 0,
				},
			},
			SymbolTable{},
			true,
		},
		{
			&ast{
				value: "OR",
				left: &ast{
					value: "A",
				},
				right: &ast{
					value: "B",
				},
			},
			SymbolTable{
				"A": 1,
				"B": 0,
			},
			true,
		},
		{
			&ast{
				value: "AND",
				left: &ast{
					value: "OR",
					left: &ast{
						value: 1,
					},
					right: &ast{
						value: 0,
					},
				},
				right: &ast{
					value: "B",
				},
			},
			SymbolTable{
				"B": 0,
			},
			false,
		},
	}

	for _, scenario := range scenarios {
		result, err := scenario.ast.Evaluate(scenario.symbol)

		if err != nil {
			t.Fatal(err)
		}

		if result != scenario.expected {
			t.Fatalf("expected %v, got %v", scenario.expected, result)
		}
	}

}
