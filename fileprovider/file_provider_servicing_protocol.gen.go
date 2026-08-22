// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Support for providing a custom service source.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServicing
type NSFileProviderServicing interface {
	objectivec.IObject

	// Asks the File Provider extension for an array of custom communication channels.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServicing/supportedServiceSources(for:completionHandler:)
	SupportedServiceSourcesForItemIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler) foundation.Progress
}

// NSFileProviderServicingObject wraps an existing Objective-C object that conforms to the NSFileProviderServicing protocol.
type NSFileProviderServicingObject struct {
	objectivec.Object
}

func (o NSFileProviderServicingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderServicingObjectFromID constructs a [NSFileProviderServicingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderServicingObjectFromID(id objc.ID) NSFileProviderServicingObject {
	return NSFileProviderServicingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the File Provider extension for an array of custom communication
// channels.
//
// itemIdentifier: The item’s identifier.
//
// completionHandler: A block that you call after gathering the service sources. You pass the
// following parameters:
//
// serviceSources: An array of service sources that lets you communicate with
// the host app. error: If an error occurs, this object contains information
// about the error; otherwise, it’s `nil`.
//
// # Return Value
//
// # An item that tracks the progress of the
//
// # Discussion
//
// The system calls this method when an app requests a list of supported
// services. Return an array of services for the specified file. An
// application with access to the file can request the supported services by
// calling the [FileManager] class’s
// [getFileProviderServicesForItem(at:completionHandler:)] method. For more
// information, see [NSFileProviderService].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServicing/supportedServiceSources(for:completionHandler:)
//
// [FileManager]: https://developer.apple.com/documentation/Foundation/FileManager
// [NSFileProviderService]: https://developer.apple.com/documentation/Foundation/NSFileProviderService
// [getFileProviderServicesForItem(at:completionHandler:)]: https://developer.apple.com/documentation/Foundation/FileManager/getFileProviderServicesForItem(at:completionHandler:)
func (o NSFileProviderServicingObject) SupportedServiceSourcesForItemIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler) foundation.Progress {
	_block1, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("supportedServiceSourcesForItemIdentifier:completionHandler:"), itemIdentifier, _block1)
	return foundation.NSProgressFromID(rv)
}
