package jaccl

import "fmt"

const (
	workIDGenerationBits = 16
	workIDSlotBits       = 2
	workIDKindBits       = 2
	workIDPeerBits       = 12
	workIDEpochBits      = 32
)

const (
	workIDGenerationMax = 1<<workIDGenerationBits - 1
	workIDSlotMax       = 1<<workIDSlotBits - 1
	workIDPeerMax       = 1<<workIDPeerBits - 1
)

type workKind uint8

const (
	workSend workKind = iota + 1
	workRecv
)

type workID struct {
	Epoch      uint32
	Peer       int
	Kind       workKind
	Slot       int
	Generation uint16
}

func (id workID) encode() (uint64, error) {
	if id.Peer < 0 || id.Peer > workIDPeerMax {
		return 0, fmt.Errorf("work peer %d: %w", id.Peer, ErrProtocol)
	}
	if id.Kind != workSend && id.Kind != workRecv {
		return 0, fmt.Errorf("work kind %d: %w", id.Kind, ErrProtocol)
	}
	if id.Slot < 0 || id.Slot > workIDSlotMax {
		return 0, fmt.Errorf("work slot %d: %w", id.Slot, ErrProtocol)
	}
	if id.Generation > workIDGenerationMax {
		return 0, fmt.Errorf("work generation %d: %w", id.Generation, ErrProtocol)
	}
	return uint64(id.Epoch)<<32 |
		uint64(id.Peer)<<20 |
		uint64(id.Kind)<<18 |
		uint64(id.Slot)<<16 |
		uint64(id.Generation), nil
}

func decodeWorkID(value uint64) workID {
	return workID{
		Epoch:      uint32(value >> 32),
		Peer:       int(value>>20) & workIDPeerMax,
		Kind:       workKind(value>>18) & 3,
		Slot:       int(value>>16) & workIDSlotMax,
		Generation: uint16(value & workIDGenerationMax),
	}
}
