// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixCopy] class.
var (
	_MPSMatrixCopyClass     MPSMatrixCopyClass
	_MPSMatrixCopyClassOnce sync.Once
)

func getMPSMatrixCopyClass() MPSMatrixCopyClass {
	_MPSMatrixCopyClassOnce.Do(func() {
		_MPSMatrixCopyClass = MPSMatrixCopyClass{class: objc.GetClass("MPSMatrixCopy")}
	})
	return _MPSMatrixCopyClass
}

// GetMPSMatrixCopyClass returns the class object for MPSMatrixCopy.
func GetMPSMatrixCopyClass() MPSMatrixCopyClass {
	return getMPSMatrixCopyClass()
}

type MPSMatrixCopyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixCopyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixCopyClass) Alloc() MPSMatrixCopy {
	rv := objc.Send[MPSMatrixCopy](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that can perform multiple matrix copy operations.
//
// # Initializers
//
//   - [MPSMatrixCopy.InitWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed]
//
// # Instance Properties
//
//   - [MPSMatrixCopy.CopyColumns]
//   - [MPSMatrixCopy.CopyRows]
//   - [MPSMatrixCopy.DestinationsAreTransposed]
//   - [MPSMatrixCopy.SourcesAreTransposed]
//
// # Instance Methods
//
//   - [MPSMatrixCopy.EncodeToCommandBufferCopyDescriptor]
//   - [MPSMatrixCopy.EncodeToCommandBufferCopyDescriptorRowPermuteIndicesRowPermuteOffsetColumnPermuteIndicesColumnPermuteOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy
type MPSMatrixCopy struct {
	MPSKernel
}

// MPSMatrixCopyFromID constructs a [MPSMatrixCopy] from an objc.ID.
//
// A class that can perform multiple matrix copy operations.
func MPSMatrixCopyFromID(id objc.ID) MPSMatrixCopy {
	return MPSMatrixCopy{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSMatrixCopy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixCopy] class.
//
// # Initializers
//
//   - [IMPSMatrixCopy.InitWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed]
//
// # Instance Properties
//
//   - [IMPSMatrixCopy.CopyColumns]
//   - [IMPSMatrixCopy.CopyRows]
//   - [IMPSMatrixCopy.DestinationsAreTransposed]
//   - [IMPSMatrixCopy.SourcesAreTransposed]
//
// # Instance Methods
//
//   - [IMPSMatrixCopy.EncodeToCommandBufferCopyDescriptor]
//   - [IMPSMatrixCopy.EncodeToCommandBufferCopyDescriptorRowPermuteIndicesRowPermuteOffsetColumnPermuteIndicesColumnPermuteOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy
type IMPSMatrixCopy interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed(device metal.MTLDevice, copyRows uint, copyColumns uint, sourcesAreTransposed bool, destinationsAreTransposed bool) MPSMatrixCopy

	// Topic: Instance Properties

	CopyColumns() uint
	CopyRows() uint
	DestinationsAreTransposed() bool
	SourcesAreTransposed() bool

	// Topic: Instance Methods

	EncodeToCommandBufferCopyDescriptor(commandBuffer metal.MTLCommandBuffer, copyDescriptor IMPSMatrixCopyDescriptor)
	EncodeToCommandBufferCopyDescriptorRowPermuteIndicesRowPermuteOffsetColumnPermuteIndicesColumnPermuteOffset(commandBuffer metal.MTLCommandBuffer, copyDescriptor IMPSMatrixCopyDescriptor, rowPermuteIndices IMPSVector, rowPermuteOffset uint, columnPermuteIndices IMPSVector, columnPermuteOffset uint)
}

// Init initializes the instance.
func (m MPSMatrixCopy) Init() MPSMatrixCopy {
	rv := objc.Send[MPSMatrixCopy](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixCopy) Autorelease() MPSMatrixCopy {
	rv := objc.Send[MPSMatrixCopy](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixCopy creates a new MPSMatrixCopy instance.
func NewMPSMatrixCopy() MPSMatrixCopy {
	class := getMPSMatrixCopyClass()
	rv := objc.Send[MPSMatrixCopy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixCopyWithCoder(aDecoder foundation.INSCoder) MPSMatrixCopy {
	instance := getMPSMatrixCopyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixCopyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/init(coder:device:)
func NewMatrixCopyWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixCopy {
	instance := getMPSMatrixCopyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixCopyFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewMatrixCopyWithDevice(device metal.MTLDevice) MPSMatrixCopy {
	instance := getMPSMatrixCopyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixCopyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/init(device:copyRows:copyColumns:sourcesAreTransposed:destinationsAreTransposed:)
func NewMatrixCopyWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed(device metal.MTLDevice, copyRows uint, copyColumns uint, sourcesAreTransposed bool, destinationsAreTransposed bool) MPSMatrixCopy {
	instance := getMPSMatrixCopyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:copyRows:copyColumns:sourcesAreTransposed:destinationsAreTransposed:"), device, copyRows, copyColumns, sourcesAreTransposed, destinationsAreTransposed)
	return MPSMatrixCopyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/init(device:copyRows:copyColumns:sourcesAreTransposed:destinationsAreTransposed:)
func (m MPSMatrixCopy) InitWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed(device metal.MTLDevice, copyRows uint, copyColumns uint, sourcesAreTransposed bool, destinationsAreTransposed bool) MPSMatrixCopy {
	rv := objc.Send[MPSMatrixCopy](m.ID, objc.Sel("initWithDevice:copyRows:copyColumns:sourcesAreTransposed:destinationsAreTransposed:"), device, copyRows, copyColumns, sourcesAreTransposed, destinationsAreTransposed)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/encode(commandBuffer:copyDescriptor:)
func (m MPSMatrixCopy) EncodeToCommandBufferCopyDescriptor(commandBuffer metal.MTLCommandBuffer, copyDescriptor IMPSMatrixCopyDescriptor) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:copyDescriptor:"), commandBuffer, copyDescriptor)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/encode(commandBuffer:copyDescriptor:rowPermuteIndices:rowPermuteOffset:columnPermuteIndices:columnPermuteOffset:)
func (m MPSMatrixCopy) EncodeToCommandBufferCopyDescriptorRowPermuteIndicesRowPermuteOffsetColumnPermuteIndicesColumnPermuteOffset(commandBuffer metal.MTLCommandBuffer, copyDescriptor IMPSMatrixCopyDescriptor, rowPermuteIndices IMPSVector, rowPermuteOffset uint, columnPermuteIndices IMPSVector, columnPermuteOffset uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:copyDescriptor:rowPermuteIndices:rowPermuteOffset:columnPermuteIndices:columnPermuteOffset:"), commandBuffer, copyDescriptor, rowPermuteIndices, rowPermuteOffset, columnPermuteIndices, columnPermuteOffset)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/copyColumns
func (m MPSMatrixCopy) CopyColumns() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("copyColumns"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/copyRows
func (m MPSMatrixCopy) CopyRows() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("copyRows"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/destinationsAreTransposed
func (m MPSMatrixCopy) DestinationsAreTransposed() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("destinationsAreTransposed"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy/sourcesAreTransposed
func (m MPSMatrixCopy) SourcesAreTransposed() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("sourcesAreTransposed"))
	return rv
}
