// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLStructMember] class.
var (
	_MTLStructMemberClass     MTLStructMemberClass
	_MTLStructMemberClassOnce sync.Once
)

func getMTLStructMemberClass() MTLStructMemberClass {
	_MTLStructMemberClassOnce.Do(func() {
		_MTLStructMemberClass = MTLStructMemberClass{class: objc.GetClass("MTLStructMember")}
	})
	return _MTLStructMemberClass
}

// GetMTLStructMemberClass returns the class object for MTLStructMember.
func GetMTLStructMemberClass() MTLStructMemberClass {
	return getMTLStructMemberClass()
}

type MTLStructMemberClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLStructMemberClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLStructMemberClass) Alloc() MTLStructMember {
	rv := objc.Send[MTLStructMember](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An instance that provides information about a field in a structure.
//
// # Overview
//
// [MTLStructMember] is part of the reflection API that allows Metal framework
// code to query details about an argument of a Metal shading language
// function. An [MTLStructMember] instance describes the data type of one
// field in a struct that is passed as an [MTLFunction] argument, which is
// represented by [MTLArgument].
//
// Don’t create [MTLStructMember] instances directly. You obtain an
// [MTLStructMember] instance from either the [MTLStructMember.Members] property or the
// [MemberByName] method of an [MTLStructType] instance. The [MTLStructMember.DataType]
// property of the [MTLStructMember] instance tells you what kind of data is
// stored in the member. Recursively drill down every struct member until you
// reach a data type that is neither a struct nor an array.
//
// # Describing the struct member
//
//   - [MTLStructMember.Name]: The name of the struct member.
//   - [MTLStructMember.DataType]: The data type of the struct member.
//   - [MTLStructMember.Offset]: The location of this member relative to the start of its struct, in bytes.
//   - [MTLStructMember.ArgumentIndex]: The index in the argument table that corresponds to the struct member.
//
// # Obtaining struct member details
//
//   - [MTLStructMember.ArrayType]: Provides a description of the underlying array when the struct member holds an array.
//   - [MTLStructMember.StructType]: Provides a description of the underlying struct when the struct member holds a struct.
//   - [MTLStructMember.PointerType]: Provides a description of the underlying pointer when the struct member holds a pointer.
//   - [MTLStructMember.TextureReferenceType]: Provides a description of the underlying texture when the struct member holds a texture.
//
// # Instance Methods
//
//   - [MTLStructMember.TensorReferenceType]: Provides a description of the underlying tensor type when this struct member holds a tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember
//
// [MTLArgument]: https://developer.apple.com/documentation/Metal/MTLArgument
type MTLStructMember struct {
	objectivec.Object
}

// MTLStructMemberFromID constructs a [MTLStructMember] from an objc.ID.
//
// An instance that provides information about a field in a structure.
func MTLStructMemberFromID(id objc.ID) MTLStructMember {
	return MTLStructMember{objectivec.Object{ID: id}}
}

// NOTE: MTLStructMember adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLStructMember] class.
//
// # Describing the struct member
//
//   - [IMTLStructMember.Name]: The name of the struct member.
//   - [IMTLStructMember.DataType]: The data type of the struct member.
//   - [IMTLStructMember.Offset]: The location of this member relative to the start of its struct, in bytes.
//   - [IMTLStructMember.ArgumentIndex]: The index in the argument table that corresponds to the struct member.
//
// # Obtaining struct member details
//
//   - [IMTLStructMember.ArrayType]: Provides a description of the underlying array when the struct member holds an array.
//   - [IMTLStructMember.StructType]: Provides a description of the underlying struct when the struct member holds a struct.
//   - [IMTLStructMember.PointerType]: Provides a description of the underlying pointer when the struct member holds a pointer.
//   - [IMTLStructMember.TextureReferenceType]: Provides a description of the underlying texture when the struct member holds a texture.
//
// # Instance Methods
//
//   - [IMTLStructMember.TensorReferenceType]: Provides a description of the underlying tensor type when this struct member holds a tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember
type IMTLStructMember interface {
	objectivec.IObject

	// Topic: Describing the struct member

	// The name of the struct member.
	Name() string
	// The data type of the struct member.
	DataType() MTLDataType
	// The location of this member relative to the start of its struct, in bytes.
	Offset() uint
	// The index in the argument table that corresponds to the struct member.
	ArgumentIndex() uint

	// Topic: Obtaining struct member details

	// Provides a description of the underlying array when the struct member holds an array.
	ArrayType() IMTLArrayType
	// Provides a description of the underlying struct when the struct member holds a struct.
	StructType() IMTLStructType
	// Provides a description of the underlying pointer when the struct member holds a pointer.
	PointerType() IMTLPointerType
	// Provides a description of the underlying texture when the struct member holds a texture.
	TextureReferenceType() IMTLTextureReferenceType

	// Topic: Instance Methods

	// Provides a description of the underlying tensor type when this struct member holds a tensor.
	TensorReferenceType() IMTLTensorReferenceType

	// An array of instances that describe the fields in the struct.
	Members() IMTLStructMember
	SetMembers(value IMTLStructMember)
}

// Init initializes the instance.
func (s MTLStructMember) Init() MTLStructMember {
	rv := objc.Send[MTLStructMember](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLStructMember) Autorelease() MTLStructMember {
	rv := objc.Send[MTLStructMember](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLStructMember creates a new MTLStructMember instance.
func NewMTLStructMember() MTLStructMember {
	class := getMTLStructMemberClass()
	rv := objc.Send[MTLStructMember](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides a description of the underlying array when the struct member holds
// an array.
//
// # Return Value
//
// An object that describes the array. If [DataType] indicates that this
// member is not an array, this method returns `nil.`
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/arrayType()
func (s MTLStructMember) ArrayType() IMTLArrayType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("arrayType"))
	return MTLArrayTypeFromID(rv)
}

// Provides a description of the underlying struct when the struct member
// holds a struct.
//
// # Return Value
//
// An object that describes the struct. If [DataType] indicates that this
// member is not a struct, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/structType()
func (s MTLStructMember) StructType() IMTLStructType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("structType"))
	return MTLStructTypeFromID(rv)
}

// Provides a description of the underlying pointer when the struct member
// holds a pointer.
//
// # Return Value
//
// An object that describes the pointer. If [DataType] indicates that this
// member isn’t a pointer, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/pointerType()
func (s MTLStructMember) PointerType() IMTLPointerType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("pointerType"))
	return MTLPointerTypeFromID(rv)
}

// Provides a description of the underlying texture when the struct member
// holds a texture.
//
// # Return Value
//
// An object that describes the texture. If [DataType] indicates that this
// member isn’t a texture, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/textureReferenceType()
func (s MTLStructMember) TextureReferenceType() IMTLTextureReferenceType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("textureReferenceType"))
	return MTLTextureReferenceTypeFromID(rv)
}

// Provides a description of the underlying tensor type when this struct
// member holds a tensor.
//
// # Return Value
//
// A description of the tensor type that this struct member holds, or `nil` if
// this struct member doesn’t hold a tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/tensorReferenceType()
func (s MTLStructMember) TensorReferenceType() IMTLTensorReferenceType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("tensorReferenceType"))
	return MTLTensorReferenceTypeFromID(rv)
}

// The name of the struct member.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/name
func (s MTLStructMember) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The data type of the struct member.
//
// # Discussion
//
// For information on possible values, see [MTLDataType]. If the value is
// [MTLDataType.array], then the [ArrayType] method returns an object that
// describes the underlying array. If the value is [MTLDataTypeStruct], then
// the [StructType] method returns an object that describes the underlying
// struct.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/dataType
//
// [MTLDataType.array]: https://developer.apple.com/documentation/Metal/MTLDataType/array
// [MTLDataType]: https://developer.apple.com/documentation/Metal/MTLDataType
func (s MTLStructMember) DataType() MTLDataType {
	rv := objc.Send[MTLDataType](s.ID, objc.Sel("dataType"))
	return MTLDataType(rv)
}

// The location of this member relative to the start of its struct, in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/offset
func (s MTLStructMember) Offset() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("offset"))
	return rv
}

// The index in the argument table that corresponds to the struct member.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructMember/argumentIndex
func (s MTLStructMember) ArgumentIndex() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("argumentIndex"))
	return rv
}

// An array of instances that describe the fields in the struct.
//
// See: https://developer.apple.com/documentation/metal/mtlstructtype/members
func (s MTLStructMember) Members() IMTLStructMember {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("members"))
	return MTLStructMemberFromID(objc.ID(rv))
}
func (s MTLStructMember) SetMembers(value IMTLStructMember) {
	objc.Send[struct{}](s.ID, objc.Sel("setMembers:"), value)
}
