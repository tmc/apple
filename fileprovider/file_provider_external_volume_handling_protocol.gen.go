// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the interface for handling external volumes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderExternalVolumeHandling
type NSFileProviderExternalVolumeHandling interface {
	objectivec.IObject

	// Determines whether to connect to a domain from another device.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderExternalVolumeHandling/shouldConnectExternalDomain(completionHandler:)
	ShouldConnectExternalDomainWithCompletionHandler(completionHandler ErrorHandler)
}

// NSFileProviderExternalVolumeHandlingObject wraps an existing Objective-C object that conforms to the NSFileProviderExternalVolumeHandling protocol.
type NSFileProviderExternalVolumeHandlingObject struct {
	objectivec.Object
}

func (o NSFileProviderExternalVolumeHandlingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderExternalVolumeHandlingObjectFromID constructs a [NSFileProviderExternalVolumeHandlingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderExternalVolumeHandlingObjectFromID(id objc.ID) NSFileProviderExternalVolumeHandlingObject {
	return NSFileProviderExternalVolumeHandlingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Determines whether to connect to a domain from another device.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderExternalVolumeHandling/shouldConnectExternalDomain(completionHandler:)
func (o NSFileProviderExternalVolumeHandlingObject) ShouldConnectExternalDomainWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("shouldConnectExternalDomainWithCompletionHandler:"), _block0)
}
