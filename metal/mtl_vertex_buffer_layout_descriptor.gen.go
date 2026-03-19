// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLVertexBufferLayoutDescriptor] class.
var (
	_MTLVertexBufferLayoutDescriptorClass     MTLVertexBufferLayoutDescriptorClass
	_MTLVertexBufferLayoutDescriptorClassOnce sync.Once
)

func getMTLVertexBufferLayoutDescriptorClass() MTLVertexBufferLayoutDescriptorClass {
	_MTLVertexBufferLayoutDescriptorClassOnce.Do(func() {
		_MTLVertexBufferLayoutDescriptorClass = MTLVertexBufferLayoutDescriptorClass{class: objc.GetClass("MTLVertexBufferLayoutDescriptor")}
	})
	return _MTLVertexBufferLayoutDescriptorClass
}

// GetMTLVertexBufferLayoutDescriptorClass returns the class object for MTLVertexBufferLayoutDescriptor.
func GetMTLVertexBufferLayoutDescriptorClass() MTLVertexBufferLayoutDescriptorClass {
	return getMTLVertexBufferLayoutDescriptorClass()
}

type MTLVertexBufferLayoutDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLVertexBufferLayoutDescriptorClass) Alloc() MTLVertexBufferLayoutDescriptor {
	rv := objc.Send[MTLVertexBufferLayoutDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that configures how a render pipeline fetches data to send to the
// vertex function.
//
// # Organizing the vertex buffer layout
//
//   - [MTLVertexBufferLayoutDescriptor.StepFunction]: The circumstances under which the vertex and its attributes are presented to the vertex function.
//   - [MTLVertexBufferLayoutDescriptor.SetStepFunction]
//   - [MTLVertexBufferLayoutDescriptor.StepRate]: The interval at which the vertex and its attributes are presented to the vertex function.
//   - [MTLVertexBufferLayoutDescriptor.SetStepRate]
//   - [MTLVertexBufferLayoutDescriptor.Stride]: The number of bytes between the first byte of two consecutive vertices in a buffer.
//   - [MTLVertexBufferLayoutDescriptor.SetStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor
type MTLVertexBufferLayoutDescriptor struct {
	objectivec.Object
}

// MTLVertexBufferLayoutDescriptorFromID constructs a [MTLVertexBufferLayoutDescriptor] from an objc.ID.
//
// An object that configures how a render pipeline fetches data to send to the
// vertex function.
func MTLVertexBufferLayoutDescriptorFromID(id objc.ID) MTLVertexBufferLayoutDescriptor {
	return MTLVertexBufferLayoutDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLVertexBufferLayoutDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLVertexBufferLayoutDescriptor] class.
//
// # Organizing the vertex buffer layout
//
//   - [IMTLVertexBufferLayoutDescriptor.StepFunction]: The circumstances under which the vertex and its attributes are presented to the vertex function.
//   - [IMTLVertexBufferLayoutDescriptor.SetStepFunction]
//   - [IMTLVertexBufferLayoutDescriptor.StepRate]: The interval at which the vertex and its attributes are presented to the vertex function.
//   - [IMTLVertexBufferLayoutDescriptor.SetStepRate]
//   - [IMTLVertexBufferLayoutDescriptor.Stride]: The number of bytes between the first byte of two consecutive vertices in a buffer.
//   - [IMTLVertexBufferLayoutDescriptor.SetStride]
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor
type IMTLVertexBufferLayoutDescriptor interface {
	objectivec.IObject

	// Topic: Organizing the vertex buffer layout

	// The circumstances under which the vertex and its attributes are presented to the vertex function.
	StepFunction() MTLVertexStepFunction
	SetStepFunction(value MTLVertexStepFunction)
	// The interval at which the vertex and its attributes are presented to the vertex function.
	StepRate() uint
	SetStepRate(value uint)
	// The number of bytes between the first byte of two consecutive vertices in a buffer.
	Stride() uint
	SetStride(value uint)

	MTLBufferLayoutStrideDynamic() int
}

// Init initializes the instance.
func (v MTLVertexBufferLayoutDescriptor) Init() MTLVertexBufferLayoutDescriptor {
	rv := objc.Send[MTLVertexBufferLayoutDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTLVertexBufferLayoutDescriptor) Autorelease() MTLVertexBufferLayoutDescriptor {
	rv := objc.Send[MTLVertexBufferLayoutDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVertexBufferLayoutDescriptor creates a new MTLVertexBufferLayoutDescriptor instance.
func NewMTLVertexBufferLayoutDescriptor() MTLVertexBufferLayoutDescriptor {
	class := getMTLVertexBufferLayoutDescriptorClass()
	rv := objc.Send[MTLVertexBufferLayoutDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The circumstances under which the vertex and its attributes are presented
// to the vertex function.
//
// # Discussion
// 
// The default value is [VertexStepFunctionPerVertex].
// 
// If `stepFunction` is [VertexStepFunctionPerVertex], the function fetches
// new attribute data based on the `[[ vertex_id ]]` attribute qualifier. The
// function fetches new attribute data each time a new vertex is processed. In
// this case, `stepRate` needs to be set to `1`, which is its default value.
// 
// If `stepFunction` is [VertexStepFunctionPerInstance], the function fetches
// new attribute data based on the `[[ instance_id ]]` attribute qualifier. In
// this case, `stepRate` needs to be greater than `0` and its value determines
// how often the function fetches new attribute data.
// 
// If `stepFunction` is [VertexStepFunctionConstant], the function fetches
// attribute data just once, and that attribute data is used for every vertex.
// In this case,`stepRate` needs to be set to `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepFunction
func (v MTLVertexBufferLayoutDescriptor) StepFunction() MTLVertexStepFunction {
	rv := objc.Send[MTLVertexStepFunction](v.ID, objc.Sel("stepFunction"))
	return MTLVertexStepFunction(rv)
}
func (v MTLVertexBufferLayoutDescriptor) SetStepFunction(value MTLVertexStepFunction) {
	objc.Send[struct{}](v.ID, objc.Sel("setStepFunction:"), value)
}
// The interval at which the vertex and its attributes are presented to the
// vertex function.
//
// # Discussion
// 
// The default value is `1`. The `stepRate` value, in conjunction with the
// [StepFunction] property, determines how often the function fetches new
// attribute data. The `stepRate` property is generally used when
// `stepFunction` is [VertexStepFunctionPerInstance]. If `stepRate` is equal
// to `1`, new attribute data is fetched for every instance; if `stepRate` is
// equal to `2`, new attribute data is fetched for every two instances, and so
// forth.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepRate
func (v MTLVertexBufferLayoutDescriptor) StepRate() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("stepRate"))
	return rv
}
func (v MTLVertexBufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setStepRate:"), value)
}
// The number of bytes between the first byte of two consecutive vertices in a
// buffer.
//
// # Discussion
// 
// Check the [Metal feature set tables (PDF)] for potential alignment
// restrictions for `stride`.
//
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stride
func (v MTLVertexBufferLayoutDescriptor) Stride() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("stride"))
	return rv
}
func (v MTLVertexBufferLayoutDescriptor) SetStride(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setStride:"), value)
}
// See: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v MTLVertexBufferLayoutDescriptor) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}

