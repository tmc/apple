// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [FSItemSetAttributesRequest] class.
var (
	_FSItemSetAttributesRequestClass     FSItemSetAttributesRequestClass
	_FSItemSetAttributesRequestClassOnce sync.Once
)

func getFSItemSetAttributesRequestClass() FSItemSetAttributesRequestClass {
	_FSItemSetAttributesRequestClassOnce.Do(func() {
		_FSItemSetAttributesRequestClass = FSItemSetAttributesRequestClass{class: objc.GetClass("FSItemSetAttributesRequest")}
	})
	return _FSItemSetAttributesRequestClass
}

// GetFSItemSetAttributesRequestClass returns the class object for FSItemSetAttributesRequest.
func GetFSItemSetAttributesRequestClass() FSItemSetAttributesRequestClass {
	return getFSItemSetAttributesRequestClass()
}

type FSItemSetAttributesRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSItemSetAttributesRequestClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSItemSetAttributesRequestClass) Alloc() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A request to set attributes on an item.
//
// # Overview
//
// Methods that take attributes use this type to receive attribute values and
// to indicate which attributes they support. The various members of the
// parent type, [FSItemAttributes], contain the values of the attributes to
// set.
//
// Modify the [FSItemSetAttributesRequest.ConsumedAttributes] property to indicate which attributes your
// file system successfully used. FSKit calls the [FSItemSetAttributesRequest.WasAttributeConsumed]
// method to determine whether the file system successfully used a given
// attribute. Only set the attributes that your file system supports.
//
// # Inspecting used attributes
//
//   - [FSItemSetAttributesRequest.ConsumedAttributes]: The attributes successfully used by the file system.
//   - [FSItemSetAttributesRequest.SetConsumedAttributes]
//   - [FSItemSetAttributesRequest.WasAttributeConsumed]: A method that indicates whether the file system used the given attribute.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest
type FSItemSetAttributesRequest struct {
	FSItemAttributes
}

// FSItemSetAttributesRequestFromID constructs a [FSItemSetAttributesRequest] from an objc.ID.
//
// A request to set attributes on an item.
func FSItemSetAttributesRequestFromID(id objc.ID) FSItemSetAttributesRequest {
	return FSItemSetAttributesRequest{FSItemAttributes: FSItemAttributesFromID(id)}
}

// NOTE: FSItemSetAttributesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSItemSetAttributesRequest] class.
//
// # Inspecting used attributes
//
//   - [IFSItemSetAttributesRequest.ConsumedAttributes]: The attributes successfully used by the file system.
//   - [IFSItemSetAttributesRequest.SetConsumedAttributes]
//   - [IFSItemSetAttributesRequest.WasAttributeConsumed]: A method that indicates whether the file system used the given attribute.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest
type IFSItemSetAttributesRequest interface {
	IFSItemAttributes

	// Topic: Inspecting used attributes

	// The attributes successfully used by the file system.
	ConsumedAttributes() FSItemAttribute
	SetConsumedAttributes(value FSItemAttribute)
	// A method that indicates whether the file system used the given attribute.
	WasAttributeConsumed(attribute FSItemAttribute) bool
}

// Init initializes the instance.
func (i FSItemSetAttributesRequest) Init() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i FSItemSetAttributesRequest) Autorelease() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemSetAttributesRequest creates a new FSItemSetAttributesRequest instance.
func NewFSItemSetAttributesRequest() FSItemSetAttributesRequest {
	class := getFSItemSetAttributesRequestClass()
	rv := objc.Send[FSItemSetAttributesRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A method that indicates whether the file system used the given attribute.
//
// attribute: The [FSItem.Attribute] to check.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/wasAttributeConsumed(_:)
//
// [FSItem.Attribute]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute
func (i FSItemSetAttributesRequest) WasAttributeConsumed(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("wasAttributeConsumed:"), attribute)
	return rv
}

// The attributes successfully used by the file system.
//
// # Discussion
//
// This property is a bit field in Objective-C and an [OptionSet] in Swift.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/consumedAttributes
//
// [OptionSet]: https://developer.apple.com/documentation/Swift/OptionSet
func (i FSItemSetAttributesRequest) ConsumedAttributes() FSItemAttribute {
	rv := objc.Send[FSItemAttribute](i.ID, objc.Sel("consumedAttributes"))
	return FSItemAttribute(rv)
}
func (i FSItemSetAttributesRequest) SetConsumedAttributes(value FSItemAttribute) {
	objc.Send[struct{}](i.ID, objc.Sel("setConsumedAttributes:"), value)
}
