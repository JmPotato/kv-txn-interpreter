package interpreter

import (
	"context"
	"fmt"

	"github.com/JmPotato/kv-txn-interpreter/transaction"
	"go.etcd.io/etcd/clientv3"
)

type EtcdTxnResponse clientv3.TxnResponse

func (etr *EtcdTxnResponse) IsSucceeded() bool {
	return etr.Succeeded
}

type EtcdTxnInterpreter struct {
	ctx    context.Context
	client *clientv3.Client
}

// NewEtcdTxnInterpreter create a new etcd transaction interpreter with the given etcd client.
func NewEtcdTxnInterpreter(ctx context.Context, client *clientv3.Client) Interpreter {
	return &EtcdTxnInterpreter{ctx, client}
}

func (eti *EtcdTxnInterpreter) Commit(txn transaction.Txn) (TxnResponse, error) {
	rawTransaction, ok := txn.(*transaction.Transaction)
	if !ok {
		return nil, fmt.Errorf("invalid txn")
	}
	etcdTxn := eti.client.Txn(eti.ctx)
	// Initialize the If comparsion.
	cmps := make([]clientv3.Cmp, 0, len(rawTransaction.Ifs))
	for _, c := range rawTransaction.Ifs {
		cmps = append(cmps, convertComparsionToEtcdCmp(c))
	}
	etcdTxn = etcdTxn.If(cmps...)
	// Initialize the Then operations.
	thenOps := make([]clientv3.Op, 0, len(rawTransaction.Thens))
	for _, op := range rawTransaction.Thens {
		thenOps = append(thenOps, convertOperationToEtcdOp(op))
	}
	etcdTxn = etcdTxn.Then(thenOps...)
	// Initialize the Else operations.
	elseOps := make([]clientv3.Op, 0, len(rawTransaction.Elses))
	for _, op := range rawTransaction.Elses {
		elseOps = append(elseOps, convertOperationToEtcdOp(op))
	}
	etcdTxn = etcdTxn.Else(elseOps...)
	resp, err := etcdTxn.Commit()
	return (*EtcdTxnResponse)(resp), err
}

func convertComparsionToEtcdCmp(comparsion *transaction.Comparison) clientv3.Cmp {
	var cmpType clientv3.Cmp
	switch comparsion.TargetType {
	case transaction.CmpCreateRevision:
		cmpType = clientv3.CreateRevision(comparsion.Key)
	case transaction.CmpValue:
		cmpType = clientv3.Value(comparsion.Key)
	}
	var cmp clientv3.Cmp
	switch comparsion.Op {
	case transaction.EQUAL:
		cmp = clientv3.Compare(cmpType, "=", comparsion.Target)
	case transaction.NOT_EQUAL:
		cmp = clientv3.Compare(cmpType, "!=", comparsion.Target)
	case transaction.GREATER:
		cmp = clientv3.Compare(cmpType, ">", comparsion.Target)
	case transaction.LESS:
		cmp = clientv3.Compare(cmpType, "<", comparsion.Target)
	}
	return cmp
}

func convertOperationToEtcdOp(operation *transaction.Operation) clientv3.Op {
	var op clientv3.Op
	switch operation.Ty {
	case transaction.OpRange:
		op = clientv3.OpGet(operation.Key)
	case transaction.OpPut:
		op = clientv3.OpPut(operation.Key, operation.Value)
	case transaction.OpDelete:
		op = clientv3.OpDelete(operation.Key)
	}
	return op
}
