// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/coregraphics"
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
		return fmt.Sprintf("skylight: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("skylight: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("skylight: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("skylight: register symbol %s: %v", name, r)
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

var _sLEventPostToPSN func(psn *ProcessSerialNumber, event coregraphics.CGEventRef) int32
var _sLEventPostToPSNErr error

func trySLEventPostToPSN(psn *ProcessSerialNumber, event coregraphics.CGEventRef) (int32, error) {
	if _sLEventPostToPSN == nil {
		return 0, symbolCallError("SLEventPostToPSN", "", _sLEventPostToPSNErr)
	}
	return _sLEventPostToPSN(psn, event), nil
}

// SLEventPostToPSN.
//
// See: https://developer.apple.com/documentation/SkyLight/SLEventPostToPSN
func SLEventPostToPSN(psn *ProcessSerialNumber, event coregraphics.CGEventRef) (int32, error) {
	return trySLEventPostToPSN(psn, event)
}

var _sLEventPostToPid func(pid int32, event coregraphics.CGEventRef) int32
var _sLEventPostToPidErr error

func trySLEventPostToPid(pid int32, event coregraphics.CGEventRef) (int32, error) {
	if _sLEventPostToPid == nil {
		return 0, symbolCallError("SLEventPostToPid", "", _sLEventPostToPidErr)
	}
	return _sLEventPostToPid(pid, event), nil
}

// SLEventPostToPid.
//
// See: https://developer.apple.com/documentation/SkyLight/SLEventPostToPid
func SLEventPostToPid(pid int32, event coregraphics.CGEventRef) (int32, error) {
	return trySLEventPostToPid(pid, event)
}

var _sLPSPostEventRecordTo func(psn *ProcessSerialNumber, record *byte) int32
var _sLPSPostEventRecordToErr error

func trySLPSPostEventRecordTo(psn *ProcessSerialNumber, record []byte) (int32, error) {
	if _sLPSPostEventRecordTo == nil {
		return 0, symbolCallError("SLPSPostEventRecordTo", "", _sLPSPostEventRecordToErr)
	}
	return _sLPSPostEventRecordTo(psn, unsafe.SliceData(record)), nil
}

// SLPSPostEventRecordTo.
//
// See: https://developer.apple.com/documentation/SkyLight/SLPSPostEventRecordTo
func SLPSPostEventRecordTo(psn *ProcessSerialNumber, record []byte) (int32, error) {
	return trySLPSPostEventRecordTo(psn, record)
}

var _sLPSGetFrontProcess func(psn *ProcessSerialNumber) int32
var _sLPSGetFrontProcessErr error

func trySLPSGetFrontProcess(psn *ProcessSerialNumber) (int32, error) {
	if _sLPSGetFrontProcess == nil {
		return 0, symbolCallError("_SLPSGetFrontProcess", "", _sLPSGetFrontProcessErr)
	}
	return _sLPSGetFrontProcess(psn), nil
}

// SLPSGetFrontProcess.
//
// See: https://developer.apple.com/documentation/SkyLight/_SLPSGetFrontProcess
func SLPSGetFrontProcess(psn *ProcessSerialNumber) (int32, error) {
	return trySLPSGetFrontProcess(psn)
}

var _sLPSSetFrontProcessWithOptions func(psn *ProcessSerialNumber, wid uint32, mode uint32) int32
var _sLPSSetFrontProcessWithOptionsErr error

func trySLPSSetFrontProcessWithOptions(psn *ProcessSerialNumber, wid uint32, mode uint32) (int32, error) {
	if _sLPSSetFrontProcessWithOptions == nil {
		return 0, symbolCallError("_SLPSSetFrontProcessWithOptions", "", _sLPSSetFrontProcessWithOptionsErr)
	}
	return _sLPSSetFrontProcessWithOptions(psn, wid, mode), nil
}

// SLPSSetFrontProcessWithOptions.
//
// See: https://developer.apple.com/documentation/SkyLight/_SLPSSetFrontProcessWithOptions
func SLPSSetFrontProcessWithOptions(psn *ProcessSerialNumber, wid uint32, mode uint32) (int32, error) {
	return trySLPSSetFrontProcessWithOptions(psn, wid, mode)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_sLEventPostToPSN, &_sLEventPostToPSNErr, frameworkHandle, "SLEventPostToPSN", "")
	registerFunc(&_sLEventPostToPid, &_sLEventPostToPidErr, frameworkHandle, "SLEventPostToPid", "")
	registerFunc(&_sLPSPostEventRecordTo, &_sLPSPostEventRecordToErr, frameworkHandle, "SLPSPostEventRecordTo", "")
	registerFunc(&_sLPSGetFrontProcess, &_sLPSGetFrontProcessErr, frameworkHandle, "_SLPSGetFrontProcess", "")
	registerFunc(&_sLPSSetFrontProcessWithOptions, &_sLPSSetFrontProcessWithOptionsErr, frameworkHandle, "_SLPSSetFrontProcessWithOptions", "")
}
