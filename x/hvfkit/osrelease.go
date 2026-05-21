//go:build darwin && arm64

package hvfkit

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	osReleaseOnce sync.Once
	osReleaseFunc func(unsafe.Pointer)
	osReleaseErr  error
)

func osRelease(object unsafe.Pointer) (err error) {
	if object == nil {
		return nil
	}
	osReleaseOnce.Do(func() {
		lib, err := purego.Dlopen("/usr/lib/libSystem.B.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			osReleaseErr = fmt.Errorf("load libSystem: %w", err)
			return
		}
		sym, err := purego.Dlsym(lib, "os_release")
		if err != nil {
			osReleaseErr = fmt.Errorf("resolve os_release: %w", err)
			return
		}
		if sym == 0 {
			osReleaseErr = fmt.Errorf("resolve os_release: symbol not found")
			return
		}
		defer func() {
			if r := recover(); r != nil {
				osReleaseErr = recoveredError("os_release", r)
			}
		}()
		purego.RegisterFunc(&osReleaseFunc, sym)
	})
	if osReleaseErr != nil {
		return osReleaseErr
	}
	defer func() {
		if r := recover(); r != nil {
			err = recoveredError("os_release", r)
		}
	}()
	osReleaseFunc(object)
	return nil
}
