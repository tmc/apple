// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A dynamically linkable representation of compiled shader code for a specific Metal device object.
//
// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary
type MTLDynamicLibrary interface {
	objectivec.IObject

	// The Metal device object that created the dynamic library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/device
	Device() MTLDevice

	// A file path for this dynamic library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/installName
	InstallName() string

	// A string that identifies the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/label
	Label() string

	// Writes the contents of the dynamic library to a file.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/serialize(to:)
	SerializeToURLError(url foundation.INSURL) (bool, error)

	// A string that identifies the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/label
	SetLabel(value string)
}

// MTLDynamicLibraryObject wraps an existing Objective-C object that conforms to the MTLDynamicLibrary protocol.
type MTLDynamicLibraryObject struct {
	objectivec.Object
}
func (o MTLDynamicLibraryObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLDynamicLibraryObjectFromID constructs a [MTLDynamicLibraryObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLDynamicLibraryObjectFromID(id objc.ID) MTLDynamicLibraryObject {
	return MTLDynamicLibraryObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The Metal device object that created the dynamic library.
//
// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/device
func (o MTLDynamicLibraryObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A file path for this dynamic library.
//
// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/installName
func (o MTLDynamicLibraryObject) InstallName() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("installName"))
	return foundation.NSStringFromID(rv).String()
	}
// A string that identifies the library.
//
// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/label
func (o MTLDynamicLibraryObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Writes the contents of the dynamic library to a file.
//
// url: The URL for the destination file.
//
// # Discussion
// 
// When the methods succeeds, the file contains a representation of the
// [MTLLibrary] from the [MTLDynamicLibrary] that creates it, as well as the
// binaries it has for the device your app is running on.
// 
// Such files may be combined with offline tools to contain the compiled code
// for multiple devices.
// 
// If this MTLDynamicLibrary was created from a file that contained compiled
// code for multiple devices, the compiled code for all other devices is not
// written (since only compiled code for the current device was loaded).
//
// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibrary/serialize(to:)
func (o MTLDynamicLibraryObject) SerializeToURLError(url foundation.INSURL) (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("serializeToURL:error:"), url)
	if err != nil {
		return false, err
	}
	return rv, nil
	}

func (o MTLDynamicLibraryObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

