//go:build darwin && arm64

package jaccl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
)

// p2pBroker accepts short-lived control sessions and matches the two endpoints
// of a direct transfer. Payload still goes only over the established UC link.
type p2pBroker struct {
	cfg         Config
	coordinator *coordinator
	closed      <-chan struct{}

	requests chan p2pRequest
	mu       sync.Mutex
	conns    map[net.Conn]bool
}

type p2pRequest struct {
	conn  net.Conn
	frame controlFrame
	op    operation
}

type p2pKey struct {
	source      int
	destination int
	epoch       uint32
}

func newP2PBroker(cfg Config, coordinator *coordinator, closed <-chan struct{}) *p2pBroker {
	b := &p2pBroker{cfg: cfg, coordinator: coordinator, closed: closed, requests: make(chan p2pRequest), conns: make(map[net.Conn]bool)}
	go b.accept()
	go b.match()
	return b
}

func (b *p2pBroker) accept() {
	for {
		conn, err := b.coordinator.listener.Accept()
		if err != nil {
			return
		}
		b.track(conn)
		go b.readRequest(conn)
	}
}

func (b *p2pBroker) local(ctx context.Context) (net.Conn, error) {
	client, server := net.Pipe()
	b.track(server)
	go b.readRequest(server)
	select {
	case <-ctx.Done():
		_ = client.Close()
		return nil, ctx.Err()
	case <-b.closed:
		_ = client.Close()
		return nil, ErrClosed
	default:
		return client, nil
	}
}

func (b *p2pBroker) readRequest(conn net.Conn) {
	frame, err := readControlFrame(conn)
	if err != nil {
		b.untrackClose(conn)
		return
	}
	op, err := decodeOperation(frame.Payload)
	if err != nil || op.Name != "transfer" || op.validateForConfig(b.cfg) != nil || frame.Kind != frameOperation || frame.GroupID != b.cfg.GroupID || frame.Incarnation != b.coordinator.incarnation || frame.Epoch == 0 || (frame.Source != op.Source && frame.Source != op.Destination) {
		b.untrackClose(conn)
		return
	}
	select {
	case b.requests <- p2pRequest{conn: conn, frame: frame, op: op}:
	case <-b.closed:
		b.untrackClose(conn)
	}
}

func (b *p2pBroker) match() {
	pending := make(map[p2pKey]p2pRequest)
	for {
		select {
		case <-b.closed:
			for _, request := range pending {
				b.untrackClose(request.conn)
			}
			return
		case request := <-b.requests:
			key := p2pKey{source: request.op.Source, destination: request.op.Destination, epoch: request.frame.Epoch}
			other, ok := pending[key]
			if !ok {
				pending[key] = request
				continue
			}
			delete(pending, key)
			if err := b.matchPair(other, request); err != nil {
				_ = b.reject(other)
				_ = b.reject(request)
				b.untrackClose(other.conn)
				b.untrackClose(request.conn)
			}
		}
	}
}

func (b *p2pBroker) matchPair(a, z p2pRequest) error {
	if a.op != z.op || !validP2PRole(a) || !validP2PRole(z) || a.frame.Source == z.frame.Source {
		return fmt.Errorf("point-to-point match: %w", ErrProtocol)
	}
	if err := b.ack(a); err != nil {
		return err
	}
	if err := b.ack(z); err != nil {
		return err
	}
	go b.relay(a.conn, z.conn, a.frame.Epoch)
	return nil
}

func validP2PRole(request p2pRequest) bool {
	return request.frame.Source == request.op.Source || request.frame.Source == request.op.Destination
}

func (b *p2pBroker) ack(request p2pRequest) error {
	payload, err := encodeOperation(request.op)
	if err != nil {
		return err
	}
	return b.reply(request, frameOperation, payload)
}

func (b *p2pBroker) reject(request p2pRequest) error {
	return b.reply(request, frameAbort, nil)
}

func (b *p2pBroker) reply(request p2pRequest, kind frameKind, payload []byte) error {
	source := 0
	if request.frame.Source == 0 {
		if request.op.Source == 0 {
			source = request.op.Destination
		} else {
			source = request.op.Source
		}
	}
	return writeControlFrame(request.conn, controlFrame{Magic: protocolMagic, Version: protocolVersion, Kind: kind, GroupID: b.cfg.GroupID, Incarnation: b.coordinator.incarnation, Source: source, Destination: request.frame.Source, Epoch: request.frame.Epoch, Payload: payload})
}

func (b *p2pBroker) relay(a, z net.Conn, epoch uint32) {
	defer b.untrackClose(a)
	defer b.untrackClose(z)
	for {
		left, err := readControlFrame(a)
		if err != nil {
			return
		}
		right, err := readControlFrame(z)
		if err != nil || left.Kind != frameCredit || right.Kind != frameCredit || left.GroupID != b.cfg.GroupID || right.GroupID != b.cfg.GroupID || left.Incarnation != b.coordinator.incarnation || right.Incarnation != b.coordinator.incarnation || left.Epoch != epoch || right.Epoch != epoch {
			return
		}
		if err := writeControlFrame(a, right); err != nil {
			return
		}
		if err := writeControlFrame(z, left); err != nil {
			return
		}
	}
}

func (b *p2pBroker) track(conn net.Conn) {
	b.mu.Lock()
	b.conns[conn] = true
	b.mu.Unlock()
}

func (b *p2pBroker) untrackClose(conn net.Conn) {
	b.mu.Lock()
	delete(b.conns, conn)
	b.mu.Unlock()
	_ = conn.Close()
}

func (b *p2pBroker) close() error {
	if b == nil {
		return nil
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	var err error
	for conn := range b.conns {
		err = errors.Join(err, conn.Close())
		delete(b.conns, conn)
	}
	return err
}

type p2pSession struct {
	conn        net.Conn
	cfg         Config
	peer        int
	epoch       uint32
	incarnation uint64
	mu          sync.Mutex
}

func (b *nativeBackend) openP2P(ctx context.Context, op operation, epoch uint32) (*p2pSession, error) {
	var conn net.Conn
	var err error
	if b.cfg.Rank == 0 {
		if b.p2p == nil {
			return nil, fmt.Errorf("open point-to-point control: %w", ErrClosed)
		}
		conn, err = b.p2p.local(ctx)
	} else {
		conn, err = (&net.Dialer{}).DialContext(ctx, "tcp", b.cfg.Coordinator)
	}
	if err != nil {
		return nil, fmt.Errorf("open point-to-point control: %w", err)
	}
	closeOnError := func(err error) (*p2pSession, error) { _ = conn.Close(); return nil, err }
	payload, err := encodeOperation(op)
	if err != nil {
		return closeOnError(err)
	}
	destination := 0
	if b.cfg.Rank == 0 {
		if op.Source == 0 {
			destination = op.Destination
		} else {
			destination = op.Source
		}
	}
	if err := writeControlFrame(conn, controlFrame{Magic: protocolMagic, Version: protocolVersion, Kind: frameOperation, GroupID: b.cfg.GroupID, Incarnation: b.incarnation(), Source: b.cfg.Rank, Destination: destination, Epoch: epoch, Payload: payload}); err != nil {
		return closeOnError(fmt.Errorf("send point-to-point request: %w", err))
	}
	ack, err := readControlFrameContext(ctx, conn)
	if err != nil {
		return closeOnError(fmt.Errorf("read point-to-point acknowledgement: %w", err))
	}
	if ack.Kind == frameAbort && ack.GroupID == b.cfg.GroupID && ack.Incarnation == b.incarnation() && ack.Destination == b.cfg.Rank && ack.Epoch == epoch {
		return closeOnError(fmt.Errorf("point-to-point rejected by coordinator: %w", ErrProtocol))
	}
	remote, err := decodeOperation(ack.Payload)
	if err != nil || ack.Kind != frameOperation || ack.GroupID != b.cfg.GroupID || ack.Incarnation != b.incarnation() || ack.Destination != b.cfg.Rank || ack.Epoch != epoch || remote != op {
		return closeOnError(fmt.Errorf("point-to-point acknowledgement: %w", ErrProtocol))
	}
	peer := op.Source
	if peer == b.cfg.Rank {
		peer = op.Destination
	}
	return &p2pSession{conn: conn, cfg: b.cfg, peer: peer, epoch: epoch, incarnation: b.incarnation()}, nil
}

func (b *nativeBackend) incarnation() uint64 {
	if b.coordinator != nil {
		return b.coordinator.incarnation
	}
	return b.peer.incarnation
}

func (s *p2pSession) exchangeCredits(ctx context.Context, local []creditRecord) ([]creditRecord, error) {
	if err := validateCreditRecords(local, []int{s.peer}, s.cfg.Size); err != nil {
		return nil, err
	}
	payload, err := json.Marshal(local)
	if err != nil {
		return nil, err
	}
	s.mu.Lock()
	err = writeControlFrame(s.conn, controlFrame{Magic: protocolMagic, Version: protocolVersion, Kind: frameCredit, GroupID: s.cfg.GroupID, Incarnation: s.incarnation, Source: s.cfg.Rank, Destination: s.peer, Epoch: s.epoch, Payload: payload})
	s.mu.Unlock()
	if err != nil {
		return nil, fmt.Errorf("send point-to-point credits: %w", err)
	}
	frame, err := readControlFrameContext(ctx, s.conn)
	if err != nil {
		return nil, fmt.Errorf("read point-to-point credits: %w", err)
	}
	if frame.Kind != frameCredit || frame.GroupID != s.cfg.GroupID || frame.Incarnation != s.incarnation || frame.Source != s.peer || frame.Destination != s.cfg.Rank || frame.Epoch != s.epoch {
		return nil, fmt.Errorf("point-to-point credit frame: %w", ErrProtocol)
	}
	var remote []creditRecord
	if err := json.Unmarshal(frame.Payload, &remote); err != nil {
		return nil, err
	}
	if err := validateCreditRecords(remote, []int{s.cfg.Rank}, s.cfg.Size); err != nil {
		return nil, err
	}
	for i := range remote {
		remote[i].Peer = s.peer
	}
	return remote, nil
}

func (s *p2pSession) close() error { return s.conn.Close() }
