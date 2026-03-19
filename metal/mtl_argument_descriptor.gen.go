// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLArgumentDescriptor] class.
var (
	_MTLArgumentDescriptorClass     MTLArgumentDescriptorClass
	_MTLArgumentDescriptorClassOnce sync.Once
)

func getMTLArgumentDescriptorClass() MTLArgumentDescriptorClass {
	_MTLArgumentDescriptorClassOnce.Do(func() {
		_MTLArgumentDescriptorClass = MTLArgumentDescriptorClass{class: objc.GetClass("MTLArgumentDescriptor")}
	})
	return _MTLArgumentDescriptorClass
}

// GetMTLArgumentDescriptorClass returns the class object for MTLArgumentDescriptor.
func GetMTLArgumentDescriptorClass() MTLArgumentDescriptorClass {
	return getMTLArgumentDescriptorClass()
}

type MTLArgumentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLArgumentDescriptorClass) Alloc() MTLArgumentDescriptor {
	rv := objc.Send[MTLArgumentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an argument within an argument buffer.
//
// # Overview
// 
// This descriptor can represent arguments within flat structures only. It can
// represent arrays of allowed argument buffer data types, but it cannot
// represent arguments within nested structures. Argument buffers with simple,
// flat structures can be represented by an array of [MTLArgumentDescriptor]
// instances. You can then use this array to create an [MTLArgumentEncoder]
// instance by calling the [NewArgumentEncoderWithArguments] method. Argument
// buffers with complex, nested structures need to define their structure in
// Metal shading language code, which can then be directly assigned to a
// specific buffer index of a function. You can then use this buffer index to
// call the [NewArgumentEncoderWithBufferIndex] method and create an
// [MTLArgumentEncoder] instance.
//
// # Setting the descriptor’s properties
//
//   - [MTLArgumentDescriptor.DataType]: The data type of the argument.
//   - [MTLArgumentDescriptor.SetDataType]
//   - [MTLArgumentDescriptor.Index]: The index ID of the argument.
//   - [MTLArgumentDescriptor.SetIndex]
//   - [MTLArgumentDescriptor.Access]: The access permissions of the argument.
//   - [MTLArgumentDescriptor.SetAccess]
//   - [MTLArgumentDescriptor.ArrayLength]: The length of an array argument.
//   - [MTLArgumentDescriptor.SetArrayLength]
//   - [MTLArgumentDescriptor.ConstantBlockAlignment]: The alignment of the constant block.
//   - [MTLArgumentDescriptor.SetConstantBlockAlignment]
//   - [MTLArgumentDescriptor.TextureType]: The texture type of a texture argument.
//   - [MTLArgumentDescriptor.SetTextureType]
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor
type MTLArgumentDescriptor struct {
	objectivec.Object
}

// MTLArgumentDescriptorFromID constructs a [MTLArgumentDescriptor] from an objc.ID.
//
// A representation of an argument within an argument buffer.
func MTLArgumentDescriptorFromID(id objc.ID) MTLArgumentDescriptor {
	return MTLArgumentDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLArgumentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLArgumentDescriptor] class.
//
// # Setting the descriptor’s properties
//
//   - [IMTLArgumentDescriptor.DataType]: The data type of the argument.
//   - [IMTLArgumentDescriptor.SetDataType]
//   - [IMTLArgumentDescriptor.Index]: The index ID of the argument.
//   - [IMTLArgumentDescriptor.SetIndex]
//   - [IMTLArgumentDescriptor.Access]: The access permissions of the argument.
//   - [IMTLArgumentDescriptor.SetAccess]
//   - [IMTLArgumentDescriptor.ArrayLength]: The length of an array argument.
//   - [IMTLArgumentDescriptor.SetArrayLength]
//   - [IMTLArgumentDescriptor.ConstantBlockAlignment]: The alignment of the constant block.
//   - [IMTLArgumentDescriptor.SetConstantBlockAlignment]
//   - [IMTLArgumentDescriptor.TextureType]: The texture type of a texture argument.
//   - [IMTLArgumentDescriptor.SetTextureType]
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor
type IMTLArgumentDescriptor interface {
	objectivec.IObject

	// Topic: Setting the descriptor’s properties

	// The data type of the argument.
	DataType() MTLDataType
	SetDataType(value MTLDataType)
	// The index ID of the argument.
	Index() uint
	SetIndex(value uint)
	// The access permissions of the argument.
	Access() MTLBindingAccess
	SetAccess(value MTLBindingAccess)
	// The length of an array argument.
	ArrayLength() uint
	SetArrayLength(value uint)
	// The alignment of the constant block.
	ConstantBlockAlignment() uint
	SetConstantBlockAlignment(value uint)
	// The texture type of a texture argument.
	TextureType() MTLTextureType
	SetTextureType(value MTLTextureType)

	MTLAttributeStrideStatic() int
}

// Init initializes the instance.
func (a MTLArgumentDescriptor) Init() MTLArgumentDescriptor {
	rv := objc.Send[MTLArgumentDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLArgumentDescriptor) Autorelease() MTLArgumentDescriptor {
	rv := objc.Send[MTLArgumentDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLArgumentDescriptor creates a new MTLArgumentDescriptor instance.
func NewMTLArgumentDescriptor() MTLArgumentDescriptor {
	class := getMTLArgumentDescriptorClass()
	rv := objc.Send[MTLArgumentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The data type of the argument.
//
// # Discussion
// 
// For a constant data argument, this value needs to match the binary format
// of the data stored in the buffer for that argument. For other parameter
// types, such as textures or samplers, specify the appropriate constant. See
// [MTLDataType] for possible values.
//
// [MTLDataType]: https://developer.apple.com/documentation/Metal/MTLDataType
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/dataType
func (a MTLArgumentDescriptor) DataType() MTLDataType {
	rv := objc.Send[MTLDataType](a.ID, objc.Sel("dataType"))
	return MTLDataType(rv)
}
func (a MTLArgumentDescriptor) SetDataType(value MTLDataType) {
	objc.Send[struct{}](a.ID, objc.Sel("setDataType:"), value)
}
// The index ID of the argument.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/index
func (a MTLArgumentDescriptor) Index() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("index"))
	return rv
}
func (a MTLArgumentDescriptor) SetIndex(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setIndex:"), value)
}
// The access permissions of the argument.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/access
func (a MTLArgumentDescriptor) Access() MTLBindingAccess {
	rv := objc.Send[MTLBindingAccess](a.ID, objc.Sel("access"))
	return MTLBindingAccess(rv)
}
func (a MTLArgumentDescriptor) SetAccess(value MTLBindingAccess) {
	objc.Send[struct{}](a.ID, objc.Sel("setAccess:"), value)
}
// The length of an array argument.
//
// # Discussion
// 
// For a nonarray argument, this value needs to be `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/arrayLength
func (a MTLArgumentDescriptor) ArrayLength() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("arrayLength"))
	return rv
}
func (a MTLArgumentDescriptor) SetArrayLength(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setArrayLength:"), value)
}
// The alignment of the constant block.
//
// # Discussion
// 
// If set, this property forces the constant block to be aligned to the
// specified value. It should be set on the first constant only, and is valid
// only if a corresponding explicit `alignas` specifier is applied to the
// constant in the Metal shader language.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/constantBlockAlignment
func (a MTLArgumentDescriptor) ConstantBlockAlignment() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("constantBlockAlignment"))
	return rv
}
func (a MTLArgumentDescriptor) SetConstantBlockAlignment(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setConstantBlockAlignment:"), value)
}
// The texture type of a texture argument.
//
// # Discussion
// 
// For a nontexture argument, this value is ignored.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/textureType
func (a MTLArgumentDescriptor) TextureType() MTLTextureType {
	rv := objc.Send[MTLTextureType](a.ID, objc.Sel("textureType"))
	return MTLTextureType(rv)
}
func (a MTLArgumentDescriptor) SetTextureType(value MTLTextureType) {
	objc.Send[struct{}](a.ID, objc.Sel("setTextureType:"), value)
}
// See: https://developer.apple.com/documentation/metal/mtlattributestridestatic
func (a MTLArgumentDescriptor) MTLAttributeStrideStatic() int {
	rv := objc.Send[int](a.ID, objc.Sel("MTLAttributeStrideStatic"))
	return rv
}

