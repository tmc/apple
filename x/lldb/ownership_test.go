//go:build darwin

package lldb

import (
	"sync"
	"testing"
	"unsafe"
)

func TestCloseAliases(t *testing.T) {
	d, err := Create()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	target := d.CreateTarget("")
	defer target.Close()
	listener := d.Listener()
	defer listener.Close()
	bp := target.BreakpointCreateByName("missing_symbol")
	defer bp.Close()
	proc := target.LaunchSimple(nil, nil, "") // no executable: constructed invalid handle
	defer proc.Close()
	thread := proc.ThreadAtIndex(0)
	defer thread.Close()
	frame := thread.FrameAtIndex(0)
	defer frame.Close()
	value := frame.FindRegister("missing_register")
	defer value.Close()
	e := newError()
	defer e.Close()

	// Method values copy each Go wrapper, just as passing it to another caller
	// does. Concurrent Close calls on those aliases must invoke the real native
	// destructor once, including for constructed invalid handles.
	for _, tt := range []struct {
		name  string
		owner *object
		close func()
		alias func()
	}{
		{"value", value.c, value.Close, value.Close},
		{"frame", frame.c, frame.Close, frame.Close},
		{"thread", thread.c, thread.Close, thread.Close},
		{"error", e.c, e.Close, e.Close},
		{"process", proc.c, proc.Close, proc.Close},
		{"breakpoint", bp.c, bp.Close, bp.Close},
		{"listener", listener.c, listener.Close, listener.Close},
		{"target", target.c, target.Close, target.Close},
		{"debugger", d.c, d.Close, d.Close},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			destroy := tt.owner.destroy
			tt.owner.destroy = func(p unsafe.Pointer) {
				calls++
				destroy(p)
			}
			var wg sync.WaitGroup
			for _, close := range []func(){tt.close, tt.alias, tt.close, tt.alias} {
				wg.Go(close)
			}
			wg.Wait()
			if calls != 1 {
				t.Fatalf("destructor called %d times, want 1", calls)
			}
		})
	}
}

func TestCloseZero(t *testing.T) {
	for _, tt := range []struct {
		name  string
		close func()
	}{
		{"debugger", Debugger{}.Close},
		{"target", Target{}.Close},
		{"listener", Listener{}.Close},
		{"breakpoint", Breakpoint{}.Close},
		{"process", Process{}.Close},
		{"thread", Thread{}.Close},
		{"frame", Frame{}.Close},
		{"value", Value{}.Close},
		{"error", Error{}.Close},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.close()
			tt.close()
		})
	}
}
