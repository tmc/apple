// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXRequest] class.
var (
	_AXRequestClass     AXRequestClass
	_AXRequestClassOnce sync.Once
)

func getAXRequestClass() AXRequestClass {
	_AXRequestClassOnce.Do(func() {
		_AXRequestClass = AXRequestClass{class: objc.GetClass("AXRequest")}
	})
	return _AXRequestClass
}

// GetAXRequestClass returns the class object for AXRequest.
func GetAXRequestClass() AXRequestClass {
	return getAXRequestClass()
}

type AXRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXRequestClass) Alloc() AXRequest {
	rv := objc.Send[AXRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [AXRequest.Technology]
//
// See: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest
type AXRequest struct {
	objectivec.Object
}

// AXRequestFromID constructs a [AXRequest] from an objc.ID.
func AXRequestFromID(id objc.ID) AXRequest {
	return AXRequest{objectivec.Object{ID: id}}
}

// NOTE: AXRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXRequest] class.
//
// # Instance Properties
//
//   - [IAXRequest.Technology]
//
// See: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest
type IAXRequest interface {
	objectivec.IObject

	// Topic: Instance Properties

	Technology() AXTechnology

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AXRequest) Init() AXRequest {
	rv := objc.Send[AXRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXRequest) Autorelease() AXRequest {
	rv := objc.Send[AXRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXRequest creates a new AXRequest instance.
func NewAXRequest() AXRequest {
	class := getAXRequestClass()
	rv := objc.Send[AXRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AXRequest) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/technology
func (a AXRequest) Technology() AXTechnology {
	rv := objc.Send[AXTechnology](a.ID, objc.Sel("technology"))
	return AXTechnology(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/current
func (_AXRequestClass AXRequestClass) CurrentRequest() AXRequest {
	rv := objc.Send[objc.ID](objc.ID(_AXRequestClass.class), objc.Sel("currentRequest"))
	return AXRequestFromID(objc.ID(rv))
}
