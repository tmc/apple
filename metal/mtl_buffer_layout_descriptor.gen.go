// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLBufferLayoutDescriptor] class.
var (
	_MTLBufferLayoutDescriptorClass     MTLBufferLayoutDescriptorClass
	_MTLBufferLayoutDescriptorClassOnce sync.Once
)

func getMTLBufferLayoutDescriptorClass() MTLBufferLayoutDescriptorClass {
	_MTLBufferLayoutDescriptorClassOnce.Do(func() {
		_MTLBufferLayoutDescriptorClass = MTLBufferLayoutDescriptorClass{class: objc.GetClass("MTLBufferLayoutDescriptor")}
	})
	return _MTLBufferLayoutDescriptorClass
}

// GetMTLBufferLayoutDescriptorClass returns the class object for MTLBufferLayoutDescriptor.
func GetMTLBufferLayoutDescriptorClass() MTLBufferLayoutDescriptorClass {
	return getMTLBufferLayoutDescriptorClass()
}

type MTLBufferLayoutDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLBufferLayoutDescriptorClass) Alloc() MTLBufferLayoutDescriptor {
	rv := objc.Send[MTLBufferLayoutDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of how a compute function fetches input data for an
// attribute.
//
// # Describing fetch behavior
//
//   - [MTLBufferLayoutDescriptor.Stride]: The number of bytes from one buffer entry to the next.
//   - [MTLBufferLayoutDescriptor.SetStride]
//   - [MTLBufferLayoutDescriptor.StepFunction]: Determines how and when compute functions fetch data.
//   - [MTLBufferLayoutDescriptor.SetStepFunction]
//   - [MTLBufferLayoutDescriptor.StepRate]: How frequently the step function should load data.
//   - [MTLBufferLayoutDescriptor.SetStepRate]
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor
type MTLBufferLayoutDescriptor struct {
	objectivec.Object
}

// MTLBufferLayoutDescriptorFromID constructs a [MTLBufferLayoutDescriptor] from an objc.ID.
//
// A description of how a compute function fetches input data for an
// attribute.
func MTLBufferLayoutDescriptorFromID(id objc.ID) MTLBufferLayoutDescriptor {
	return MTLBufferLayoutDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLBufferLayoutDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLBufferLayoutDescriptor] class.
//
// # Describing fetch behavior
//
//   - [IMTLBufferLayoutDescriptor.Stride]: The number of bytes from one buffer entry to the next.
//   - [IMTLBufferLayoutDescriptor.SetStride]
//   - [IMTLBufferLayoutDescriptor.StepFunction]: Determines how and when compute functions fetch data.
//   - [IMTLBufferLayoutDescriptor.SetStepFunction]
//   - [IMTLBufferLayoutDescriptor.StepRate]: How frequently the step function should load data.
//   - [IMTLBufferLayoutDescriptor.SetStepRate]
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor
type IMTLBufferLayoutDescriptor interface {
	objectivec.IObject

	// Topic: Describing fetch behavior

	// The number of bytes from one buffer entry to the next.
	Stride() uint
	SetStride(value uint)
	// Determines how and when compute functions fetch data.
	StepFunction() MTLStepFunction
	SetStepFunction(value MTLStepFunction)
	// How frequently the step function should load data.
	StepRate() uint
	SetStepRate(value uint)

	// The organization of input and output data for the next kernel call.
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)
}

// Init initializes the instance.
func (b MTLBufferLayoutDescriptor) Init() MTLBufferLayoutDescriptor {
	rv := objc.Send[MTLBufferLayoutDescriptor](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MTLBufferLayoutDescriptor) Autorelease() MTLBufferLayoutDescriptor {
	rv := objc.Send[MTLBufferLayoutDescriptor](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLBufferLayoutDescriptor creates a new MTLBufferLayoutDescriptor instance.
func NewMTLBufferLayoutDescriptor() MTLBufferLayoutDescriptor {
	class := getMTLBufferLayoutDescriptorClass()
	rv := objc.Send[MTLBufferLayoutDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The number of bytes from one buffer entry to the next.
//
// # Discussion
// 
// The default value is `1`. See [Metal feature set tables (PDF)] for `stride`
// alignment restrictions on GPU architectures.
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stride
func (b MTLBufferLayoutDescriptor) Stride() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("stride"))
	return rv
}
func (b MTLBufferLayoutDescriptor) SetStride(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setStride:"), value)
}

// Determines how and when compute functions fetch data.
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepFunction
func (b MTLBufferLayoutDescriptor) StepFunction() MTLStepFunction {
	rv := objc.Send[MTLStepFunction](b.ID, objc.Sel("stepFunction"))
	return MTLStepFunction(rv)
}
func (b MTLBufferLayoutDescriptor) SetStepFunction(value MTLStepFunction) {
	objc.Send[struct{}](b.ID, objc.Sel("setStepFunction:"), value)
}

// How frequently the step function should load data.
//
// # Discussion
// 
// The interpretation of this value depends on the setting of `stepFunction`.
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepRate
func (b MTLBufferLayoutDescriptor) StepRate() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("stepRate"))
	return rv
}
func (b MTLBufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setStepRate:"), value)
}

// The organization of input and output data for the next kernel call.
//
// See: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b MTLBufferLayoutDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("stageInputDescriptor"))
	return MTLStageInputOutputDescriptorFromID(objc.ID(rv))
}
func (b MTLBufferLayoutDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[struct{}](b.ID, objc.Sel("setStageInputDescriptor:"), value)
}

