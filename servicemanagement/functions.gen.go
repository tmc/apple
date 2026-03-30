// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ServiceManagement: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: ServiceManagement: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ServiceManagement: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _sMCopyAllJobDictionaries func(domain corefoundation.CFStringRef) corefoundation.CFArrayRef

// SMCopyAllJobDictionaries copies the job description dictionaries for all jobs in the specified domain.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMCopyAllJobDictionaries(_:)
func SMCopyAllJobDictionaries(domain corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _sMCopyAllJobDictionaries == nil {
		panic("ServiceManagement: symbol SMCopyAllJobDictionaries not loaded")
	}
	return _sMCopyAllJobDictionaries(domain)
}

var _sMJobCopyDictionary func(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef) corefoundation.CFDictionaryRef

// SMJobCopyDictionary copies the job description dictionary for the specified job label.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMJobCopyDictionary(_:_:)
func SMJobCopyDictionary(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	if _sMJobCopyDictionary == nil {
		panic("ServiceManagement: symbol SMJobCopyDictionary not loaded")
	}
	return _sMJobCopyDictionary(domain, jobLabel)
}

var _sMJobRemove func(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef, auth unsafe.Pointer, wait bool, outError *corefoundation.CFErrorRef) bool

// SMJobRemove removes the job with the specified label from the specified domain.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMJobRemove(_:_:_:_:_:)
func SMJobRemove(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef, auth unsafe.Pointer, wait bool, outError *corefoundation.CFErrorRef) bool {
	if _sMJobRemove == nil {
		panic("ServiceManagement: symbol SMJobRemove not loaded")
	}
	return _sMJobRemove(domain, jobLabel, auth, wait, outError)
}

var _sMJobSubmit func(domain corefoundation.CFStringRef, job corefoundation.CFDictionaryRef, auth unsafe.Pointer, outError *corefoundation.CFErrorRef) bool

// SMJobSubmit submits the specified job to the specified domain.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMJobSubmit(_:_:_:_:)
func SMJobSubmit(domain corefoundation.CFStringRef, job corefoundation.CFDictionaryRef, auth unsafe.Pointer, outError *corefoundation.CFErrorRef) bool {
	if _sMJobSubmit == nil {
		panic("ServiceManagement: symbol SMJobSubmit not loaded")
	}
	return _sMJobSubmit(domain, job, auth, outError)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_sMCopyAllJobDictionaries, frameworkHandle, "SMCopyAllJobDictionaries")
	registerFunc(&_sMJobCopyDictionary, frameworkHandle, "SMJobCopyDictionary")
	registerFunc(&_sMJobRemove, frameworkHandle, "SMJobRemove")
	registerFunc(&_sMJobSubmit, frameworkHandle, "SMJobSubmit")
}
