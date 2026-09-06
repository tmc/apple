package jaccl

import (
	"errors"
	"testing"
)

func TestDestinationRecordsPreserveRingWireIdentity(t *testing.T) {
	config := Config{
		Rank:        2,
		Size:        3,
		GroupID:     "test",
		Coordinator: "127.0.0.1:1",
		Port:        1,
		Topology:    Ring(3),
		PreferRing:  true,
		Devices: [][][]string{
			{nil, {"r01a", "r01b"}, {"r02a", "r02b"}},
			{{"r10a", "r10b"}, nil, {"r12a", "r12b"}},
			{{"r20a", "r20b"}, {"r21a", "r21b"}, nil},
		},
	}
	records := []destinationRecord{
		{Peer: 1, Direction: -1, Wire: 0, Data: []byte("wire-0")},
		{Peer: 1, Direction: -1, Wire: 1, Data: []byte("wire-1")},
		{Peer: 0, Direction: 1, Wire: 0, Data: []byte("wire-0")},
		{Peer: 0, Direction: 1, Wire: 1, Data: []byte("wire-1")},
	}
	if err := validateDestinationRecords(records, config, 2); err != nil {
		t.Fatal(err)
	}
	records[3].Wire = 0
	if err := validateDestinationRecords(records, config, 2); !errors.Is(err, ErrProtocol) {
		t.Fatalf("duplicate wire error = %v, want ErrProtocol", err)
	}

	inbound := inboundDestinations([][]destinationRecord{
		nil,
		{{Peer: 2, Direction: 1, Wire: 0, Data: []byte("wire-0")}, {Peer: 2, Direction: 1, Wire: 1, Data: []byte("wire-1")}},
	}, 2)
	if len(inbound) != 2 {
		t.Fatalf("inbound = %#v", inbound)
	}
	for wire, record := range inbound {
		if record.Peer != 1 || record.Direction != -1 || record.Wire != wire {
			t.Fatalf("inbound[%d] = %#v", wire, record)
		}
	}
}
