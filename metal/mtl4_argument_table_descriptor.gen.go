// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4ArgumentTableDescriptor] class.
var (
	_MTL4ArgumentTableDescriptorClass     MTL4ArgumentTableDescriptorClass
	_MTL4ArgumentTableDescriptorClassOnce sync.Once
)

func getMTL4ArgumentTableDescriptorClass() MTL4ArgumentTableDescriptorClass {
	_MTL4ArgumentTableDescriptorClassOnce.Do(func() {
		_MTL4ArgumentTableDescriptorClass = MTL4ArgumentTableDescriptorClass{class: objc.GetClass("MTL4ArgumentTableDescriptor")}
	})
	return _MTL4ArgumentTableDescriptorClass
}

// GetMTL4ArgumentTableDescriptorClass returns the class object for MTL4ArgumentTableDescriptor.
func GetMTL4ArgumentTableDescriptorClass() MTL4ArgumentTableDescriptorClass {
	return getMTL4ArgumentTableDescriptorClass()
}

type MTL4ArgumentTableDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4ArgumentTableDescriptorClass) Alloc() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups parameters for the creation of a Metal argument table.
//
// # Overview
// 
// Argument tables provide resource bindings to your Metal pipeline states.
//
// # Instance Properties
//
//   - [MTL4ArgumentTableDescriptor.InitializeBindings]: Configures whether Metal initializes the bindings to nil values upon creation of argument table.
//   - [MTL4ArgumentTableDescriptor.SetInitializeBindings]
//   - [MTL4ArgumentTableDescriptor.Label]: Assigns an optional label with the argument table for debug purposes.
//   - [MTL4ArgumentTableDescriptor.SetLabel]
//   - [MTL4ArgumentTableDescriptor.MaxBufferBindCount]: Determines the number of buffer-binding slots for the argument table.
//   - [MTL4ArgumentTableDescriptor.SetMaxBufferBindCount]
//   - [MTL4ArgumentTableDescriptor.MaxSamplerStateBindCount]: Determines the number of sampler state-binding slots for the argument table.
//   - [MTL4ArgumentTableDescriptor.SetMaxSamplerStateBindCount]
//   - [MTL4ArgumentTableDescriptor.MaxTextureBindCount]: Determines the number of texture-binding slots for the argument table.
//   - [MTL4ArgumentTableDescriptor.SetMaxTextureBindCount]
//   - [MTL4ArgumentTableDescriptor.SupportAttributeStrides]: Controls whether Metal should reserve memory for attribute strides in the argument table.
//   - [MTL4ArgumentTableDescriptor.SetSupportAttributeStrides]
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor
type MTL4ArgumentTableDescriptor struct {
	objectivec.Object
}

// MTL4ArgumentTableDescriptorFromID constructs a [MTL4ArgumentTableDescriptor] from an objc.ID.
//
// Groups parameters for the creation of a Metal argument table.
func MTL4ArgumentTableDescriptorFromID(id objc.ID) MTL4ArgumentTableDescriptor {
	return MTL4ArgumentTableDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4ArgumentTableDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4ArgumentTableDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4ArgumentTableDescriptor.InitializeBindings]: Configures whether Metal initializes the bindings to nil values upon creation of argument table.
//   - [IMTL4ArgumentTableDescriptor.SetInitializeBindings]
//   - [IMTL4ArgumentTableDescriptor.Label]: Assigns an optional label with the argument table for debug purposes.
//   - [IMTL4ArgumentTableDescriptor.SetLabel]
//   - [IMTL4ArgumentTableDescriptor.MaxBufferBindCount]: Determines the number of buffer-binding slots for the argument table.
//   - [IMTL4ArgumentTableDescriptor.SetMaxBufferBindCount]
//   - [IMTL4ArgumentTableDescriptor.MaxSamplerStateBindCount]: Determines the number of sampler state-binding slots for the argument table.
//   - [IMTL4ArgumentTableDescriptor.SetMaxSamplerStateBindCount]
//   - [IMTL4ArgumentTableDescriptor.MaxTextureBindCount]: Determines the number of texture-binding slots for the argument table.
//   - [IMTL4ArgumentTableDescriptor.SetMaxTextureBindCount]
//   - [IMTL4ArgumentTableDescriptor.SupportAttributeStrides]: Controls whether Metal should reserve memory for attribute strides in the argument table.
//   - [IMTL4ArgumentTableDescriptor.SetSupportAttributeStrides]
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor
type IMTL4ArgumentTableDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Configures whether Metal initializes the bindings to nil values upon creation of argument table.
	InitializeBindings() bool
	SetInitializeBindings(value bool)
	// Assigns an optional label with the argument table for debug purposes.
	Label() string
	SetLabel(value string)
	// Determines the number of buffer-binding slots for the argument table.
	MaxBufferBindCount() uint
	SetMaxBufferBindCount(value uint)
	// Determines the number of sampler state-binding slots for the argument table.
	MaxSamplerStateBindCount() uint
	SetMaxSamplerStateBindCount(value uint)
	// Determines the number of texture-binding slots for the argument table.
	MaxTextureBindCount() uint
	SetMaxTextureBindCount(value uint)
	// Controls whether Metal should reserve memory for attribute strides in the argument table.
	SupportAttributeStrides() bool
	SetSupportAttributeStrides(value bool)

	MTL4CommandQueueErrorDomain() string
}

// Init initializes the instance.
func (m MTL4ArgumentTableDescriptor) Init() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4ArgumentTableDescriptor) Autorelease() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4ArgumentTableDescriptor creates a new MTL4ArgumentTableDescriptor instance.
func NewMTL4ArgumentTableDescriptor() MTL4ArgumentTableDescriptor {
	class := getMTL4ArgumentTableDescriptorClass()
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Configures whether Metal initializes the bindings to nil values upon
// creation of argument table.
//
// # Discussion
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/initializeBindings
func (m MTL4ArgumentTableDescriptor) InitializeBindings() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("initializeBindings"))
	return rv
}
func (m MTL4ArgumentTableDescriptor) SetInitializeBindings(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setInitializeBindings:"), value)
}

// Assigns an optional label with the argument table for debug purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/label
func (m MTL4ArgumentTableDescriptor) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4ArgumentTableDescriptor) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Determines the number of buffer-binding slots for the argument table.
//
// # Discussion
// 
// The maximum value of this parameter is 31.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxBufferBindCount
func (m MTL4ArgumentTableDescriptor) MaxBufferBindCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxBufferBindCount"))
	return rv
}
func (m MTL4ArgumentTableDescriptor) SetMaxBufferBindCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxBufferBindCount:"), value)
}

// Determines the number of sampler state-binding slots for the argument
// table.
//
// # Discussion
// 
// The maximum value of this parameter is 16.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxSamplerStateBindCount
func (m MTL4ArgumentTableDescriptor) MaxSamplerStateBindCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxSamplerStateBindCount"))
	return rv
}
func (m MTL4ArgumentTableDescriptor) SetMaxSamplerStateBindCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxSamplerStateBindCount:"), value)
}

// Determines the number of texture-binding slots for the argument table.
//
// # Discussion
// 
// The maximum value of this parameter is 128.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxTextureBindCount
func (m MTL4ArgumentTableDescriptor) MaxTextureBindCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTextureBindCount"))
	return rv
}
func (m MTL4ArgumentTableDescriptor) SetMaxTextureBindCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTextureBindCount:"), value)
}

// Controls whether Metal should reserve memory for attribute strides in the
// argument table.
//
// # Discussion
// 
// Set this value to true if you intend to provide dynamic attribute strides
// when binding vertex array buffers to the argument table by calling
// [SetAddressAttributeStrideAtIndex]
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/supportAttributeStrides
func (m MTL4ArgumentTableDescriptor) SupportAttributeStrides() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportAttributeStrides"))
	return rv
}
func (m MTL4ArgumentTableDescriptor) SetSupportAttributeStrides(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportAttributeStrides:"), value)
}

// See: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m MTL4ArgumentTableDescriptor) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

