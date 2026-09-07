//go:build darwin

package objcinspect_test

import (
	"errors"
	"reflect"
	"testing"

	basepurego "github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcinspect"
)

// loadFoundation makes NSString available. objc.String resolves the NSString
// class once, on its first call anywhere in the process, so any test that
// builds a string must ensure Foundation is loaded before that first call.
func loadFoundation(t *testing.T) objc.ID {
	t.Helper()
	_, err := basepurego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation",
		basepurego.RTLD_LAZY|basepurego.RTLD_GLOBAL)
	if err != nil {
		t.Fatalf("load Foundation: %v", err)
	}
	s := objc.ID(objc.String("hello"))
	if s == 0 {
		t.Fatal("Foundation NSString unavailable")
	}
	return s
}

func TestSendCheckedDisabledPassesThrough(t *testing.T) {
	s := loadFoundation(t)
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(false)) // force off, restore prior

	n, err := objcinspect.SendChecked[uint64](s, objc.Sel("length"), "unexpected arg")
	if err != nil {
		t.Fatalf("disabled SendChecked returned error: %v", err)
	}
	if n != 5 {
		t.Fatalf("length = %d, want 5", n)
	}
}

func TestSendCheckedGoodCall(t *testing.T) {
	s := loadFoundation(t)
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))

	n, err := objcinspect.SendChecked[uint64](s, objc.Sel("length"))
	if err != nil {
		t.Fatalf("good call errored: %v", err)
	}
	if n != 5 {
		t.Fatalf("length = %d, want 5", n)
	}
}

func TestSendCheckedCatchesArity(t *testing.T) {
	s := loadFoundation(t)
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))

	_, err := objcinspect.SendChecked[uint64](s, objc.Sel("length"), "extra")
	if err == nil {
		t.Fatal("expected an arity error, got nil (call would have dispatched)")
	}
	t.Logf("caught before dispatch: %v", err)
}

func TestSendCheckedCatchesFloatIntCrossing(t *testing.T) {
	s := loadFoundation(t)
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))

	// length returns an integer (NSUInteger); asking for it as float64 would
	// read the result from the wrong register file.
	_, err := objcinspect.SendChecked[float64](s, objc.Sel("length"))
	if err == nil {
		t.Fatal("expected a float/int register mismatch, got nil")
	}
	t.Logf("caught register-file mismatch: %v", err)
}

func TestSendCheckedCatchesMissingSelector(t *testing.T) {
	s := loadFoundation(t)
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))

	_, err := objcinspect.SendChecked[uint64](s, objc.Sel("noSuchSelectorXYZ"))
	if !errors.Is(err, objcinspect.ErrSelectorNotFound) {
		t.Fatalf("error = %v, want ErrSelectorNotFound", err)
	}
	t.Logf("caught missing selector: %v", err)
}

func TestCallCheckedVoidPath(t *testing.T) {
	s := loadFoundation(t)
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))

	// hash returns a value but we ignore it via CallChecked; no return-type
	// constraint means no false positive here.
	if err := objcinspect.CallChecked(s, objc.Sel("hash")); err != nil {
		t.Fatalf("CallChecked errored: %v", err)
	}
}

// TestConformanceSweep shows the "moved left" use: a table of a package's
// intended call sites, verified against the live runtime in one test.
func TestConformanceSweep(t *testing.T) {
	s := loadFoundation(t)
	u := reflect.TypeFor[uint64]()

	good := []objcinspect.Call{
		{ID: s, Sel: objc.Sel("length"), Return: u},
		{ID: s, Sel: objc.Sel("hash"), Return: u},
		{ID: s, Sel: objc.Sel("isEqualToString:"), Return: reflect.TypeFor[bool](), Args: []any{s}},
	}
	if err := objcinspect.CheckAll(good...); err != nil {
		t.Fatalf("conformance sweep of valid sites failed: %v", err)
	}

	bad := []objcinspect.Call{
		{ID: s, Sel: objc.Sel("length"), Return: u, Args: []any{"extra"}},    // wrong arity
		{ID: s, Sel: objc.Sel("length"), Return: reflect.TypeFor[float64]()}, // wrong register file
	}
	if err := objcinspect.CheckAll(bad...); err == nil {
		t.Fatal("conformance sweep passed sites that should fail")
	} else {
		t.Logf("sweep flagged bad sites:\n%v", err)
	}
}

func TestCallCheckedRejectsBeforeDispatch(t *testing.T) {
	s := loadFoundation(t)
	defer objcinspect.SetEnabled(objcinspect.SetEnabled(true))
	if err := objcinspect.CallChecked(s, objc.Sel("noSuchSelectorXYZ")); !errors.Is(err, objcinspect.ErrSelectorNotFound) {
		t.Fatalf("error = %v, want ErrSelectorNotFound", err)
	}
}

func TestCheckAllEmpty(t *testing.T) {
	if err := objcinspect.CheckAll(); err != nil {
		t.Fatal(err)
	}
}
