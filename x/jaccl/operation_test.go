package jaccl

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"
)

func TestOperationValidation(t *testing.T) {
	tests := []struct {
		name string
		op   operation
		want error
	}{
		{"barrier", operation{Name: "barrier"}, nil},
		{"gather", operation{Name: "all gather", Length: 1}, nil},
		{"reduce", operation{Name: "all reduce", Length: 4, DType: Int32, Reduce: Sum}, nil},
		{"unknown", operation{Name: "send", Length: 1}, ErrProtocol},
		{"negative length", operation{Name: "all gather", Length: -1}, ErrProtocol},
		{"bad dtype", operation{Name: "all reduce", Length: 4, DType: 99, Reduce: Sum}, ErrProtocol},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := encodeOperation(test.op)
			if !errors.Is(err, test.want) {
				t.Fatalf("encodeOperation(%+v) error = %v, want %v", test.op, err, test.want)
			}
		})
	}
}

func TestOperationAgreementRejectsMismatch(t *testing.T) {
	left, right := net.Pipe()
	defer left.Close()
	defer right.Close()
	configs := []Config{
		{Rank: 0, Size: 2, GroupID: "group", Coordinator: "unused", Port: 1, Topology: Mesh(2)},
		{Rank: 1, Size: 2, GroupID: "group", Coordinator: "unused", Port: 1, Topology: Mesh(2)},
	}
	root := &coordinator{incarnation: 1, topology: [][]int{{1}, {0}}, peers: map[int]*controlPeer{1: {conn: left, rank: 1, groupID: "group", incarnation: 1}}}
	peer := &controlPeer{conn: right, rank: 0, groupID: "group", incarnation: 1}
	rootErr := make(chan error, 1)
	go func() { rootErr <- root.agree(context.Background(), configs[0], 1, operation{Name: "barrier"}) }()
	err := peer.agree(context.Background(), configs[1], 1, operation{Name: "all gather", Length: 1})
	if !errors.Is(err, ErrProtocol) {
		t.Fatalf("peer agreement error = %v, want ErrProtocol", err)
	}
	select {
	case err := <-rootErr:
		if !errors.Is(err, ErrProtocol) {
			t.Fatalf("root agreement error = %v, want ErrProtocol", err)
		}
	case <-time.After(time.Second):
		t.Fatal("root mismatch did not return")
	}
}

func TestDecodeOperationRejectsInvalid(t *testing.T) {
	_, err := decodeOperation([]byte(`{"name":"all reduce","length":4,"dtype":99,"reduce":1}`))
	if !errors.Is(err, ErrProtocol) {
		t.Fatalf("decodeOperation error = %v, want ErrProtocol", err)
	}
}
