//go:build darwin

package e5rt

import "testing"

// TestOpenResolves checks that the e5rt_* entry points resolve. It calls none
// of them: nothing is compiled, loaded, or dispatched on the Neural Engine.
func TestOpenResolves(t *testing.T) {
	lib, err := Open()
	if lib == nil {
		t.Skipf("Espresso unavailable: %v", err)
	}
	if err != nil {
		t.Logf("partial resolution: %v", err)
	}
	for _, name := range Symbols {
		sym, err := lib.Sym(name)
		if err != nil {
			t.Errorf("Sym(%q) = %v, want an address", name, err)
			continue
		}
		if sym == 0 {
			t.Errorf("Sym(%q) = 0, want a nonzero address", name)
		}
	}
	if got, want := len(lib.Resolved()), len(Symbols); got != want {
		t.Errorf("resolved %d symbols, want %d", got, want)
	}
}

// TestSymMissIsAnError is the negative control for TestOpenResolves: a name
// that cannot exist must be reported as unresolved, so a pass above is
// evidence rather than an instrument that cannot fail.
func TestSymMissIsAnError(t *testing.T) {
	lib, err := Open()
	if lib == nil {
		t.Skipf("Espresso unavailable: %v", err)
	}
	if _, err := lib.Sym("e5rt_this_symbol_does_not_exist"); err == nil {
		t.Fatal("Sym of a nonexistent name succeeded, want an error")
	}
}

// TestLookupReachesOutsideSymbols checks the live-dlsym escape hatch on a name
// this package deliberately does not list, with a negative control.
func TestLookupReachesOutsideSymbols(t *testing.T) {
	lib, err := Open()
	if lib == nil {
		t.Skipf("Espresso unavailable: %v", err)
	}
	// The older submit entry point, which [Lib.SubmitAsync] supersedes: exported
	// by Espresso, used by nothing here, and deliberately absent from Symbols.
	const name = "e5rt_execution_stream_async_submit"
	if _, err := lib.Sym(name); err == nil {
		t.Fatalf("Sym(%q) resolved; the test needs a name outside Symbols", name)
	}
	sym, err := lib.Lookup(name)
	if err != nil {
		t.Fatalf("Lookup(%q) = %v, want an address", name, err)
	}
	if sym == 0 {
		t.Fatalf("Lookup(%q) = 0, want a nonzero address", name)
	}
	if again, err := lib.Lookup(name); err != nil || again != sym {
		t.Errorf("Lookup(%q) again = %#x, %v, want %#x, nil", name, again, err, sym)
	}
	if _, err := lib.Lookup("e5rt_this_symbol_does_not_exist"); err == nil {
		t.Error("Lookup of a nonexistent name succeeded, want an error")
	}
}

// TestAsyncEventCreate drives the parts of the event family that need no model
// and no Neural Engine work. It records two things the recovered header does not
// say.
//
// The third argument of e5rt_async_event_create, which ANEForge always passes as
// 0 and calls an initial value, accepts only 0. Every nonzero value is tried
// under a fresh name, so a rejection is not a name collision, and the names are
// then reused deliberately to show they are labels rather than keys.
func TestAsyncEventCreate(t *testing.T) {
	lib, err := Open()
	if lib == nil {
		t.Skipf("Espresso unavailable: %v", err)
	}
	event, err := lib.AsyncEventCreate("test")
	if err != nil {
		t.Fatalf("AsyncEventCreate: %v", err)
	}
	if event == 0 {
		t.Fatal("AsyncEventCreate returned a null handle with no error")
	}
	defer lib.AsyncEventRelease(event)

	if v, err := lib.AsyncEventLastSignaledValue(event); err != nil || v != 0 {
		t.Errorf("a fresh event reports %d, %v, want 0, nil", v, err)
	}

	// Two live events sharing a name are two events, so the name is not a key.
	second, err := lib.AsyncEventCreate("test")
	if err != nil {
		t.Fatalf("a second event with the same name: %v", err)
	}
	defer lib.AsyncEventRelease(second)
	if second == event {
		t.Errorf("two live events named %q share the handle %#x", "test", event)
	}

	if _, err := lib.AsyncEventCreate(""); err == nil {
		t.Error("AsyncEventCreate with an empty name succeeded, want an error")
	}
}

// TestAsyncEventSignalValue verifies the standalone signal and reader pair.
func TestAsyncEventSignalValue(t *testing.T) {
	lib, err := Open()
	if lib == nil {
		t.Skipf("Espresso unavailable: %v", err)
	}
	event, err := lib.AsyncEventCreate("test")
	if err != nil {
		t.Fatalf("AsyncEventCreate: %v", err)
	}
	defer lib.AsyncEventRelease(event)

	if err := lib.AsyncEventSignal(event, 7); err != nil {
		t.Fatalf("AsyncEventSignal: %v", err)
	}
	if v, err := lib.AsyncEventLastSignaledValue(event); err != nil || v != 7 {
		t.Errorf("after signal the event reports %d, %v, want 7, nil", v, err)
	}
	if err := lib.AsyncEventSyncWait(event, 7, 1_000_000); err != nil {
		t.Errorf("AsyncEventSyncWait for reached value: %v", err)
	}
	if v, err := lib.AsyncEventActiveFutureValue(event); err != nil || v != 0 {
		t.Fatalf("initial active future value = %d, %v, want 0, nil", v, err)
	}
	if err := lib.AsyncEventSetActiveFutureValue(event, 8); err != nil {
		t.Errorf("AsyncEventSetActiveFutureValue: %v", err)
	}
	if v, err := lib.AsyncEventActiveFutureValue(event); err != nil || v != 8 {
		t.Errorf("active future value after setting = %d, %v, want 8, nil", v, err)
	}
	if v, err := lib.AsyncEventLastSignaledValue(event); err != nil || v != 7 {
		t.Errorf("after setting an active future value the event reports %d, %v, want 7, nil", v, err)
	}
}

// TestComputeDeviceMaskRoundTrips reads back each mask it writes. The getter
// and the setter are checked against each other, so this is evidence that the
// setter stores the value it is given and that both take the mask in the
// position this package passes it — a wrong argument order would not round-trip.
//
// The distinct readings are what make it discriminating: a getter that returned
// a constant, or a setter that ignored its argument, would collapse them.
func TestComputeDeviceMaskRoundTrips(t *testing.T) {
	lib, err := Open()
	if lib == nil {
		t.Skipf("Espresso unavailable: %v", err)
	}
	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		t.Fatalf("CompilerOptionsCreate: %v", err)
	}
	defer lib.CompilerOptionsRelease(options)

	seen := make(map[uint64]bool)
	for _, mask := range []uint64{ComputeDeviceANE, ComputeDeviceCPU, ComputeDeviceGPU, ComputeDeviceCPU | ComputeDeviceANE} {
		if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, mask); err != nil {
			t.Fatalf("set mask %#x: %v", mask, err)
		}
		got, err := lib.CompilerOptionsGetComputeDeviceTypesMask(options)
		if err != nil {
			t.Fatalf("get mask after setting %#x: %v", mask, err)
		}
		if got != mask {
			t.Errorf("set %#x, read back %#x", mask, got)
		}
		seen[got] = true
	}
	if len(seen) < 2 {
		t.Errorf("every mask read back as the same value; the getter cannot discriminate")
	}
}

func TestStatusErr(t *testing.T) {
	if err := Status(0).Err("op"); err != nil {
		t.Errorf("Status(0).Err = %v, want nil", err)
	}
	if err := Status(-1).Err("op"); err == nil {
		t.Error("Status(-1).Err = nil, want an error")
	}
}

func TestErrorString(t *testing.T) {
	lib, err := Open()
	if lib == nil {
		t.Skipf("Espresso unavailable: %v", err)
	}
	for status := Status(0); status <= 6; status++ {
		text, err := lib.ErrorString(status)
		if err != nil {
			t.Errorf("ErrorString(%d): %v", status, err)
			continue
		}
		if text == "" {
			t.Errorf("ErrorString(%d) = empty string", status)
		}
	}
}
