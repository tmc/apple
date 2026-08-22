// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenDriver] class.
var (
	_TKTokenDriverClass     TKTokenDriverClass
	_TKTokenDriverClassOnce sync.Once
)

func getTKTokenDriverClass() TKTokenDriverClass {
	_TKTokenDriverClassOnce.Do(func() {
		_TKTokenDriverClass = TKTokenDriverClass{class: objc.GetClass("TKTokenDriver")}
	})
	return _TKTokenDriverClass
}

// GetTKTokenDriverClass returns the class object for TKTokenDriver.
func GetTKTokenDriverClass() TKTokenDriverClass {
	return getTKTokenDriverClass()
}

type TKTokenDriverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenDriverClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenDriverClass) Alloc() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A base class for building token drivers.
//
// # Overview
//
// When using the [TKTokenDriver] class, implement the [TKTokenDriverDelegate]
// protocol with the [TokenDriverTokenForConfigurationError] method, which the
// system invokes when it requests the creation of a token instance. After you
// create the token driver, it can examine
// [TKTokenConfiguration.KeychainItems] and
// [TKTokenConfiguration.ConfigurationData] to implement your desired
// functionality.
//
// An implementation can also access its associated token configuration using
// the [TKTokenConfiguration] property.
//
// # Responding to Token Creation
//
//   - [TKTokenDriver.Delegate]: The token driver delegate.
//   - [TKTokenDriver.SetDelegate]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver
type TKTokenDriver struct {
	objectivec.Object
}

// TKTokenDriverFromID constructs a [TKTokenDriver] from an objc.ID.
//
// A base class for building token drivers.
func TKTokenDriverFromID(id objc.ID) TKTokenDriver {
	return TKTokenDriver{objectivec.Object{ID: id}}
}

// NOTE: TKTokenDriver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenDriver] class.
//
// # Responding to Token Creation
//
//   - [ITKTokenDriver.Delegate]: The token driver delegate.
//   - [ITKTokenDriver.SetDelegate]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver
type ITKTokenDriver interface {
	objectivec.IObject

	// Topic: Responding to Token Creation

	// The token driver delegate.
	Delegate() TKTokenDriverDelegate
	SetDelegate(value TKTokenDriverDelegate)
}

// Init initializes the instance.
func (t TKTokenDriver) Init() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenDriver) Autorelease() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenDriver creates a new TKTokenDriver instance.
func NewTKTokenDriver() TKTokenDriver {
	class := getTKTokenDriverClass()
	rv := objc.Send[TKTokenDriver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The token driver delegate.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/delegate
func (t TKTokenDriver) Delegate() TKTokenDriverDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return TKTokenDriverDelegateObjectFromID(rv)
}
func (t TKTokenDriver) SetDelegate(value TKTokenDriverDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
