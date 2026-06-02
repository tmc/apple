package rdma

import (
	"strings"
	"testing"
)

func TestQPChainCloseZero(t *testing.T) {
	var chain QPChain
	if err := chain.Close(); err != nil {
		t.Fatalf("Close zero: %v", err)
	}
	if err := (*QPChain)(nil).Close(); err != nil {
		t.Fatalf("Close nil: %v", err)
	}
}

func TestQPChainCloseOrderAndIdempotence(t *testing.T) {
	var order []string
	record := func(name string) func(uintptr) (int, error) {
		return func(uintptr) (int, error) {
			order = append(order, name)
			return 0, nil
		}
	}
	defer func(q, m, c, p, x func(uintptr) (int, error)) {
		destroyQP, deregMR, destroyCQ, deallocPD, closeContext = q, m, c, p, x
	}(destroyQP, deregMR, destroyCQ, deallocPD, closeContext)

	destroyQP = record("qp")
	deregMR = record("mr")
	destroyCQ = record("cq")
	deallocPD = record("pd")
	closeContext = record("context")

	chain := &QPChain{Context: 1, PD: 2, MR: 3, CQ: 4, QP: 5}
	if err := chain.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if got, want := strings.Join(order, ","), "qp,mr,cq,pd,context"; got != want {
		t.Fatalf("close order = %s, want %s", got, want)
	}
	if chain.Context != 0 || chain.PD != 0 || chain.MR != 0 || chain.CQ != 0 || chain.QP != 0 {
		t.Fatalf("handles not cleared after Close: %+v", chain)
	}
	// A second Close must not re-invoke any provider (double-free safety).
	if err := chain.Close(); err != nil {
		t.Fatalf("second Close: %v", err)
	}
	if got, want := strings.Join(order, ","), "qp,mr,cq,pd,context"; got != want {
		t.Fatalf("second Close changed order to %s, want %s", got, want)
	}
}
