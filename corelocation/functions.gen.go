// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"fmt"

	"github.com/ebitengine/purego"
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
		return fmt.Sprintf("CoreLocation: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreLocation: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreLocation: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreLocation: register symbol %s: %v", name, r)
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

var _cLLocationCoordinate2DIsValid func(coord CLLocationCoordinate2D) bool
var _cLLocationCoordinate2DIsValidErr error

func tryCLLocationCoordinate2DIsValid(coord CLLocationCoordinate2D) (bool, error) {
	if _cLLocationCoordinate2DIsValid == nil {
		return false, symbolCallError("CLLocationCoordinate2DIsValid", "10.7", _cLLocationCoordinate2DIsValidErr)
	}
	return _cLLocationCoordinate2DIsValid(coord), nil
}

// CLLocationCoordinate2DIsValid returns a Boolean value indicating whether the specified coordinate is valid.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationCoordinate2DIsValid(_:)
func CLLocationCoordinate2DIsValid(coord CLLocationCoordinate2D) bool {
	result, callErr := tryCLLocationCoordinate2DIsValid(coord)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cLLocationCoordinate2DMake func(latitude CLLocationDegrees, longitude CLLocationDegrees) CLLocationCoordinate2D
var _cLLocationCoordinate2DMakeErr error

func tryCLLocationCoordinate2DMake(latitude CLLocationDegrees, longitude CLLocationDegrees) (CLLocationCoordinate2D, error) {
	if _cLLocationCoordinate2DMake == nil {
		return CLLocationCoordinate2D{}, symbolCallError("CLLocationCoordinate2DMake", "10.7", _cLLocationCoordinate2DMakeErr)
	}
	return _cLLocationCoordinate2DMake(latitude, longitude), nil
}

// CLLocationCoordinate2DMake formats a latitude and longitude value into a coordinate data structure format.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationCoordinate2DMake(_:_:)
func CLLocationCoordinate2DMake(latitude CLLocationDegrees, longitude CLLocationDegrees) CLLocationCoordinate2D {
	result, callErr := tryCLLocationCoordinate2DMake(latitude, longitude)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cLLocationCoordinate2DIsValid, &_cLLocationCoordinate2DIsValidErr, frameworkHandle, "CLLocationCoordinate2DIsValid", "10.7")
	registerFunc(&_cLLocationCoordinate2DMake, &_cLLocationCoordinate2DMakeErr, frameworkHandle, "CLLocationCoordinate2DMake", "10.7")
}
