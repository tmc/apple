// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSharedDirectory] class.
var (
	_VZSharedDirectoryClass     VZSharedDirectoryClass
	_VZSharedDirectoryClassOnce sync.Once
)

func getVZSharedDirectoryClass() VZSharedDirectoryClass {
	_VZSharedDirectoryClassOnce.Do(func() {
		_VZSharedDirectoryClass = VZSharedDirectoryClass{class: objc.GetClass("VZSharedDirectory")}
	})
	return _VZSharedDirectoryClass
}

// GetVZSharedDirectoryClass returns the class object for VZSharedDirectory.
func GetVZSharedDirectoryClass() VZSharedDirectoryClass {
	return getVZSharedDirectoryClass()
}

type VZSharedDirectoryClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSharedDirectoryClass) Alloc() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A directory on the host that you can expose to a guest.
//
// # Overview
// 
// This exposes a directory from the host file system to the guest.
//
// # Creating a Shared Directory
//
//   - [VZSharedDirectory.InitWithURLReadOnly]: Initialize with a host directory.
//
// # Accessing Directory Properties
//
//   - [VZSharedDirectory.URL]: A file URL to a directory on the host system to expose to the guest.
//   - [VZSharedDirectory.ReadOnly]: A Boolean value that indicates whether the directory is read-only to the guest.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory
type VZSharedDirectory struct {
	objectivec.Object
}

// VZSharedDirectoryFromID constructs a [VZSharedDirectory] from an objc.ID.
//
// A directory on the host that you can expose to a guest.
func VZSharedDirectoryFromID(id objc.ID) VZSharedDirectory {
	return VZSharedDirectory{objectivec.Object{ID: id}}
}
// NOTE: VZSharedDirectory adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZSharedDirectory] class.
//
// # Creating a Shared Directory
//
//   - [IVZSharedDirectory.InitWithURLReadOnly]: Initialize with a host directory.
//
// # Accessing Directory Properties
//
//   - [IVZSharedDirectory.URL]: A file URL to a directory on the host system to expose to the guest.
//   - [IVZSharedDirectory.ReadOnly]: A Boolean value that indicates whether the directory is read-only to the guest.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory
type IVZSharedDirectory interface {
	objectivec.IObject

	// Topic: Creating a Shared Directory

	// Initialize with a host directory.
	InitWithURLReadOnly(url foundation.INSURL, readOnly bool) VZSharedDirectory

	// Topic: Accessing Directory Properties

	// A file URL to a directory on the host system to expose to the guest.
	URL() foundation.INSURL
	// A Boolean value that indicates whether the directory is read-only to the guest.
	ReadOnly() bool
}

// Init initializes the instance.
func (s VZSharedDirectory) Init() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSharedDirectory) Autorelease() VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSharedDirectory creates a new VZSharedDirectory instance.
func NewVZSharedDirectory() VZSharedDirectory {
	class := getVZSharedDirectoryClass()
	rv := objc.Send[VZSharedDirectory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize with a host directory.
//
// url: A local file URL to expose to the guest.
//
// readOnly: A Boolean value that indicates whether to expose the directory as read-only
// to the guest.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/init(url:readOnly:)
func NewSharedDirectoryWithURLReadOnly(url foundation.INSURL, readOnly bool) VZSharedDirectory {
	instance := getVZSharedDirectoryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:readOnly:"), url, readOnly)
	return VZSharedDirectoryFromID(rv)
}

// Initialize with a host directory.
//
// url: A local file URL to expose to the guest.
//
// readOnly: A Boolean value that indicates whether to expose the directory as read-only
// to the guest.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/init(url:readOnly:)
func (s VZSharedDirectory) InitWithURLReadOnly(url foundation.INSURL, readOnly bool) VZSharedDirectory {
	rv := objc.Send[VZSharedDirectory](s.ID, objc.Sel("initWithURL:readOnly:"), url, readOnly)
	return rv
}

// A file URL to a directory on the host system to expose to the guest.
//
// # Discussion
// 
// The URL must point to an existing directory path in the host file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/url
func (s VZSharedDirectory) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the directory is read-only to the
// guest.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSharedDirectory/isReadOnly
func (s VZSharedDirectory) ReadOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isReadOnly"))
	return rv
}

