package expression

import "testing"

func BenchmarkParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parser, _ := NewParser("((A) OR (C)) OR D")

		_, _ = parser.Eval(map[string]int{
			"A": 1,
			"C": 0,
			"D": 0,
		})
	}
}

func TestSimpleParser(t *testing.T) {
	parser, err := NewParser("(A AND C) OR D")

	if err != nil {
		t.Fatal(err)
	}

	evaluate, err := parser.Eval(map[string]int{
		"A": 1,
		"C": 0,
		"D": 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	if evaluate == false {
		t.Error("evaluate is not 1")
	}
}

func TestParserError(t *testing.T) {
	parser, err := NewParser("A AND B")

	if err != nil {
		t.Fatal(err)
	}

	_, err = parser.Eval(map[string]int{
		"B": 1,
	})

	if err == nil {
		t.Fatal("expected error")
	}
}
