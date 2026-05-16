// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSModuleIdentity] class.
var (
	_FSModuleIdentityClass     FSModuleIdentityClass
	_FSModuleIdentityClassOnce sync.Once
)

func getFSModuleIdentityClass() FSModuleIdentityClass {
	_FSModuleIdentityClassOnce.Do(func() {
		_FSModuleIdentityClass = FSModuleIdentityClass{class: objc.GetClass("FSModuleIdentity")}
	})
	return _FSModuleIdentityClass
}

// GetFSModuleIdentityClass returns the class object for FSModuleIdentity.
func GetFSModuleIdentityClass() FSModuleIdentityClass {
	return getFSModuleIdentityClass()
}

type FSModuleIdentityClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSModuleIdentityClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSModuleIdentityClass) Alloc() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An installed file system module.
//
// # Accessing module properties
//
//   - [FSModuleIdentity.BundleIdentifier]: The module’s bundle identifier.
//   - [FSModuleIdentity.Url]: The module’s URL.
//   - [FSModuleIdentity.Enabled]: A Boolean value that indicates if the module is enabled.
//
// See: https://developer.apple.com/documentation/FSKit/FSModuleIdentity
type FSModuleIdentity struct {
	objectivec.Object
}

// FSModuleIdentityFromID constructs a [FSModuleIdentity] from an objc.ID.
//
// An installed file system module.
func FSModuleIdentityFromID(id objc.ID) FSModuleIdentity {
	return FSModuleIdentity{objectivec.Object{ID: id}}
}

// NOTE: FSModuleIdentity adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSModuleIdentity] class.
//
// # Accessing module properties
//
//   - [IFSModuleIdentity.BundleIdentifier]: The module’s bundle identifier.
//   - [IFSModuleIdentity.Url]: The module’s URL.
//   - [IFSModuleIdentity.Enabled]: A Boolean value that indicates if the module is enabled.
//
// See: https://developer.apple.com/documentation/FSKit/FSModuleIdentity
type IFSModuleIdentity interface {
	objectivec.IObject

	// Topic: Accessing module properties

	// The module’s bundle identifier.
	BundleIdentifier() string
	// The module’s URL.
	Url() foundation.INSURL
	// A Boolean value that indicates if the module is enabled.
	Enabled() bool
}

// Init initializes the instance.
func (m FSModuleIdentity) Init() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m FSModuleIdentity) Autorelease() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSModuleIdentity creates a new FSModuleIdentity instance.
func NewFSModuleIdentity() FSModuleIdentity {
	class := getFSModuleIdentityClass()
	rv := objc.Send[FSModuleIdentity](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The module’s bundle identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/bundleIdentifier
func (m FSModuleIdentity) BundleIdentifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The module’s URL.
//
// See: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/url
func (m FSModuleIdentity) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// A Boolean value that indicates if the module is enabled.
//
// See: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/isEnabled
func (m FSModuleIdentity) Enabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}
