// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZSingleDirectoryShare] class.
var (
	_VZSingleDirectoryShareClass     VZSingleDirectoryShareClass
	_VZSingleDirectoryShareClassOnce sync.Once
)

func getVZSingleDirectoryShareClass() VZSingleDirectoryShareClass {
	_VZSingleDirectoryShareClassOnce.Do(func() {
		_VZSingleDirectoryShareClass = VZSingleDirectoryShareClass{class: objc.GetClass("VZSingleDirectoryShare")}
	})
	return _VZSingleDirectoryShareClass
}

// GetVZSingleDirectoryShareClass returns the class object for VZSingleDirectoryShare.
func GetVZSingleDirectoryShareClass() VZSingleDirectoryShareClass {
	return getVZSingleDirectoryShareClass()
}

type VZSingleDirectoryShareClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSingleDirectoryShareClass) Alloc() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that defines the directory share for a single directory.
//
// # Overview
// 
// This directory share exposes a single directory from the host file system
// to the guest.
//
// # Creating a directory share
//
//   - [VZSingleDirectoryShare.InitWithDirectory]: Creates a directory share with a directory that you specify on the host.
//
// # Accessing directory properties
//
//   - [VZSingleDirectoryShare.Directory]: The directory on the host to share with the guest VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare
type VZSingleDirectoryShare struct {
	VZDirectoryShare
}

// VZSingleDirectoryShareFromID constructs a [VZSingleDirectoryShare] from an objc.ID.
//
// An object that defines the directory share for a single directory.
func VZSingleDirectoryShareFromID(id objc.ID) VZSingleDirectoryShare {
	return VZSingleDirectoryShare{VZDirectoryShare: VZDirectoryShareFromID(id)}
}
// NOTE: VZSingleDirectoryShare adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZSingleDirectoryShare] class.
//
// # Creating a directory share
//
//   - [IVZSingleDirectoryShare.InitWithDirectory]: Creates a directory share with a directory that you specify on the host.
//
// # Accessing directory properties
//
//   - [IVZSingleDirectoryShare.Directory]: The directory on the host to share with the guest VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare
type IVZSingleDirectoryShare interface {
	IVZDirectoryShare

	// Topic: Creating a directory share

	// Creates a directory share with a directory that you specify on the host.
	InitWithDirectory(directory IVZSharedDirectory) VZSingleDirectoryShare

	// Topic: Accessing directory properties

	// The directory on the host to share with the guest VM.
	Directory() IVZSharedDirectory
}





// Init initializes the instance.
func (s VZSingleDirectoryShare) Init() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSingleDirectoryShare) Autorelease() VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSingleDirectoryShare creates a new VZSingleDirectoryShare instance.
func NewVZSingleDirectoryShare() VZSingleDirectoryShare {
	class := getVZSingleDirectoryShareClass()
	rv := objc.Send[VZSingleDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a directory share with a directory that you specify on the host.
//
// directory: The directory to share.
//
// # Discussion
// 
// Use this method to share a directory on the host system with the guest VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare/init(directory:)
func NewSingleDirectoryShareWithDirectory(directory IVZSharedDirectory) VZSingleDirectoryShare {
	instance := getVZSingleDirectoryShareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDirectory:"), directory)
	return VZSingleDirectoryShareFromID(rv)
}







// Creates a directory share with a directory that you specify on the host.
//
// directory: The directory to share.
//
// # Discussion
// 
// Use this method to share a directory on the host system with the guest VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare/init(directory:)
func (s VZSingleDirectoryShare) InitWithDirectory(directory IVZSharedDirectory) VZSingleDirectoryShare {
	rv := objc.Send[VZSingleDirectoryShare](s.ID, objc.Sel("initWithDirectory:"), directory)
	return rv
}












// The directory on the host to share with the guest VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSingleDirectoryShare/directory
func (s VZSingleDirectoryShare) Directory() IVZSharedDirectory {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("directory"))
	return VZSharedDirectoryFromID(objc.ID(rv))
}























