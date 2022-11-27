package expression

import (
	"errors"
)

type semantic struct {
	tokens  []*token
	pointer int
}

func newSemantic(tokens []*token) *semantic {
	return &semantic{tokens: tokens, pointer: 0}
}

func (s *semantic) peekToken(next ...int) *token {
	sum := 0

	if len(next) > 0 {
		sum = next[0] - 1
	}

	return s.tokens[s.pointer+sum]
}

func (s *semantic) popToken() *token {
	token := s.tokens[s.pointer]
	s.pointer++
	return token
}

func (s *semantic) getSemanticTree() (*ast, error) {
	if len(s.tokens) == 0 {
		return nil, errors.New("expression empty")
	}

	exp, err := s.expression()
	if err != nil {
		return nil, err
	}

	right, err := s.continueExpression()
	if err != nil {
		return nil, err
	}

	if right != nil {
		right.left = exp

		return right, nil
	}

	return exp, nil
}

func (s *semantic) expression() (*ast, error) {
	nextToken := s.popToken()

	if nextToken.value == "(" {
		expression, err := s.expression()
		if err != nil {
			return nil, err
		}

		nextToken = s.peekToken()

		if nextToken.tokenType == logicalOperation {
			expressionContinue, err := s.continueExpression()
			if err != nil {
				return nil, err
			}

			if expressionContinue == nil {
				return &ast{value: nextToken.value}, nil
			}

			expressionContinue.left = expression

			return expressionContinue, nil
		}

		if nextToken.value != ")" {
			return nil, errors.New("expected )")
		}
		s.popToken()

		return expression, nil
	}

	if nextToken.tokenType != variable {
		return nil, errors.New("expected variable")
	}

	expression, err := s.continueExpression()
	if err != nil {
		return nil, err
	}

	if expression == nil {
		return &ast{value: nextToken.value}, nil
	}

	expression.left = &ast{value: nextToken.value}

	return expression, nil
}
func (s *semantic) continueExpression() (*ast, error) {
	operator := s.peekToken()

	if operator.tokenType != logicalOperation {
		return nil, nil
	}
	s.popToken()

	exp, err := s.expression()
	if err != nil {
		return nil, err
	}

	right, err := s.continueExpression()
	if err != nil {
		return nil, err
	}

	if right != nil {
		right.left = exp

		return &ast{
			value: operator.value,
			right: right,
			left:  nil,
		}, nil
	}

	return &ast{
		value: operator.value,
		right: exp,
		left:  nil,
	}, nil
}
