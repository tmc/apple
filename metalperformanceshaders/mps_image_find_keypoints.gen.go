// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageFindKeypoints] class.
var (
	_MPSImageFindKeypointsClass     MPSImageFindKeypointsClass
	_MPSImageFindKeypointsClassOnce sync.Once
)

func getMPSImageFindKeypointsClass() MPSImageFindKeypointsClass {
	_MPSImageFindKeypointsClassOnce.Do(func() {
		_MPSImageFindKeypointsClass = MPSImageFindKeypointsClass{class: objc.GetClass("MPSImageFindKeypoints")}
	})
	return _MPSImageFindKeypointsClass
}

// GetMPSImageFindKeypointsClass returns the class object for MPSImageFindKeypoints.
func GetMPSImageFindKeypointsClass() MPSImageFindKeypointsClass {
	return getMPSImageFindKeypointsClass()
}

type MPSImageFindKeypointsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageFindKeypointsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageFindKeypointsClass) Alloc() MPSImageFindKeypoints {
	rv := objc.Send[MPSImageFindKeypoints](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that is used to find a list of keypoints.
//
// # Overview
//
// This kernel is used to find a list of keypoints whose values are greater
// than the [minimumThresholdValue] in [MPSImageKeypointRangeInfo]. The
// keypoints are generated for a specified region in the image. The pixel
// format of the source image must be [MTLPixelFormat.r8Unorm].
//
// # Initializers
//
//   - [MPSImageFindKeypoints.InitWithDeviceInfo]
//
// # Instance Properties
//
//   - [MPSImageFindKeypoints.KeypointRangeInfo]
//
// # Instance Methods
//
//   - [MPSImageFindKeypoints.EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints
//
// [MPSImageKeypointRangeInfo]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageKeypointRangeInfo
// [MTLPixelFormat.r8Unorm]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Unorm
// [minimumThresholdValue]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageKeypointRangeInfo/minimumThresholdValue
type MPSImageFindKeypoints struct {
	MPSKernel
}

// MPSImageFindKeypointsFromID constructs a [MPSImageFindKeypoints] from an objc.ID.
//
// A kernel that is used to find a list of keypoints.
func MPSImageFindKeypointsFromID(id objc.ID) MPSImageFindKeypoints {
	return MPSImageFindKeypoints{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSImageFindKeypoints adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageFindKeypoints] class.
//
// # Initializers
//
//   - [IMPSImageFindKeypoints.InitWithDeviceInfo]
//
// # Instance Properties
//
//   - [IMPSImageFindKeypoints.KeypointRangeInfo]
//
// # Instance Methods
//
//   - [IMPSImageFindKeypoints.EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints
type IMPSImageFindKeypoints interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceInfo(device metal.MTLDevice, info *MPSImageKeypointRangeInfo) MPSImageFindKeypoints

	// Topic: Instance Properties

	KeypointRangeInfo() MPSImageKeypointRangeInfo

	// Topic: Instance Methods

	EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, regions *metal.MTLRegion, numberOfRegions uint, keypointCountBuffer metal.MTLBuffer, keypointCountBufferOffset uint, keypointDataBuffer metal.MTLBuffer, keypointDataBufferOffset uint)
}

// Init initializes the instance.
func (i MPSImageFindKeypoints) Init() MPSImageFindKeypoints {
	rv := objc.Send[MPSImageFindKeypoints](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageFindKeypoints) Autorelease() MPSImageFindKeypoints {
	rv := objc.Send[MPSImageFindKeypoints](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageFindKeypoints creates a new MPSImageFindKeypoints instance.
func NewMPSImageFindKeypoints() MPSImageFindKeypoints {
	class := getMPSImageFindKeypointsClass()
	rv := objc.Send[MPSImageFindKeypoints](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageFindKeypointsWithCoder(aDecoder foundation.INSCoder) MPSImageFindKeypoints {
	instance := getMPSImageFindKeypointsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageFindKeypointsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints/init(coder:device:)
func NewImageFindKeypointsWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageFindKeypoints {
	instance := getMPSImageFindKeypointsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageFindKeypointsFromID(rv)
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
func NewImageFindKeypointsWithDevice(device metal.MTLDevice) MPSImageFindKeypoints {
	instance := getMPSImageFindKeypointsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageFindKeypointsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints/init(device:info:)
func NewImageFindKeypointsWithDeviceInfo(device metal.MTLDevice, info *MPSImageKeypointRangeInfo) MPSImageFindKeypoints {
	instance := getMPSImageFindKeypointsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:info:"), device, unsafe.Pointer(info))
	return MPSImageFindKeypointsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints/init(device:info:)
func (i MPSImageFindKeypoints) InitWithDeviceInfo(device metal.MTLDevice, info *MPSImageKeypointRangeInfo) MPSImageFindKeypoints {
	rv := objc.Send[MPSImageFindKeypoints](i.ID, objc.Sel("initWithDevice:info:"), device, unsafe.Pointer(info))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints/encode(to:sourceTexture:regions:numberOfRegions:keypointCount:keypointCountBufferOffset:keypointDataBuffer:keypointDataBufferOffset:)
func (i MPSImageFindKeypoints) EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, regions *metal.MTLRegion, numberOfRegions uint, keypointCountBuffer metal.MTLBuffer, keypointCountBufferOffset uint, keypointDataBuffer metal.MTLBuffer, keypointDataBufferOffset uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:regions:numberOfRegions:keypointCountBuffer:keypointCountBufferOffset:keypointDataBuffer:keypointDataBufferOffset:"), commandBuffer, source, unsafe.Pointer(regions), numberOfRegions, keypointCountBuffer, keypointCountBufferOffset, keypointDataBuffer, keypointDataBufferOffset)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints/keypointRangeInfo
func (i MPSImageFindKeypoints) KeypointRangeInfo() MPSImageKeypointRangeInfo {
	rv := objc.Send[MPSImageKeypointRangeInfo](i.ID, objc.Sel("keypointRangeInfo"))
	return MPSImageKeypointRangeInfo(rv)
}
