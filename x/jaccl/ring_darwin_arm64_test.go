//go:build darwin && arm64

package jaccl

import "testing"

func TestRingWireRange(t *testing.T) {
	tests := []struct {
		length int
		wires  int
		want   [][2]int
	}{
		{length: 0, wires: 2, want: [][2]int{{0, 0}, {0, 0}}},
		{length: 1, wires: 4, want: [][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 1}}},
		{length: 10, wires: 3, want: [][2]int{{0, 3}, {3, 6}, {6, 10}}},
	}
	for _, test := range tests {
		for wire, want := range test.want {
			start, end := ringWireRange(test.length, wire, test.wires)
			if start != want[0] || end != want[1] {
				t.Fatalf("range(%d, %d, %d) = %d, %d, want %d, %d", test.length, wire, test.wires, start, end, want[0], want[1])
			}
		}
	}
}

func TestPrimaryLinksPreservesTwoRankDirections(t *testing.T) {
	left := &nativeLink{}
	right := &nativeLink{}
	links, ringLeft, ringRight, err := primaryLinks(Config{Rank: 0, Size: 2, GroupID: "test", Coordinator: "127.0.0.1:1", Port: 1, Topology: Ring(2)}, map[linkKey]*nativeLink{
		{Peer: 1, Direction: -1, Wire: 0}: left,
		{Peer: 1, Direction: 1, Wire: 0}:  right,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(ringLeft) != 1 || ringLeft[0] != left || len(ringRight) != 1 || ringRight[0] != right {
		t.Fatalf("ring links = left %#v right %#v", ringLeft, ringRight)
	}
	if links[1] != right {
		t.Fatalf("public two-rank link = %p, want right %p", links[1], right)
	}
	backend := &nativeBackend{cfg: Config{Rank: 0, Size: 2}, ring: true, links: links, ringLeft: ringLeft, ringRight: ringRight}
	if got := backend.pointToPointLinks(1, true); len(got) != 1 || got[0] != left {
		t.Fatalf("two-rank send links = %#v, want left", got)
	}
	if got := backend.pointToPointLinks(1, false); len(got) != 1 || got[0] != right {
		t.Fatalf("two-rank recv links = %#v, want right", got)
	}
}

func TestPointToPointWireRange(t *testing.T) {
	want := [][2]int{{0, 4}, {4, 8}, {8, 10}}
	for wire, bounds := range want {
		start, end := pointToPointWireRange(10, wire, len(want))
		if start != bounds[0] || end != bounds[1] {
			t.Fatalf("range(10, %d, 3) = %d, %d, want %d, %d", wire, start, end, bounds[0], bounds[1])
		}
	}
}
