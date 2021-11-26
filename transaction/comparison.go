package transaction

import "github.com/JmPotato/kv-txn-interpreter/pkg"

// CmpTargetType represents the target type for comparison.
type CmpTargetType int

const (
	CmpCreateRevision CmpTargetType = iota
	CmpValue
)

// CmpOp represents the comparison operation.
type CmpOp int

const (
	EQUAL CmpOp = iota
	GREATER
	LESS
	NOT_EQUAL
)

type Comparison struct {
	Op  CmpOp
	Key string

	TargetType CmpTargetType
	Target     interface{}
}

func Value(Key string) *Comparison {
	return &Comparison{Key: Key, TargetType: CmpValue}
}

func CreateRevision(Key string) *Comparison {
	return &Comparison{Key: Key, TargetType: CmpCreateRevision}
}

// Compare is used to create a new comparison operation.
func Compare(cmp *Comparison, opStr string, v interface{}) *Comparison {
	newCmp := &Comparison{}

	switch opStr {
	case "=":
		newCmp.Op = EQUAL
	case "!=":
		newCmp.Op = NOT_EQUAL
	case ">":
		newCmp.Op = GREATER
	case "<":
		newCmp.Op = LESS
	default:
		panic("unknown comparison operator")
	}

	switch cmp.TargetType {
	case CmpValue:
		newCmp.Target = pkg.ConvertToString(v)
	case CmpCreateRevision:
		newCmp.Target = pkg.ConvertToInt64(v)
	default:
		panic("Unknown comparsion target type")
	}

	return newCmp
}
