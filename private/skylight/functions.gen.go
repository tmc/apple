// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
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

var _cGEventSetWindowLocation func(event coregraphics.CGEventRef, point corefoundation.CGPoint)
var _cGEventSetWindowLocationErr error

func tryCGEventSetWindowLocation(event coregraphics.CGEventRef, point corefoundation.CGPoint) error {
	if _cGEventSetWindowLocation == nil {
		return symbolCallError("CGEventSetWindowLocation", "", _cGEventSetWindowLocationErr)
	}
	_cGEventSetWindowLocation(event, point)
	return nil
}

// CGEventSetWindowLocation.
//
// See: https://developer.apple.com/documentation/SkyLight/CGEventSetWindowLocation
func CGEventSetWindowLocation(event coregraphics.CGEventRef, point corefoundation.CGPoint) error {
	return tryCGEventSetWindowLocation(event, point)
}

var _cGSMainConnectionID func() CGSConnectionID
var _cGSMainConnectionIDErr error

func tryCGSMainConnectionID() (CGSConnectionID, error) {
	if _cGSMainConnectionID == nil {
		return *new(CGSConnectionID), symbolCallError("CGSMainConnectionID", "", _cGSMainConnectionIDErr)
	}
	return _cGSMainConnectionID(), nil
}

// CGSMainConnectionID.
//
// See: https://developer.apple.com/documentation/SkyLight/CGSMainConnectionID
func CGSMainConnectionID() (CGSConnectionID, error) {
	return tryCGSMainConnectionID()
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

var _sLEventSetAuthenticationMessage func(event coregraphics.CGEventRef, message objectivec.Object)
var _sLEventSetAuthenticationMessageErr error

func trySLEventSetAuthenticationMessage(event coregraphics.CGEventRef, message objectivec.Object) error {
	if _sLEventSetAuthenticationMessage == nil {
		return symbolCallError("SLEventSetAuthenticationMessage", "", _sLEventSetAuthenticationMessageErr)
	}
	_sLEventSetAuthenticationMessage(event, message)
	return nil
}

// SLEventSetAuthenticationMessage.
//
// See: https://developer.apple.com/documentation/SkyLight/SLEventSetAuthenticationMessage
func SLEventSetAuthenticationMessage(event coregraphics.CGEventRef, message objectivec.Object) error {
	return trySLEventSetAuthenticationMessage(event, message)
}

var _sLEventSetIntegerValueField func(event coregraphics.CGEventRef, field coregraphics.CGEventField, value int64)
var _sLEventSetIntegerValueFieldErr error

func trySLEventSetIntegerValueField(event coregraphics.CGEventRef, field coregraphics.CGEventField, value int64) error {
	if _sLEventSetIntegerValueField == nil {
		return symbolCallError("SLEventSetIntegerValueField", "", _sLEventSetIntegerValueFieldErr)
	}
	_sLEventSetIntegerValueField(event, field, value)
	return nil
}

// SLEventSetIntegerValueField.
//
// See: https://developer.apple.com/documentation/SkyLight/SLEventSetIntegerValueField
func SLEventSetIntegerValueField(event coregraphics.CGEventRef, field coregraphics.CGEventField, value int64) error {
	return trySLEventSetIntegerValueField(event, field, value)
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

var _sLSGetConnectionPSN func(cid CGSConnectionID, psn *ProcessSerialNumber) coregraphics.CGError
var _sLSGetConnectionPSNErr error

func trySLSGetConnectionPSN(cid CGSConnectionID, psn *ProcessSerialNumber) (coregraphics.CGError, error) {
	if _sLSGetConnectionPSN == nil {
		return *new(coregraphics.CGError), symbolCallError("SLSGetConnectionPSN", "", _sLSGetConnectionPSNErr)
	}
	return _sLSGetConnectionPSN(cid, psn), nil
}

// SLSGetConnectionPSN.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSGetConnectionPSN
func SLSGetConnectionPSN(cid CGSConnectionID, psn *ProcessSerialNumber) (coregraphics.CGError, error) {
	return trySLSGetConnectionPSN(cid, psn)
}

var _sLSGetWindowOwner func(cid CGSConnectionID, wid coregraphics.CGWindowID, owner *CGSConnectionID) coregraphics.CGError
var _sLSGetWindowOwnerErr error

func trySLSGetWindowOwner(cid CGSConnectionID, wid coregraphics.CGWindowID, owner *CGSConnectionID) (coregraphics.CGError, error) {
	if _sLSGetWindowOwner == nil {
		return *new(coregraphics.CGError), symbolCallError("SLSGetWindowOwner", "", _sLSGetWindowOwnerErr)
	}
	return _sLSGetWindowOwner(cid, wid, owner), nil
}

// SLSGetWindowOwner.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSGetWindowOwner
func SLSGetWindowOwner(cid CGSConnectionID, wid coregraphics.CGWindowID, owner *CGSConnectionID) (coregraphics.CGError, error) {
	return trySLSGetWindowOwner(cid, wid, owner)
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
	registerFunc(&_cGEventSetWindowLocation, &_cGEventSetWindowLocationErr, frameworkHandle, "CGEventSetWindowLocation", "")
	registerFunc(&_cGSMainConnectionID, &_cGSMainConnectionIDErr, frameworkHandle, "CGSMainConnectionID", "")
	registerFunc(&_sLEventPostToPSN, &_sLEventPostToPSNErr, frameworkHandle, "SLEventPostToPSN", "")
	registerFunc(&_sLEventPostToPid, &_sLEventPostToPidErr, frameworkHandle, "SLEventPostToPid", "")
	registerFunc(&_sLEventSetAuthenticationMessage, &_sLEventSetAuthenticationMessageErr, frameworkHandle, "SLEventSetAuthenticationMessage", "")
	registerFunc(&_sLEventSetIntegerValueField, &_sLEventSetIntegerValueFieldErr, frameworkHandle, "SLEventSetIntegerValueField", "")
	registerFunc(&_sLPSPostEventRecordTo, &_sLPSPostEventRecordToErr, frameworkHandle, "SLPSPostEventRecordTo", "")
	registerFunc(&_sLSGetConnectionPSN, &_sLSGetConnectionPSNErr, frameworkHandle, "SLSGetConnectionPSN", "")
	registerFunc(&_sLSGetWindowOwner, &_sLSGetWindowOwnerErr, frameworkHandle, "SLSGetWindowOwner", "")
	registerFunc(&_sLPSGetFrontProcess, &_sLPSGetFrontProcessErr, frameworkHandle, "_SLPSGetFrontProcess", "")
	registerFunc(&_sLPSSetFrontProcessWithOptions, &_sLPSSetFrontProcessWithOptionsErr, frameworkHandle, "_SLPSSetFrontProcessWithOptions", "")
}
