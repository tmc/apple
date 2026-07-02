package security

import "testing"

// TestPackageInitUnderRace guards against a checkptr regression in the generated
// global initializer. The CSSM GUID/OID globals are 12-byte structs, not
// pointers; reading them as *unsafe.Pointer trips the race detector's checkptr
// alignment guard at init time, killing any -race binary that imports this
// package before user code runs. The generator must store the symbol address
// (unsafe.Pointer(ptr)), not dereference its first word.
//
// This test does nothing but exist: building it under `go test -race` runs the
// package init, which is where the fault occurred. Run with -race to exercise.
func TestPackageInitUnderRace(t *testing.T) {
	_ = KSecRandomDefault
}
