//go:build darwin && arm64

package jaccl

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/tmc/apple/rdma"
)

// nativeBackend owns a fully connected set of native UC links. Its operation
// mutex gives the control connection a single reader and makes epochs globally
// ordered for this rank.
type nativeBackend struct {
	cfg Config

	devices      map[string]*nativeDevice
	links        map[int]*nativeLink
	allLinks     map[linkKey]*nativeLink
	destinations map[linkKey]nativeDestination
	ring         bool
	ringLeft     []*nativeLink
	ringRight    []*nativeLink
	coordinator  *coordinator
	peer         *controlPeer
	p2p          *p2pBroker

	closed      chan struct{}
	closeOnce   sync.Once
	closeErr    error
	operationMu sync.Mutex
	epoch       uint32
	p2pEpoch    map[int]uint32
}

func (b *nativeBackend) beginClose() {
	if b == nil {
		return
	}
	b.closeOnce.Do(func() {
		close(b.closed)
		b.closeErr = b.closeControl()
	})
}

func (b *nativeBackend) close() error {
	if b == nil {
		return nil
	}
	b.beginClose()
	b.operationMu.Lock()
	defer b.operationMu.Unlock()
	var errs []error
	errs = append(errs, b.closeErr)
	for _, key := range sortedLinkKeys(b.allLinks) {
		errs = append(errs, b.allLinks[key].Close())
		b.allLinks[key] = nil
	}
	for _, name := range sortedDeviceNames(b.devices) {
		errs = append(errs, b.devices[name].Close())
		b.devices[name] = nil
	}
	return errors.Join(errs...)
}

func (b *nativeBackend) closeControl() error {
	if b.p2p != nil {
		err := b.p2p.close()
		b.p2p = nil
		if b.coordinator != nil {
			coordinator := b.coordinator
			b.coordinator = nil
			return errors.Join(err, coordinator.close())
		}
		return err
	}
	if b.coordinator != nil {
		err := b.coordinator.close()
		b.coordinator = nil
		return err
	}
	if b.peer != nil {
		err := b.peer.close()
		b.peer = nil
		return err
	}
	return nil
}

func (b *nativeBackend) send(ctx context.Context, dst int, src []byte) error {
	if b == nil || dst == b.cfg.Rank || b.links[dst] == nil {
		return fmt.Errorf("send to rank %d: %w", dst, ErrProtocol)
	}
	op := operation{Name: "transfer", Length: len(src), Source: b.cfg.Rank, Destination: dst}
	return b.pointToPoint(ctx, op, dst, src, nil, true)
}

func (b *nativeBackend) recv(ctx context.Context, src int, dst []byte) error {
	if b == nil || src == b.cfg.Rank || b.links[src] == nil {
		return fmt.Errorf("recv from rank %d: %w", src, ErrProtocol)
	}
	op := operation{Name: "transfer", Length: len(dst), Source: src, Destination: b.cfg.Rank}
	return b.pointToPoint(ctx, op, src, nil, dst, false)
}

func (b *nativeBackend) pointToPoint(ctx context.Context, op operation, peer int, src, dst []byte, sending bool) error {
	if ctx == nil {
		ctx = context.Background()
	}
	b.operationMu.Lock()
	defer b.operationMu.Unlock()
	select {
	case <-b.closed:
		return ErrClosed
	default:
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if b.p2pEpoch[peer] == ^uint32(0) {
		return fmt.Errorf("point-to-point epoch for rank %d exhausted: %w", peer, ErrProtocol)
	}
	b.p2pEpoch[peer]++
	links := b.pointToPointLinks(peer, sending)
	if len(links) == 0 {
		return fmt.Errorf("point-to-point link for rank %d: %w", peer, ErrProtocol)
	}
	if !b.ring || len(links) == 1 {
		return b.pointToPointWire(ctx, op, b.p2pEpoch[peer], peer, links[0], src, dst, sending)
	}
	length := len(src)
	if !sending {
		length = len(dst)
	}
	if length == 0 {
		return b.pointToPointWire(ctx, op, b.p2pEpoch[peer], peer, links[0], src, dst, sending)
	}
	for wire, link := range links {
		start, end := pointToPointWireRange(length, wire, len(links))
		if start == end {
			continue
		}
		wireOp := op
		wireOp.Length = end - start
		if sending {
			if err := b.pointToPointWire(ctx, wireOp, b.p2pEpoch[peer], peer, link, src[start:end], nil, true); err != nil {
				return fmt.Errorf("point-to-point wire %d: %w", wire, err)
			}
		} else if err := b.pointToPointWire(ctx, wireOp, b.p2pEpoch[peer], peer, link, nil, dst[start:end], false); err != nil {
			return fmt.Errorf("point-to-point wire %d: %w", wire, err)
		}
	}
	return nil
}

func (b *nativeBackend) pointToPointWire(ctx context.Context, op operation, epoch uint32, peer int, link *nativeLink, src, dst []byte, sending bool) error {
	session, err := b.openP2P(ctx, op, epoch)
	if err != nil {
		return err
	}
	defer session.close()
	return b.transferLink(ctx, epoch, peer, link, src, dst, sending, session.exchangeCredits)
}

func (b *nativeBackend) pointToPointLinks(peer int, sending bool) []*nativeLink {
	if b == nil || !b.ring || len(b.ringLeft) == 0 || len(b.ringRight) == 0 {
		if b == nil || b.links[peer] == nil {
			return nil
		}
		return []*nativeLink{b.links[peer]}
	}
	left := (b.cfg.Rank + b.cfg.Size - 1) % b.cfg.Size
	right := (b.cfg.Rank + 1) % b.cfg.Size
	if sending {
		// Standalone JACCL intentionally biases a two-rank send toward left.
		if peer == left {
			return b.ringLeft
		}
		if peer == right {
			return b.ringRight
		}
		return nil
	}
	// Standalone JACCL intentionally biases a two-rank receive toward right.
	if peer == right {
		return b.ringRight
	}
	if peer == left {
		return b.ringLeft
	}
	return nil
}

func pointToPointWireRange(length, wire, wires int) (int, int) {
	bytesPerWire := (length + wires - 1) / wires
	start := min(length, wire*bytesPerWire)
	end := min(length, (wire+1)*bytesPerWire)
	return start, end
}

func (b *nativeBackend) barrier(ctx context.Context) error {
	return b.collective(ctx, operation{Name: "barrier"}, func(uint32) error { return nil })
}

func (b *nativeBackend) allGather(ctx context.Context, dst, src []byte) error {
	return b.collective(ctx, operation{Name: "all gather", Length: len(src)}, func(epoch uint32) error {
		if b.cfg.Size == 1 {
			copy(dst, src)
			return nil
		}
		if b.ring {
			return b.ringGather(ctx, epoch, dst, src)
		}
		copy(dst[b.cfg.Rank*len(src):], src)
		return b.exchange(ctx, epoch, src, func(peer, offset, length int, data []byte) error {
			start := peer*len(src) + offset
			copy(dst[start:start+length], data)
			return nil
		})
	})
}

func (b *nativeBackend) allReduce(ctx context.Context, dst, src []byte, dtype DType, op ReduceOp) error {
	return b.collective(ctx, operation{Name: "all reduce", Length: len(src), DType: dtype, Reduce: op}, func(epoch uint32) error {
		if b.cfg.Size == 1 {
			copy(dst, src)
			return nil
		}
		if b.ring {
			gathered := make([]byte, b.cfg.Size*len(src))
			if err := b.ringGather(ctx, epoch, gathered, src); err != nil {
				return err
			}
			return reduceRing(dst, gathered, b.cfg.Size, len(src), dtype, op, len(b.ringLeft))
		}
		gathered := make([]byte, b.cfg.Size*len(src))
		copy(gathered[b.cfg.Rank*len(src):], src)
		if err := b.exchange(ctx, epoch, src, func(peer, offset, length int, data []byte) error {
			start := peer*len(src) + offset
			copy(gathered[start:start+length], data)
			return nil
		}); err != nil {
			return err
		}
		if meshUsesRingReduction(b.cfg.Size, len(src), dtype) {
			return reduceRing(dst, gathered, b.cfg.Size, len(src), dtype, op, 1)
		}
		return reduceRankMajor(dst, gathered, b.cfg.Size, len(src), dtype, op)
	})
}

func (b *nativeBackend) collective(ctx context.Context, op operation, fn func(uint32) error) error {
	if b == nil {
		return ErrClosed
	}
	if ctx == nil {
		ctx = context.Background()
	}
	b.operationMu.Lock()
	defer b.operationMu.Unlock()
	select {
	case <-b.closed:
		return ErrClosed
	default:
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if b.epoch == ^uint32(0) {
		return fmt.Errorf("operation epoch exhausted: %w", ErrProtocol)
	}
	b.epoch++
	epoch := b.epoch
	var err error
	if b.coordinator != nil {
		err = b.coordinator.agree(ctx, b.cfg, epoch, op)
	} else {
		err = b.peer.agree(ctx, b.cfg, epoch, op)
	}
	if err != nil {
		return fmt.Errorf("agree %s epoch %d: %w", op.Name, epoch, err)
	}
	if err := fn(epoch); err != nil {
		return fmt.Errorf("%s epoch %d: %w", op.Name, epoch, err)
	}
	// Agreement before the payload admits matching work. Agreement after it
	// keeps a faster rank from closing its control sessions while a peer is
	// completing the same collective.
	if b.coordinator != nil {
		err = b.coordinator.agree(ctx, b.cfg, epoch, op)
	} else {
		err = b.peer.agree(ctx, b.cfg, epoch, op)
	}
	if err != nil {
		return fmt.Errorf("complete %s epoch %d: %w", op.Name, epoch, err)
	}
	return nil
}

func (b *nativeBackend) exchange(ctx context.Context, epoch uint32, src []byte, consume func(peer, offset, length int, data []byte) error) error {
	if len(src) == 0 {
		return nil
	}
	chunks := (len(src) + nativeSlotBytes - 1) / nativeSlotBytes
	if chunks > workIDGenerationMax+1 {
		return fmt.Errorf("payload length %d needs %d generations: %w", len(src), chunks, ErrProtocol)
	}
	states := make(map[int]*nativeReceiveState, len(b.links))
	credits := make([]creditRecord, 0, len(b.links))
	for _, peer := range sortedLinkPeers(b.links) {
		state, err := b.armReceive(peer, epoch)
		if err != nil {
			return err
		}
		states[peer] = state
		credits = append(credits, state.credit)
	}
	for chunk := 0; chunk < chunks; chunk++ {
		inbound, err := b.exchangeCredits(ctx, epoch, credits)
		if err != nil {
			return err
		}
		if err := b.validateInboundCredits(inbound, uint16(chunk)); err != nil {
			return err
		}
		if chunk > 0 {
			for _, peer := range sortedLinkPeers(b.links) {
				id, err := states[peer].slots.takeCredit()
				if err != nil {
					return fmt.Errorf("reserve receive rank %d chunk %d: %w", peer, chunk, err)
				}
				if id != states[peer].recvID {
					return fmt.Errorf("receive credit rank %d chunk %d id=%#x want=%#x: %w", peer, chunk, id, states[peer].recvID, ErrProtocol)
				}
			}
		}
		offset := chunk * nativeSlotBytes
		length := min(nativeSlotBytes, len(src)-offset)
		for _, peer := range sortedLinkPeers(b.links) {
			id, err := workID{Epoch: epoch, Peer: peer, Kind: workSend, Slot: 0, Generation: uint16(chunk)}.encode()
			if err != nil {
				return err
			}
			if err := b.links[peer].postSend(0, src[offset:offset+length], id); err != nil {
				return fmt.Errorf("post send rank %d chunk %d: %w", peer, chunk, err)
			}
			states[peer].sendID = id
		}
		for _, peer := range sortedLinkPeers(b.links) {
			state := states[peer]
			if err := b.links[peer].waitPair(ctx, b.closed, state.sendID, state.recvID, length); err != nil {
				return fmt.Errorf("wait rank %d chunk %d: %w", peer, chunk, err)
			}
			slot, err := state.slots.complete(state.recvID)
			if err != nil {
				return fmt.Errorf("complete rank %d chunk %d: %w", peer, chunk, err)
			}
			if err := consume(peer, offset, length, b.links[peer].recv[slot*nativeSlotBytes:slot*nativeSlotBytes+length]); err != nil {
				return err
			}
			if err := state.slots.consume(slot); err != nil {
				return fmt.Errorf("consume rank %d chunk %d: %w", peer, chunk, err)
			}
			if chunk+1 < chunks {
				id, err := state.slots.repost(slot)
				if err != nil {
					return fmt.Errorf("repost rank %d chunk %d: %w", peer, chunk, err)
				}
				if err := b.links[peer].postRecv(slot, nativeSlotBytes, id); err != nil {
					return fmt.Errorf("post repost rank %d chunk %d: %w", peer, chunk, err)
				}
				state.recvID = id
				state.credit = creditRecord{Peer: peer, Slot: slot, Generation: decodeWorkID(id).Generation}
			}
		}
		credits = credits[:0]
		if chunk+1 < chunks {
			for _, peer := range sortedLinkPeers(b.links) {
				credits = append(credits, states[peer].credit)
			}
		}
	}
	// The final credit exchange finishes the control-plane handshake without
	// leaving a receive work request that could consume the next operation.
	_, err := b.exchangeCredits(ctx, epoch, credits)
	return err
}

// transfer carries one matched Send or Recv over its directly connected UC
// link. The coordinator brokers only control-session matching and credits;
// payload never traverses it.
func (b *nativeBackend) transferLink(ctx context.Context, epoch uint32, peer int, link *nativeLink, src, dst []byte, sending bool, exchange func(context.Context, []creditRecord) ([]creditRecord, error)) error {
	length := len(src)
	if !sending {
		length = len(dst)
	}
	if length == 0 {
		return nil
	}
	chunks := (length + nativeSlotBytes - 1) / nativeSlotBytes
	if chunks > workIDGenerationMax+1 {
		return fmt.Errorf("transfer length %d needs %d generations: %w", length, chunks, ErrProtocol)
	}
	var state *nativeReceiveState
	var credits []creditRecord
	if !sending {
		var err error
		state, err = b.armReceiveLink(peer, link, epoch)
		if err != nil {
			return err
		}
		credits = []creditRecord{state.credit}
	}
	for chunk := 0; chunk < chunks; chunk++ {
		inbound, err := exchange(ctx, credits)
		if err != nil {
			return err
		}
		offset := chunk * nativeSlotBytes
		chunkLength := min(nativeSlotBytes, length-offset)
		if sending {
			if len(inbound) != 1 || inbound[0].Peer != peer || inbound[0].Slot != 0 || inbound[0].Generation != uint16(chunk) {
				return fmt.Errorf("transfer chunk %d received %#v credits: %w", chunk, inbound, ErrProtocol)
			}
			id, err := workID{Epoch: epoch, Peer: peer, Kind: workSend, Slot: 0, Generation: uint16(chunk)}.encode()
			if err != nil {
				return err
			}
			if err := link.postSend(0, src[offset:offset+chunkLength], id); err != nil {
				return fmt.Errorf("post send chunk %d: %w", chunk, err)
			}
			if err := link.waitCompletion(ctx, b.closed, id, rdma.IBV_WC_SEND, 0); err != nil {
				return fmt.Errorf("wait send chunk %d: %w", chunk, err)
			}
			continue
		}
		if len(inbound) != 0 {
			return fmt.Errorf("receiver got %#v credits: %w", inbound, ErrProtocol)
		}
		if chunk > 0 {
			id, err := state.slots.takeCredit()
			if err != nil || id != state.recvID {
				return fmt.Errorf("reserve receive chunk %d: %w", chunk, errors.Join(err, ErrProtocol))
			}
		}
		if err := link.waitCompletion(ctx, b.closed, state.recvID, rdma.IBV_WC_RECV, chunkLength); err != nil {
			return fmt.Errorf("wait receive chunk %d: %w", chunk, err)
		}
		slot, err := state.slots.complete(state.recvID)
		if err != nil {
			return fmt.Errorf("complete receive chunk %d: %w", chunk, err)
		}
		copy(dst[offset:offset+chunkLength], link.recv[slot*nativeSlotBytes:slot*nativeSlotBytes+chunkLength])
		if err := state.slots.consume(slot); err != nil {
			return fmt.Errorf("consume receive chunk %d: %w", chunk, err)
		}
		if chunk+1 < chunks {
			id, err := state.slots.repost(slot)
			if err != nil {
				return fmt.Errorf("repost receive chunk %d: %w", chunk, err)
			}
			if err := link.postRecv(slot, nativeSlotBytes, id); err != nil {
				return fmt.Errorf("post receive chunk %d: %w", chunk, err)
			}
			state.recvID = id
			credits = []creditRecord{{Peer: peer, Slot: slot, Generation: decodeWorkID(id).Generation}}
		} else {
			credits = nil
		}
	}
	_, err := exchange(ctx, credits)
	return err
}

type nativeReceiveState struct {
	slots  *receiveSlots
	recvID uint64
	sendID uint64
	credit creditRecord
}

func (b *nativeBackend) armReceive(peer int, epoch uint32) (*nativeReceiveState, error) {
	link := b.links[peer]
	if link == nil {
		return nil, fmt.Errorf("receive link to rank %d: %w", peer, ErrProtocol)
	}
	return b.armReceiveLink(peer, link, epoch)
}

func (b *nativeBackend) armReceiveLink(peer int, link *nativeLink, epoch uint32) (*nativeReceiveState, error) {
	slots, err := newReceiveSlots(peer, epoch, nativeSlotDepth)
	if err != nil {
		return nil, err
	}
	ids, err := slots.postInitial()
	if err != nil {
		return nil, err
	}
	if len(ids) != 1 {
		return nil, fmt.Errorf("receive slots depth %d: %w", len(ids), ErrProtocol)
	}
	if err := link.postRecv(0, nativeSlotBytes, ids[0]); err != nil {
		return nil, fmt.Errorf("post initial receive rank %d: %w", peer, err)
	}
	if _, err := slots.takeCredit(); err != nil {
		return nil, fmt.Errorf("reserve initial receive rank %d: %w", peer, err)
	}
	return &nativeReceiveState{
		slots: slots, recvID: ids[0],
		credit: creditRecord{Peer: peer, Slot: 0, Generation: decodeWorkID(ids[0]).Generation},
	}, nil
}

func (b *nativeBackend) exchangeCredits(ctx context.Context, epoch uint32, credits []creditRecord) ([]creditRecord, error) {
	if b.coordinator != nil {
		return b.coordinator.exchangeCredits(ctx, b.cfg, epoch, credits)
	}
	return b.peer.exchangeCredits(ctx, b.cfg, epoch, credits)
}

func (b *nativeBackend) validateInboundCredits(credits []creditRecord, generation uint16) error {
	if len(credits) != len(b.links) {
		return fmt.Errorf("received %d credits, want %d: %w", len(credits), len(b.links), ErrProtocol)
	}
	seen := make(map[int]bool, len(credits))
	for _, credit := range credits {
		if b.links[credit.Peer] == nil || credit.Slot != 0 || credit.Generation != generation || seen[credit.Peer] {
			return fmt.Errorf("credit from rank %d slot %d: %w", credit.Peer, credit.Slot, ErrProtocol)
		}
		seen[credit.Peer] = true
	}
	return nil
}

func sortedLinkPeers(links map[int]*nativeLink) []int {
	peers := make([]int, 0, len(links))
	for peer := range links {
		peers = append(peers, peer)
	}
	sort.Ints(peers)
	return peers
}

func sortedLinkKeys(links map[linkKey]*nativeLink) []linkKey {
	keys := make([]linkKey, 0, len(links))
	for key := range links {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Peer != keys[j].Peer {
			return keys[i].Peer < keys[j].Peer
		}
		if keys[i].Direction != keys[j].Direction {
			return keys[i].Direction < keys[j].Direction
		}
		return keys[i].Wire < keys[j].Wire
	})
	return keys
}

func sortedDeviceNames(devices map[string]*nativeDevice) []string {
	names := make([]string, 0, len(devices))
	for name := range devices {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (b *nativeBackend) primaryDevice() *nativeDevice {
	if b == nil || len(b.devices) != 1 {
		return nil
	}
	for _, device := range b.devices {
		return device
	}
	return nil
}
