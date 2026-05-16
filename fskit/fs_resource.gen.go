// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSResource] class.
var (
	_FSResourceClass     FSResourceClass
	_FSResourceClassOnce sync.Once
)

func getFSResourceClass() FSResourceClass {
	_FSResourceClassOnce.Do(func() {
		_FSResourceClass = FSResourceClass{class: objc.GetClass("FSResource")}
	})
	return _FSResourceClass
}

// GetFSResourceClass returns the class object for FSResource.
func GetFSResourceClass() FSResourceClass {
	return getFSResourceClass()
}

type FSResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSResourceClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSResourceClass) Alloc() FSResource {
	rv := objc.Send[FSResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An abstract resource a file system uses to provide data for a volume.
//
// # Overview
//
// [FSResource] is a base class to represent the various possible sources of
// data for a file system. These range from dedicated storage devices like
// hard drives and flash storage to network connections, and beyond.
// Subclasses define behavior specific to a given kind of resource, such as
// [FSBlockDeviceResource] for disk partition (IOMedia) file systems. These
// file systems are typical disk file systems such as HFS, APFS, ExFAT,
// ext2fs, or NTFS.
//
// A resource’s type also determines its life cycle. Resources based on
// block storage devices come into being when the system probes the media
// underlying the volumes and container. Other kinds of resources, like those
// based on URLs, might have different life cycles. For example, a resource
// based on a `//` URL might initialize when a person uses the “Connect to
// server” command in the macOS Finder.
//
// # Proxying resources
//
// Some resources, like [FSBlockDeviceResource], come in proxy and non-proxy
// variants. This addresses the issue that opening an external device like
// `/dev/disk2s1` requires an entitlement. Proxy resources allow unentitled
// clients of FSKit to describe which disk an [FSBlockDeviceResource] should
// represent. This allows, for example, the `mount(8)` tool to mount FSKit
// file systems on block devices when run as root. The tool uses a proxy when
// executing a command like `mount -t ffs /dev/disk2s1 /some/path`, which
// prevents leaking privileged resource access.
//
// # Creating proxies
//
//   - [FSResource.MakeProxy]: Creates a proxy object of this resource.
//
// # Revoking the resource
//
//   - [FSResource.Revoke]: Revokes the resource.
//   - [FSResource.Revoked]: A Boolean value that indicates whether the resource is revoked.
//
// See: https://developer.apple.com/documentation/FSKit/FSResource
type FSResource struct {
	objectivec.Object
}

// FSResourceFromID constructs a [FSResource] from an objc.ID.
//
// An abstract resource a file system uses to provide data for a volume.
func FSResourceFromID(id objc.ID) FSResource {
	return FSResource{objectivec.Object{ID: id}}
}

// NOTE: FSResource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSResource] class.
//
// # Creating proxies
//
//   - [IFSResource.MakeProxy]: Creates a proxy object of this resource.
//
// # Revoking the resource
//
//   - [IFSResource.Revoke]: Revokes the resource.
//   - [IFSResource.Revoked]: A Boolean value that indicates whether the resource is revoked.
//
// See: https://developer.apple.com/documentation/FSKit/FSResource
type IFSResource interface {
	objectivec.IObject

	// Topic: Creating proxies

	// Creates a proxy object of this resource.
	MakeProxy() IFSResource

	// Topic: Revoking the resource

	// Revokes the resource.
	Revoke()
	// A Boolean value that indicates whether the resource is revoked.
	Revoked() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r FSResource) Init() FSResource {
	rv := objc.Send[FSResource](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r FSResource) Autorelease() FSResource {
	rv := objc.Send[FSResource](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSResource creates a new FSResource instance.
func NewFSResource() FSResource {
	class := getFSResourceClass()
	rv := objc.Send[FSResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a proxy object of this resource.
//
// # Discussion
//
// If you create a proxy from a proxy resource, this method returns a copy of
// the proxy.
//
// See: https://developer.apple.com/documentation/FSKit/FSResource/makeProxy()
func (r FSResource) MakeProxy() IFSResource {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("makeProxy"))
	return FSResourceFromID(rv)
}

// Revokes the resource.
//
// # Discussion
//
// This method works by stripping away any underlying privileges associated
// with the resource. This effectively disconnects this object from its
// underlying resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSResource/revoke()
func (r FSResource) Revoke() {
	objc.Send[objc.ID](r.ID, objc.Sel("revoke"))
}
func (r FSResource) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether the resource is revoked.
//
// # Discussion
//
// If this is a proxy resource, the value of this property is always `true`
// (Swift) or [YES] (Objective-C).
//
// See: https://developer.apple.com/documentation/FSKit/FSResource/isRevoked
func (r FSResource) Revoked() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isRevoked"))
	return rv
}
