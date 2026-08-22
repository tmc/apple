// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [TKSmartCardTokenDriver] class.
var (
	_TKSmartCardTokenDriverClass     TKSmartCardTokenDriverClass
	_TKSmartCardTokenDriverClassOnce sync.Once
)

func getTKSmartCardTokenDriverClass() TKSmartCardTokenDriverClass {
	_TKSmartCardTokenDriverClassOnce.Do(func() {
		_TKSmartCardTokenDriverClass = TKSmartCardTokenDriverClass{class: objc.GetClass("TKSmartCardTokenDriver")}
	})
	return _TKSmartCardTokenDriverClass
}

// GetTKSmartCardTokenDriverClass returns the class object for TKSmartCardTokenDriver.
func GetTKSmartCardTokenDriverClass() TKSmartCardTokenDriverClass {
	return getTKSmartCardTokenDriverClass()
}

type TKSmartCardTokenDriverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardTokenDriverClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardTokenDriverClass) Alloc() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// The driver that acts as an entry point for smart card app extensions.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenDriver
type TKSmartCardTokenDriver struct {
	TKTokenDriver
}

// TKSmartCardTokenDriverFromID constructs a [TKSmartCardTokenDriver] from an objc.ID.
//
// The driver that acts as an entry point for smart card app extensions.
func TKSmartCardTokenDriverFromID(id objc.ID) TKSmartCardTokenDriver {
	return TKSmartCardTokenDriver{TKTokenDriver: TKTokenDriverFromID(id)}
}

// NOTE: TKSmartCardTokenDriver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardTokenDriver] class.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenDriver
type ITKSmartCardTokenDriver interface {
	ITKTokenDriver
}

// Init initializes the instance.
func (t TKSmartCardTokenDriver) Init() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardTokenDriver) Autorelease() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenDriver creates a new TKSmartCardTokenDriver instance.
func NewTKSmartCardTokenDriver() TKSmartCardTokenDriver {
	class := getTKSmartCardTokenDriverClass()
	rv := objc.Send[TKSmartCardTokenDriver](objc.ID(class.class), objc.Sel("new"))
	return rv
}
