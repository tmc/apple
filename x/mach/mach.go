package mach

import (
	"fmt"
	"sync"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/kernel"
)

// Thread is a Mach thread port name.
type Thread uint32

var (
	loadOnce sync.Once
	loadErr  error

	pthreadSelf       func() uintptr
	pthreadMachThread func(uintptr) uint32
	timebaseInfo      func(*[2]uint32) int32
)

func load() error {
	loadOnce.Do(func() {
		for _, s := range []struct {
			name string
			fn   any
		}{
			{"pthread_self", &pthreadSelf},
			{"pthread_mach_thread_np", &pthreadMachThread},
			{"mach_timebase_info", &timebaseInfo},
		} {
			sym, err := purego.Dlsym(purego.RTLD_DEFAULT, s.name)
			if err != nil || sym == 0 {
				loadErr = fmt.Errorf("mach: resolve %s: %w", s.name, err)
				return
			}
			purego.RegisterFunc(s.fn, sym)
		}
	})
	return loadErr
}

// ThreadSelf returns the Mach thread port of the calling OS thread.
//
// It resolves the port via pthread_mach_thread_np(pthread_self()) rather than
// mach_thread_self: the pthread route borrows the port cached in the pthread
// structure, so no port right is allocated and nothing needs to be
// deallocated. The result is only meaningful for the current OS thread; pin
// it with runtime.LockOSThread before using the port.
func ThreadSelf() (Thread, error) {
	if err := load(); err != nil {
		return 0, err
	}
	return Thread(pthreadMachThread(pthreadSelf())), nil
}

// kernError converts a kern_return_t to an error, nil on KERN_SUCCESS.
func kernError(op string, kr kernel.Kern_return_t) error {
	if kr == 0 {
		return nil
	}
	return fmt.Errorf("mach: %s: kern_return %d", op, kr)
}
