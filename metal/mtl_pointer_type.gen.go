// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLPointerType] class.
var (
	_MTLPointerTypeClass     MTLPointerTypeClass
	_MTLPointerTypeClassOnce sync.Once
)

func getMTLPointerTypeClass() MTLPointerTypeClass {
	_MTLPointerTypeClassOnce.Do(func() {
		_MTLPointerTypeClass = MTLPointerTypeClass{class: objc.GetClass("MTLPointerType")}
	})
	return _MTLPointerTypeClass
}

// GetMTLPointerTypeClass returns the class object for MTLPointerType.
func GetMTLPointerTypeClass() MTLPointerTypeClass {
	return getMTLPointerTypeClass()
}

type MTLPointerTypeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLPointerTypeClass) Alloc() MTLPointerType {
	rv := objc.Send[MTLPointerType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A description of a pointer.
//
// # Describing the pointer elements
//
//   - [MTLPointerType.Alignment]: The required byte alignment in memory for the element data.
//   - [MTLPointerType.DataSize]: The size, in bytes, of the element data.
//   - [MTLPointerType.ElementType]: The data type of the element data.
//   - [MTLPointerType.Access]: The function’s read/write access to the element data.
//   - [MTLPointerType.ElementIsArgumentBuffer]: A Boolean value that indicates whether the element is an argument buffer.
//
// # Obtaining details for complex pointer elements
//
//   - [MTLPointerType.ElementArrayType]: Provides a description of the underlying array when the pointer points to an array.
//   - [MTLPointerType.ElementStructType]: Provides a description of the underlying struct when the pointer points to a struct.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType
type MTLPointerType struct {
	MTLType
}

// MTLPointerTypeFromID constructs a [MTLPointerType] from an objc.ID.
//
// A description of a pointer.
func MTLPointerTypeFromID(id objc.ID) MTLPointerType {
	return MTLPointerType{MTLType: MTLTypeFromID(id)}
}
// NOTE: MTLPointerType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLPointerType] class.
//
// # Describing the pointer elements
//
//   - [IMTLPointerType.Alignment]: The required byte alignment in memory for the element data.
//   - [IMTLPointerType.DataSize]: The size, in bytes, of the element data.
//   - [IMTLPointerType.ElementType]: The data type of the element data.
//   - [IMTLPointerType.Access]: The function’s read/write access to the element data.
//   - [IMTLPointerType.ElementIsArgumentBuffer]: A Boolean value that indicates whether the element is an argument buffer.
//
// # Obtaining details for complex pointer elements
//
//   - [IMTLPointerType.ElementArrayType]: Provides a description of the underlying array when the pointer points to an array.
//   - [IMTLPointerType.ElementStructType]: Provides a description of the underlying struct when the pointer points to a struct.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType
type IMTLPointerType interface {
	IMTLType

	// Topic: Describing the pointer elements

	// The required byte alignment in memory for the element data.
	Alignment() uint
	// The size, in bytes, of the element data.
	DataSize() uint
	// The data type of the element data.
	ElementType() MTLDataType
	// The function’s read/write access to the element data.
	Access() MTLBindingAccess
	// A Boolean value that indicates whether the element is an argument buffer.
	ElementIsArgumentBuffer() bool

	// Topic: Obtaining details for complex pointer elements

	// Provides a description of the underlying array when the pointer points to an array.
	ElementArrayType() IMTLArrayType
	// Provides a description of the underlying struct when the pointer points to a struct.
	ElementStructType() IMTLStructType
}





// Init initializes the instance.
func (p MTLPointerType) Init() MTLPointerType {
	rv := objc.Send[MTLPointerType](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MTLPointerType) Autorelease() MTLPointerType {
	rv := objc.Send[MTLPointerType](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLPointerType creates a new MTLPointerType instance.
func NewMTLPointerType() MTLPointerType {
	class := getMTLPointerTypeClass()
	rv := objc.Send[MTLPointerType](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Provides a description of the underlying array when the pointer points to
// an array.
//
// # Return Value
// 
// An object that describes the array. If the pointer does not point to an
// array, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType/elementArrayType()
func (p MTLPointerType) ElementArrayType() IMTLArrayType {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("elementArrayType"))
	return MTLArrayTypeFromID(rv)
}

// Provides a description of the underlying struct when the pointer points to
// a struct.
//
// # Return Value
// 
// An object that describes the struct. If the pointer does not point to an
// struct, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType/elementStructType()
func (p MTLPointerType) ElementStructType() IMTLStructType {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("elementStructType"))
	return MTLStructTypeFromID(rv)
}












// The required byte alignment in memory for the element data.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType/alignment
func (p MTLPointerType) Alignment() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("alignment"))
	return rv
}



// The size, in bytes, of the element data.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType/dataSize
func (p MTLPointerType) DataSize() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("dataSize"))
	return rv
}



// The data type of the element data.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType/elementType
func (p MTLPointerType) ElementType() MTLDataType {
	rv := objc.Send[MTLDataType](p.ID, objc.Sel("elementType"))
	return MTLDataType(rv)
}



// The function’s read/write access to the element data.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType/access
func (p MTLPointerType) Access() MTLBindingAccess {
	rv := objc.Send[MTLBindingAccess](p.ID, objc.Sel("access"))
	return MTLBindingAccess(rv)
}



// A Boolean value that indicates whether the element is an argument buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLPointerType/elementIsArgumentBuffer
func (p MTLPointerType) ElementIsArgumentBuffer() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("elementIsArgumentBuffer"))
	return rv
}























