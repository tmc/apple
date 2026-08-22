package mach

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// The bootstrap server (launchd) is the rendezvous namespace: two unrelated
// processes have no shared port namespace, so one registers a send right
// under a name and the other looks it up. bootstrap.h is not a documented
// framework surface, so these are bound here rather than generated.

var (
	bootstrapOnce sync.Once
	bootstrapErr  error

	bootstrapPort     Port
	bootstrapLookUp   func(bp uint32, name string, sp *uint32) int32
	bootstrapRegister func(bp uint32, name string, sp uint32) int32
)

func loadBootstrap() error {
	bootstrapOnce.Do(func() {
		sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "bootstrap_port")
		if err != nil || sym == 0 {
			bootstrapErr = fmt.Errorf("mach: resolve bootstrap_port: %w", err)
			return
		}
		// bootstrap_port is a global mach_port_t variable, not a function:
		// dlsym returns the variable's address. The address is C static
		// data, not a Go pointer, so the uintptr round-trip is not a GC
		// hazard.
		addr := sym
		bootstrapPort = Port(**(**uint32)(unsafe.Pointer(&addr)))
		for _, s := range []struct {
			name string
			fn   any
		}{
			{"bootstrap_look_up", &bootstrapLookUp},
			{"bootstrap_register", &bootstrapRegister},
		} {
			fsym, err := purego.Dlsym(purego.RTLD_DEFAULT, s.name)
			if err != nil || fsym == 0 {
				bootstrapErr = fmt.Errorf("mach: resolve %s: %w", s.name, err)
				return
			}
			purego.RegisterFunc(s.fn, fsym)
		}
	})
	return bootstrapErr
}

func bootstrapError(op, name string, code int32) error {
	if code == 0 {
		return nil
	}
	// bootstrap errors are 1100+ (BOOTSTRAP_NOT_PRIVILEGED etc).
	return fmt.Errorf("mach: %s(%q): bootstrap error %d", op, name, code)
}

// BootstrapRegister registers a send right for p under name in this
// session's bootstrap namespace, making it visible to BootstrapLookUp in
// other processes. p must carry a send right (see MakeSendRight).
//
// bootstrap_register is deprecated in favor of launchd job submission, but
// remains the only route for an unprivileged process to publish a port at
// runtime, which is exactly the rendezvous case this package serves.
func BootstrapRegister(name string, p Port) error {
	if err := loadBootstrap(); err != nil {
		return err
	}
	return bootstrapError("bootstrap_register", name, bootstrapRegister(uint32(bootstrapPort), name, uint32(p)))
}

// BootstrapLookUp returns a send right to the port registered under name.
// The caller owns the returned right and balances it with Deallocate.
func BootstrapLookUp(name string) (Port, error) {
	if err := loadBootstrap(); err != nil {
		return 0, err
	}
	var sp uint32
	if err := bootstrapError("bootstrap_look_up", name, bootstrapLookUp(uint32(bootstrapPort), name, &sp)); err != nil {
		return 0, err
	}
	return Port(sp), nil
}
