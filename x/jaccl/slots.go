package jaccl

import "fmt"

// nativeSlotDepth is deliberately one in the first native transport. A
// one-slot ring makes each advertised UC credit identify the only outstanding
// receive, so the provider's receive-queue order cannot make a credit alias a
// different buffer. Increasing it is a later throughput change that requires
// an explicit receive-order proof.
const nativeSlotDepth = 1

type receiveSlotState uint8

const (
	receiveSlotIdle receiveSlotState = iota
	receiveSlotPosted
	receiveSlotInFlight
	receiveSlotCompleted
	receiveSlotConsumed
)

type receiveSlot struct {
	state      receiveSlotState
	generation uint16
}

// receiveSlots models the ownership side of the UC receive-ring invariant.
// Credits are available only after a successful receive post.
type receiveSlots struct {
	peer  int
	epoch uint32
	slots []receiveSlot

	credits int
}

func newReceiveSlots(peer int, epoch uint32, depth int) (*receiveSlots, error) {
	if peer < 0 || peer > workIDPeerMax {
		return nil, fmt.Errorf("receive peer %d: %w", peer, ErrProtocol)
	}
	if depth <= 0 || depth > workIDSlotMax+1 {
		return nil, fmt.Errorf("receive depth %d: %w", depth, ErrProtocol)
	}
	return &receiveSlots{peer: peer, epoch: epoch, slots: make([]receiveSlot, depth)}, nil
}

func (r *receiveSlots) postInitial() ([]uint64, error) {
	ids := make([]uint64, len(r.slots))
	for slot := range r.slots {
		id, err := r.post(slot)
		if err != nil {
			return nil, err
		}
		ids[slot] = id
	}
	return ids, nil
}

func (r *receiveSlots) post(slot int) (uint64, error) {
	if r == nil || slot < 0 || slot >= len(r.slots) {
		return 0, fmt.Errorf("receive post slot %d: %w", slot, ErrProtocol)
	}
	s := &r.slots[slot]
	if s.state != receiveSlotIdle && s.state != receiveSlotConsumed {
		return 0, fmt.Errorf("receive post slot %d state %d: %w", slot, s.state, ErrProtocol)
	}
	s.state = receiveSlotPosted
	id, err := workID{Epoch: r.epoch, Peer: r.peer, Kind: workRecv, Slot: slot, Generation: s.generation}.encode()
	if err != nil {
		return 0, err
	}
	r.credits++
	return id, nil
}

func (r *receiveSlots) takeCredit() (uint64, error) {
	if r == nil || r.credits == 0 {
		return 0, fmt.Errorf("receive credit unavailable: %w", ErrProtocol)
	}
	for slot := range r.slots {
		s := &r.slots[slot]
		if s.state != receiveSlotPosted {
			continue
		}
		s.state = receiveSlotInFlight
		r.credits--
		return workID{Epoch: r.epoch, Peer: r.peer, Kind: workRecv, Slot: slot, Generation: s.generation}.encode()
	}
	return 0, fmt.Errorf("receive credit has no posted slot: %w", ErrProtocol)
}

func (r *receiveSlots) complete(value uint64) (int, error) {
	if r == nil {
		return 0, fmt.Errorf("receive completion: %w", ErrProtocol)
	}
	id := decodeWorkID(value)
	if id.Epoch != r.epoch || id.Peer != r.peer || id.Kind != workRecv || id.Slot < 0 || id.Slot >= len(r.slots) {
		return 0, fmt.Errorf("receive completion %x: %w", value, ErrProtocol)
	}
	s := &r.slots[id.Slot]
	if s.state != receiveSlotInFlight || s.generation != id.Generation {
		return 0, fmt.Errorf("receive completion slot %d generation %d state %d: %w", id.Slot, id.Generation, s.state, ErrProtocol)
	}
	s.state = receiveSlotCompleted
	return id.Slot, nil
}

func (r *receiveSlots) consume(slot int) error {
	if r == nil || slot < 0 || slot >= len(r.slots) {
		return fmt.Errorf("receive consume slot %d: %w", slot, ErrProtocol)
	}
	s := &r.slots[slot]
	if s.state != receiveSlotCompleted {
		return fmt.Errorf("receive consume slot %d state %d: %w", slot, s.state, ErrProtocol)
	}
	s.state = receiveSlotConsumed
	return nil
}

func (r *receiveSlots) repost(slot int) (uint64, error) {
	if r == nil || slot < 0 || slot >= len(r.slots) {
		return 0, fmt.Errorf("receive repost slot %d: %w", slot, ErrProtocol)
	}
	s := &r.slots[slot]
	if s.state != receiveSlotConsumed {
		return 0, fmt.Errorf("receive repost slot %d state %d before consumption: %w", slot, s.state, ErrProtocol)
	}
	if s.generation == workIDGenerationMax {
		return 0, fmt.Errorf("receive repost slot %d generation exhausted: %w", slot, ErrProtocol)
	}
	s.generation++
	return r.post(slot)
}
