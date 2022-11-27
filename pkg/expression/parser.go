package expression

import "errors"

type SymbolTable = map[string]int

type Parser struct {
	ast *ast
}

func NewParser(expression string) (*Parser, error) {
	tokens, err := newLexer(expression).getAllTokens()
	if err != nil {
		return nil, err
	}

	s := newSemantic(tokens)

	tree, err := s.getSemanticTree()
	if err != nil {
		return nil, err
	}

	return &Parser{
		ast: tree,
	}, nil
}

func (p *Parser) Eval(sb SymbolTable) (bool, error) {
	if p.ast == nil {
		return false, errors.New("parser not initialized")
	}

	return p.ast.Evaluate(sb)
}
