// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Support for enumerating the file provider’s contents.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerating
type NSFileProviderEnumerating interface {
	objectivec.IObject

	// Tells the file provider to return an enumerator for the provided directory.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerating/enumerator(for:request:)
	EnumeratorForContainerItemIdentifierRequestError(containerItemIdentifier NSFileProviderItemIdentifier, request INSFileProviderRequest) (NSFileProviderEnumerator, error)
}

// NSFileProviderEnumeratingObject wraps an existing Objective-C object that conforms to the NSFileProviderEnumerating protocol.
type NSFileProviderEnumeratingObject struct {
	objectivec.Object
}

func (o NSFileProviderEnumeratingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderEnumeratingObjectFromID constructs a [NSFileProviderEnumeratingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderEnumeratingObjectFromID(id objc.ID) NSFileProviderEnumeratingObject {
	return NSFileProviderEnumeratingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the file provider to return an enumerator for the provided directory.
//
// containerItemIdentifier: The item identifier for the directory.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// # Return Value
//
// An enumerator for the specified directory.
//
// # Discussion
//
// The system calls this method to request an enumerator for the specified
// item.
//
// Possible item identifiers include:
//
// [rootContainer]: The system passes this identifier when the user begins
// browsing your file provider’s content. A directory’s [ItemIdentifier]:
// The system requests a new enumerator each time the user opens a new
// directory. [workingSet]: The system can request an enumerator so that it
// can sync the working set in the background. A document’s
// [ItemIdentifier]: The system subscribes to live updates by requesting an
// enumerator for a document. The [trashContainer] directory.: The system
// passes this identifier when the user browses the contents of the trash. If
// your File Provider extension doesn’t support moving items to the trash,
// your implementation should throw or return an error.
//
// Your implementation should create and return an [NSFileProviderEnumerator]
// object that provides the requested content.
//
// # Handle Errors
//
// If you can’t return the requested enumerator, you must throw an error in
// Swift, or if you return nil in Objective-C, you must set the `error` out
// parameter.
//
// If the `containerItemIdentifier` parameter is [trashContainer] and your
// extension doesn’t support trashing items, then it should fail with the
// [NSFeatureUnsupportedError] error code from the [NSCocoaErrorDomain]
// domain. Additionally, make sure the items managed by your File Provider
// extension don’t have the [NSFileProviderItemCapabilitiesAllowsTrashing]
// capability enabled.
//
// If the `containerItemIdentifier` parameter doesn’t exist in your remote
// storage, you should fail with an [NSFileProviderError.Code.noSuchItem]
// error. The system then attempts to delete the item from disk.
//
// If you pass [NSFileProviderError.Code.notAuthenticated] or
// [NSFileProviderError.Code.serverUnreachable] to the handler, the system
// presents an appropriate alert to the user, but doesn’t try to access the
// metadata until triggered again by the user.
//
// The system considers any other errors to be transient, and automatically
// retries the method call.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerating/enumerator(for:request:)
//
// [NSCocoaErrorDomain]: https://developer.apple.com/documentation/Foundation/NSCocoaErrorDomain
// [NSFeatureUnsupportedError]: https://developer.apple.com/documentation/Foundation/NSFeatureUnsupportedError-swift.var
// [NSFileProviderError.Code.noSuchItem]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
// [NSFileProviderError.Code.notAuthenticated]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/notAuthenticated
// [NSFileProviderError.Code.serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/serverUnreachable
// [rootContainer]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/rootContainer
// [trashContainer]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/trashContainer
// [workingSet]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/workingSet
func (o NSFileProviderEnumeratingObject) EnumeratorForContainerItemIdentifierRequestError(containerItemIdentifier NSFileProviderItemIdentifier, request INSFileProviderRequest) (NSFileProviderEnumerator, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("enumeratorForContainerItemIdentifier:request:error:"), objc.String(string(containerItemIdentifier)), request)
	if err != nil {
		return nil, err
	}
	return NSFileProviderEnumeratorObjectFromID(rv), nil
}
