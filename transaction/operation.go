package transaction

type OperationType int

const (
	OpRange OperationType = iota
	OpPut
	OpDelete
)

// Operation is the key-value operation inside a transaction .
type Operation struct {
	Ty OperationType

	Key   string
	Value string
}

func Get(key string) *Operation {
	return &Operation{
		Ty:  OpRange,
		Key: key,
	}
}

func Put(key, val string) *Operation {
	return &Operation{
		Ty:    OpPut,
		Key:   key,
		Value: val,
	}
}

func Delete(key string) *Operation {
	return &Operation{
		Ty:  OpDelete,
		Key: key,
	}
}
