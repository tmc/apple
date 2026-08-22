//go:build darwin

package objc

import (
	"os"
	"regexp"
	"testing"
)

// The objc_slowpath tag disables the fast path by controlling whether
// initFastSend is called: fastsend_enabled.gen.go calls it, and
// fastsend_disabled.gen.go does not. An unconditional init in objc.gen.go
// defeats that, because it runs under every tag combination.
//
// This has regressed once already. The tag was made functional by hand-editing
// the unconditional init out of the generated objc.gen.go; the template still
// emitted it, so the next regeneration restored it and the tag silently went
// inert while both build-tagged files stayed in place, looking correct.
//
// The check is on the source rather than on behavior because a behavioral
// assertion only holds in one tag world, and the world where this breaks is
// the default one, where nobody passes the tag.
func TestFastSendInitIsBuildTagged(t *testing.T) {
	src, err := os.ReadFile("objc.gen.go")
	if err != nil {
		t.Fatalf("read generated source: %v", err)
	}
	unconditional := regexp.MustCompile(`(?m)^func init\(\) \{\n\tinitFastSend\(\)\n\}`)
	if unconditional.Match(src) {
		t.Error("objc.gen.go calls initFastSend from an unconditional init; " +
			"the objc_slowpath build tag is inert and the fast path is armed " +
			"under every tag. It belongs in fastsend_enabled.gen.go.")
	}
	// Guard the guard: if initFastSend is renamed or removed, the pattern above
	// stops matching for a reason that has nothing to do with the defect.
	if !regexp.MustCompile(`func initFastSend\(\)`).Match(src) {
		t.Fatal("initFastSend is not defined in objc.gen.go; this test no longer " +
			"checks what it claims to")
	}
}
