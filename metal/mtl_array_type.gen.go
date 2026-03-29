// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLArrayType] class.
var (
	_MTLArrayTypeClass     MTLArrayTypeClass
	_MTLArrayTypeClassOnce sync.Once
)

func getMTLArrayTypeClass() MTLArrayTypeClass {
	_MTLArrayTypeClassOnce.Do(func() {
		_MTLArrayTypeClass = MTLArrayTypeClass{class: objc.GetClass("MTLArrayType")}
	})
	return _MTLArrayTypeClass
}

// GetMTLArrayTypeClass returns the class object for MTLArrayType.
func GetMTLArrayTypeClass() MTLArrayTypeClass {
	return getMTLArrayTypeClass()
}

type MTLArrayTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLArrayTypeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLArrayTypeClass) Alloc() MTLArrayType {
	rv := objc.Send[MTLArrayType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of an array.
//
// # Overview
// 
// An [MTLArrayType] instance provides details about an array parameter.
// Don’t create [MTLArrayType] instances directly; other reflection
// instances contain properties to determine if a parameter is an array and to
// obtain the [MTLArrayType] instance that describes the array.
//
// # Describing the array elements
//
//   - [MTLArrayType.ArrayLength]: The number of elements in the array.
//   - [MTLArrayType.ElementType]: The data type of the array’s elements.
//   - [MTLArrayType.Stride]: The stride between array elements, in bytes.
//   - [MTLArrayType.ArgumentIndexStride]: The stride, in bytes, between argument indices.
//
// # Obtaining details for complex array elements
//
//   - [MTLArrayType.ElementArrayType]: Provides a description of the underlying type when an array holds other arrays as its elements.
//   - [MTLArrayType.ElementStructType]: Provides a description of the underlying struct type when an array holds structs as its elements.
//   - [MTLArrayType.ElementPointerType]: Provides a description of the underlying pointer type when an array holds pointers as its elements.
//   - [MTLArrayType.ElementTextureReferenceType]: Provides a description of the underlying texture type when an array holds textures as its elements.
//
// # Instance Methods
//
//   - [MTLArrayType.ElementTensorReferenceType]: Provides a description of the underlying tensor type when this array holds tensors as its elements.
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType
type MTLArrayType struct {
	MTLType
}

// MTLArrayTypeFromID constructs a [MTLArrayType] from an objc.ID.
//
// A description of an array.
func MTLArrayTypeFromID(id objc.ID) MTLArrayType {
	return MTLArrayType{MTLType: MTLTypeFromID(id)}
}
// NOTE: MTLArrayType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLArrayType] class.
//
// # Describing the array elements
//
//   - [IMTLArrayType.ArrayLength]: The number of elements in the array.
//   - [IMTLArrayType.ElementType]: The data type of the array’s elements.
//   - [IMTLArrayType.Stride]: The stride between array elements, in bytes.
//   - [IMTLArrayType.ArgumentIndexStride]: The stride, in bytes, between argument indices.
//
// # Obtaining details for complex array elements
//
//   - [IMTLArrayType.ElementArrayType]: Provides a description of the underlying type when an array holds other arrays as its elements.
//   - [IMTLArrayType.ElementStructType]: Provides a description of the underlying struct type when an array holds structs as its elements.
//   - [IMTLArrayType.ElementPointerType]: Provides a description of the underlying pointer type when an array holds pointers as its elements.
//   - [IMTLArrayType.ElementTextureReferenceType]: Provides a description of the underlying texture type when an array holds textures as its elements.
//
// # Instance Methods
//
//   - [IMTLArrayType.ElementTensorReferenceType]: Provides a description of the underlying tensor type when this array holds tensors as its elements.
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType
type IMTLArrayType interface {
	IMTLType

	// Topic: Describing the array elements

	// The number of elements in the array.
	ArrayLength() uint
	// The data type of the array’s elements.
	ElementType() MTLDataType
	// The stride between array elements, in bytes.
	Stride() uint
	// The stride, in bytes, between argument indices.
	ArgumentIndexStride() uint

	// Topic: Obtaining details for complex array elements

	// Provides a description of the underlying type when an array holds other arrays as its elements.
	ElementArrayType() IMTLArrayType
	// Provides a description of the underlying struct type when an array holds structs as its elements.
	ElementStructType() IMTLStructType
	// Provides a description of the underlying pointer type when an array holds pointers as its elements.
	ElementPointerType() IMTLPointerType
	// Provides a description of the underlying texture type when an array holds textures as its elements.
	ElementTextureReferenceType() IMTLTextureReferenceType

	// Topic: Instance Methods

	// Provides a description of the underlying tensor type when this array holds tensors as its elements.
	ElementTensorReferenceType() IMTLTensorReferenceType
}

// Init initializes the instance.
func (a MTLArrayType) Init() MTLArrayType {
	rv := objc.Send[MTLArrayType](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLArrayType) Autorelease() MTLArrayType {
	rv := objc.Send[MTLArrayType](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLArrayType creates a new MTLArrayType instance.
func NewMTLArrayType() MTLArrayType {
	class := getMTLArrayTypeClass()
	rv := objc.Send[MTLArrayType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides a description of the underlying type when an array holds other
// arrays as its elements.
//
// # Return Value
// 
// Returns an object that describes an array. If the array elements aren’t
// arrays, this method returns `nil`.
//
// # Discussion
// 
// Use this method if [ElementType] is [DataTypeArray].
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/element()
func (a MTLArrayType) ElementArrayType() IMTLArrayType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("elementArrayType"))
	return MTLArrayTypeFromID(rv)
}
// Provides a description of the underlying struct type when an array holds
// structs as its elements.
//
// # Return Value
// 
// An object that describes the struct. If the array elements aren’t
// structs, this method returns `nil`.
//
// # Discussion
// 
// Use this method if [ElementType] is [DataTypeStruct].
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/elementStructType()
func (a MTLArrayType) ElementStructType() IMTLStructType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("elementStructType"))
	return MTLStructTypeFromID(rv)
}
// Provides a description of the underlying pointer type when an array holds
// pointers as its elements.
//
// # Return Value
// 
// An object that describes the pointer. If the array elements aren’t
// pointers, this method returns `nil`.
//
// # Discussion
// 
// Use this method if [ElementType] is [DataTypePointer].
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/elementPointerType()
func (a MTLArrayType) ElementPointerType() IMTLPointerType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("elementPointerType"))
	return MTLPointerTypeFromID(rv)
}
// Provides a description of the underlying texture type when an array holds
// textures as its elements.
//
// # Return Value
// 
// An object that describes the texture. If the array elements aren’t
// textures, this method returns `nil`.
//
// # Discussion
// 
// Use this method if [ElementType] is [DataTypeTexture].
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/elementTextureReferenceType()
func (a MTLArrayType) ElementTextureReferenceType() IMTLTextureReferenceType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("elementTextureReferenceType"))
	return MTLTextureReferenceTypeFromID(rv)
}
// Provides a description of the underlying tensor type when this array holds
// tensors as its elements.
//
// # Return Value
// 
// A description of the tensor type that this array holds, or `nil` if this
// struct member doesn’t hold a tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/elementTensorReferenceType()
func (a MTLArrayType) ElementTensorReferenceType() IMTLTensorReferenceType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("elementTensorReferenceType"))
	return MTLTensorReferenceTypeFromID(rv)
}

// The number of elements in the array.
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/arrayLength
func (a MTLArrayType) ArrayLength() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("arrayLength"))
	return rv
}
// The data type of the array’s elements.
//
// # Discussion
// 
// For information on possible values, see [MTLDataType].
//
// [MTLDataType]: https://developer.apple.com/documentation/Metal/MTLDataType
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/elementType
func (a MTLArrayType) ElementType() MTLDataType {
	rv := objc.Send[MTLDataType](a.ID, objc.Sel("elementType"))
	return MTLDataType(rv)
}
// The stride between array elements, in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/stride
func (a MTLArrayType) Stride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("stride"))
	return rv
}
// The stride, in bytes, between argument indices.
//
// See: https://developer.apple.com/documentation/Metal/MTLArrayType/argumentIndexStride
func (a MTLArrayType) ArgumentIndexStride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("argumentIndexStride"))
	return rv
}

