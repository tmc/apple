// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Support for suppressing user-interaction alerts.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderUserInteractionSuppressing
type NSFileProviderUserInteractionSuppressing interface {
	objectivec.IObject

	// Asks the File Provider extension if the user suppressed the specified interaction.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderUserInteractionSuppressing/isInteractionSuppressed(forIdentifier:)
	IsInteractionSuppressedForIdentifier(suppressionIdentifier string) bool

	// Tells the File Provider extension that the user wants to suppress the user interaction.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderUserInteractionSuppressing/setInteractionSuppressed(_:forIdentifier:)
	SetInteractionSuppressedForIdentifier(suppression bool, suppressionIdentifier string)
}

// NSFileProviderUserInteractionSuppressingObject wraps an existing Objective-C object that conforms to the NSFileProviderUserInteractionSuppressing protocol.
type NSFileProviderUserInteractionSuppressingObject struct {
	objectivec.Object
}

func (o NSFileProviderUserInteractionSuppressingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderUserInteractionSuppressingObjectFromID constructs a [NSFileProviderUserInteractionSuppressingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderUserInteractionSuppressingObjectFromID(id objc.ID) NSFileProviderUserInteractionSuppressingObject {
	return NSFileProviderUserInteractionSuppressingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the File Provider extension if the user suppressed the specified
// interaction.
//
// suppressionIdentifier: A unique identifier for the user interaction.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderUserInteractionSuppressing/isInteractionSuppressed(forIdentifier:)
func (o NSFileProviderUserInteractionSuppressingObject) IsInteractionSuppressedForIdentifier(suppressionIdentifier string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isInteractionSuppressedForIdentifier:"), objc.String(suppressionIdentifier))
	return rv
}

// Tells the File Provider extension that the user wants to suppress the user
// interaction.
//
// suppression: A Boolean value that indicates whether the user wants to suppress the
// specified user interaction.
//
// suppressionIdentifier: A unique identifier for the user interaction.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderUserInteractionSuppressing/setInteractionSuppressed(_:forIdentifier:)
func (o NSFileProviderUserInteractionSuppressingObject) SetInteractionSuppressedForIdentifier(suppression bool, suppressionIdentifier string) {
	objc.Send[struct{}](o.ID, objc.Sel("setInteractionSuppressed:forIdentifier:"), suppression, objc.String(suppressionIdentifier))
}
