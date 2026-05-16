// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSItemGetAttributesRequest] class.
var (
	_FSItemGetAttributesRequestClass     FSItemGetAttributesRequestClass
	_FSItemGetAttributesRequestClassOnce sync.Once
)

func getFSItemGetAttributesRequestClass() FSItemGetAttributesRequestClass {
	_FSItemGetAttributesRequestClassOnce.Do(func() {
		_FSItemGetAttributesRequestClass = FSItemGetAttributesRequestClass{class: objc.GetClass("FSItemGetAttributesRequest")}
	})
	return _FSItemGetAttributesRequestClass
}

// GetFSItemGetAttributesRequestClass returns the class object for FSItemGetAttributesRequest.
func GetFSItemGetAttributesRequestClass() FSItemGetAttributesRequestClass {
	return getFSItemGetAttributesRequestClass()
}

type FSItemGetAttributesRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSItemGetAttributesRequestClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSItemGetAttributesRequestClass) Alloc() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A request to get attributes from an item.
//
// # Overview
//
// Methods that retrieve attributes use this type and inspect the
// [FSItemGetAttributesRequest.WantedAttributes] property to determine which attributes to provide. FSKit
// calls the [FSItemGetAttributesRequest.IsAttributeWanted] method to determine whether the request
// requires a given attribute.
//
// # Inspecting requested attributes
//
//   - [FSItemGetAttributesRequest.WantedAttributes]: The attributes requested by the request.
//   - [FSItemGetAttributesRequest.SetWantedAttributes]
//   - [FSItemGetAttributesRequest.IsAttributeWanted]: A method that indicates whether the request wants given attribute.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest
type FSItemGetAttributesRequest struct {
	objectivec.Object
}

// FSItemGetAttributesRequestFromID constructs a [FSItemGetAttributesRequest] from an objc.ID.
//
// A request to get attributes from an item.
func FSItemGetAttributesRequestFromID(id objc.ID) FSItemGetAttributesRequest {
	return FSItemGetAttributesRequest{objectivec.Object{ID: id}}
}

// NOTE: FSItemGetAttributesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSItemGetAttributesRequest] class.
//
// # Inspecting requested attributes
//
//   - [IFSItemGetAttributesRequest.WantedAttributes]: The attributes requested by the request.
//   - [IFSItemGetAttributesRequest.SetWantedAttributes]
//   - [IFSItemGetAttributesRequest.IsAttributeWanted]: A method that indicates whether the request wants given attribute.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest
type IFSItemGetAttributesRequest interface {
	objectivec.IObject

	// Topic: Inspecting requested attributes

	// The attributes requested by the request.
	WantedAttributes() FSItemAttribute
	SetWantedAttributes(value FSItemAttribute)
	// A method that indicates whether the request wants given attribute.
	IsAttributeWanted(attribute FSItemAttribute) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i FSItemGetAttributesRequest) Init() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i FSItemGetAttributesRequest) Autorelease() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemGetAttributesRequest creates a new FSItemGetAttributesRequest instance.
func NewFSItemGetAttributesRequest() FSItemGetAttributesRequest {
	class := getFSItemGetAttributesRequestClass()
	rv := objc.Send[FSItemGetAttributesRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A method that indicates whether the request wants given attribute.
//
// attribute: The [FSItem.Attribute] to check.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/isAttributeWanted(_:)
//
// [FSItem.Attribute]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute
func (i FSItemGetAttributesRequest) IsAttributeWanted(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isAttributeWanted:"), attribute)
	return rv
}
func (i FSItemGetAttributesRequest) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The attributes requested by the request.
//
// # Discussion
//
// This property is a bit field in Objective-C and an [OptionSet] in Swift.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/wantedAttributes
//
// [OptionSet]: https://developer.apple.com/documentation/Swift/OptionSet
func (i FSItemGetAttributesRequest) WantedAttributes() FSItemAttribute {
	rv := objc.Send[FSItemAttribute](i.ID, objc.Sel("wantedAttributes"))
	return FSItemAttribute(rv)
}
func (i FSItemGetAttributesRequest) SetWantedAttributes(value FSItemAttribute) {
	objc.Send[struct{}](i.ID, objc.Sel("setWantedAttributes:"), value)
}
