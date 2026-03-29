// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
var (
	_VZLinuxRosettaAbstractSocketCachingOptionsClass     VZLinuxRosettaAbstractSocketCachingOptionsClass
	_VZLinuxRosettaAbstractSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaAbstractSocketCachingOptionsClass() VZLinuxRosettaAbstractSocketCachingOptionsClass {
	_VZLinuxRosettaAbstractSocketCachingOptionsClassOnce.Do(func() {
		_VZLinuxRosettaAbstractSocketCachingOptionsClass = VZLinuxRosettaAbstractSocketCachingOptionsClass{class: objc.GetClass("VZLinuxRosettaAbstractSocketCachingOptions")}
	})
	return _VZLinuxRosettaAbstractSocketCachingOptionsClass
}

// GetVZLinuxRosettaAbstractSocketCachingOptionsClass returns the class object for VZLinuxRosettaAbstractSocketCachingOptions.
func GetVZLinuxRosettaAbstractSocketCachingOptionsClass() VZLinuxRosettaAbstractSocketCachingOptionsClass {
	return getVZLinuxRosettaAbstractSocketCachingOptionsClass()
}

type VZLinuxRosettaAbstractSocketCachingOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxRosettaAbstractSocketCachingOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaAbstractSocketCachingOptionsClass) Alloc() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Caching options for an abstract socket.
//
// # Overview
// 
// Use this object to configure Rosetta to communicate with the Rosetta daemon
// using an abstract socket.
//
// # Initalizers
//
//   - [VZLinuxRosettaAbstractSocketCachingOptions.InitWithNameError]: Initialize options to set on a Rosetta directory share.
//
// # Accessing the socket name
//
//   - [VZLinuxRosettaAbstractSocketCachingOptions.Name]: The name of the abstract socket that Rosetta uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions
type VZLinuxRosettaAbstractSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaAbstractSocketCachingOptionsFromID constructs a [VZLinuxRosettaAbstractSocketCachingOptions] from an objc.ID.
//
// Caching options for an abstract socket.
func VZLinuxRosettaAbstractSocketCachingOptionsFromID(id objc.ID) VZLinuxRosettaAbstractSocketCachingOptions {
	return VZLinuxRosettaAbstractSocketCachingOptions{VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFromID(id)}
}
// Ensure VZLinuxRosettaAbstractSocketCachingOptions implements IVZLinuxRosettaAbstractSocketCachingOptions.
var _ IVZLinuxRosettaAbstractSocketCachingOptions = VZLinuxRosettaAbstractSocketCachingOptions{}

// An interface definition for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
//
// # Initalizers
//
//   - [IVZLinuxRosettaAbstractSocketCachingOptions.InitWithNameError]: Initialize options to set on a Rosetta directory share.
//
// # Accessing the socket name
//
//   - [IVZLinuxRosettaAbstractSocketCachingOptions.Name]: The name of the abstract socket that Rosetta uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions
type IVZLinuxRosettaAbstractSocketCachingOptions interface {
	IVZLinuxRosettaCachingOptions

	// Topic: Initalizers

	// Initialize options to set on a Rosetta directory share.
	InitWithNameError(name string) (VZLinuxRosettaAbstractSocketCachingOptions, error)

	// Topic: Accessing the socket name

	// The name of the abstract socket that Rosetta uses.
	Name() string
}

// Init initializes the instance.
func (l VZLinuxRosettaAbstractSocketCachingOptions) Init() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l VZLinuxRosettaAbstractSocketCachingOptions) Autorelease() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaAbstractSocketCachingOptions creates a new VZLinuxRosettaAbstractSocketCachingOptions instance.
func NewVZLinuxRosettaAbstractSocketCachingOptions() VZLinuxRosettaAbstractSocketCachingOptions {
	class := getVZLinuxRosettaAbstractSocketCachingOptionsClass()
	rv := objc.Send[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize options to set on a Rosetta directory share.
//
// name: This is the name of the abstract socket that Rosetta uses.
//
// error: If not `nil`, assigned with the error if the initialization fails.
//
// # Discussion
// 
// The `sockaddr_un` structure in Linux defines the maximum allowed length of
// `name`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/initWithName:error:
func NewLinuxRosettaAbstractSocketCachingOptionsWithNameError(name string) (VZLinuxRosettaAbstractSocketCachingOptions, error) {
	var errorPtr objc.ID
	instance := getVZLinuxRosettaAbstractSocketCachingOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:error:"), objc.String(name), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZLinuxRosettaAbstractSocketCachingOptions{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZLinuxRosettaAbstractSocketCachingOptionsFromID(rv), nil
}

// Initialize options to set on a Rosetta directory share.
//
// name: This is the name of the abstract socket that Rosetta uses.
//
// error: If not `nil`, assigned with the error if the initialization fails.
//
// # Discussion
// 
// The `sockaddr_un` structure in Linux defines the maximum allowed length of
// `name`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/initWithName:error:
func (l VZLinuxRosettaAbstractSocketCachingOptions) InitWithNameError(name string) (VZLinuxRosettaAbstractSocketCachingOptions, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("initWithName:error:"), objc.String(name), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZLinuxRosettaAbstractSocketCachingOptions{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZLinuxRosettaAbstractSocketCachingOptionsFromID(rv), nil

}

// The name of the abstract socket that Rosetta uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/name
func (l VZLinuxRosettaAbstractSocketCachingOptions) Name() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The maximum length of name that the framework allows.
//
// # Discussion
// 
// The `sockaddr_un` structure in Linux defines the maximum length for the
// path [Name].
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAbstractSocketCachingOptions/maximumNameLength
func (_VZLinuxRosettaAbstractSocketCachingOptionsClass VZLinuxRosettaAbstractSocketCachingOptionsClass) MaximumNameLength() uint {
	rv := objc.Send[uint](objc.ID(_VZLinuxRosettaAbstractSocketCachingOptionsClass.class), objc.Sel("maximumNameLength"))
	return rv
}

