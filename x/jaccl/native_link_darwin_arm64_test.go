//go:build darwin && arm64

package jaccl

import (
	"context"
	"errors"
	"testing"

	"github.com/tmc/apple/rdma"
)

func TestValidateCompletion(t *testing.T) {
	tests := []struct {
		name       string
		completion rdma.IbvWC
		want       error
	}{
		{"send", rdma.IbvWC{WRID: 1, Opcode: rdma.IBV_WC_SEND, Status: rdma.IBV_WC_SUCCESS}, nil},
		{"receive", rdma.IbvWC{WRID: 2, Opcode: rdma.IBV_WC_RECV, Status: rdma.IBV_WC_SUCCESS, ByteLen: 8}, nil},
		{"stale", rdma.IbvWC{WRID: 3, Opcode: rdma.IBV_WC_SEND, Status: rdma.IBV_WC_SUCCESS}, ErrProtocol},
		{"status", rdma.IbvWC{WRID: 1, Opcode: rdma.IBV_WC_SEND, Status: rdma.IBV_WC_LOC_PROT_ERR}, ErrProtocol},
		{"short", rdma.IbvWC{WRID: 2, Opcode: rdma.IBV_WC_RECV, Status: rdma.IBV_WC_SUCCESS, ByteLen: 7}, ErrProtocol},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wantID, opcode, length := uint64(1), int32(rdma.IBV_WC_SEND), 0
			if test.name == "receive" || test.name == "short" {
				wantID, opcode, length = 2, rdma.IBV_WC_RECV, 8
			}
			err := validateCompletion(test.completion, wantID, opcode, length)
			if !errors.Is(err, test.want) {
				t.Fatalf("validateCompletion error = %v, want %v", err, test.want)
			}
		})
	}
}

func TestNativeMemory(t *testing.T) {
	buffer, err := newNativeMemory(4096)
	if err != nil {
		t.Fatal(err)
	}
	buffer[0] = 1
	buffer[len(buffer)-1] = 2
	if err := releaseNativeMemory(buffer)(); err != nil {
		t.Fatal(err)
	}
}

func TestCompletionFailurePoisonsGroup(t *testing.T) {
	backend := completionFailureBackend{}
	group := &Group{rank: 0, size: 2, backend: backend}
	if err := group.Send(context.Background(), 1, nil); !errors.Is(err, ErrProtocol) {
		t.Fatalf("first send error = %v, want ErrProtocol", err)
	}
	if err := group.Send(context.Background(), 1, nil); !errors.Is(err, ErrPoisoned) {
		t.Fatalf("second send error = %v, want ErrPoisoned", err)
	}
}

type completionFailureBackend struct{}

func (completionFailureBackend) beginClose()  {}
func (completionFailureBackend) close() error { return nil }
func (completionFailureBackend) send(context.Context, int, []byte) error {
	return validateCompletion(rdma.IbvWC{WRID: 1, Opcode: rdma.IBV_WC_RECV, Status: rdma.IBV_WC_SUCCESS, ByteLen: 1}, 1, rdma.IBV_WC_RECV, 2)
}
func (completionFailureBackend) recv(context.Context, int, []byte) error { return nil }
func (completionFailureBackend) barrier(context.Context) error           { return nil }
func (completionFailureBackend) allGather(context.Context, []byte, []byte) error {
	return nil
}
func (completionFailureBackend) allReduce(context.Context, []byte, []byte, DType, ReduceOp) error {
	return nil
}
