package jaccl

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

type testBackend struct {
	closed chan struct{}
	once   sync.Once
}

func newTestBackend() *testBackend {
	return &testBackend{closed: make(chan struct{})}
}

func (b *testBackend) beginClose()  { b.once.Do(func() { close(b.closed) }) }
func (b *testBackend) close() error { return nil }
func (b *testBackend) send(context.Context, int, []byte) error {
	<-b.closed
	return ErrClosed
}
func (b *testBackend) recv(context.Context, int, []byte) error { return nil }
func (b *testBackend) barrier(context.Context) error           { return nil }
func (b *testBackend) allGather(context.Context, []byte, []byte) error {
	return nil
}
func (b *testBackend) allReduce(context.Context, []byte, []byte, DType, ReduceOp) error {
	return nil
}

func TestCloseInterruptsOperation(t *testing.T) {
	backend := newTestBackend()
	g := &Group{rank: 0, size: 2, backend: backend}
	returned := make(chan error, 1)
	go func() { returned <- g.Send(context.Background(), 1, []byte("x")) }()

	select {
	case <-time.After(time.Second):
		t.Fatal("send did not reach backend")
	case <-time.After(time.Millisecond):
	}

	closed := make(chan error, 1)
	go func() { closed <- g.Close() }()
	select {
	case err := <-returned:
		if !errors.Is(err, ErrClosed) {
			t.Fatalf("send error = %v, want ErrClosed", err)
		}
	case <-time.After(time.Second):
		t.Fatal("send did not return after Close")
	}
	select {
	case err := <-closed:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("Close did not return")
	}
}

func TestGroupArgumentValidation(t *testing.T) {
	g := &Group{rank: 0, size: 2, backend: newTestBackend()}
	if err := g.Send(context.Background(), 0, nil); !errors.Is(err, ErrProtocol) {
		t.Fatalf("self send error = %v, want ErrProtocol", err)
	}
	if err := g.Recv(context.Background(), 2, nil); !errors.Is(err, ErrProtocol) {
		t.Fatalf("bad receive error = %v, want ErrProtocol", err)
	}
	if err := g.AllGather(context.Background(), make([]byte, 3), make([]byte, 2)); !errors.Is(err, ErrProtocol) {
		t.Fatalf("all gather error = %v, want ErrProtocol", err)
	}
	if err := g.AllReduce(context.Background(), make([]byte, 2), make([]byte, 2), Int32, Sum); !errors.Is(err, ErrProtocol) {
		t.Fatalf("all reduce error = %v, want ErrProtocol", err)
	}
	data := make([]byte, 8)
	if err := g.AllReduce(context.Background(), data[1:5], data[:4], Uint8, Sum); !errors.Is(err, ErrProtocol) {
		t.Fatalf("partially overlapping all reduce error = %v, want ErrProtocol", err)
	}
	if err := g.AllReduce(context.Background(), data[:4], data[:4], Uint8, Sum); err != nil {
		t.Fatalf("in-place all reduce error = %v", err)
	}
	var nilGroup *Group
	if err := nilGroup.Send(context.Background(), 0, nil); !errors.Is(err, ErrClosed) {
		t.Fatalf("nil group send error = %v, want ErrClosed", err)
	}
}

func TestGroupPoisonsAfterBackendFailure(t *testing.T) {
	backend := newTestBackend()
	g := &Group{rank: 0, size: 2, backend: backend}
	backend.beginClose()
	if err := g.Send(context.Background(), 1, nil); !errors.Is(err, ErrClosed) {
		t.Fatalf("send error = %v, want ErrClosed", err)
	}

	failed := &failingBackend{}
	g = &Group{rank: 0, size: 2, backend: failed}
	if err := g.Send(context.Background(), 1, nil); !errors.Is(err, ErrProtocol) {
		t.Fatalf("first send error = %v, want ErrProtocol", err)
	}
	if err := g.Send(context.Background(), 1, nil); !errors.Is(err, ErrPoisoned) {
		t.Fatalf("second send error = %v, want ErrPoisoned", err)
	}
}

type failingBackend struct{}

func (failingBackend) beginClose()                                     {}
func (failingBackend) close() error                                    { return nil }
func (failingBackend) send(context.Context, int, []byte) error         { return ErrProtocol }
func (failingBackend) recv(context.Context, int, []byte) error         { return nil }
func (failingBackend) barrier(context.Context) error                   { return nil }
func (failingBackend) allGather(context.Context, []byte, []byte) error { return nil }
func (failingBackend) allReduce(context.Context, []byte, []byte, DType, ReduceOp) error {
	return nil
}
