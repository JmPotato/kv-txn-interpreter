package interpreter

import "github.com/JmPotato/kv-txn-interpreter/transaction"

// TODO: add more detailed response info.
type TxnResponse interface {
	IsSucceeded() bool
}

// Interpreter is used to interpret the raw Transaction and do the real work inside a KV storage.
type Interpreter interface {
	Commit(txn transaction.Txn) (TxnResponse, error)
}
