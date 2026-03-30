// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLStencilDescriptor] class.
var (
	_MTLStencilDescriptorClass     MTLStencilDescriptorClass
	_MTLStencilDescriptorClassOnce sync.Once
)

func getMTLStencilDescriptorClass() MTLStencilDescriptorClass {
	_MTLStencilDescriptorClassOnce.Do(func() {
		_MTLStencilDescriptorClass = MTLStencilDescriptorClass{class: objc.GetClass("MTLStencilDescriptor")}
	})
	return _MTLStencilDescriptorClass
}

// GetMTLStencilDescriptorClass returns the class object for MTLStencilDescriptor.
func GetMTLStencilDescriptorClass() MTLStencilDescriptorClass {
	return getMTLStencilDescriptorClass()
}

type MTLStencilDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLStencilDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLStencilDescriptorClass) Alloc() MTLStencilDescriptor {
	rv := objc.Send[MTLStencilDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the front-facing or back-facing stencil operations
// of a depth and stencil state object.
//
// # Overview
//
// A stencil test is a comparison between a masked reference value and a
// masked value stored in a stencil attachment. (A value is by performing a
// logical AND operation on it with the [MTLStencilDescriptor.ReadMask] value.) The
// [MTLStencilDescriptor] object defines how to update the contents of the
// stencil attachment, based on the results of the stencil test and the depth
// test.
//
// The [MTLStencilDescriptor.StencilCompareFunction] property defines the stencil test. The
// [MTLStencilDescriptor.StencilFailureOperation], [MTLStencilDescriptor.DepthFailureOperation], and
// [MTLStencilDescriptor.DepthStencilPassOperation] properties specify what to do to a stencil
// value stored in the stencil attachment for three different test outcomes:
// if the stencil test fails, if the stencil test passes and the depth test
// fails, or if both stencil and depth tests succeed, respectively.
// [MTLStencilDescriptor.WriteMask] determines which stencil bits can be modified as the result of
// a stencil operation.
//
// # Configuring stencil functions and operations
//
//   - [MTLStencilDescriptor.StencilFailureOperation]: The operation that is performed to update the values in the stencil attachment when the stencil test fails.
//   - [MTLStencilDescriptor.SetStencilFailureOperation]
//   - [MTLStencilDescriptor.DepthFailureOperation]: The operation that is performed to update the values in the stencil attachment when the stencil test passes, but the depth test fails.
//   - [MTLStencilDescriptor.SetDepthFailureOperation]
//   - [MTLStencilDescriptor.DepthStencilPassOperation]: The operation that is performed to update the values in the stencil attachment when both the stencil test and the depth test pass.
//   - [MTLStencilDescriptor.SetDepthStencilPassOperation]
//   - [MTLStencilDescriptor.StencilCompareFunction]: The comparison that is performed between the masked reference value and a masked value in the stencil attachment.
//   - [MTLStencilDescriptor.SetStencilCompareFunction]
//
// # Configuring stencil bit mask properties
//
//   - [MTLStencilDescriptor.ReadMask]: A bitmask that determines from which bits that stencil comparison tests can read.
//   - [MTLStencilDescriptor.SetReadMask]
//   - [MTLStencilDescriptor.WriteMask]: A bitmask that determines to which bits that stencil operations can write.
//   - [MTLStencilDescriptor.SetWriteMask]
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor
type MTLStencilDescriptor struct {
	objectivec.Object
}

// MTLStencilDescriptorFromID constructs a [MTLStencilDescriptor] from an objc.ID.
//
// An object that defines the front-facing or back-facing stencil operations
// of a depth and stencil state object.
func MTLStencilDescriptorFromID(id objc.ID) MTLStencilDescriptor {
	return MTLStencilDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLStencilDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLStencilDescriptor] class.
//
// # Configuring stencil functions and operations
//
//   - [IMTLStencilDescriptor.StencilFailureOperation]: The operation that is performed to update the values in the stencil attachment when the stencil test fails.
//   - [IMTLStencilDescriptor.SetStencilFailureOperation]
//   - [IMTLStencilDescriptor.DepthFailureOperation]: The operation that is performed to update the values in the stencil attachment when the stencil test passes, but the depth test fails.
//   - [IMTLStencilDescriptor.SetDepthFailureOperation]
//   - [IMTLStencilDescriptor.DepthStencilPassOperation]: The operation that is performed to update the values in the stencil attachment when both the stencil test and the depth test pass.
//   - [IMTLStencilDescriptor.SetDepthStencilPassOperation]
//   - [IMTLStencilDescriptor.StencilCompareFunction]: The comparison that is performed between the masked reference value and a masked value in the stencil attachment.
//   - [IMTLStencilDescriptor.SetStencilCompareFunction]
//
// # Configuring stencil bit mask properties
//
//   - [IMTLStencilDescriptor.ReadMask]: A bitmask that determines from which bits that stencil comparison tests can read.
//   - [IMTLStencilDescriptor.SetReadMask]
//   - [IMTLStencilDescriptor.WriteMask]: A bitmask that determines to which bits that stencil operations can write.
//   - [IMTLStencilDescriptor.SetWriteMask]
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor
type IMTLStencilDescriptor interface {
	objectivec.IObject

	// Topic: Configuring stencil functions and operations

	// The operation that is performed to update the values in the stencil attachment when the stencil test fails.
	StencilFailureOperation() MTLStencilOperation
	SetStencilFailureOperation(value MTLStencilOperation)
	// The operation that is performed to update the values in the stencil attachment when the stencil test passes, but the depth test fails.
	DepthFailureOperation() MTLStencilOperation
	SetDepthFailureOperation(value MTLStencilOperation)
	// The operation that is performed to update the values in the stencil attachment when both the stencil test and the depth test pass.
	DepthStencilPassOperation() MTLStencilOperation
	SetDepthStencilPassOperation(value MTLStencilOperation)
	// The comparison that is performed between the masked reference value and a masked value in the stencil attachment.
	StencilCompareFunction() MTLCompareFunction
	SetStencilCompareFunction(value MTLCompareFunction)

	// Topic: Configuring stencil bit mask properties

	// A bitmask that determines from which bits that stencil comparison tests can read.
	ReadMask() uint32
	SetReadMask(value uint32)
	// A bitmask that determines to which bits that stencil operations can write.
	WriteMask() uint32
	SetWriteMask(value uint32)
}

// Init initializes the instance.
func (s MTLStencilDescriptor) Init() MTLStencilDescriptor {
	rv := objc.Send[MTLStencilDescriptor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLStencilDescriptor) Autorelease() MTLStencilDescriptor {
	rv := objc.Send[MTLStencilDescriptor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLStencilDescriptor creates a new MTLStencilDescriptor instance.
func NewMTLStencilDescriptor() MTLStencilDescriptor {
	class := getMTLStencilDescriptorClass()
	rv := objc.Send[MTLStencilDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The operation that is performed to update the values in the stencil
// attachment when the stencil test fails.
//
// # Discussion
//
// The default value is [MTLStencilOperationKeep], which does not change the
// current stencil value. For more information on possible values, see
// [MTLStencilOperation].
//
// When the stencil test fails for a pixel, its incoming color, depth, or
// stencil values are discarded.
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/stencilFailureOperation
//
// [MTLStencilOperation]: https://developer.apple.com/documentation/Metal/MTLStencilOperation
func (s MTLStencilDescriptor) StencilFailureOperation() MTLStencilOperation {
	rv := objc.Send[MTLStencilOperation](s.ID, objc.Sel("stencilFailureOperation"))
	return MTLStencilOperation(rv)
}
func (s MTLStencilDescriptor) SetStencilFailureOperation(value MTLStencilOperation) {
	objc.Send[struct{}](s.ID, objc.Sel("setStencilFailureOperation:"), value)
}

// The operation that is performed to update the values in the stencil
// attachment when the stencil test passes, but the depth test fails.
//
// # Discussion
//
// The default value is [MTLStencilOperationKeep], which does not change the
// current stencil value. For more information on possible values, see
// [MTLStencilOperation].
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/depthFailureOperation
//
// [MTLStencilOperation]: https://developer.apple.com/documentation/Metal/MTLStencilOperation
func (s MTLStencilDescriptor) DepthFailureOperation() MTLStencilOperation {
	rv := objc.Send[MTLStencilOperation](s.ID, objc.Sel("depthFailureOperation"))
	return MTLStencilOperation(rv)
}
func (s MTLStencilDescriptor) SetDepthFailureOperation(value MTLStencilOperation) {
	objc.Send[struct{}](s.ID, objc.Sel("setDepthFailureOperation:"), value)
}

// The operation that is performed to update the values in the stencil
// attachment when both the stencil test and the depth test pass.
//
// # Discussion
//
// The default value is [MTLStencilOperationKeep], which does not change the
// current stencil value. For more information on possible values, see
// [MTLStencilOperation].
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/depthStencilPassOperation
//
// [MTLStencilOperation]: https://developer.apple.com/documentation/Metal/MTLStencilOperation
func (s MTLStencilDescriptor) DepthStencilPassOperation() MTLStencilOperation {
	rv := objc.Send[MTLStencilOperation](s.ID, objc.Sel("depthStencilPassOperation"))
	return MTLStencilOperation(rv)
}
func (s MTLStencilDescriptor) SetDepthStencilPassOperation(value MTLStencilOperation) {
	objc.Send[struct{}](s.ID, objc.Sel("setDepthStencilPassOperation:"), value)
}

// The comparison that is performed between the masked reference value and a
// masked value in the stencil attachment.
//
// # Discussion
//
// For example, if `stencilCompareFunction` is [MTLCompareFunctionLess], then
// the stencil test passes if the masked reference value is less than the
// masked stored stencil value. The default value is
// [MTLCompareFunction.always], which indicates that the stencil test always
// passes.
//
// The stored stencil value and the reference value are both by performing a
// logical AND operation with the [ReadMask] value before the comparison takes
// place. For more information on possible values, see [MTLCompareFunction].
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/stencilCompareFunction
//
// [MTLCompareFunction.always]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/always
// [MTLCompareFunction]: https://developer.apple.com/documentation/Metal/MTLCompareFunction
func (s MTLStencilDescriptor) StencilCompareFunction() MTLCompareFunction {
	rv := objc.Send[MTLCompareFunction](s.ID, objc.Sel("stencilCompareFunction"))
	return MTLCompareFunction(rv)
}
func (s MTLStencilDescriptor) SetStencilCompareFunction(value MTLCompareFunction) {
	objc.Send[struct{}](s.ID, objc.Sel("setStencilCompareFunction:"), value)
}

// A bitmask that determines from which bits that stencil comparison tests can
// read.
//
// # Discussion
//
// The [ReadMask] bits are used for logical AND operations to both the stored
// stencil value and the reference value.
//
// The least significant bits of the read mask are used. The default value is
// all ones. A logical AND operation with the default [ReadMask] does not
// change the value.
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/readMask
func (s MTLStencilDescriptor) ReadMask() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("readMask"))
	return rv
}
func (s MTLStencilDescriptor) SetReadMask(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setReadMask:"), value)
}

// A bitmask that determines to which bits that stencil operations can write.
//
// # Discussion
//
// [WriteMask] are used for logical AND operations to values that are going to
// be written into a stencil attachment as the result of a stencil operation.
//
// The least significant bits of the write mask are used. The default value is
// all ones. A logical AND operation with the default [WriteMask] does not
// change the value.
//
// See: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/writeMask
func (s MTLStencilDescriptor) WriteMask() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("writeMask"))
	return rv
}
func (s MTLStencilDescriptor) SetWriteMask(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setWriteMask:"), value)
}
