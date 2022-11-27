package expression

import (
	"fmt"
)

type ast struct {
	value any
	left  *ast
	right *ast
}

func (a *ast) Evaluate(sb SymbolTable) (bool, error) {
	if v, ok := a.value.(int); ok {
		return v != 0, nil
	}

	if v, ok := a.value.(string); ok {
		if operator, ok := allLogicalOperations[v]; ok {
			left, err := a.left.Evaluate(sb)

			if err != nil {
				return false, err
			}

			right, err := a.right.Evaluate(sb)

			if err != nil {
				return false, err
			}

			switch operator {
			case OR:
				return left || right, nil
			case AND:
				return left && right, nil
			}
		}

		if valueSymbol, ok := sb[v]; ok {
			return valueSymbol != 0, nil
		}

		return false, fmt.Errorf("symbol %s not found", v)
	}

	return false, fmt.Errorf("invalid value %v", a.value)
}
