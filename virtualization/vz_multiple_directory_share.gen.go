// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZMultipleDirectoryShare] class.
var (
	_VZMultipleDirectoryShareClass     VZMultipleDirectoryShareClass
	_VZMultipleDirectoryShareClassOnce sync.Once
)

func getVZMultipleDirectoryShareClass() VZMultipleDirectoryShareClass {
	_VZMultipleDirectoryShareClassOnce.Do(func() {
		_VZMultipleDirectoryShareClass = VZMultipleDirectoryShareClass{class: objc.GetClass("VZMultipleDirectoryShare")}
	})
	return _VZMultipleDirectoryShareClass
}

// GetVZMultipleDirectoryShareClass returns the class object for VZMultipleDirectoryShare.
func GetVZMultipleDirectoryShareClass() VZMultipleDirectoryShareClass {
	return getVZMultipleDirectoryShareClass()
}

type VZMultipleDirectoryShareClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMultipleDirectoryShareClass) Alloc() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that describes a directory share for multiple directories.
//
// # Overview
// 
// This directory share exposes multiple directories from the host file system
// to the guest VM.
//
// # Creating a directory share
//
//   - [VZMultipleDirectoryShare.InitWithDirectories]: Creates the directory share with a set of directories on the host.
//
// # Accessing the shared directories
//
//   - [VZMultipleDirectoryShare.Directories]: The directories on the host to expose to the guest.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare
type VZMultipleDirectoryShare struct {
	VZDirectoryShare
}

// VZMultipleDirectoryShareFromID constructs a [VZMultipleDirectoryShare] from an objc.ID.
//
// An object that describes a directory share for multiple directories.
func VZMultipleDirectoryShareFromID(id objc.ID) VZMultipleDirectoryShare {
	return VZMultipleDirectoryShare{VZDirectoryShare: VZDirectoryShareFromID(id)}
}
// NOTE: VZMultipleDirectoryShare adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMultipleDirectoryShare] class.
//
// # Creating a directory share
//
//   - [IVZMultipleDirectoryShare.InitWithDirectories]: Creates the directory share with a set of directories on the host.
//
// # Accessing the shared directories
//
//   - [IVZMultipleDirectoryShare.Directories]: The directories on the host to expose to the guest.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare
type IVZMultipleDirectoryShare interface {
	IVZDirectoryShare

	// Topic: Creating a directory share

	// Creates the directory share with a set of directories on the host.
	InitWithDirectories(directories foundation.INSDictionary) VZMultipleDirectoryShare

	// Topic: Accessing the shared directories

	// The directories on the host to expose to the guest.
	Directories() foundation.INSDictionary
}





// Init initializes the instance.
func (m VZMultipleDirectoryShare) Init() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMultipleDirectoryShare) Autorelease() VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMultipleDirectoryShare creates a new VZMultipleDirectoryShare instance.
func NewVZMultipleDirectoryShare() VZMultipleDirectoryShare {
	class := getVZMultipleDirectoryShareClass()
	rv := objc.Send[VZMultipleDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates the directory share with a set of directories on the host.
//
// directories: Directories on the host to expose to the guest VM by name.
//
// # Discussion
// 
// The dictionary string keys are the names for the directory. The keys must
// be valid names or the system raises an exception and the app exits.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/init(directories:)
func NewMultipleDirectoryShareWithDirectories(directories foundation.INSDictionary) VZMultipleDirectoryShare {
	instance := getVZMultipleDirectoryShareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDirectories:"), directories)
	return VZMultipleDirectoryShareFromID(rv)
}







// Creates the directory share with a set of directories on the host.
//
// directories: Directories on the host to expose to the guest VM by name.
//
// # Discussion
// 
// The dictionary string keys are the names for the directory. The keys must
// be valid names or the system raises an exception and the app exits.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/init(directories:)
func (m VZMultipleDirectoryShare) InitWithDirectories(directories foundation.INSDictionary) VZMultipleDirectoryShare {
	rv := objc.Send[VZMultipleDirectoryShare](m.ID, objc.Sel("initWithDirectories:"), directories)
	return rv
}





// Transforms a string to be a valid directory name.
//
// name: The name to transform.
//
// # Return Value
// 
// Returns a String with the canonicalized name, or `nil` if there was an
// error.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/canonicalizedName(from:)
func (_VZMultipleDirectoryShareClass VZMultipleDirectoryShareClass) CanonicalizedNameFromName(name string) string {
	rv := objc.Send[objc.ID](objc.ID(_VZMultipleDirectoryShareClass.class), objc.Sel("canonicalizedNameFromName:"), objc.String(name))
	return foundation.NSStringFromID(rv).String()
}

// Check if a name is a valid directory name.
//
// name: The name to validate.
//
// # Discussion
// 
// The name must not be empty, have characters unsafe for file systems, be
// longer than `NAME_MAX`, or using a reserved name such as the Unix “.”
// or “..” current and parent directory filenames.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/validateName(_:)
func (_VZMultipleDirectoryShareClass VZMultipleDirectoryShareClass) ValidateNameError(name string) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_VZMultipleDirectoryShareClass.class), objc.Sel("validateName:error:"), objc.String(name), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateName:error: returned NO with nil NSError")
	}
	return rv, nil

}








// The directories on the host to expose to the guest.
//
// # Discussion
// 
// The dictionary string keys are the names for the directory. The keys must
// be valid names or the system raises an exception.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMultipleDirectoryShare/directories
func (m VZMultipleDirectoryShare) Directories() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("directories"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}























