// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: mlruntime: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: mlruntime: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: mlruntime: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _mLRSanitizeErrorSymbol uintptr

// MLRSanitizeError has an opaque C signature in discovered metadata.
// Call MLRSanitizeErrorSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRSanitizeError
func MLRSanitizeError() {
	panic("mlruntime: symbol MLRSanitizeError has opaque signature; use MLRSanitizeErrorSymbol() and a typed manual wrapper")
}

// MLRSanitizeErrorSymbol returns the raw symbol address for MLRSanitizeError.
func MLRSanitizeErrorSymbol() uintptr {
	if _mLRSanitizeErrorSymbol == 0 {
		return 0
	}
	return _mLRSanitizeErrorSymbol
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerSymbol(&_mLRSanitizeErrorSymbol, frameworkHandle, "MLRSanitizeError")
	}

