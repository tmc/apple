// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// You implement this protocol when a class’s objects have subcomponents that can be discarded when not being used, thereby giving an application a smaller memory footprint.
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent
type NSDiscardableContent interface {
	objectivec.IObject

	// Returns a Boolean value indicating whether the discardable contents are still available and have been successfully accessed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/beginContentAccess()
	BeginContentAccess() bool

	// Called if the discardable contents are no longer being accessed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/endContentAccess()
	EndContentAccess()

	// Called to discard the contents of the receiver if the value of the accessed counter is 0.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/discardContentIfPossible()
	DiscardContentIfPossible()

	// Returns a Boolean value indicating whether the content has been discarded.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/isContentDiscarded()
	IsContentDiscarded() bool
}

// NSDiscardableContentObject wraps an existing Objective-C object that conforms to the NSDiscardableContent protocol.
type NSDiscardableContentObject struct {
	objectivec.Object
}
func (o NSDiscardableContentObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSDiscardableContentObjectFromID constructs a [NSDiscardableContentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSDiscardableContentObjectFromID(id objc.ID) NSDiscardableContentObject {
	return NSDiscardableContentObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a Boolean value indicating whether the discardable contents are
// still available and have been successfully accessed.
//
// # Return Value
// 
// YES if the discardable contents are still available and have now been
// successfully accessed; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Discussion
// 
// Call this method if the object’s memory is needed or is about to be used.
// This method increments the counter variable, thus protecting the object’s
// memory from possibly being discarded. The implementing class may decide
// that this method will try to recreate the contents if they have been
// discarded and return YES if the re-creation was successful. Implementors of
// this protocol should raise exceptions if the [NSDiscardableContent] objects
// are used when the `beginContentAccess` method has not been called on them.
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/beginContentAccess()

func (o NSDiscardableContentObject) BeginContentAccess() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("beginContentAccess"))
	return rv
	}

// Called if the discardable contents are no longer being accessed.
//
// # Discussion
// 
// This method decrements the counter variable of the object, which will
// usually bring the value of the counter variable back down to 0, which
// allows the discardable contents of the object to be thrown away if
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/endContentAccess()

func (o NSDiscardableContentObject) EndContentAccess() {
	
	objc.Send[struct{}](o.ID, objc.Sel("endContentAccess"))
	}

// Called to discard the contents of the receiver if the value of the accessed
// counter is 0.
//
// # Discussion
// 
// This method should only discard the contents of the object if the value of
// the accessed counter is 0. Otherwise, it should do nothing.
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/discardContentIfPossible()

func (o NSDiscardableContentObject) DiscardContentIfPossible() {
	
	objc.Send[struct{}](o.ID, objc.Sel("discardContentIfPossible"))
	}

// Returns a Boolean value indicating whether the content has been discarded.
//
// # Return Value
// 
// [true] if the content has been discarded; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDiscardableContent/isContentDiscarded()

func (o NSDiscardableContentObject) IsContentDiscarded() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isContentDiscarded"))
	return rv
	}

