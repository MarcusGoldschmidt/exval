package expression

//go:generate stringer -type=LogicalOperations -output=logical_operations_string.go
type LogicalOperations int

const (
	OR LogicalOperations = iota
	AND
)

var allLogicalOperations = map[string]LogicalOperations{
	"OR":  OR,
	"AND": AND,
}

var allSymbols = map[string]struct{}{
	"(": {},
	")": {},
}
