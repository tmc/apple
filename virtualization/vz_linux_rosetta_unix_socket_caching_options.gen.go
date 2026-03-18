// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZLinuxRosettaUnixSocketCachingOptions] class.
var (
	_VZLinuxRosettaUnixSocketCachingOptionsClass     VZLinuxRosettaUnixSocketCachingOptionsClass
	_VZLinuxRosettaUnixSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaUnixSocketCachingOptionsClass() VZLinuxRosettaUnixSocketCachingOptionsClass {
	_VZLinuxRosettaUnixSocketCachingOptionsClassOnce.Do(func() {
		_VZLinuxRosettaUnixSocketCachingOptionsClass = VZLinuxRosettaUnixSocketCachingOptionsClass{class: objc.GetClass("VZLinuxRosettaUnixSocketCachingOptions")}
	})
	return _VZLinuxRosettaUnixSocketCachingOptionsClass
}

// GetVZLinuxRosettaUnixSocketCachingOptionsClass returns the class object for VZLinuxRosettaUnixSocketCachingOptions.
func GetVZLinuxRosettaUnixSocketCachingOptionsClass() VZLinuxRosettaUnixSocketCachingOptionsClass {
	return getVZLinuxRosettaUnixSocketCachingOptionsClass()
}

type VZLinuxRosettaUnixSocketCachingOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaUnixSocketCachingOptionsClass) Alloc() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents caching options for a UNIX domain socket.
//
// # Overview
// 
// This object configures Rosetta to communicate with the Rosetta daemon using
// a UNIX domain socket.
//
// # Initializers
//
//   - [VZLinuxRosettaUnixSocketCachingOptions.InitWithPathError]: Creates a new Rosetta caching options object for a UNIX domain socket with the path you specify.
//
// # Accessing the socket path
//
//   - [VZLinuxRosettaUnixSocketCachingOptions.Path]: The path to the UNIX domain socket that Rosetta uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions
type VZLinuxRosettaUnixSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaUnixSocketCachingOptionsFromID constructs a [VZLinuxRosettaUnixSocketCachingOptions] from an objc.ID.
//
// An object that represents caching options for a UNIX domain socket.
func VZLinuxRosettaUnixSocketCachingOptionsFromID(id objc.ID) VZLinuxRosettaUnixSocketCachingOptions {
	return VZLinuxRosettaUnixSocketCachingOptions{VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFromID(id)}
}
// Ensure VZLinuxRosettaUnixSocketCachingOptions implements IVZLinuxRosettaUnixSocketCachingOptions.
var _ IVZLinuxRosettaUnixSocketCachingOptions = VZLinuxRosettaUnixSocketCachingOptions{}

// An interface definition for the [VZLinuxRosettaUnixSocketCachingOptions] class.
//
// # Initializers
//
//   - [IVZLinuxRosettaUnixSocketCachingOptions.InitWithPathError]: Creates a new Rosetta caching options object for a UNIX domain socket with the path you specify.
//
// # Accessing the socket path
//
//   - [IVZLinuxRosettaUnixSocketCachingOptions.Path]: The path to the UNIX domain socket that Rosetta uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions
type IVZLinuxRosettaUnixSocketCachingOptions interface {
	IVZLinuxRosettaCachingOptions

	// Topic: Initializers

	// Creates a new Rosetta caching options object for a UNIX domain socket with the path you specify.
	InitWithPathError(path string) (VZLinuxRosettaUnixSocketCachingOptions, error)

	// Topic: Accessing the socket path

	// The path to the UNIX domain socket that Rosetta uses.
	Path() string
}

// Init initializes the instance.
func (l VZLinuxRosettaUnixSocketCachingOptions) Init() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l VZLinuxRosettaUnixSocketCachingOptions) Autorelease() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaUnixSocketCachingOptions creates a new VZLinuxRosettaUnixSocketCachingOptions instance.
func NewVZLinuxRosettaUnixSocketCachingOptions() VZLinuxRosettaUnixSocketCachingOptions {
	class := getVZLinuxRosettaUnixSocketCachingOptionsClass()
	rv := objc.Send[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new Rosetta caching options object for a UNIX domain socket with
// the path you specify.
//
// path: The path of the UNIX Domain Socket that Rosetta uses.
//
// error: If not `nil`, assigned with the error if the initialization fails.
//
// # Discussion
// 
// You can optionally configure Rosetta to use cached translations from the
// Rosetta translation daemon communicating through a UNIX domain socket.
// 
// If `path` length exceeds [MaximumPathLength] in UTF-8 bytes, the framework
// returns `nil` and sets the `error` value, if available.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/initWithPath:error:
func NewLinuxRosettaUnixSocketCachingOptionsWithPathError(path string) (VZLinuxRosettaUnixSocketCachingOptions, error) {
	var errorPtr objc.ID
	instance := getVZLinuxRosettaUnixSocketCachingOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZLinuxRosettaUnixSocketCachingOptionsFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VZLinuxRosettaUnixSocketCachingOptionsFromID(rv), nil
}

// Creates a new Rosetta caching options object for a UNIX domain socket with
// the path you specify.
//
// path: The path of the UNIX Domain Socket that Rosetta uses.
//
// error: If not `nil`, assigned with the error if the initialization fails.
//
// # Discussion
// 
// You can optionally configure Rosetta to use cached translations from the
// Rosetta translation daemon communicating through a UNIX domain socket.
// 
// If `path` length exceeds [MaximumPathLength] in UTF-8 bytes, the framework
// returns `nil` and sets the `error` value, if available.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/initWithPath:error:
func (l VZLinuxRosettaUnixSocketCachingOptions) InitWithPathError(path string) (VZLinuxRosettaUnixSocketCachingOptions, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("initWithPath:error:"), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZLinuxRosettaUnixSocketCachingOptions{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZLinuxRosettaUnixSocketCachingOptionsFromID(rv), nil

}

// The path to the UNIX domain socket that Rosetta uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/path
func (l VZLinuxRosettaUnixSocketCachingOptions) Path() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("path"))
	return foundation.NSStringFromID(rv).String()
}

// The maximum allowed length of the path to the UNIX domain socket.
//
// # Discussion
// 
// The `sockaddr_un` structure in Linux defines the maximum length for this
// path.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaUnixSocketCachingOptions/maximumPathLength
func (_VZLinuxRosettaUnixSocketCachingOptionsClass VZLinuxRosettaUnixSocketCachingOptionsClass) MaximumPathLength() uint {
	rv := objc.Send[uint](objc.ID(_VZLinuxRosettaUnixSocketCachingOptionsClass.class), objc.Sel("maximumPathLength"))
	return rv
}

