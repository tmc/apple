package rdma

import (
	"fmt"
	"runtime"
	"sync"
)

var rdmaProviderMu sync.Mutex

func rdmaProviderCall[T any](fn func() T) T {
	rdmaProviderMu.Lock()
	defer rdmaProviderMu.Unlock()
	return fn()
}

func rdmaProviderCall0(fn func()) {
	rdmaProviderMu.Lock()
	defer rdmaProviderMu.Unlock()
	fn()
}

func rdmaKeepAlive(v any) {
	runtime.KeepAlive(v)
}

func rdmaSetFinalizer(obj any, finalizer any) {
	runtime.SetFinalizer(obj, finalizer)
}

func rdmaNilHandleError(name string) error {
	return fmt.Errorf("rdma: nil %s handle", name)
}

func rdmaNilPointerError(name string) error {
	return fmt.Errorf("rdma: nil %s pointer", name)
}
