// Copyright 2026 The tmc/apple Authors. All rights reserved.

package xpc

import "testing"

// The omissions ledger is a difference, not a list: it is the members of the
// XPC Swift surface that the generated Go surface does not cover. Whether it is
// COMPLETE depends on the population it was subtracted from, and that
// population is whichever member documents the generating run had loaded. A
// member whose doc page was missing is absent from the ledger and from the
// covered surface alike, and until this file nothing reported the denominator.
//
// This is the third column: the ledger's 11 records are not wrong, but "11" was
// not checked against any stated intent. It is now pinned against the
// population that produced it, so a run that sees a different surface fails
// here instead of quietly re-ledgering.
//
// The figures below were reproduced independently of the generator, by
// classifying the cached XPC doc pages under the three Swift roots
// (XPCListener, XPCSession, XPCReceivedMessage) with the same rules
// classifyMember applies:
//
//	67 member doc pages -> 56 covered, 11 omitted
//
// and the resulting 11 identifiers are set-identical to xpcSwiftOmissions.
//
// A second defensible population gives a different answer, which is why the
// denominator has to be recorded rather than assumed: taking the roots' own
// topicSections identifiers as the population gives 52 members and 10
// omissions. The difference is Apple's disambiguated overloads — three
// hash-suffixed init(endpoint:targetQueue:options:incomingMessageHandler:...)
// pages collapse to one canonical identifier in a topic section. Both numbers
// describe the same surface; they count spellings differently.
const (
	wantSwiftMemberPopulation = 67
	wantSwiftOmissions        = 11
)

func TestSwiftOmissionLedgerPopulation(t *testing.T) {
	// A census that examined nothing passes by never looking.
	if len(xpcSwiftMemberPopulation) == 0 {
		t.Fatalf("xpcSwiftMemberPopulation is empty: the generator emitted no denominator, so no verdict below is about the surface")
	}
	total := 0
	for name, n := range xpcSwiftMemberPopulation {
		if n == 0 {
			t.Errorf("root type %s classified 0 members: its doc pages were not in the loaded set, so any omission of its members is UNMEASURED rather than absent", name)
		}
		total += n
	}
	t.Logf("swift member population: %d members across %d root types; %d omitted", total, len(xpcSwiftMemberPopulation), len(xpcSwiftOmissions))

	if total != wantSwiftMemberPopulation {
		t.Errorf("population is %d members, want %d: the generating run saw a different surface than the one the ledger was checked against, "+
			"so re-derive the ledger rather than trusting its length", total, wantSwiftMemberPopulation)
	}
	if got := len(xpcSwiftOmissions); got != wantSwiftOmissions {
		t.Errorf("ledger has %d records, want %d", got, wantSwiftOmissions)
	}
	if len(xpcSwiftOmissions) > total {
		t.Errorf("ledger has %d records but the population is %d: the ledger cannot exceed what was classified", len(xpcSwiftOmissions), total)
	}

	// Every ledgered omission must belong to one of the root types the
	// population counts, or the two numbers are not about the same surface.
	for _, om := range xpcSwiftOmissions {
		matched := false
		for name := range xpcSwiftMemberPopulation {
			if containsSubstring(om.Identifier, "/XPC"+name+"/") {
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("omission %s belongs to no root type in xpcSwiftMemberPopulation: the ledger and the denominator describe different surfaces", om.Identifier)
		}
	}
}

// TestSwiftOmissionPopulationGateIsSensitive is the mutation control: the same
// checks run against a fabricated population and must report each defect.
func TestSwiftOmissionPopulationGateIsSensitive(t *testing.T) {
	cases := []struct {
		name string
		pop  map[string]int
		want bool // true if the shape must be reported as bad
	}{
		{"empty", map[string]int{}, true},
		{"root with zero members", map[string]int{"Session": 60, "Listener": 0, "ReceivedMessage": 7}, true},
		{"wrong total", map[string]int{"Session": 1, "Listener": 1, "ReceivedMessage": 1}, true},
		{"the real one", xpcSwiftMemberPopulation, false},
	}
	for _, c := range cases {
		bad := len(c.pop) == 0
		total := 0
		for _, n := range c.pop {
			if n == 0 {
				bad = true
			}
			total += n
		}
		if total != wantSwiftMemberPopulation {
			bad = true
		}
		if bad != c.want {
			t.Errorf("%s: reported bad=%v, want %v (total %d)", c.name, bad, c.want, total)
		}
	}
}

func containsSubstring(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
