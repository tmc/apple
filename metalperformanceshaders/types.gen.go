// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/metal"
)

// C struct types

// MPSAxisAlignedBoundingBox - An axis-aligned bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAxisAlignedBoundingBox-c.struct
type MPSAxisAlignedBoundingBox struct {
	Min float32
	Max float32
}

// MPSCustomKernelArgumentCount - A structure that contains the number of destination, source, and broadcaset textures used by a custom kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCustomKernelArgumentCount
type MPSCustomKernelArgumentCount struct {
	DestinationTextureCount uint
	SourceTextureCount      uint
	BroadcastTextureCount   uint
}

// MPSCustomKernelInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCustomKernelInfo
type MPSCustomKernelInfo struct {
	ClipOrigin                 uint16
	ClipSize                   uint16
	DestinationFeatureChannels uint16
	DestImageArraySize         uint16
	SourceImageCount           uint16
	ThreadgroupSize            uint16
	SubbatchIndex              uint16
	SubbatchStride             uint16
	Idiv                       MPSIntegerDivisionParams
}

// MPSCustomKernelSourceInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCustomKernelSourceInfo
type MPSCustomKernelSourceInfo struct {
	KernelOrigin         int16
	KernelPhase          uint16
	KernelSize           uint16
	Offset               int16
	Stride               uint16
	DilationRate         uint16
	FeatureChannelOffset uint16
	FeatureChannels      uint16
	ImageArrayOffset     uint16
	ImageArraySize       uint16
}

// MPSDimensionSlice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDimensionSlice
type MPSDimensionSlice struct {
	Start  uint
	Length uint
}

// MPSFunctions_AABB
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunctions_AABB
type MPSFunctions_AABB struct {
	Max float32
	Min float32 // maximum representable per channel values

}

// MPSImageCoordinate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCoordinate
type MPSImageCoordinate struct {
	X       uint
	Y       uint
	Channel uint
}

// MPSImageHistogramInfo - The information used to compute the histogram channels of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramInfo
type MPSImageHistogramInfo struct {
	NumberOfHistogramEntries uint    // Specifies the number of histogram entries (bins) for each channel.
	HistogramForAlpha        bool    // Specifies whether the histogram for the alpha channel should be computed or not.
	MinPixelValue            float32 // Specifies the minimum pixel value. Any pixel value less than this will be clipped to this value (for the purposes of histogram calculation), and assigned to the first histogram entry. This minimum value is applied to each of the four channels separately.
	MaxPixelValue            float32 // Specifies the maximum pixel value.  Any pixel value greater than this will be clipped to this value (for the purposes of histogram calculation), and assigned to the first histogram entry. This maximum value is applied to each of the four channels separately.

}

// MPSImageKeypointData - A structure that specifies keypoint information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageKeypointData
type MPSImageKeypointData struct {
	KeypointCoordinate uint16
	KeypointColorValue float32
}

// MPSImageKeypointRangeInfo - A structure that specifies information to find the keypoints in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageKeypointRangeInfo
type MPSImageKeypointRangeInfo struct {
	MaximumKeypoints      uint
	MinimumThresholdValue float32
}

// MPSImageReadWriteParams - Parameters that control reading and writing of a particular set of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReadWriteParams
type MPSImageReadWriteParams struct {
	FeatureChannelOffset               uint
	NumberOfFeatureChannelsToReadWrite uint
}

// MPSImageRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageRegion
type MPSImageRegion struct {
	Offset MPSImageCoordinate
	Size   MPSImageCoordinate
}

// MPSIntegerDivisionParams - Parameters that define the parts of a division operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntegerDivisionParams
type MPSIntegerDivisionParams struct {
	Divisor uint16
	Recip   uint16
	Addend  uint16
	Shift   uint16
}

// MPSIntersectionDistance - An intersection result that contains the distance from the ray origin to the intersection point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistance
type MPSIntersectionDistance struct {
	Distance float32
}

// MPSIntersectionDistancePrimitiveIndex - An intersection result that contains the distance from the ray origin to the intersection point, and the index of the intersected primitive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndex
type MPSIntersectionDistancePrimitiveIndex struct {
	Distance       float32
	PrimitiveIndex uint32
}

// MPSIntersectionDistancePrimitiveIndexBufferIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndex
type MPSIntersectionDistancePrimitiveIndexBufferIndex struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
}

// MPSIntersectionDistancePrimitiveIndexBufferIndexCoordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexBufferIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
	Coordinates    float32
}

// MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndex
type MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndex struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
	InstanceIndex  uint32
}

// MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
	InstanceIndex  uint32
	Coordinates    float32
}

// MPSIntersectionDistancePrimitiveIndexCoordinates - An intersection result that contains the origin-intersection distance, intersected primitive index, and intersection point coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	Coordinates    float32
}

// MPSIntersectionDistancePrimitiveIndexInstanceIndex - An intersection result that contains the origin-intersection distance, and intersected primitive and instance indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexInstanceIndex
type MPSIntersectionDistancePrimitiveIndexInstanceIndex struct {
	Distance       float32
	PrimitiveIndex uint32
	InstanceIndex  uint32
}

// MPSIntersectionDistancePrimitiveIndexInstanceIndexCoordinates - An intersection result that contains the origin-intersection distance, intersected primitive and instance indices, and intersection point coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexInstanceIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexInstanceIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	InstanceIndex  uint32
	Coordinates    float32
}

// MPSMatrixCopyOffsets - A description of matrix copy operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyOffsets
type MPSMatrixCopyOffsets struct {
	SourceRowOffset         uint32
	SourceColumnOffset      uint32
	DestinationRowOffset    uint32
	DestinationColumnOffset uint32
}

// MPSMatrixOffset - A description of row and column offsets into a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixOffset
type MPSMatrixOffset struct {
	RowOffset    uint32
	ColumnOffset uint32
}

// MPSNDArrayOffsets
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayOffsets
type MPSNDArrayOffsets struct {
	Dimensions [16]int
}

// MPSNDArraySizes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArraySizes
type MPSNDArraySizes struct {
	Dimensions [16]uint
}

// MPSOffset - A signed coordinate with x, y, and z components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSOffset
type MPSOffset struct {
	X int // The horizontal component of the offset, in pixels.
	Y int // The vertical component of the offset, in pixels.
	Z int // The depth component of the offset, in pixels.

}

// MPSOrigin - A position in an image used as the source origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSOrigin
type MPSOrigin struct {
	X float64 // The x coordinate of the position, in pixels.
	Y float64 // The y coordinate of the position, in pixels.
	Z float64 // The z coordinate of the position, in pixels.

}

// MPSPackedFloat3 - A packed three-element vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPackedFloat3-c.struct
type MPSPackedFloat3 struct {
	X float32
	Y float32
	Z float32
}

// MPSRayOriginDirection - A 3D ray with an origin and a direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayOriginDirection
type MPSRayOriginDirection struct {
	Origin    float32
	Direction float32
}

// MPSRayOriginMaskDirectionMaxDistance - A 3D ray with an origin, a direction, and a mask to filter out intersections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayOriginMaskDirectionMaxDistance
type MPSRayOriginMaskDirectionMaxDistance struct {
	Origin      MPSPackedFloat3
	Mask        uint32
	Direction   MPSPackedFloat3
	MaxDistance float32
}

// MPSRayOriginMinDistanceDirectionMaxDistance - A 3D ray with an origin, a direction, and an intersection distance range from the origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayOriginMinDistanceDirectionMaxDistance
type MPSRayOriginMinDistanceDirectionMaxDistance struct {
	Origin      MPSPackedFloat3
	MinDistance float32
	Direction   MPSPackedFloat3
	MaxDistance float32
}

// MPSRayPackedOriginDirection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayPackedOriginDirection
type MPSRayPackedOriginDirection struct {
	Origin    MPSPackedFloat3
	Direction MPSPackedFloat3
}

// MPSRegion - A region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRegion
type MPSRegion struct {
	Origin MPSOrigin // The top-left corner of the region.
	Size   MPSSize   // The size of the region.

}

// MPSScaleTransform - A transform matrix for explicit resampling control with a Lanczos kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSScaleTransform
type MPSScaleTransform struct {
	ScaleX     float64 // The horizontal scale factor.
	ScaleY     float64 // The vertical scale factor.
	TranslateX float64 // The horizontal translation factor.
	TranslateY float64 // The vertical translation factor.

}

// MPSSize - A size of a region in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSize
type MPSSize struct {
	Width  float64 // The width of the region, in pixels.
	Height float64 // The height of the region, in pixels.
	Depth  float64 // The depth of the region, in pixels.

}

// MPSStateTextureInfo - An encapsulation of a texture’s dimensions, format, type, and usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateTextureInfo
type MPSStateTextureInfo struct {
	Width       uint
	Height      uint
	Depth       uint
	ArrayLength uint
	PixelFormat metal.MTLPixelFormat
	TextureType metal.MTLTextureType
	Usage       metal.MTLTextureUsage
	_reserved   [4]uint
}
