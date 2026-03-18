// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRasterizationRateMapDescriptor] class.
var (
	_MTLRasterizationRateMapDescriptorClass     MTLRasterizationRateMapDescriptorClass
	_MTLRasterizationRateMapDescriptorClassOnce sync.Once
)

func getMTLRasterizationRateMapDescriptorClass() MTLRasterizationRateMapDescriptorClass {
	_MTLRasterizationRateMapDescriptorClassOnce.Do(func() {
		_MTLRasterizationRateMapDescriptorClass = MTLRasterizationRateMapDescriptorClass{class: objc.GetClass("MTLRasterizationRateMapDescriptor")}
	})
	return _MTLRasterizationRateMapDescriptorClass
}

// GetMTLRasterizationRateMapDescriptorClass returns the class object for MTLRasterizationRateMapDescriptor.
func GetMTLRasterizationRateMapDescriptorClass() MTLRasterizationRateMapDescriptorClass {
	return getMTLRasterizationRateMapDescriptorClass()
}

type MTLRasterizationRateMapDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRasterizationRateMapDescriptorClass) Alloc() MTLRasterizationRateMapDescriptor {
	rv := objc.Send[MTLRasterizationRateMapDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that you use to configure new rasterization rate maps.
//
// # Overview
// 
// To create a new rate map, first create an
// [MTLRasterizationRateMapDescriptor] instance and set its property values.
// Then, create a new rasterization rate-map by calling an [MTLDevice]
// instance’s [NewRasterizationRateMapWithDescriptor] method.
// 
// When creating a rate map, Metal copies into it property values from the
// descriptor. You can reuse a descrptor by modifying its property values,
// which doesn’t affect the other rate-map instances that already exist.
//
// # Identifying the rate map
//
//   - [MTLRasterizationRateMapDescriptor.Label]: A string used to identify the rate map you create with the descriptor.
//   - [MTLRasterizationRateMapDescriptor.SetLabel]
//
// # Configuring the viewport size
//
//   - [MTLRasterizationRateMapDescriptor.ScreenSize]: The size of the viewport coordinate system, in logical pixels.
//   - [MTLRasterizationRateMapDescriptor.SetScreenSize]
//
// # Configuring the rate map layers
//
//   - [MTLRasterizationRateMapDescriptor.LayerCount]: The number of layers in the rate map.
//   - [MTLRasterizationRateMapDescriptor.LayerAtIndex]: Returns the layer description for a layer in the rate map.
//   - [MTLRasterizationRateMapDescriptor.SetLayerAtIndex]: Sets a configuration for a layer rate map.
//   - [MTLRasterizationRateMapDescriptor.Layers]: The rasterization rates for one or more layers in the rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor
type MTLRasterizationRateMapDescriptor struct {
	objectivec.Object
}

// MTLRasterizationRateMapDescriptorFromID constructs a [MTLRasterizationRateMapDescriptor] from an objc.ID.
//
// An object that you use to configure new rasterization rate maps.
func MTLRasterizationRateMapDescriptorFromID(id objc.ID) MTLRasterizationRateMapDescriptor {
	return MTLRasterizationRateMapDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLRasterizationRateMapDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRasterizationRateMapDescriptor] class.
//
// # Identifying the rate map
//
//   - [IMTLRasterizationRateMapDescriptor.Label]: A string used to identify the rate map you create with the descriptor.
//   - [IMTLRasterizationRateMapDescriptor.SetLabel]
//
// # Configuring the viewport size
//
//   - [IMTLRasterizationRateMapDescriptor.ScreenSize]: The size of the viewport coordinate system, in logical pixels.
//   - [IMTLRasterizationRateMapDescriptor.SetScreenSize]
//
// # Configuring the rate map layers
//
//   - [IMTLRasterizationRateMapDescriptor.LayerCount]: The number of layers in the rate map.
//   - [IMTLRasterizationRateMapDescriptor.LayerAtIndex]: Returns the layer description for a layer in the rate map.
//   - [IMTLRasterizationRateMapDescriptor.SetLayerAtIndex]: Sets a configuration for a layer rate map.
//   - [IMTLRasterizationRateMapDescriptor.Layers]: The rasterization rates for one or more layers in the rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor
type IMTLRasterizationRateMapDescriptor interface {
	objectivec.IObject

	// Topic: Identifying the rate map

	// A string used to identify the rate map you create with the descriptor.
	Label() string
	SetLabel(value string)

	// Topic: Configuring the viewport size

	// The size of the viewport coordinate system, in logical pixels.
	ScreenSize() MTLSize
	SetScreenSize(value MTLSize)

	// Topic: Configuring the rate map layers

	// The number of layers in the rate map.
	LayerCount() uint
	// Returns the layer description for a layer in the rate map.
	LayerAtIndex(layerIndex uint) IMTLRasterizationRateLayerDescriptor
	// Sets a configuration for a layer rate map.
	SetLayerAtIndex(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint)
	// The rasterization rates for one or more layers in the rate map.
	Layers() IMTLRasterizationRateLayerArray
}

// Init initializes the instance.
func (r MTLRasterizationRateMapDescriptor) Init() MTLRasterizationRateMapDescriptor {
	rv := objc.Send[MTLRasterizationRateMapDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRasterizationRateMapDescriptor) Autorelease() MTLRasterizationRateMapDescriptor {
	rv := objc.Send[MTLRasterizationRateMapDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRasterizationRateMapDescriptor creates a new MTLRasterizationRateMapDescriptor instance.
func NewMTLRasterizationRateMapDescriptor() MTLRasterizationRateMapDescriptor {
	class := getMTLRasterizationRateMapDescriptorClass()
	rv := objc.Send[MTLRasterizationRateMapDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the layer description for a layer in the rate map.
//
// layerIndex: The entry to return.
//
// # Return Value
// 
// The [MTLRasterizationRateLayerDescriptor] instance for the given index, or
// `nil` if you haven’t set an instance for this index.
//
// # Discussion
// 
// Calling this method is equivalent to using array subscript syntax.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layer(at:)
func (r MTLRasterizationRateMapDescriptor) LayerAtIndex(layerIndex uint) IMTLRasterizationRateLayerDescriptor {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("layerAtIndex:"), layerIndex)
	return MTLRasterizationRateLayerDescriptorFromID(rv)
}

// Sets a configuration for a layer rate map.
//
// layer: A description of a layer to add to the rate map descriptor. Use `nil` to
// remove the layer at that index.
//
// layerIndex: The index to put the new layer description in.
//
// # Discussion
// 
// Calling this method is equivalent to using array subscript syntax.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/setLayer(_:at:)
func (r MTLRasterizationRateMapDescriptor) SetLayerAtIndex(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint) {
	objc.Send[objc.ID](r.ID, objc.Sel("setLayer:atIndex:"), layer, layerIndex)
}

// Creates a rate map descriptor with a given size and identifier.
//
// screenSize: The logical size, in pixels, of the viewport coordinate system.
//
// # Return Value
// 
// A descriptor object whose [ScreenSize] is set to the provided size. You
// need to add at least one layer rate map to the descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:
func (_MTLRasterizationRateMapDescriptorClass MTLRasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSize(screenSize MTLSize) MTLRasterizationRateMapDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLRasterizationRateMapDescriptorClass.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:"), screenSize)
	return MTLRasterizationRateMapDescriptorFromID(rv)
}

// Creates a rate map descriptor with a single rate layer.
//
// screenSize: The logical size, in pixels, of the viewport coordinate system.
//
// layer: A descriptor for the rate layer to create.
//
// # Return Value
// 
// A descriptor object whose [ScreenSize] is set to the provided size. Layer 0
// in the rate map is set to the provided layer descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:layer:
func (_MTLRasterizationRateMapDescriptorClass MTLRasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSizeLayer(screenSize MTLSize, layer IMTLRasterizationRateLayerDescriptor) MTLRasterizationRateMapDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLRasterizationRateMapDescriptorClass.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:layer:"), screenSize, layer)
	return MTLRasterizationRateMapDescriptorFromID(rv)
}

// Creates a rate map descriptor with a set of layer descriptors.
//
// screenSize: The logical size, in pixels, of the viewport coordinate system.
//
// layerCount: The number of array elements in `layers`.
//
// layers: An array of rate layer descriptors for the rate map’s layers.
//
// # Return Value
// 
// A descriptor object whose [ScreenSize] is set to the provided size and
// whose rate map layers are set to the array you provided.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:layerCount:layers:
func (_MTLRasterizationRateMapDescriptorClass MTLRasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSizeLayerCountLayers(screenSize MTLSize, layerCount uint, layers IMTLRasterizationRateLayerDescriptor) MTLRasterizationRateMapDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLRasterizationRateMapDescriptorClass.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:layerCount:layers:"), screenSize, layerCount, layers)
	return MTLRasterizationRateMapDescriptorFromID(rv)
}

// A string used to identify the rate map you create with the descriptor.
//
// # Discussion
// 
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/label
func (r MTLRasterizationRateMapDescriptor) Label() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (r MTLRasterizationRateMapDescriptor) SetLabel(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The size of the viewport coordinate system, in logical pixels.
//
// # Discussion
// 
// Metal ignores the depth component of this property.
// 
// The viewport coordinate system’s origin is always at `(0,0)` and this
// property determines its size.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/screenSize
func (r MTLRasterizationRateMapDescriptor) ScreenSize() MTLSize {
	rv := objc.Send[MTLSize](r.ID, objc.Sel("screenSize"))
	return MTLSize(rv)
}
func (r MTLRasterizationRateMapDescriptor) SetScreenSize(value MTLSize) {
	objc.Send[struct{}](r.ID, objc.Sel("setScreenSize:"), value)
}

// The number of layers in the rate map.
//
// # Discussion
// 
// The value of this property is dynamically determined based on how many
// layers you’ve added to the descriptor. To add a new layer, call
// [SetLayerAtIndex] or use the subscripting operator to assign a layer.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layerCount
func (r MTLRasterizationRateMapDescriptor) LayerCount() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("layerCount"))
	return rv
}

// The rasterization rates for one or more layers in the rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layers
func (r MTLRasterizationRateMapDescriptor) Layers() IMTLRasterizationRateLayerArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("layers"))
	return MTLRasterizationRateLayerArrayFromID(objc.ID(rv))
}

