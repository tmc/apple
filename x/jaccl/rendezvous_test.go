package jaccl

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"
)

func TestRendezvousHandshake(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	address := listener.Addr().String()
	configs := []Config{
		{Rank: 0, Size: 2, GroupID: "group", Coordinator: address, Port: 1, Topology: Mesh(2)},
		{Rank: 1, Size: 2, GroupID: "group", Coordinator: address, Port: 1, Topology: Mesh(2)},
	}
	roots := make(chan *coordinator, 1)
	rootErr := make(chan error, 1)
	go func() {
		coordinator, err := listenRendezvousWithListener(context.Background(), configs[0], listener)
		if err != nil {
			rootErr <- err
			return
		}
		roots <- coordinator
	}()
	peer, err := dialRendezvous(context.Background(), configs[1])
	if err != nil {
		t.Fatal(err)
	}
	defer peer.close()
	select {
	case err := <-rootErr:
		t.Fatal(err)
	case root := <-roots:
		defer root.close()
		if root.incarnation == 0 {
			t.Fatal("incarnation is zero")
		}
		if peer.incarnation != root.incarnation {
			t.Fatalf("peer incarnation = %d, want %d", peer.incarnation, root.incarnation)
		}
		rootDestinations := make(chan []destinationRecord, 1)
		rootDestinationErr := make(chan error, 1)
		go func() {
			destinations, err := root.exchangeDestinations(context.Background(), configs[0], []destinationRecord{{Peer: 1, Data: []byte("from-zero")}})
			if err != nil {
				rootDestinationErr <- err
				return
			}
			rootDestinations <- destinations
		}()
		clientDestinations, err := peer.exchangeDestinations(context.Background(), configs[1], []destinationRecord{{Peer: 0, Data: []byte("from-one")}})
		if err != nil {
			t.Fatal(err)
		}
		select {
		case err := <-rootDestinationErr:
			t.Fatal(err)
		case destinations := <-rootDestinations:
			if len(destinations) != 1 || destinations[0].Peer != 1 || string(destinations[0].Data) != "from-one" {
				t.Fatalf("root destinations = %#v", destinations)
			}
		case <-time.After(time.Second):
			t.Fatal("root destination exchange did not finish")
		}
		if len(clientDestinations) != 1 || clientDestinations[0].Peer != 0 || string(clientDestinations[0].Data) != "from-zero" {
			t.Fatalf("client destinations = %#v", clientDestinations)
		}
		readyErr := make(chan error, 1)
		go func() { readyErr <- root.ready(context.Background(), configs[0]) }()
		if err := peer.ready(context.Background(), configs[1]); err != nil {
			t.Fatal(err)
		}
		select {
		case err := <-readyErr:
			if err != nil {
				t.Fatal(err)
			}
		case <-time.After(time.Second):
			t.Fatal("ready handshake did not finish")
		}
		operationErr := make(chan error, 1)
		operation := operation{Name: "all reduce", Length: 8, DType: Int32, Reduce: Sum}
		go func() { operationErr <- root.agree(context.Background(), configs[0], 1, operation) }()
		if err := peer.agree(context.Background(), configs[1], 1, operation); err != nil {
			t.Fatal(err)
		}
		select {
		case err := <-operationErr:
			if err != nil {
				t.Fatal(err)
			}
		case <-time.After(time.Second):
			t.Fatal("operation agreement did not finish")
		}
		rootCredits := make(chan []creditRecord, 1)
		creditErr := make(chan error, 1)
		go func() {
			credits, err := root.exchangeCredits(context.Background(), configs[0], 1, []creditRecord{{Peer: 1, Slot: 0}})
			if err != nil {
				creditErr <- err
				return
			}
			rootCredits <- credits
		}()
		clientCredits, err := peer.exchangeCredits(context.Background(), configs[1], 1, []creditRecord{{Peer: 0, Slot: 0}})
		if err != nil {
			t.Fatal(err)
		}
		if len(clientCredits) != 1 || clientCredits[0].Peer != 0 || clientCredits[0].Slot != 0 {
			t.Fatalf("client credits = %#v", clientCredits)
		}
		select {
		case err := <-creditErr:
			t.Fatal(err)
		case credits := <-rootCredits:
			if len(credits) != 1 || credits[0].Peer != 1 || credits[0].Slot != 0 {
				t.Fatalf("root credits = %#v", credits)
			}
		case <-time.After(time.Second):
			t.Fatal("credit exchange did not finish")
		}
	case <-time.After(time.Second):
		t.Fatal("coordinator did not finish")
	}
}

func TestReciprocalTopology(t *testing.T) {
	if err := validateReciprocalTopology([][]int{{1}, {0, 2}, {1}}); err != nil {
		t.Fatal(err)
	}
	if err := validateReciprocalTopology([][]int{{1}, {}, {}}); !errors.Is(err, ErrProtocol) {
		t.Fatalf("asymmetric topology error = %v, want ErrProtocol", err)
	}
}
