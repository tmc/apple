// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRasterizationRateLayerDescriptor] class.
var (
	_MTLRasterizationRateLayerDescriptorClass     MTLRasterizationRateLayerDescriptorClass
	_MTLRasterizationRateLayerDescriptorClassOnce sync.Once
)

func getMTLRasterizationRateLayerDescriptorClass() MTLRasterizationRateLayerDescriptorClass {
	_MTLRasterizationRateLayerDescriptorClassOnce.Do(func() {
		_MTLRasterizationRateLayerDescriptorClass = MTLRasterizationRateLayerDescriptorClass{class: objc.GetClass("MTLRasterizationRateLayerDescriptor")}
	})
	return _MTLRasterizationRateLayerDescriptorClass
}

// GetMTLRasterizationRateLayerDescriptorClass returns the class object for MTLRasterizationRateLayerDescriptor.
func GetMTLRasterizationRateLayerDescriptorClass() MTLRasterizationRateLayerDescriptorClass {
	return getMTLRasterizationRateLayerDescriptorClass()
}

type MTLRasterizationRateLayerDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLRasterizationRateLayerDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRasterizationRateLayerDescriptorClass) Alloc() MTLRasterizationRateLayerDescriptor {
	rv := objc.Send[MTLRasterizationRateLayerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The minimum rasterization rates to apply to sections of a layer in the
// render target.
//
// # Overview
//
// Use a layer map to divide the logical viewport coordinate system into a 2D
// grid of equal-sized rectangles, and choose different rasterization rates
// for each cell.
//
// Specify rasterization rates using floating-point numbers between `0.0` and
// `1.0`, inclusive. A rate of `1.0` represents the normal rasterization rate,
// where each logical unit is equal to a physical pixel; a rate of `0.5` means
// that two logical units equate to one physical pixel, and so on. A value of
// `0.0` means that the GPU renders at its lowest quality level. When you
// create the map, the device object chooses the nearest rasterization rate
// supported by the GPU that meets or exceeds the rate you specified.
//
// In the layer map, you provide separate rasterization rates for the grid’s
// rows and columns. The horizontal rates specify a horizontal rasterization
// rate for each column, and the vertical rates specify a vertical
// rasterization rate for each row. Each cell calculates its physical size in
// pixels by using the logical size of cells in the map, the horizontal rate
// from the cell’s column, and the vertical rate from its row.
//
// # Creating a layer rasterization rate descriptor
//
//   - [MTLRasterizationRateLayerDescriptor.InitWithSampleCount]: Initializes the layer map with an empty grid.
//
// # Inspecting the layer rate function parameters
//
//   - [MTLRasterizationRateLayerDescriptor.SampleCount]: The number of rows and columns in the layer map.
//   - [MTLRasterizationRateLayerDescriptor.MaxSampleCount]: The maximum number of rows and columns in the layer map.
//   - [MTLRasterizationRateLayerDescriptor.Horizontal]: The horizontal rasterization rates for the layer map’s rows.
//   - [MTLRasterizationRateLayerDescriptor.Vertical]: The vertical rasterization rates for the layer map’s rows.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor
type MTLRasterizationRateLayerDescriptor struct {
	objectivec.Object
}

// MTLRasterizationRateLayerDescriptorFromID constructs a [MTLRasterizationRateLayerDescriptor] from an objc.ID.
//
// The minimum rasterization rates to apply to sections of a layer in the
// render target.
func MTLRasterizationRateLayerDescriptorFromID(id objc.ID) MTLRasterizationRateLayerDescriptor {
	return MTLRasterizationRateLayerDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLRasterizationRateLayerDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRasterizationRateLayerDescriptor] class.
//
// # Creating a layer rasterization rate descriptor
//
//   - [IMTLRasterizationRateLayerDescriptor.InitWithSampleCount]: Initializes the layer map with an empty grid.
//
// # Inspecting the layer rate function parameters
//
//   - [IMTLRasterizationRateLayerDescriptor.SampleCount]: The number of rows and columns in the layer map.
//   - [IMTLRasterizationRateLayerDescriptor.MaxSampleCount]: The maximum number of rows and columns in the layer map.
//   - [IMTLRasterizationRateLayerDescriptor.Horizontal]: The horizontal rasterization rates for the layer map’s rows.
//   - [IMTLRasterizationRateLayerDescriptor.Vertical]: The vertical rasterization rates for the layer map’s rows.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor
type IMTLRasterizationRateLayerDescriptor interface {
	objectivec.IObject

	// Topic: Creating a layer rasterization rate descriptor

	// Initializes the layer map with an empty grid.
	InitWithSampleCount(sampleCount MTLSize) MTLRasterizationRateLayerDescriptor

	// Topic: Inspecting the layer rate function parameters

	// The number of rows and columns in the layer map.
	SampleCount() MTLSize
	// The maximum number of rows and columns in the layer map.
	MaxSampleCount() MTLSize
	// The horizontal rasterization rates for the layer map’s rows.
	Horizontal() IMTLRasterizationRateSampleArray
	// The vertical rasterization rates for the layer map’s rows.
	Vertical() IMTLRasterizationRateSampleArray

	// A pointer to the storage for the layer map’s vertical rasterization rates.
	VerticalSampleStorage() unsafe.Pointer
	// Initializes the layer map with the provided grid size and rasterization rates.
	InitWithSampleCountHorizontalVertical(sampleCount MTLSize, horizontal unsafe.Pointer, vertical unsafe.Pointer) MTLRasterizationRateLayerDescriptor
}

// Init initializes the instance.
func (r MTLRasterizationRateLayerDescriptor) Init() MTLRasterizationRateLayerDescriptor {
	rv := objc.Send[MTLRasterizationRateLayerDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRasterizationRateLayerDescriptor) Autorelease() MTLRasterizationRateLayerDescriptor {
	rv := objc.Send[MTLRasterizationRateLayerDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRasterizationRateLayerDescriptor creates a new MTLRasterizationRateLayerDescriptor instance.
func NewMTLRasterizationRateLayerDescriptor() MTLRasterizationRateLayerDescriptor {
	class := getMTLRasterizationRateLayerDescriptorClass()
	rv := objc.Send[MTLRasterizationRateLayerDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the layer map with an empty grid.
//
// sampleCount: The size of the grid. Specify the width and height to determine the number
// of columns and rows in the layer map. The initializer ignores the depth
// component.
//
// # Return Value
//
// A layer descriptor with a grid of the specified size. All of the
// rasterization rates are set to `0.0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/init(sampleCount:)
func NewRasterizationRateLayerDescriptorWithSampleCount(sampleCount MTLSize) MTLRasterizationRateLayerDescriptor {
	instance := getMTLRasterizationRateLayerDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSampleCount:"), sampleCount)
	return MTLRasterizationRateLayerDescriptorFromID(rv)
}

// Initializes the layer map with the provided grid size and rasterization
// rates.
//
// sampleCount: The size of the grid. Specify the width and height to determine the number
// of columns and rows in the layer map. The initializer ignores the depth
// component.
//
// horizontal: The rasterization rates for the layer map’s columns. There needs to be at
// least as many samples as the width you specified in `sampleCount`.
//
// vertical: The rasterization rates for the layer map’s columns. There needs to be at
// least as many samples as the height you specified in `sampleCount`.
//
// # Return Value
//
// A layer descriptor with a grid of the specified size. The layer descriptor
// copies the rasterization rates.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/initWithSampleCount:horizontal:vertical:
func NewRasterizationRateLayerDescriptorWithSampleCountHorizontalVertical(sampleCount MTLSize, horizontal unsafe.Pointer, vertical unsafe.Pointer) MTLRasterizationRateLayerDescriptor {
	instance := getMTLRasterizationRateLayerDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSampleCount:horizontal:vertical:"), sampleCount, horizontal, vertical)
	return MTLRasterizationRateLayerDescriptorFromID(rv)
}

// Initializes the layer map with an empty grid.
//
// sampleCount: The size of the grid. Specify the width and height to determine the number
// of columns and rows in the layer map. The initializer ignores the depth
// component.
//
// # Return Value
//
// A layer descriptor with a grid of the specified size. All of the
// rasterization rates are set to `0.0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/init(sampleCount:)
func (r MTLRasterizationRateLayerDescriptor) InitWithSampleCount(sampleCount MTLSize) MTLRasterizationRateLayerDescriptor {
	rv := objc.Send[MTLRasterizationRateLayerDescriptor](r.ID, objc.Sel("initWithSampleCount:"), sampleCount)
	return rv
}

// Initializes the layer map with the provided grid size and rasterization
// rates.
//
// sampleCount: The size of the grid. Specify the width and height to determine the number
// of columns and rows in the layer map. The initializer ignores the depth
// component.
//
// horizontal: The rasterization rates for the layer map’s columns. There needs to be at
// least as many samples as the width you specified in `sampleCount`.
//
// vertical: The rasterization rates for the layer map’s columns. There needs to be at
// least as many samples as the height you specified in `sampleCount`.
//
// # Return Value
//
// A layer descriptor with a grid of the specified size. The layer descriptor
// copies the rasterization rates.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/initWithSampleCount:horizontal:vertical:
func (r MTLRasterizationRateLayerDescriptor) InitWithSampleCountHorizontalVertical(sampleCount MTLSize, horizontal unsafe.Pointer, vertical unsafe.Pointer) MTLRasterizationRateLayerDescriptor {
	rv := objc.Send[MTLRasterizationRateLayerDescriptor](r.ID, objc.Sel("initWithSampleCount:horizontal:vertical:"), sampleCount, horizontal, vertical)
	return rv
}

// The number of rows and columns in the layer map.
//
// # Discussion
//
// The [SampleCount] property splits the logical viewport coordinate space
// into a 2D grid of equal-sized cells. Its [depth] value is always `0`.
//
// The default value is the same as [MaxSampleCount].
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/sampleCount
//
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
func (r MTLRasterizationRateLayerDescriptor) SampleCount() MTLSize {
	rv := objc.Send[MTLSize](r.ID, objc.Sel("sampleCount"))
	return MTLSize(rv)
}

// The maximum number of rows and columns in the layer map.
//
// # Discussion
//
// Its [depth] value is always `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/maxSampleCount
//
// [depth]: https://developer.apple.com/documentation/Metal/MTLSize/depth
func (r MTLRasterizationRateLayerDescriptor) MaxSampleCount() MTLSize {
	rv := objc.Send[MTLSize](r.ID, objc.Sel("maxSampleCount"))
	return MTLSize(rv)
}

// The horizontal rasterization rates for the layer map’s rows.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/horizontal
func (r MTLRasterizationRateLayerDescriptor) Horizontal() IMTLRasterizationRateSampleArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("horizontal"))
	return MTLRasterizationRateSampleArrayFromID(objc.ID(rv))
}

// The vertical rasterization rates for the layer map’s rows.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/vertical
func (r MTLRasterizationRateLayerDescriptor) Vertical() IMTLRasterizationRateSampleArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("vertical"))
	return MTLRasterizationRateSampleArrayFromID(objc.ID(rv))
}

// A pointer to the storage for the layer map’s vertical rasterization
// rates.
//
// # Discussion
//
// Points to the first element in the array of vertical rasterization rates.
// The number of elements is equal to the [height] value of [SampleCount].
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/verticalSampleStorage
//
// [height]: https://developer.apple.com/documentation/Metal/MTLSize/height
func (r MTLRasterizationRateLayerDescriptor) VerticalSampleStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r.ID, objc.Sel("verticalSampleStorage"))
	return rv
}
