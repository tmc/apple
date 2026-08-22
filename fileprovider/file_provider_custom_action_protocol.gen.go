// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Support for custom actions.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderCustomAction
type NSFileProviderCustomAction interface {
	objectivec.IObject

	// Tells the File Provider extension to perform a custom action.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderCustomAction/performAction(identifier:onItemsWithIdentifiers:completionHandler:)
	PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler(actionIdentifier NSFileProviderExtensionActionIdentifier, itemIdentifiers []string, completionHandler ErrorHandler) foundation.Progress
}

// NSFileProviderCustomActionObject wraps an existing Objective-C object that conforms to the NSFileProviderCustomAction protocol.
type NSFileProviderCustomActionObject struct {
	objectivec.Object
}

func (o NSFileProviderCustomActionObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderCustomActionObjectFromID constructs a [NSFileProviderCustomActionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderCustomActionObjectFromID(id objc.ID) NSFileProviderCustomActionObject {
	return NSFileProviderCustomActionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the File Provider extension to perform a custom action.
//
// actionIdentifier: The identifier for the requested custom action from the extension’s
// `Info.Plist()` file.
//
// itemIdentifiers: A list of item identifiers affected by the action.
//
// completionHandler: A block that you call after completing the specified action. You pass the
// following parameters:
//
// error: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Return Value
//
// An item that tracks your extension’s progress.
//
// # Discussion
//
// Define the custom actions in the File Provider Extension’s `Info.Plist()`
// file, under the [NSExtensionFileProviderActions] key. The format of this
// key is identical to actions defined for a [File Provider UI] extension. For
// more information, see `Adding Actions to the Context Menu`.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderCustomAction/performAction(identifier:onItemsWithIdentifiers:completionHandler:)
//
// [File Provider UI]: https://developer.apple.com/documentation/FileProviderUI
func (o NSFileProviderCustomActionObject) PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler(actionIdentifier NSFileProviderExtensionActionIdentifier, itemIdentifiers []string, completionHandler ErrorHandler) foundation.Progress {
	_block2, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("performActionWithIdentifier:onItemsWithIdentifiers:completionHandler:"), actionIdentifier, itemIdentifiers, _block2)
	return foundation.NSProgressFromID(rv)
}
