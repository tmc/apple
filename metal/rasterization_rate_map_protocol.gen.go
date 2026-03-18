// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A compiled read-only instance that determines how to apply variable rasterization rates when rendering.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap
type MTLRasterizationRateMap interface {
	objectivec.IObject

	// The device object that created the rate map.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/device
	Device() MTLDevice

	// A string that identifies the rate map.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/label
	Label() string

	// The number of layers in the rate map.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/layerCount
	LayerCount() uint

	// The logical size, in pixels, of the viewport coordinate system.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/screenSize
	ScreenSize() MTLSize

	// Returns the dimensions, in pixels, of the area in the render target affected by the rasterization rate map.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/physicalSize(layer:)
	PhysicalSizeForLayer(layerIndex uint) MTLSize

	// The granularity, in physical pixels, at which the rasterization rate varies.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/physicalGranularity
	PhysicalGranularity() MTLSize

	// Converts a point in logical viewport coordinates to the corresponding physical coordinates in a render layer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/physicalCoordinates(screenCoordinates:layer:)
	MapScreenToPhysicalCoordinatesForLayer(screenCoordinates MTLCoordinate2D, layerIndex uint) MTLCoordinate2D

	// Converts a point in physical coordinates inside a layer to its corresponding logical viewport coordinates.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/screenCoordinates(physicalCoordinates:layer:)
	MapPhysicalToScreenCoordinatesForLayer(physicalCoordinates MTLCoordinate2D, layerIndex uint) MTLCoordinate2D

	// The size and alignment requirements to contain the coordinate transformation information in this rate map.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/parameterDataSizeAndAlign
	ParameterBufferSizeAndAlign() MTLSizeAndAlign

	// Copies the parameter data into the provided buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/copyParameterData(buffer:offset:)
	CopyParameterDataToBufferOffset(buffer MTLBuffer, offset uint)
}

// MTLRasterizationRateMapObject wraps an existing Objective-C object that conforms to the MTLRasterizationRateMap protocol.
type MTLRasterizationRateMapObject struct {
	objectivec.Object
}
func (o MTLRasterizationRateMapObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLRasterizationRateMapObjectFromID constructs a [MTLRasterizationRateMapObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLRasterizationRateMapObjectFromID(id objc.ID) MTLRasterizationRateMapObject {
	return MTLRasterizationRateMapObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device object that created the rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/device

func (o MTLRasterizationRateMapObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// A string that identifies the rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/label

func (o MTLRasterizationRateMapObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// The number of layers in the rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/layerCount

func (o MTLRasterizationRateMapObject) LayerCount() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("layerCount"))
	return rv
	}

// The logical size, in pixels, of the viewport coordinate system.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/screenSize

func (o MTLRasterizationRateMapObject) ScreenSize() MTLSize {
	
	rv := objc.Send[MTLSize](o.ID, objc.Sel("screenSize"))
	return rv
	}

// Returns the dimensions, in pixels, of the area in the render target
// affected by the rasterization rate map.
//
// layerIndex: The index of the layer.
//
// # Return Value
// 
// The dimensions, in pixels, of the area in the render target affected by the
// rasterization rate map.
//
// # Discussion
// 
// Your render targets should be at least as large as the physical size
// returned by this method. Each layer may have different rasterization rates
// and therefore different physical size requirements.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/physicalSize(layer:)

func (o MTLRasterizationRateMapObject) PhysicalSizeForLayer(layerIndex uint) MTLSize {
	
	rv := objc.Send[MTLSize](o.ID, objc.Sel("physicalSizeForLayer:"), layerIndex)
	return rv
	}

// The granularity, in physical pixels, at which the rasterization rate
// varies.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/physicalGranularity

func (o MTLRasterizationRateMapObject) PhysicalGranularity() MTLSize {
	
	rv := objc.Send[MTLSize](o.ID, objc.Sel("physicalGranularity"))
	return rv
	}

// Converts a point in logical viewport coordinates to the corresponding
// physical coordinates in a render layer.
//
// screenCoordinates: A point in viewport coordinates.
//
// layerIndex: The index of the rate map to use.
//
// # Return Value
// 
// A point in the layer’s physical coordinate system corresponding to the
// source point.
//
// # Discussion
// 
// The returned coordinates are always less than or equal to the input
// coordinates because the rasterization rate never exceeds 1:1 in any region.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/physicalCoordinates(screenCoordinates:layer:)

func (o MTLRasterizationRateMapObject) MapScreenToPhysicalCoordinatesForLayer(screenCoordinates MTLCoordinate2D, layerIndex uint) MTLCoordinate2D {
	
	rv := objc.Send[MTLCoordinate2D](o.ID, objc.Sel("mapScreenToPhysicalCoordinates:forLayer:"), screenCoordinates, layerIndex)
	return rv
	}

// Converts a point in physical coordinates inside a layer to its
// corresponding logical viewport coordinates.
//
// physicalCoordinates: A point in layer coordinates.
//
// layerIndex: The index of the rate map to use.
//
// # Return Value
// 
// A point in the view coordinates corresponding to the source point.
//
// # Discussion
// 
// The returned coordinates are always greater than or equal to the input
// coordinates because the rasterization rate never exceeds 1:1 in any region.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/screenCoordinates(physicalCoordinates:layer:)

func (o MTLRasterizationRateMapObject) MapPhysicalToScreenCoordinatesForLayer(physicalCoordinates MTLCoordinate2D, layerIndex uint) MTLCoordinate2D {
	
	rv := objc.Send[MTLCoordinate2D](o.ID, objc.Sel("mapPhysicalToScreenCoordinates:forLayer:"), physicalCoordinates, layerIndex)
	return rv
	}

// The size and alignment requirements to contain the coordinate
// transformation information in this rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/parameterDataSizeAndAlign

func (o MTLRasterizationRateMapObject) ParameterBufferSizeAndAlign() MTLSizeAndAlign {
	
	rv := objc.Send[MTLSizeAndAlign](o.ID, objc.Sel("parameterBufferSizeAndAlign"))
	return rv
	}

// Copies the parameter data into the provided buffer.
//
// buffer: The buffer instance to copy the data into. It needs to have an
// [StorageModeShared] storage mode, and there needs to be enough room in the
// buffer to store the data.
//
// offset: The location in the buffer to copy the data to. The offset needs to be a
// multiple of the parameter alignment.
//
// # Discussion
// 
// To convert coordinate values inside your shader, pass the rate map data
// into the shader in an [MTLBuffer] instance. The
// [ParameterBufferSizeAndAlign] property provides the size and alignment
// requirements for the buffer.
// 
// You can convert between screen space and physical fragment space by binding
// the buffer to the shader with type `rasterization_rate_map_data`, then
// constructing `rasterization_rate_map_decoder` with the buffer data. For
// more details, see the “Variable Rasterization Rate” section of the
// [Metal Shading Language Specification].
//
// [Metal Shading Language Specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMap/copyParameterData(buffer:offset:)

func (o MTLRasterizationRateMapObject) CopyParameterDataToBufferOffset(buffer MTLBuffer, offset uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyParameterDataToBuffer:offset:"), buffer, offset)
	}

