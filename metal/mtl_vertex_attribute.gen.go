// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLVertexAttribute] class.
var (
	_MTLVertexAttributeClass     MTLVertexAttributeClass
	_MTLVertexAttributeClassOnce sync.Once
)

func getMTLVertexAttributeClass() MTLVertexAttributeClass {
	_MTLVertexAttributeClassOnce.Do(func() {
		_MTLVertexAttributeClass = MTLVertexAttributeClass{class: objc.GetClass("MTLVertexAttribute")}
	})
	return _MTLVertexAttributeClass
}

// GetMTLVertexAttributeClass returns the class object for MTLVertexAttribute.
func GetMTLVertexAttributeClass() MTLVertexAttributeClass {
	return getMTLVertexAttributeClass()
}

type MTLVertexAttributeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLVertexAttributeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLVertexAttributeClass) Alloc() MTLVertexAttribute {
	rv := objc.Send[MTLVertexAttribute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An instance that represents an attribute of a vertex function.
//
// # Overview
// 
// An [MTLVertexAttribute] instance represents an attribute for per-vertex
// input in a vertex function. You use vertex attribute instances to inspect
// the inputs of a vertex function by examining the [MTLVertexAttribute.VertexAttributes]
// property of the corresponding [MTLFunction] instance.
//
// # Describing the attribute
//
//   - [MTLVertexAttribute.Name]: The name of the attribute.
//   - [MTLVertexAttribute.AttributeIndex]: The index of the attribute, as declared in Metal shader source code.
//   - [MTLVertexAttribute.AttributeType]: The data type for the attribute, as declared in Metal shader source code.
//   - [MTLVertexAttribute.Active]: A Boolean value that indicates whether this vertex attribute is active.
//   - [MTLVertexAttribute.PatchControlPointData]: A Boolean value that indicates whether this vertex attribute represents control point data.
//   - [MTLVertexAttribute.PatchData]: A Boolean value that indicates whether this vertex attribute represents patch data.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute
type MTLVertexAttribute struct {
	objectivec.Object
}

// MTLVertexAttributeFromID constructs a [MTLVertexAttribute] from an objc.ID.
//
// An instance that represents an attribute of a vertex function.
func MTLVertexAttributeFromID(id objc.ID) MTLVertexAttribute {
	return MTLVertexAttribute{objectivec.Object{ID: id}}
}
// NOTE: MTLVertexAttribute adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLVertexAttribute] class.
//
// # Describing the attribute
//
//   - [IMTLVertexAttribute.Name]: The name of the attribute.
//   - [IMTLVertexAttribute.AttributeIndex]: The index of the attribute, as declared in Metal shader source code.
//   - [IMTLVertexAttribute.AttributeType]: The data type for the attribute, as declared in Metal shader source code.
//   - [IMTLVertexAttribute.Active]: A Boolean value that indicates whether this vertex attribute is active.
//   - [IMTLVertexAttribute.PatchControlPointData]: A Boolean value that indicates whether this vertex attribute represents control point data.
//   - [IMTLVertexAttribute.PatchData]: A Boolean value that indicates whether this vertex attribute represents patch data.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute
type IMTLVertexAttribute interface {
	objectivec.IObject

	// Topic: Describing the attribute

	// The name of the attribute.
	Name() string
	// The index of the attribute, as declared in Metal shader source code.
	AttributeIndex() uint
	// The data type for the attribute, as declared in Metal shader source code.
	AttributeType() MTLDataType
	// A Boolean value that indicates whether this vertex attribute is active.
	Active() bool
	// A Boolean value that indicates whether this vertex attribute represents control point data.
	PatchControlPointData() bool
	// A Boolean value that indicates whether this vertex attribute represents patch data.
	PatchData() bool

	// An array that describes the vertex input attributes to a vertex function.
	VertexAttributes() IMTLVertexAttribute
	SetVertexAttributes(value IMTLVertexAttribute)
}

// Init initializes the instance.
func (v MTLVertexAttribute) Init() MTLVertexAttribute {
	rv := objc.Send[MTLVertexAttribute](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTLVertexAttribute) Autorelease() MTLVertexAttribute {
	rv := objc.Send[MTLVertexAttribute](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVertexAttribute creates a new MTLVertexAttribute instance.
func NewMTLVertexAttribute() MTLVertexAttribute {
	class := getMTLVertexAttributeClass()
	rv := objc.Send[MTLVertexAttribute](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of the attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/name
func (v MTLVertexAttribute) Name() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The index of the attribute, as declared in Metal shader source code.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/attributeIndex
func (v MTLVertexAttribute) AttributeIndex() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("attributeIndex"))
	return rv
}
// The data type for the attribute, as declared in Metal shader source code.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/attributeType
func (v MTLVertexAttribute) AttributeType() MTLDataType {
	rv := objc.Send[MTLDataType](v.ID, objc.Sel("attributeType"))
	return MTLDataType(rv)
}
// A Boolean value that indicates whether this vertex attribute is active.
//
// # Discussion
// 
// If [false], this attribute is inactive and can be ignored.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/isActive
func (v MTLVertexAttribute) Active() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isActive"))
	return rv
}
// A Boolean value that indicates whether this vertex attribute represents
// control point data.
//
// # Discussion
// 
// This value is always [false] if the vertex function is not a
// post-tessellation vertex function.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/isPatchControlPointData
func (v MTLVertexAttribute) PatchControlPointData() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isPatchControlPointData"))
	return rv
}
// A Boolean value that indicates whether this vertex attribute represents
// patch data.
//
// # Discussion
// 
// This value is always [false] if the vertex function is not a
// post-tessellation vertex function.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttribute/isPatchData
func (v MTLVertexAttribute) PatchData() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isPatchData"))
	return rv
}
// An array that describes the vertex input attributes to a vertex function.
//
// See: https://developer.apple.com/documentation/metal/mtlfunction/vertexattributes
func (v MTLVertexAttribute) VertexAttributes() IMTLVertexAttribute {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("vertexAttributes"))
	return MTLVertexAttributeFromID(objc.ID(rv))
}
func (v MTLVertexAttribute) SetVertexAttributes(value IMTLVertexAttribute) {
	objc.Send[struct{}](v.ID, objc.Sel("setVertexAttributes:"), value)
}

