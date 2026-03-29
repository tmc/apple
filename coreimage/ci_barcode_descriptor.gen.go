// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIBarcodeDescriptor] class.
var (
	_CIBarcodeDescriptorClass     CIBarcodeDescriptorClass
	_CIBarcodeDescriptorClassOnce sync.Once
)

func getCIBarcodeDescriptorClass() CIBarcodeDescriptorClass {
	_CIBarcodeDescriptorClassOnce.Do(func() {
		_CIBarcodeDescriptorClass = CIBarcodeDescriptorClass{class: objc.GetClass("CIBarcodeDescriptor")}
	})
	return _CIBarcodeDescriptorClass
}

// GetCIBarcodeDescriptorClass returns the class object for CIBarcodeDescriptor.
func GetCIBarcodeDescriptorClass() CIBarcodeDescriptorClass {
	return getCIBarcodeDescriptorClass()
}

type CIBarcodeDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIBarcodeDescriptorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIBarcodeDescriptorClass) Alloc() CIBarcodeDescriptor {
	rv := objc.Send[CIBarcodeDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class that represents a machine-readable code’s
// attributes.
//
// # Overview
// 
// Subclasses encapsulate the formal specification and fields specific to a
// code type. Each subclass is sufficient to recreate the unique symbol
// exactly as seen or used with a custom parser.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarcodeDescriptor
type CIBarcodeDescriptor struct {
	objectivec.Object
}

// CIBarcodeDescriptorFromID constructs a [CIBarcodeDescriptor] from an objc.ID.
//
// An abstract base class that represents a machine-readable code’s
// attributes.
func CIBarcodeDescriptorFromID(id objc.ID) CIBarcodeDescriptor {
	return CIBarcodeDescriptor{objectivec.Object{ID: id}}
}
// NOTE: CIBarcodeDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIBarcodeDescriptor] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarcodeDescriptor
type ICIBarcodeDescriptor interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (b CIBarcodeDescriptor) Init() CIBarcodeDescriptor {
	rv := objc.Send[CIBarcodeDescriptor](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b CIBarcodeDescriptor) Autorelease() CIBarcodeDescriptor {
	rv := objc.Send[CIBarcodeDescriptor](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIBarcodeDescriptor creates a new CIBarcodeDescriptor instance.
func NewCIBarcodeDescriptor() CIBarcodeDescriptor {
	class := getCIBarcodeDescriptorClass()
	rv := objc.Send[CIBarcodeDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (b CIBarcodeDescriptor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

