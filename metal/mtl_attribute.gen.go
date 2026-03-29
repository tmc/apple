// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAttribute] class.
var (
	_MTLAttributeClass     MTLAttributeClass
	_MTLAttributeClassOnce sync.Once
)

func getMTLAttributeClass() MTLAttributeClass {
	_MTLAttributeClassOnce.Do(func() {
		_MTLAttributeClass = MTLAttributeClass{class: objc.GetClass("MTLAttribute")}
	})
	return _MTLAttributeClass
}

// GetMTLAttributeClass returns the class object for MTLAttribute.
func GetMTLAttributeClass() MTLAttributeClass {
	return getMTLAttributeClass()
}

type MTLAttributeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLAttributeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAttributeClass) Alloc() MTLAttribute {
	rv := objc.Send[MTLAttribute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an attribute defined in the stage-in argument for
// a shader.
//
// # Reading an attribute’s properties
//
//   - [MTLAttribute.Name]: The name of the attribute.
//   - [MTLAttribute.AttributeIndex]: The index of the attribute, as declared in Metal shader source code.
//   - [MTLAttribute.AttributeType]: The data type for the attribute, as declared in Metal shader source code.
//   - [MTLAttribute.Active]: A Boolean value that indicates whether the attribute is active.
//   - [MTLAttribute.PatchControlPointData]: A Boolean value that indicates whether the attribute represents control point data.
//   - [MTLAttribute.PatchData]: A Boolean value that indicates whether the attribute represents tessellation patch data.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute
type MTLAttribute struct {
	objectivec.Object
}

// MTLAttributeFromID constructs a [MTLAttribute] from an objc.ID.
//
// An object that describes an attribute defined in the stage-in argument for
// a shader.
func MTLAttributeFromID(id objc.ID) MTLAttribute {
	return MTLAttribute{objectivec.Object{ID: id}}
}
// NOTE: MTLAttribute adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAttribute] class.
//
// # Reading an attribute’s properties
//
//   - [IMTLAttribute.Name]: The name of the attribute.
//   - [IMTLAttribute.AttributeIndex]: The index of the attribute, as declared in Metal shader source code.
//   - [IMTLAttribute.AttributeType]: The data type for the attribute, as declared in Metal shader source code.
//   - [IMTLAttribute.Active]: A Boolean value that indicates whether the attribute is active.
//   - [IMTLAttribute.PatchControlPointData]: A Boolean value that indicates whether the attribute represents control point data.
//   - [IMTLAttribute.PatchData]: A Boolean value that indicates whether the attribute represents tessellation patch data.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute
type IMTLAttribute interface {
	objectivec.IObject

	// Topic: Reading an attribute’s properties

	// The name of the attribute.
	Name() string
	// The index of the attribute, as declared in Metal shader source code.
	AttributeIndex() uint
	// The data type for the attribute, as declared in Metal shader source code.
	AttributeType() MTLDataType
	// A Boolean value that indicates whether the attribute is active.
	Active() bool
	// A Boolean value that indicates whether the attribute represents control point data.
	PatchControlPointData() bool
	// A Boolean value that indicates whether the attribute represents tessellation patch data.
	PatchData() bool
}

// Init initializes the instance.
func (a MTLAttribute) Init() MTLAttribute {
	rv := objc.Send[MTLAttribute](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAttribute) Autorelease() MTLAttribute {
	rv := objc.Send[MTLAttribute](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAttribute creates a new MTLAttribute instance.
func NewMTLAttribute() MTLAttribute {
	class := getMTLAttributeClass()
	rv := objc.Send[MTLAttribute](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of the attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute/name
func (a MTLAttribute) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The index of the attribute, as declared in Metal shader source code.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute/attributeIndex
func (a MTLAttribute) AttributeIndex() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("attributeIndex"))
	return rv
}
// The data type for the attribute, as declared in Metal shader source code.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute/attributeType
func (a MTLAttribute) AttributeType() MTLDataType {
	rv := objc.Send[MTLDataType](a.ID, objc.Sel("attributeType"))
	return MTLDataType(rv)
}
// A Boolean value that indicates whether the attribute is active.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute/isActive
func (a MTLAttribute) Active() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isActive"))
	return rv
}
// A Boolean value that indicates whether the attribute represents control
// point data.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute/isPatchControlPointData
func (a MTLAttribute) PatchControlPointData() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPatchControlPointData"))
	return rv
}
// A Boolean value that indicates whether the attribute represents
// tessellation patch data.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttribute/isPatchData
func (a MTLAttribute) PatchData() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPatchData"))
	return rv
}

