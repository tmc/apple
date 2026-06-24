//go:build darwin && arm64

package main

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

const extensionFoundationTextBase = 0x227a54000

const (
	exRunningSharedInstanceOffset = 0x227aa1400 - extensionFoundationTextBase
	exRunningResumeOffset         = 0x227aa2b08 - extensionFoundationTextBase
)

type dlInfo struct {
	fileName uintptr
	base     uintptr
	symName  uintptr
	symAddr  uintptr
}

func callSwiftSharedInstance(fn uintptr) uintptr
func callSwiftResume(self, fn uintptr)

func tryPrivateRunningExtensionResume() error {
	handle, err := purego.Dlopen("/System/Library/Frameworks/ExtensionFoundation.framework/ExtensionFoundation", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("open ExtensionFoundation: %w", err)
	}
	mainSym, err := purego.Dlsym(handle, "_$s19ExtensionFoundation03AppA0PAAE4mainyyKFZ")
	if err != nil {
		return fmt.Errorf("resolve AppExtension.main: %w", err)
	}
	libSystem, err := purego.Dlopen("/usr/lib/libSystem.B.dylib", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("open libSystem: %w", err)
	}
	dladdrSym, err := purego.Dlsym(libSystem, "dladdr")
	if err != nil {
		return fmt.Errorf("resolve dladdr: %w", err)
	}
	var dladdr func(uintptr, *dlInfo) int
	purego.RegisterFunc(&dladdr, dladdrSym)
	var info dlInfo
	if dladdr(mainSym, &info) == 0 || info.base == 0 {
		return fmt.Errorf("dladdr ExtensionFoundation base from %#x failed", mainSym)
	}
	sharedFn := info.base + exRunningSharedInstanceOffset
	resumeFn := info.base + exRunningResumeOffset
	extensionLog(fmt.Sprintf("private resume: base=%#x shared=%#x resume=%#x", info.base, sharedFn, resumeFn))
	instance := callSwiftSharedInstance(sharedFn)
	if instance == 0 {
		return fmt.Errorf("private sharedInstance returned nil")
	}
	extensionLog(fmt.Sprintf("private resume: sharedInstance=%#x", instance))
	callSwiftResume(instance, resumeFn)
	extensionLog("private resume returned")
	_ = unsafe.Sizeof(info)
	return nil
}
