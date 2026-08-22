package rdma

import (
	"strings"
	"testing"

	"github.com/tmc/apple/rdma"
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
	recordQP := func(name string) func(rdma.RDMAQP) (int32, error) {
		return func(rdma.RDMAQP) (int32, error) {
			order = append(order, name)
			return 0, nil
		}
	}
	recordMR := func(name string) func(rdma.RDMAMR) (int32, error) {
		return func(rdma.RDMAMR) (int32, error) {
			order = append(order, name)
			return 0, nil
		}
	}
	recordCQ := func(name string) func(rdma.RDMACQ) (int32, error) {
		return func(rdma.RDMACQ) (int32, error) {
			order = append(order, name)
			return 0, nil
		}
	}
	recordPD := func(name string) func(rdma.RDMAPD) (int32, error) {
		return func(rdma.RDMAPD) (int32, error) {
			order = append(order, name)
			return 0, nil
		}
	}
	recordCtx := func(name string) func(rdma.RDMAContext) (int32, error) {
		return func(rdma.RDMAContext) (int32, error) {
			order = append(order, name)
			return 0, nil
		}
	}

	oldQ, oldM, oldC, oldP, oldX := destroyQP, deregMR, destroyCQ, deallocPD, closeContext
	defer func() {
		destroyQP, deregMR, destroyCQ, deallocPD, closeContext = oldQ, oldM, oldC, oldP, oldX
	}()

	destroyQP = recordQP("qp")
	deregMR = recordMR("mr")
	destroyCQ = recordCQ("cq")
	deallocPD = recordPD("pd")
	closeContext = recordCtx("context")

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
