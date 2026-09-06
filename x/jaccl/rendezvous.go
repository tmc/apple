package jaccl

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"sync"
)

type helloPayload struct {
	Size  int   `json:"size"`
	Peers []int `json:"peers"`
}

type controlPeer struct {
	conn net.Conn
	mu   sync.Mutex

	rank        int
	groupID     string
	incarnation uint64
}

func (p *controlPeer) send(f controlFrame) error {
	if p == nil || p.conn == nil {
		return fmt.Errorf("control peer: %w", ErrClosed)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	return writeControlFrame(p.conn, f)
}

func (p *controlPeer) close() error {
	if p == nil || p.conn == nil {
		return nil
	}
	return p.conn.Close()
}

// coordinator owns control connections accepted by rank zero. It is setup-only
// at this stage: later link and operation messages use the same framed peers.
type coordinator struct {
	listener    net.Listener
	peers       map[int]*controlPeer
	incarnation uint64
	topology    [][]int
}

func (c *coordinator) close() error {
	if c == nil {
		return nil
	}
	var stack closeStack
	if c.listener != nil {
		stack.add(c.listener.Close)
	}
	for _, peer := range c.peers {
		stack.add(peer.close)
	}
	return stack.close()
}

func listenRendezvous(ctx context.Context, cfg Config) (*coordinator, error) {
	if cfg.Rank != 0 {
		return nil, fmt.Errorf("listen rendezvous rank %d: %w", cfg.Rank, ErrProtocol)
	}
	listener, err := net.Listen("tcp", cfg.Coordinator)
	if err != nil {
		return nil, fmt.Errorf("listen coordinator %q: %w", cfg.Coordinator, err)
	}
	return listenRendezvousWithListener(ctx, cfg, listener)
}

func listenRendezvousWithListener(ctx context.Context, cfg Config, listener net.Listener) (*coordinator, error) {
	if cfg.Rank != 0 || listener == nil {
		return nil, fmt.Errorf("listen rendezvous: %w", ErrProtocol)
	}
	incarnation, err := randomIncarnation()
	if err != nil {
		_ = listener.Close()
		return nil, err
	}
	coordinator := &coordinator{listener: listener, peers: make(map[int]*controlPeer), incarnation: incarnation}
	localPeers, _ := cfg.validate()
	topology := make([][]int, cfg.Size)
	topology[0] = localPeers
	for len(coordinator.peers) != cfg.Size-1 {
		conn, err := acceptContext(ctx, listener)
		if err != nil {
			_ = coordinator.close()
			return nil, err
		}
		peer, hello, err := acceptHello(conn, cfg)
		if err != nil {
			_ = conn.Close()
			_ = coordinator.close()
			return nil, err
		}
		if coordinator.peers[peer.rank] != nil {
			_ = conn.Close()
			_ = coordinator.close()
			return nil, fmt.Errorf("duplicate rank %d: %w", peer.rank, ErrProtocol)
		}
		topology[peer.rank] = hello.Peers
		coordinator.peers[peer.rank] = peer
	}
	if err := validateReciprocalTopology(topology); err != nil {
		_ = coordinator.close()
		return nil, err
	}
	coordinator.topology = topology
	for rank, peer := range coordinator.peers {
		if err := peer.send(controlFrame{
			Magic:       protocolMagic,
			Version:     protocolVersion,
			Kind:        frameHello,
			GroupID:     cfg.GroupID,
			Incarnation: incarnation,
			Source:      0,
			Destination: rank,
		}); err != nil {
			_ = coordinator.close()
			return nil, fmt.Errorf("acknowledge rank %d: %w", rank, err)
		}
	}
	return coordinator, nil
}

func dialRendezvous(ctx context.Context, cfg Config) (*controlPeer, error) {
	if cfg.Rank == 0 {
		return nil, fmt.Errorf("dial rendezvous rank zero: %w", ErrProtocol)
	}
	dialer := net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", cfg.Coordinator)
	if err != nil {
		return nil, fmt.Errorf("dial coordinator %q: %w", cfg.Coordinator, err)
	}
	peers, _ := cfg.validate()
	payload, err := json.Marshal(helloPayload{Size: cfg.Size, Peers: peers})
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("encode rendezvous hello: %w", err)
	}
	if err := writeControlFrame(conn, controlFrame{
		Magic:       protocolMagic,
		Version:     protocolVersion,
		Kind:        frameHello,
		GroupID:     cfg.GroupID,
		Source:      cfg.Rank,
		Destination: 0,
		Payload:     payload,
	}); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("send rendezvous hello: %w", err)
	}
	ack, err := readControlFrameContext(ctx, conn)
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("read rendezvous acknowledgement: %w", err)
	}
	if ack.Kind != frameHello || ack.GroupID != cfg.GroupID || ack.Source != 0 || ack.Destination != cfg.Rank || ack.Incarnation == 0 {
		_ = conn.Close()
		return nil, fmt.Errorf("rendezvous acknowledgement: %w", ErrProtocol)
	}
	return &controlPeer{conn: conn, rank: 0, groupID: cfg.GroupID, incarnation: ack.Incarnation}, nil
}

func acceptHello(conn net.Conn, cfg Config) (*controlPeer, helloPayload, error) {
	hello, err := readControlFrame(conn)
	if err != nil {
		return nil, helloPayload{}, fmt.Errorf("read rendezvous hello: %w", err)
	}
	if hello.Kind != frameHello || hello.GroupID != cfg.GroupID || hello.Destination != 0 || hello.Source <= 0 || hello.Source >= cfg.Size || hello.Incarnation != 0 {
		return nil, helloPayload{}, fmt.Errorf("rendezvous hello: %w", ErrProtocol)
	}
	var payload helloPayload
	if err := json.Unmarshal(hello.Payload, &payload); err != nil {
		return nil, helloPayload{}, fmt.Errorf("decode rendezvous hello: %w", err)
	}
	if payload.Size != cfg.Size || !validPeerSet(payload.Peers, hello.Source, cfg.Size) {
		return nil, helloPayload{}, fmt.Errorf("rendezvous hello topology: %w", ErrProtocol)
	}
	return &controlPeer{conn: conn, rank: hello.Source, groupID: cfg.GroupID}, payload, nil
}

func acceptContext(ctx context.Context, listener net.Listener) (net.Conn, error) {
	type result struct {
		conn net.Conn
		err  error
	}
	accepted := make(chan result, 1)
	go func() {
		conn, err := listener.Accept()
		accepted <- result{conn: conn, err: err}
	}()
	select {
	case result := <-accepted:
		if result.err != nil {
			return nil, fmt.Errorf("accept coordinator connection: %w", result.err)
		}
		return result.conn, nil
	case <-ctx.Done():
		_ = listener.Close()
		return nil, ctx.Err()
	}
}

func validateReciprocalTopology(topology [][]int) error {
	for rank, peers := range topology {
		if !validPeerSet(peers, rank, len(topology)) {
			return fmt.Errorf("topology rank %d: %w", rank, ErrProtocol)
		}
		for _, peer := range peers {
			if !containsRank(topology[peer], rank) {
				return fmt.Errorf("topology %d -> %d is not reciprocal: %w", rank, peer, ErrProtocol)
			}
		}
	}
	return nil
}

func validPeerSet(peers []int, rank, size int) bool {
	if rank < 0 || rank >= size {
		return false
	}
	seen := make(map[int]bool, len(peers))
	for _, peer := range peers {
		if peer < 0 || peer >= size || peer == rank || seen[peer] {
			return false
		}
		seen[peer] = true
	}
	return true
}

func containsRank(peers []int, want int) bool {
	for _, peer := range peers {
		if peer == want {
			return true
		}
	}
	return false
}

func sortedPeers(peers []int) []int {
	copy := append([]int(nil), peers...)
	sort.Ints(copy)
	return copy
}

func randomIncarnation() (uint64, error) {
	var data [8]byte
	if _, err := cryptorand.Read(data[:]); err != nil {
		return 0, fmt.Errorf("generate group incarnation: %w", err)
	}
	value := uint64(0)
	for _, b := range data {
		value = value<<8 | uint64(b)
	}
	if value == 0 {
		value = 1
	}
	return value, nil
}
