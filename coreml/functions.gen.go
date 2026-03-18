// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreML: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: CoreML: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreML: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _mLAllComputeDevices func() []unsafe.Pointer

// MLAllComputeDevices returns an array that contains all of the compute devices that are accessible.
//
// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDevices
func MLAllComputeDevices() []unsafe.Pointer {
	if _mLAllComputeDevices == nil {
		panic("CoreML: symbol MLAllComputeDevices not loaded")
	}
	return _mLAllComputeDevices()
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_mLAllComputeDevices, frameworkHandle, "MLAllComputeDevices")
	}

