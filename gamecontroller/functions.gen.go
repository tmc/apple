// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
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
		return fmt.Sprintf("GameController: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("GameController: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("GameController: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("GameController: register symbol %s: %v", name, r)
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

var _gCInputArcadeButtonName func(row int, column int) GCInputButtonName
var _gCInputArcadeButtonNameErr error

func tryGCInputArcadeButtonName(row int, column int) (GCInputButtonName, error) {
	if _gCInputArcadeButtonName == nil {
		return *new(GCInputButtonName), symbolCallError("GCInputArcadeButtonName", "13.0", _gCInputArcadeButtonNameErr)
	}
	return _gCInputArcadeButtonName(row, column), nil
}

// GCInputArcadeButtonName returns the name of the arcade stick button at the specified location.
//
// See: https://developer.apple.com/documentation/GameController/GCInputArcadeButtonName
func GCInputArcadeButtonName(row int, column int) GCInputButtonName {
	result, callErr := tryGCInputArcadeButtonName(row, column)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gCInputBackLeftButton func(position int) GCInputButtonName
var _gCInputBackLeftButtonErr error

func tryGCInputBackLeftButton(position int) (GCInputButtonName, error) {
	if _gCInputBackLeftButton == nil {
		return *new(GCInputButtonName), symbolCallError("GCInputBackLeftButton", "14.4", _gCInputBackLeftButtonErr)
	}
	return _gCInputBackLeftButton(position), nil
}

// GCInputBackLeftButton returns the name of the back left button at the specified location.
//
// See: https://developer.apple.com/documentation/GameController/GCInputBackLeftButton
func GCInputBackLeftButton(position int) GCInputButtonName {
	result, callErr := tryGCInputBackLeftButton(position)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gCInputBackRightButton func(position int) GCInputButtonName
var _gCInputBackRightButtonErr error

func tryGCInputBackRightButton(position int) (GCInputButtonName, error) {
	if _gCInputBackRightButton == nil {
		return *new(GCInputButtonName), symbolCallError("GCInputBackRightButton", "14.4", _gCInputBackRightButtonErr)
	}
	return _gCInputBackRightButton(position), nil
}

// GCInputBackRightButton returns the name of the back right button at the specified location.
//
// See: https://developer.apple.com/documentation/GameController/GCInputBackRightButton
func GCInputBackRightButton(position int) GCInputButtonName {
	result, callErr := tryGCInputBackRightButton(position)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGCPoint2 func(point GCPoint2) *foundation.NSString
var _nSStringFromGCPoint2Err error

func tryNSStringFromGCPoint2(point GCPoint2) (*foundation.NSString, error) {
	if _nSStringFromGCPoint2 == nil {
		return nil, symbolCallError("NSStringFromGCPoint2", "14.3", _nSStringFromGCPoint2Err)
	}
	return _nSStringFromGCPoint2(point), nil
}

// NSStringFromGCPoint2 returns a string representation of a point.
//
// See: https://developer.apple.com/documentation/GameController/NSStringFromGCPoint2(_:)
func NSStringFromGCPoint2(point GCPoint2) *foundation.NSString {
	result, callErr := tryNSStringFromGCPoint2(point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_gCInputArcadeButtonName, &_gCInputArcadeButtonNameErr, frameworkHandle, "GCInputArcadeButtonName", "13.0")
	registerFunc(&_gCInputBackLeftButton, &_gCInputBackLeftButtonErr, frameworkHandle, "GCInputBackLeftButton", "14.4")
	registerFunc(&_gCInputBackRightButton, &_gCInputBackRightButtonErr, frameworkHandle, "GCInputBackRightButton", "14.4")
	registerFunc(&_nSStringFromGCPoint2, &_nSStringFromGCPoint2Err, frameworkHandle, "NSStringFromGCPoint2", "14.3")
}
