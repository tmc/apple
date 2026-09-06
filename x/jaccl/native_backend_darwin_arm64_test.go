//go:build darwin && arm64

package jaccl

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestCollectiveWaitsForPeerCompletion(t *testing.T) {
	rootConn, peerConn := net.Pipe()
	defer rootConn.Close()
	defer peerConn.Close()
	configs := []Config{
		{Rank: 0, Size: 2, GroupID: "group", Coordinator: "unused", Port: 1, Topology: Mesh(2)},
		{Rank: 1, Size: 2, GroupID: "group", Coordinator: "unused", Port: 1, Topology: Mesh(2)},
	}
	root := &nativeBackend{
		cfg: configs[0],
		coordinator: &coordinator{incarnation: 1, topology: [][]int{{1}, {0}}, peers: map[int]*controlPeer{
			1: {conn: rootConn, rank: 1, groupID: "group", incarnation: 1},
		}},
		closed: make(chan struct{}),
	}
	peer := &nativeBackend{
		cfg:    configs[1],
		peer:   &controlPeer{conn: peerConn, rank: 0, groupID: "group", incarnation: 1},
		closed: make(chan struct{}),
	}
	entered := make(chan struct{})
	release := make(chan struct{})
	rootResult := make(chan error, 1)
	peerResult := make(chan error, 1)
	go func() {
		rootResult <- root.collective(context.Background(), operation{Name: "barrier"}, func(uint32) error { return nil })
	}()
	go func() {
		peerResult <- peer.collective(context.Background(), operation{Name: "barrier"}, func(uint32) error {
			close(entered)
			<-release
			return nil
		})
	}()
	select {
	case <-entered:
	case <-time.After(time.Second):
		t.Fatal("peer did not enter collective body")
	}
	select {
	case err := <-rootResult:
		t.Fatalf("root returned before peer completion: %v", err)
	default:
	}
	close(release)
	for _, result := range []<-chan error{rootResult, peerResult} {
		select {
		case err := <-result:
			if err != nil {
				t.Fatal(err)
			}
		case <-time.After(time.Second):
			t.Fatal("collective did not complete")
		}
	}
}
