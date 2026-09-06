package jaccl

import (
	"errors"
	"testing"
)

func TestWorkIDRoundTrip(t *testing.T) {
	want := workID{Epoch: 17, Peer: 9, Kind: workRecv, Slot: 3, Generation: workIDGenerationMax}
	value, err := want.encode()
	if err != nil {
		t.Fatal(err)
	}
	if got := decodeWorkID(value); got != want {
		t.Fatalf("decodeWorkID(%#x) = %+v, want %+v", value, got, want)
	}
}

func TestWorkIDRejectsOutOfRange(t *testing.T) {
	tests := []workID{
		{Peer: workIDPeerMax + 1, Kind: workSend},
		{Peer: 1, Kind: 0},
		{Peer: 1, Kind: workSend, Slot: workIDSlotMax + 1},
	}
	for _, id := range tests {
		if _, err := id.encode(); !errors.Is(err, ErrProtocol) {
			t.Fatalf("encode(%+v) error = %v, want ErrProtocol", id, err)
		}
	}
}
