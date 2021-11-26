package transaction

type Txn interface {
	If(cs ...Comparison) Txn
	Then(ops ...Operation) Txn
	Else(ops ...Operation) Txn
}

// Transaction is used to build a structure to represent a transaction,
// which will be interpreted later by a certain KV storage to run the transaction.
type Transaction struct {
	ifFlag   bool
	thenFlag bool
	elseFlag bool
	Ifs      []*Comparison
	Thens    []*Operation
	Elses    []*Operation
}

// NewTransaction creates a new Transaction.
func NewTransaction() Txn {
	return &Transaction{}
}

func (txn *Transaction) If(cs ...Comparison) Txn {
	if txn.ifFlag {
		panic("Transaction.If() can not be called twice!")
	}
	if txn.thenFlag {
		panic("Transaction.If() can not be called after Transaction.Then()!")
	}
	if txn.elseFlag {
		panic("Transaction.If() can not be called after Transaction.Else()!")
	}
	txn.ifFlag = true
	for _, c := range cs {
		txn.Ifs = append(txn.Ifs, &c)
	}
	return txn
}

func (txn *Transaction) Then(ops ...Operation) Txn {
	if txn.thenFlag {
		panic("Transaction.Then() can not be called twice!")
	}
	if txn.elseFlag {
		panic("Transaction.Then() can not be called after Transaction.Else()!")
	}
	txn.thenFlag = true
	for _, op := range ops {
		txn.Thens = append(txn.Thens, &op)
	}
	return txn
}

func (txn *Transaction) Else(ops ...Operation) Txn {
	if txn.elseFlag {
		panic("Transaction.Else() can not be called twice!")
	}
	txn.elseFlag = true
	for _, op := range ops {
		txn.Elses = append(txn.Elses, &op)
	}
	return txn
}
