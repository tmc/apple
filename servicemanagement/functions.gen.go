// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("ServiceManagement: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("ServiceManagement: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("ServiceManagement: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("ServiceManagement: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _sMCopyAllJobDictionaries func(domain corefoundation.CFStringRef) corefoundation.CFArrayRef
var _sMCopyAllJobDictionariesErr error

func trySMCopyAllJobDictionaries(domain corefoundation.CFStringRef) (corefoundation.CFArrayRef, error) {
	if _sMCopyAllJobDictionaries == nil {
		return 0, symbolCallError("SMCopyAllJobDictionaries", "10.6", _sMCopyAllJobDictionariesErr)
	}
	return _sMCopyAllJobDictionaries(domain), nil
}

// SMCopyAllJobDictionaries copies the job description dictionaries for all jobs in the specified domain.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMCopyAllJobDictionaries(_:)
func SMCopyAllJobDictionaries(domain corefoundation.CFStringRef) corefoundation.CFArrayRef {
	result, callErr := trySMCopyAllJobDictionaries(domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sMJobCopyDictionary func(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef) corefoundation.CFDictionaryRef
var _sMJobCopyDictionaryErr error

func trySMJobCopyDictionary(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef) (corefoundation.CFDictionaryRef, error) {
	if _sMJobCopyDictionary == nil {
		return 0, symbolCallError("SMJobCopyDictionary", "10.6", _sMJobCopyDictionaryErr)
	}
	return _sMJobCopyDictionary(domain, jobLabel), nil
}

// SMJobCopyDictionary copies the job description dictionary for the specified job label.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMJobCopyDictionary(_:_:)
func SMJobCopyDictionary(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	result, callErr := trySMJobCopyDictionary(domain, jobLabel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sMJobRemove func(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef, auth unsafe.Pointer, wait bool, outError *corefoundation.CFErrorRef) bool
var _sMJobRemoveErr error

func trySMJobRemove(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef, auth unsafe.Pointer, wait bool, outError *corefoundation.CFErrorRef) (bool, error) {
	if _sMJobRemove == nil {
		return false, symbolCallError("SMJobRemove", "10.6", _sMJobRemoveErr)
	}
	return _sMJobRemove(domain, jobLabel, auth, wait, outError), nil
}

// SMJobRemove removes the job with the specified label from the specified domain.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMJobRemove(_:_:_:_:_:)
func SMJobRemove(domain corefoundation.CFStringRef, jobLabel corefoundation.CFStringRef, auth unsafe.Pointer, wait bool, outError *corefoundation.CFErrorRef) bool {
	result, callErr := trySMJobRemove(domain, jobLabel, auth, wait, outError)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sMJobSubmit func(domain corefoundation.CFStringRef, job corefoundation.CFDictionaryRef, auth unsafe.Pointer, outError *corefoundation.CFErrorRef) bool
var _sMJobSubmitErr error

func trySMJobSubmit(domain corefoundation.CFStringRef, job corefoundation.CFDictionaryRef, auth unsafe.Pointer, outError *corefoundation.CFErrorRef) (bool, error) {
	if _sMJobSubmit == nil {
		return false, symbolCallError("SMJobSubmit", "10.6", _sMJobSubmitErr)
	}
	return _sMJobSubmit(domain, job, auth, outError), nil
}

// SMJobSubmit submits the specified job to the specified domain.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMJobSubmit(_:_:_:_:)
func SMJobSubmit(domain corefoundation.CFStringRef, job corefoundation.CFDictionaryRef, auth unsafe.Pointer, outError *corefoundation.CFErrorRef) bool {
	result, callErr := trySMJobSubmit(domain, job, auth, outError)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_sMCopyAllJobDictionaries, &_sMCopyAllJobDictionariesErr, frameworkHandle, "SMCopyAllJobDictionaries", "10.6")
	registerFunc(&_sMJobCopyDictionary, &_sMJobCopyDictionaryErr, frameworkHandle, "SMJobCopyDictionary", "10.6")
	registerFunc(&_sMJobRemove, &_sMJobRemoveErr, frameworkHandle, "SMJobRemove", "10.6")
	registerFunc(&_sMJobSubmit, &_sMJobSubmitErr, frameworkHandle, "SMJobSubmit", "10.6")
}
