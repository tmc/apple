// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLPipelineBufferDescriptor] class.
var (
	_MTLPipelineBufferDescriptorClass     MTLPipelineBufferDescriptorClass
	_MTLPipelineBufferDescriptorClassOnce sync.Once
)

func getMTLPipelineBufferDescriptorClass() MTLPipelineBufferDescriptorClass {
	_MTLPipelineBufferDescriptorClassOnce.Do(func() {
		_MTLPipelineBufferDescriptorClass = MTLPipelineBufferDescriptorClass{class: objc.GetClass("MTLPipelineBufferDescriptor")}
	})
	return _MTLPipelineBufferDescriptorClass
}

// GetMTLPipelineBufferDescriptorClass returns the class object for MTLPipelineBufferDescriptor.
func GetMTLPipelineBufferDescriptorClass() MTLPipelineBufferDescriptorClass {
	return getMTLPipelineBufferDescriptorClass()
}

type MTLPipelineBufferDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLPipelineBufferDescriptorClass) Alloc() MTLPipelineBufferDescriptor {
	rv := objc.Send[MTLPipelineBufferDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The mutability options for a buffer that a render or compute pipeline uses.
//
// # Overview
// 
// Metal can perform additional optimizations if you guarantee that neither
// the CPU nor the GPU modify a buffer’s contents before starting a pass.
// Use immutable buffers as much as possible to take advantage of Metal
// optimizations.
// 
// To declare that a buffer is immutable, set the [MTLPipelineBufferDescriptor.Mutability] property of
// their associated [MTLPipelineBufferDescriptor] object to
// [MutabilityImmutable].
//
// # Setting buffer mutability
//
//   - [MTLPipelineBufferDescriptor.Mutability]: A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer.
//   - [MTLPipelineBufferDescriptor.SetMutability]
//
// See: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor
type MTLPipelineBufferDescriptor struct {
	objectivec.Object
}

// MTLPipelineBufferDescriptorFromID constructs a [MTLPipelineBufferDescriptor] from an objc.ID.
//
// The mutability options for a buffer that a render or compute pipeline uses.
func MTLPipelineBufferDescriptorFromID(id objc.ID) MTLPipelineBufferDescriptor {
	return MTLPipelineBufferDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLPipelineBufferDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLPipelineBufferDescriptor] class.
//
// # Setting buffer mutability
//
//   - [IMTLPipelineBufferDescriptor.Mutability]: A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer.
//   - [IMTLPipelineBufferDescriptor.SetMutability]
//
// See: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor
type IMTLPipelineBufferDescriptor interface {
	objectivec.IObject

	// Topic: Setting buffer mutability

	// A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer.
	Mutability() MTLMutability
	SetMutability(value MTLMutability)
}

// Init initializes the instance.
func (p MTLPipelineBufferDescriptor) Init() MTLPipelineBufferDescriptor {
	rv := objc.Send[MTLPipelineBufferDescriptor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MTLPipelineBufferDescriptor) Autorelease() MTLPipelineBufferDescriptor {
	rv := objc.Send[MTLPipelineBufferDescriptor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLPipelineBufferDescriptor creates a new MTLPipelineBufferDescriptor instance.
func NewMTLPipelineBufferDescriptor() MTLPipelineBufferDescriptor {
	class := getMTLPipelineBufferDescriptorClass()
	rv := objc.Send[MTLPipelineBufferDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A mutability option that determines whether you can update a buffer’s
// contents before related commands use the buffer.
//
// # Discussion
// 
// The default value is [MutabilityDefault].
// 
// If you don’t explicitly declare mutability, Metal uses the following
// default behaviors:
// 
// - Regular buffers are mutable by default, and Metal treats
// [MutabilityDefault] as if it were [MutabilityMutable]. - Argument buffers
// are immutable by default, and Metal treats [MutabilityDefault] as if it
// were [MutabilityImmutable].
//
// See: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor/mutability
func (p MTLPipelineBufferDescriptor) Mutability() MTLMutability {
	rv := objc.Send[MTLMutability](p.ID, objc.Sel("mutability"))
	return MTLMutability(rv)
}
func (p MTLPipelineBufferDescriptor) SetMutability(value MTLMutability) {
	objc.Send[struct{}](p.ID, objc.Sel("setMutability:"), value)
}

