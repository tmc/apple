//go:build darwin

package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

func callNSExtensionMain() error {
	handle, err := purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("open Foundation: %w", err)
	}
	sym, err := purego.Dlsym(handle, "NSExtensionMain")
	if err != nil {
		return fmt.Errorf("resolve NSExtensionMain: %w", err)
	}
	var nsextensionMain func(int32, uintptr) int32
	purego.RegisterFunc(&nsextensionMain, sym)

	var cstrings [][]byte
	var argv []uintptr
	for _, arg := range os.Args {
		cstrings = append(cstrings, append([]byte(arg), 0))
	}
	for i := range cstrings {
		argv = append(argv, uintptr(unsafe.Pointer(&cstrings[i][0])))
	}
	argv = append(argv, 0)

	status := nsextensionMain(int32(len(cstrings)), uintptr(unsafe.Pointer(&argv[0])))
	extensionLog(fmt.Sprintf("NSExtensionMain returned %d", status))
	if status != 0 {
		return fmt.Errorf("NSExtensionMain returned %d", status)
	}
	corefoundation.CFRunLoopRun()
	return nil
}
