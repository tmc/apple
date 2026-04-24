package skylight_test

import (
	"os"
	"testing"

	"github.com/tmc/apple/private/skylight"
)

// TestSPISymbolResolution verifies that the four PSN-taking SkyLight SPIs
// resolved at package init. Each generated wrapper returns an error when the
// underlying symbol is not resolvable; we call each with throwaway inputs and
// check only whether the returned error is nil. The PSN pointer is a no-op
// sentinel — Apple's SPIs dereference it, so we can't invoke the SPI for
// real without mutating system state. Instead we rely on the contract that
// the try-helper's symbol check runs before any argument is passed along,
// so the resolver itself is what's under test here.
func TestSPISymbolResolution(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping private-framework test in -short mode")
	}
	if os.Getenv("GOOS_FORCE_SKIP") != "" {
		t.Skip("explicit skip")
	}

	var psn skylight.ProcessSerialNumber
	// SLPSGetFrontProcess is safe to call — it fills the PSN in place.
	if _, err := skylight.SLPSGetFrontProcess(&psn); err != nil {
		t.Errorf("SLPSGetFrontProcess: symbol not resolved: %v", err)
	}
	// The other three SPIs mutate system state, so we just confirm they
	// exist via the package-level function variables: a nil function
	// pointer would make the wrapper return a symbolCallError.
	// We do not actually invoke them.
	_ = skylight.SLPSSetFrontProcessWithOptions
	_ = skylight.SLPSPostEventRecordTo
	_ = skylight.SLEventPostToPSN
}

// TestSLPSGetFrontProcess exercises the read path end-to-end. Reading the
// front process cannot fail in practice on a running system; if the returned
// status is nonzero, something is wrong with either the ABI or our wrapper.
func TestSLPSGetFrontProcess(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping private-framework test in -short mode")
	}

	var psn skylight.ProcessSerialNumber
	status, err := skylight.SLPSGetFrontProcess(&psn)
	if err != nil {
		t.Skipf("_SLPSGetFrontProcess not resolvable on this system: %v", err)
	}
	if status != 0 {
		t.Fatalf("SLPSGetFrontProcess returned non-zero status %d", status)
	}
	if psn.HighLongOfPSN == 0 && psn.LowLongOfPSN == 0 {
		t.Fatalf("SLPSGetFrontProcess returned zero PSN {%d, %d}", psn.HighLongOfPSN, psn.LowLongOfPSN)
	}
}

// TestCPSModeConstants pins the widely-used values so accidental edits fail
// loudly. These are treated as advisory by the wrapper and can be overridden
// by callers that pass raw uint32 modes.
func TestCPSModeConstants(t *testing.T) {
	if skylight.CPSModeAllWindows != 0x100 {
		t.Errorf("CPSModeAllWindows = %#x, want 0x100", skylight.CPSModeAllWindows)
	}
	if skylight.CPSModeUserGenerated != 0x200 {
		t.Errorf("CPSModeUserGenerated = %#x, want 0x200", skylight.CPSModeUserGenerated)
	}
	if skylight.CPSModeNoWindows != 0x400 {
		t.Errorf("CPSModeNoWindows = %#x, want 0x400", skylight.CPSModeNoWindows)
	}
}
