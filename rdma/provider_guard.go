package rdma

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	// rdmaProviderMu serializes Apple RDMA provider calls in this process. The
	// provider has shown process-wide wedge behavior under overlapping calls; this
	// guard is separate from errno handling, which is protected by LockOSThread.
	rdmaProviderMu   sync.Mutex
	rdmaErrnoOnce    sync.Once
	rdmaErrorPointer func() unsafe.Pointer
	rdmaContextNames sync.Map
)

func rdmaProviderCall[T any](fn func() T) T {
	v, _, _ := rdmaProviderCallWithErrno(fn)
	return v
}

func rdmaProviderCallWithErrno[T any](fn func() T) (T, int, bool) {
	rdmaProviderMu.Lock()
	defer rdmaProviderMu.Unlock()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	rdmaSetErrno(0)
	v := fn()
	errno, ok := rdmaErrno()
	return v, errno, ok
}

func rdmaProviderCall0(fn func()) {
	rdmaProviderCall(func() struct{} {
		fn()
		return struct{}{}
	})
}

func rdmaKeepAlive(v any) {
	runtime.KeepAlive(v)
}

func rdmaErrno() (int, bool) {
	rdmaInitErrno()
	if rdmaErrorPointer == nil {
		return 0, false
	}
	return int(*(*int32)(rdmaErrorPointer())), true
}

func rdmaSetErrno(v int32) {
	rdmaInitErrno()
	if rdmaErrorPointer != nil {
		*(*int32)(rdmaErrorPointer()) = v
	}
}

func rdmaInitErrno() {
	rdmaErrnoOnce.Do(func() {
		sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "__error")
		if err != nil || sym == 0 {
			return
		}
		purego.RegisterFunc(&rdmaErrorPointer, sym)
	})
}

func rdmaRememberContext(ctx RDMAContext, device string) {
	if ctx == 0 || device == "" {
		return
	}
	rdmaContextNames.Store(ctx, device)
}

func rdmaForgetContext(ctx RDMAContext) {
	if ctx != 0 {
		rdmaContextNames.Delete(ctx)
	}
}

func rdmaContextDevice(ctx RDMAContext) string {
	if ctx == 0 {
		return ""
	}
	if v, ok := rdmaContextNames.Load(ctx); ok {
		if name, ok := v.(string); ok {
			return name
		}
	}
	return ""
}
