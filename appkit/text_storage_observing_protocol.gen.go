// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Optional methods that delegates implement to handle editing and transaction processing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving
type NSTextStorageObserving interface {
	objectivec.IObject

	// TextStorage protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/textStorage
	TextStorage() NSTextStorage

	// PerformEditingTransactionForTextStorageUsingBlock protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/performEditingTransaction(for:using:)
	PerformEditingTransactionForTextStorageUsingBlock(textStorage NSTextStorage, transaction VoidHandler)

	// ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/processEditing(for:edited:range:changeInLength:invalidatedRange:)
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage NSTextStorage, editMask NSTextStorageEditActions, newCharRange foundation.NSRange, delta int, invalidatedCharRange foundation.NSRange)

	// textStorage protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/textStorage
	SetTextStorage(value NSTextStorage)
}

// NSTextStorageObservingObject wraps an existing Objective-C object that conforms to the NSTextStorageObserving protocol.
type NSTextStorageObservingObject struct {
	objectivec.Object
}
func (o NSTextStorageObservingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextStorageObservingObjectFromID constructs a [NSTextStorageObservingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextStorageObservingObjectFromID(id objc.ID) NSTextStorageObservingObject {
	return NSTextStorageObservingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/textStorage
func (o NSTextStorageObservingObject) TextStorage() NSTextStorage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textStorage"))
	return NSTextStorageFromID(rv)
	}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/performEditingTransaction(for:using:)
func (o NSTextStorageObservingObject) PerformEditingTransactionForTextStorageUsingBlock(textStorage NSTextStorage, transaction VoidHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("performEditingTransactionForTextStorage:usingBlock:"), textStorage, transaction)
	}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/processEditing(for:edited:range:changeInLength:invalidatedRange:)
func (o NSTextStorageObservingObject) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage NSTextStorage, editMask NSTextStorageEditActions, newCharRange foundation.NSRange, delta int, invalidatedCharRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), textStorage, editMask, newCharRange, delta, invalidatedCharRange)
	}

func (o NSTextStorageObservingObject) SetTextStorage(value NSTextStorage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextStorage:"), value)
}

