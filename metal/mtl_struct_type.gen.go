// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLStructType] class.
var (
	_MTLStructTypeClass     MTLStructTypeClass
	_MTLStructTypeClassOnce sync.Once
)

func getMTLStructTypeClass() MTLStructTypeClass {
	_MTLStructTypeClassOnce.Do(func() {
		_MTLStructTypeClass = MTLStructTypeClass{class: objc.GetClass("MTLStructType")}
	})
	return _MTLStructTypeClass
}

// GetMTLStructTypeClass returns the class object for MTLStructType.
func GetMTLStructTypeClass() MTLStructTypeClass {
	return getMTLStructTypeClass()
}

type MTLStructTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLStructTypeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLStructTypeClass) Alloc() MTLStructType {
	rv := objc.Send[MTLStructType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a structure.
//
// # Overview
// 
// [MTLStructType] is part of the reflection API that allows Metal framework
// code to query details of a struct that is passed as an argument of a Metal
// shading language function. Don’t create [MTLStructType] instances
// directly; instead query the [bufferStructType] property of an [MTLArgument]
// instance, or call the [StructType] method for an [MTLStructMember]
// instance. To examine the details of the struct, you can recursively drill
// down the [MTLStructType.Members] property of the [MTLStructType] instance, which contains
// details about struct members, each of which is represented by an
// [MTLStructMember] instance.
//
// [MTLArgument]: https://developer.apple.com/documentation/Metal/MTLArgument
// [bufferStructType]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferStructType
//
// # Obtaining information about struct members
//
//   - [MTLStructType.Members]: An array of instances that describe the fields in the struct.
//   - [MTLStructType.MemberByName]: Provides a representation of a struct member.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructType
type MTLStructType struct {
	MTLType
}

// MTLStructTypeFromID constructs a [MTLStructType] from an objc.ID.
//
// A description of a structure.
func MTLStructTypeFromID(id objc.ID) MTLStructType {
	return MTLStructType{MTLType: MTLTypeFromID(id)}
}
// NOTE: MTLStructType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLStructType] class.
//
// # Obtaining information about struct members
//
//   - [IMTLStructType.Members]: An array of instances that describe the fields in the struct.
//   - [IMTLStructType.MemberByName]: Provides a representation of a struct member.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructType
type IMTLStructType interface {
	IMTLType

	// Topic: Obtaining information about struct members

	// An array of instances that describe the fields in the struct.
	Members() []MTLStructMember
	// Provides a representation of a struct member.
	MemberByName(name string) IMTLStructMember

	// A description of the structure data of a buffer argument.
	BufferStructType() IMTLStructType
	SetBufferStructType(value IMTLStructType)
}

// Init initializes the instance.
func (s MTLStructType) Init() MTLStructType {
	rv := objc.Send[MTLStructType](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLStructType) Autorelease() MTLStructType {
	rv := objc.Send[MTLStructType](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLStructType creates a new MTLStructType instance.
func NewMTLStructType() MTLStructType {
	class := getMTLStructTypeClass()
	rv := objc.Send[MTLStructType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides a representation of a struct member.
//
// name: The name of a member in the struct.
//
// # Return Value
// 
// An object that represents the named struct member. If `name` does not match
// a member name, `nil` is returned.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructType/memberByName(_:)
func (s MTLStructType) MemberByName(name string) IMTLStructMember {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("memberByName:"), objc.String(name))
	return MTLStructMemberFromID(rv)
}

// An array of instances that describe the fields in the struct.
//
// # Discussion
// 
// Each array element in [Members] is an [MTLStructMember] instance that
// corresponds to one of the fields in the struct.
//
// See: https://developer.apple.com/documentation/Metal/MTLStructType/members
func (s MTLStructType) Members() []MTLStructMember {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("members"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLStructMember {
		return MTLStructMemberFromID(id)
	})
}
// A description of the structure data of a buffer argument.
//
// See: https://developer.apple.com/documentation/metal/mtlargument/bufferstructtype
func (s MTLStructType) BufferStructType() IMTLStructType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bufferStructType"))
	return MTLStructTypeFromID(objc.ID(rv))
}
func (s MTLStructType) SetBufferStructType(value IMTLStructType) {
	objc.Send[struct{}](s.ID, objc.Sel("setBufferStructType:"), value)
}

