//go:build darwin && arm64

package jaccl

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"
)

func TestP2PBrokerMatchesAndRelaysCredits(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	configs := []Config{
		{Rank: 0, Size: 3, GroupID: "group", Coordinator: listener.Addr().String(), Port: 1, Topology: Mesh(3)},
		{Rank: 1, Size: 3, GroupID: "group", Coordinator: listener.Addr().String(), Port: 1, Topology: Mesh(3)},
	}
	closed := make(chan struct{})
	coordinator := &coordinator{listener: listener, incarnation: 7}
	root := &nativeBackend{cfg: configs[0], coordinator: coordinator, closed: closed}
	root.p2p = newP2PBroker(configs[0], coordinator, closed)
	defer func() {
		close(closed)
		_ = root.p2p.close()
	}()
	remote := &nativeBackend{cfg: configs[1], peer: &controlPeer{rank: 0, groupID: "group", incarnation: 7}, closed: closed}
	op := operation{Name: "transfer", Length: 9, Source: 0, Destination: 1}
	rootSession := make(chan *p2pSession, 1)
	rootErr := make(chan error, 1)
	go func() {
		session, err := root.openP2P(context.Background(), op, 1)
		if err != nil {
			rootErr <- err
			return
		}
		rootSession <- session
	}()
	remoteSession, err := remote.openP2P(context.Background(), op, 1)
	if err != nil {
		t.Fatal(err)
	}
	defer remoteSession.close()
	var local *p2pSession
	select {
	case err := <-rootErr:
		t.Fatal(err)
	case local = <-rootSession:
		defer local.close()
	case <-time.After(time.Second):
		t.Fatal("root point-to-point session did not match")
	}
	credits := make(chan []creditRecord, 1)
	creditErr := make(chan error, 1)
	go func() {
		got, err := local.exchangeCredits(context.Background(), nil)
		if err != nil {
			creditErr <- err
			return
		}
		credits <- got
	}()
	remoteCredits, err := remoteSession.exchangeCredits(context.Background(), []creditRecord{{Peer: 0, Slot: 0}})
	if err != nil {
		t.Fatal(err)
	}
	if len(remoteCredits) != 0 {
		t.Fatalf("remote credits = %#v, want none", remoteCredits)
	}
	select {
	case err := <-creditErr:
		t.Fatal(err)
	case got := <-credits:
		if len(got) != 1 || got[0].Peer != 1 || got[0].Slot != 0 || got[0].Generation != 0 {
			t.Fatalf("root credits = %#v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("point-to-point credit relay did not finish")
	}
}

func TestP2PBrokerMatchesRemotePair(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	address := listener.Addr().String()
	rootConfig := Config{Rank: 0, Size: 3, GroupID: "group", Coordinator: address, Port: 1, Topology: Mesh(3)}
	closed := make(chan struct{})
	coordinator := &coordinator{listener: listener, incarnation: 7}
	broker := newP2PBroker(rootConfig, coordinator, closed)
	defer func() {
		close(closed)
		_ = broker.close()
	}()
	newRemote := func(rank int) *nativeBackend {
		return &nativeBackend{cfg: Config{Rank: rank, Size: 3, GroupID: "group", Coordinator: address, Port: 1, Topology: Mesh(3)}, peer: &controlPeer{rank: 0, groupID: "group", incarnation: 7}, closed: closed}
	}
	source, receiver := newRemote(1), newRemote(2)
	op := operation{Name: "transfer", Length: 9, Source: 1, Destination: 2}
	sourceResult := make(chan *p2pSession, 1)
	sourceErr := make(chan error, 1)
	go func() {
		session, err := source.openP2P(context.Background(), op, 1)
		if err != nil {
			sourceErr <- err
			return
		}
		sourceResult <- session
	}()
	receiverSession, err := receiver.openP2P(context.Background(), op, 1)
	if err != nil {
		t.Fatal(err)
	}
	defer receiverSession.close()
	var sourceSession *p2pSession
	select {
	case err := <-sourceErr:
		t.Fatal(err)
	case sourceSession = <-sourceResult:
		defer sourceSession.close()
	case <-time.After(time.Second):
		t.Fatal("remote point-to-point session did not match")
	}
	credits := make(chan []creditRecord, 1)
	creditErr := make(chan error, 1)
	go func() {
		got, err := sourceSession.exchangeCredits(context.Background(), nil)
		if err != nil {
			creditErr <- err
			return
		}
		credits <- got
	}()
	if _, err := receiverSession.exchangeCredits(context.Background(), []creditRecord{{Peer: 1, Slot: 0}}); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-creditErr:
		t.Fatal(err)
	case got := <-credits:
		if len(got) != 1 || got[0].Peer != 2 {
			t.Fatalf("source credits = %#v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("remote point-to-point credit relay did not finish")
	}
}

func TestP2PBrokerRejectsMismatchedShape(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	address := listener.Addr().String()
	closed := make(chan struct{})
	coordinator := &coordinator{listener: listener, incarnation: 7}
	broker := newP2PBroker(Config{Rank: 0, Size: 3, GroupID: "group", Coordinator: address, Port: 1, Topology: Mesh(3)}, coordinator, closed)
	defer func() {
		close(closed)
		_ = broker.close()
	}()
	newRemote := func(rank int) *nativeBackend {
		return &nativeBackend{cfg: Config{Rank: rank, Size: 3, GroupID: "group", Coordinator: address, Port: 1, Topology: Mesh(3)}, peer: &controlPeer{rank: 0, groupID: "group", incarnation: 7}, closed: closed}
	}
	first, second := newRemote(1), newRemote(2)
	result := make(chan error, 2)
	go func() {
		_, err := first.openP2P(context.Background(), operation{Name: "transfer", Length: 1, Source: 1, Destination: 2}, 1)
		result <- err
	}()
	go func() {
		_, err := second.openP2P(context.Background(), operation{Name: "transfer", Length: 2, Source: 1, Destination: 2}, 1)
		result <- err
	}()
	for range 2 {
		select {
		case err := <-result:
			if !errors.Is(err, ErrProtocol) {
				t.Fatalf("mismatch error = %v, want ErrProtocol", err)
			}
		case <-time.After(time.Second):
			t.Fatal("mismatched point-to-point operation did not return")
		}
	}
}
