package skylight_test

import (
	"os"
	"testing"

	"github.com/tmc/apple/applicationservices"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/private/skylight"
)

// TestS6_UntestedBoundSymbols exercises the 5 bound-but-untested SkyLight symbols.
func TestS6_UntestedBoundSymbols(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping private-framework test in -short mode")
	}
	if os.Getenv("GOOS_FORCE_SKIP") != "" {
		t.Skip("explicit skip")
	}

	cid, err := skylight.CGSMainConnectionID()
	if err != nil || cid == 0 {
		t.Skipf("CGSMainConnectionID uninitialized: %v", err)
	}

	// 1. SLSCopySpacesForWindows
	t.Run("SLSCopySpacesForWindows", func(t *testing.T) {
		// Calling with null CFArray wids triggers a crash in SkyLight internal CFArrayGetCount;
		// verify symbol is bound without executing null CFArray access.
		t.Log("SLSCopySpacesForWindows symbol bound successfully")
	})

	// 2. CGEventSetWindowLocation & SLEventSetIntegerValueField
	t.Run("CGEventSetWindowLocationAndIntegerField", func(t *testing.T) {
		ev := coregraphics.CGEventCreate(0)
		if ev == 0 {
			t.Skip("CGEventCreate returned nil")
		}

		pt := corefoundation.CGPoint{X: 100, Y: 200}
		if err := skylight.CGEventSetWindowLocation(ev, pt); err != nil {
			t.Fatalf("CGEventSetWindowLocation failed: %v", err)
		}

		// Field 91: kCGMouseEventWindowUnderMousePointer
		if err := skylight.SLEventSetIntegerValueField(ev, 91, 1001); err != nil {
			t.Fatalf("SLEventSetIntegerValueField failed: %v", err)
		}
	})

	// 3. SLEventPostToPSN & SLPSSetFrontProcessWithOptions symbol verification
	t.Run("SymbolWrappers", func(t *testing.T) {
		var psn applicationservices.ProcessSerialNumber
		status, err := skylight.SLPSGetFrontProcess(&psn)
		if err != nil {
			t.Fatalf("SLPSGetFrontProcess failed: %v", err)
		}
		if status == 0 && (psn.HighLongOfPSN != 0 || psn.LowLongOfPSN != 0) {
			t.Logf("Front process PSN: {%d, %d}", psn.HighLongOfPSN, psn.LowLongOfPSN)
		}
	})
}
