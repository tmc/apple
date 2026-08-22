// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Support for decorating items.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemDecorating
type NSFileProviderItemDecorating interface {
	objectivec.IObject

	// Asks the item for an array of decorations.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemDecorating/decorations
	Decorations() []string
}

// NSFileProviderItemDecoratingObject wraps an existing Objective-C object that conforms to the NSFileProviderItemDecorating protocol.
type NSFileProviderItemDecoratingObject struct {
	objectivec.Object
}

func (o NSFileProviderItemDecoratingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderItemDecoratingObjectFromID constructs a [NSFileProviderItemDecoratingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderItemDecoratingObjectFromID(id objc.ID) NSFileProviderItemDecoratingObject {
	return NSFileProviderItemDecoratingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the item for an array of decorations.
//
// # Discussion
//
// The system calls this method to request the item’s decorations. Your
// implementation should return an array of
// [NSFileProviderItemDecorationIdentifier] instances.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemDecorating/decorations
func (o NSFileProviderItemDecoratingObject) Decorations() []string {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("decorations"))
	return objc.ConvertSliceToStrings(rvIDs)
}
