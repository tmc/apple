package main

import "testing"

// TestResolveAMFI re-verifies, against the live Objective-C runtime in our own
// process, the pinned class/selector/ivar facts this tool patches blind to.
// This is a safe, read-only probe requiring no privileges. If Apple ships a
// build whose layout differs, resolveAMFI returns an error naming the ivar, and
// this test fails loudly rather than letting the tool write to a wrong offset.
func TestResolveAMFI(t *testing.T) {
	a, err := resolveAMFI()
	if err != nil {
		t.Fatalf("resolveAMFI: %v", err)
	}
	if a.imp == 0 {
		t.Fatal("validateWithError: IMP is nil")
	}
	for _, c := range []struct {
		name string
		got  int64
		want int64
	}{
		{"_isValid", a.offIsValid, ivarIsValid},
		{"_isApple", a.offIsApple, ivarIsApple},
		{"_shouldUnrestrict", a.offShouldUnrestrict, ivarShouldUnrestrict},
		{"_hasRestrictedEntitlements", a.offHasRestrictedEntitlements, ivarHasRestrictedEntitlements},
		{"_signingIdentifier", a.offSigningIdentifier, ivarSigningIdentifier},
	} {
		if c.got != c.want {
			t.Errorf("ivar %s offset = %d, want pinned %d", c.name, c.got, c.want)
		}
	}
	t.Logf("live: %s", a.summary())
	t.Logf("live: bootarg state resolved=%v raw=%#x amfi_get_out_of_my_way=%v", a.boot.resolved, a.boot.raw, a.boot.getOutOfMyWay)
}
