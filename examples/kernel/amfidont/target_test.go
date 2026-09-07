package main

import "testing"

func TestTranslateIMP(t *testing.T) {
	const localIMP = 0x23d7a5e78
	// Equal slides: the target IMP equals the local IMP.
	got, err := translateIMP(localIMP, 0x1a4000, 0x1a4000)
	if err != nil {
		t.Fatalf("equal slides: unexpected error %v", err)
	}
	if got != localIMP {
		t.Errorf("equal slides: got %#x, want %#x", got, localIMP)
	}
	// Differing slides must be an error, not a computed arbitrary address.
	if _, err := translateIMP(localIMP, 0x1a4000, 0x2b0000); err == nil {
		t.Error("differing slides: want error, got nil")
	}
}
