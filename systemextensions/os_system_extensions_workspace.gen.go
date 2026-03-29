// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSSystemExtensionsWorkspace] class.
var (
	_OSSystemExtensionsWorkspaceClass     OSSystemExtensionsWorkspaceClass
	_OSSystemExtensionsWorkspaceClassOnce sync.Once
)

func getOSSystemExtensionsWorkspaceClass() OSSystemExtensionsWorkspaceClass {
	_OSSystemExtensionsWorkspaceClassOnce.Do(func() {
		_OSSystemExtensionsWorkspaceClass = OSSystemExtensionsWorkspaceClass{class: objc.GetClass("OSSystemExtensionsWorkspace")}
	})
	return _OSSystemExtensionsWorkspaceClass
}

// GetOSSystemExtensionsWorkspaceClass returns the class object for OSSystemExtensionsWorkspace.
func GetOSSystemExtensionsWorkspaceClass() OSSystemExtensionsWorkspaceClass {
	return getOSSystemExtensionsWorkspaceClass()
}

type OSSystemExtensionsWorkspaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSSystemExtensionsWorkspaceClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSSystemExtensionsWorkspaceClass) Alloc() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

//
// # Overview
//
// # Instance Methods
//
//   - [OSSystemExtensionsWorkspace.AddObserverError]
//   - [OSSystemExtensionsWorkspace.RemoveObserver]
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace
type OSSystemExtensionsWorkspace struct {
	objectivec.Object
}

// OSSystemExtensionsWorkspaceFromID constructs a [OSSystemExtensionsWorkspace] from an objc.ID.
func OSSystemExtensionsWorkspaceFromID(id objc.ID) OSSystemExtensionsWorkspace {
	return OSSystemExtensionsWorkspace{objectivec.Object{ID: id}}
}
// NOTE: OSSystemExtensionsWorkspace adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSSystemExtensionsWorkspace] class.
//
// # Instance Methods
//
//   - [IOSSystemExtensionsWorkspace.AddObserverError]
//   - [IOSSystemExtensionsWorkspace.RemoveObserver]
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace
type IOSSystemExtensionsWorkspace interface {
	objectivec.IObject

	// Topic: Instance Methods

	AddObserverError(observer OSSystemExtensionsWorkspaceObserver) (bool, error)
	RemoveObserver(observer OSSystemExtensionsWorkspaceObserver)
}

// Init initializes the instance.
func (o OSSystemExtensionsWorkspace) Init() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSSystemExtensionsWorkspace) Autorelease() OSSystemExtensionsWorkspace {
	rv := objc.Send[OSSystemExtensionsWorkspace](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionsWorkspace creates a new OSSystemExtensionsWorkspace instance.
func NewOSSystemExtensionsWorkspace() OSSystemExtensionsWorkspace {
	class := getOSSystemExtensionsWorkspaceClass()
	rv := objc.Send[OSSystemExtensionsWorkspace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// # Discussion
// 
// Start observing changes to System Extension(s) which are enabled or ready
// to be enabled.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/addObserver(_:)
func (o OSSystemExtensionsWorkspace) AddObserverError(observer OSSystemExtensionsWorkspaceObserver) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("addObserver:error:"), observer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addObserver:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// # Discussion
// 
// Stop observing changes to System Extension(s).
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/removeObserver(_:)
func (o OSSystemExtensionsWorkspace) RemoveObserver(observer OSSystemExtensionsWorkspaceObserver) {
	objc.Send[objc.ID](o.ID, objc.Sel("removeObserver:"), observer)
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/shared
func (_OSSystemExtensionsWorkspaceClass OSSystemExtensionsWorkspaceClass) SharedWorkspace() OSSystemExtensionsWorkspace {
	rv := objc.Send[objc.ID](objc.ID(_OSSystemExtensionsWorkspaceClass.class), objc.Sel("sharedWorkspace"))
	return OSSystemExtensionsWorkspaceFromID(objc.ID(rv))
}

