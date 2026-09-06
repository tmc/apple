package jaccl

import (
	"errors"
	"testing"
)

func TestReceiveSlotsRequireConsumptionBeforeRepost(t *testing.T) {
	r, err := newReceiveSlots(1, 3, 2)
	if err != nil {
		t.Fatal(err)
	}
	ids, err := r.postInitial()
	if err != nil {
		t.Fatal(err)
	}
	if r.credits != 2 {
		t.Fatalf("credits = %d, want 2", r.credits)
	}
	if _, err := r.takeCredit(); err != nil {
		t.Fatal(err)
	}
	if _, err := r.takeCredit(); err != nil {
		t.Fatal(err)
	}
	if _, err := r.takeCredit(); !errors.Is(err, ErrProtocol) {
		t.Fatalf("third credit error = %v, want ErrProtocol", err)
	}
	slot, err := r.complete(ids[0])
	if err != nil || slot != 0 {
		t.Fatalf("complete = %d, %v, want 0, nil", slot, err)
	}
	if _, err := r.repost(slot); !errors.Is(err, ErrProtocol) {
		t.Fatalf("premature repost error = %v, want ErrProtocol", err)
	}
	if err := r.consume(slot); err != nil {
		t.Fatal(err)
	}
	next, err := r.repost(slot)
	if err != nil {
		t.Fatal(err)
	}
	if next == ids[0] {
		t.Fatalf("reposted work id = %x, want a new generation", next)
	}
	if _, err := r.takeCredit(); err != nil {
		t.Fatalf("credit after successful repost: %v", err)
	}
}

func TestReceiveSlotsRejectsStaleCompletion(t *testing.T) {
	r, err := newReceiveSlots(1, 3, 1)
	if err != nil {
		t.Fatal(err)
	}
	ids, err := r.postInitial()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := r.takeCredit(); err != nil {
		t.Fatal(err)
	}
	if _, err := r.complete(ids[0]); err != nil {
		t.Fatal(err)
	}
	if err := r.consume(0); err != nil {
		t.Fatal(err)
	}
	if _, err := r.repost(0); err != nil {
		t.Fatal(err)
	}
	if _, err := r.complete(ids[0]); !errors.Is(err, ErrProtocol) {
		t.Fatalf("stale completion error = %v, want ErrProtocol", err)
	}
}
