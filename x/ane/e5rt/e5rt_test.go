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
	const name = "e5rt_async_event_create"
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

func TestStatusErr(t *testing.T) {
	if err := Status(0).Err("op"); err != nil {
		t.Errorf("Status(0).Err = %v, want nil", err)
	}
	if err := Status(-1).Err("op"); err == nil {
		t.Error("Status(-1).Err = nil, want an error")
	}
}
