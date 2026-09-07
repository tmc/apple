package main

import (
	"errors"
	"testing"
)

func TestDecide(t *testing.T) {
	const want = "com.apple.covevm"
	cases := []struct {
		name      string
		self      uint64
		identity  string
		identErr  error
		wantPatch bool
	}{
		{"match", 0x1000, want, nil, true},
		{"null validator", 0, want, nil, false},
		{"unreadable identity", 0x1000, "", errors.New("boom"), false},
		{"mismatch", 0x1000, "com.example.other", nil, false},
		{"empty identity no error", 0x1000, "", nil, false},
		{"null validator beats a match", 0, want, nil, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := decide(c.self, c.identity, c.identErr, want)
			if got.patch != c.wantPatch {
				t.Errorf("decide(self=%#x, %q, %v) patch = %v (%s), want %v",
					c.self, c.identity, c.identErr, got.patch, got.reason, c.wantPatch)
			}
			if got.reason == "" {
				t.Error("decision has no reason")
			}
		})
	}
}
