// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRasterizationRateLayerArray] class.
var (
	_MTLRasterizationRateLayerArrayClass     MTLRasterizationRateLayerArrayClass
	_MTLRasterizationRateLayerArrayClassOnce sync.Once
)

func getMTLRasterizationRateLayerArrayClass() MTLRasterizationRateLayerArrayClass {
	_MTLRasterizationRateLayerArrayClassOnce.Do(func() {
		_MTLRasterizationRateLayerArrayClass = MTLRasterizationRateLayerArrayClass{class: objc.GetClass("MTLRasterizationRateLayerArray")}
	})
	return _MTLRasterizationRateLayerArrayClass
}

// GetMTLRasterizationRateLayerArrayClass returns the class object for MTLRasterizationRateLayerArray.
func GetMTLRasterizationRateLayerArrayClass() MTLRasterizationRateLayerArrayClass {
	return getMTLRasterizationRateLayerArrayClass()
}

type MTLRasterizationRateLayerArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRasterizationRateLayerArrayClass) Alloc() MTLRasterizationRateLayerArray {
	rv := objc.Send[MTLRasterizationRateLayerArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Descriptions for the rasterization rates to apply to the set of layers in a
// rate map.
//
// # Accessing members of the array
//
//   - [MTLRasterizationRateLayerArray.ObjectAtIndexedSubscript]: Retrieves the sample value at the specified index.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerArray
type MTLRasterizationRateLayerArray struct {
	objectivec.Object
}

// MTLRasterizationRateLayerArrayFromID constructs a [MTLRasterizationRateLayerArray] from an objc.ID.
//
// Descriptions for the rasterization rates to apply to the set of layers in a
// rate map.
func MTLRasterizationRateLayerArrayFromID(id objc.ID) MTLRasterizationRateLayerArray {
	return MTLRasterizationRateLayerArray{objectivec.Object{ID: id}}
}
// NOTE: MTLRasterizationRateLayerArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRasterizationRateLayerArray] class.
//
// # Accessing members of the array
//
//   - [IMTLRasterizationRateLayerArray.ObjectAtIndexedSubscript]: Retrieves the sample value at the specified index.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerArray
type IMTLRasterizationRateLayerArray interface {
	objectivec.IObject

	// Topic: Accessing members of the array

	// Retrieves the sample value at the specified index.
	ObjectAtIndexedSubscript(layerIndex uint) IMTLRasterizationRateLayerDescriptor

	// The number of layers in the rate map.
	LayerCount() int
	SetLayerCount(value int)
	// The rasterization rates for one or more layers in the rate map.
	Layers() IMTLRasterizationRateLayerArray
	SetLayers(value IMTLRasterizationRateLayerArray)
	// Stores a sample value at the specified index.
	SetObjectAtIndexedSubscript(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint)
}

// Init initializes the instance.
func (r MTLRasterizationRateLayerArray) Init() MTLRasterizationRateLayerArray {
	rv := objc.Send[MTLRasterizationRateLayerArray](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRasterizationRateLayerArray) Autorelease() MTLRasterizationRateLayerArray {
	rv := objc.Send[MTLRasterizationRateLayerArray](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRasterizationRateLayerArray creates a new MTLRasterizationRateLayerArray instance.
func NewMTLRasterizationRateLayerArray() MTLRasterizationRateLayerArray {
	class := getMTLRasterizationRateLayerArrayClass()
	rv := objc.Send[MTLRasterizationRateLayerArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the sample value at the specified index.
//
// layerIndex: The index of the sample you want to retrieve.
//
// # Return Value
// 
// An [NSNumber] instance describing the value of the sample at the specified
// index, or `0` if the index is out of range.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerArray/subscript(_:)
func (r MTLRasterizationRateLayerArray) ObjectAtIndexedSubscript(layerIndex uint) IMTLRasterizationRateLayerDescriptor {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("objectAtIndexedSubscript:"), layerIndex)
	return MTLRasterizationRateLayerDescriptorFromID(rv)
}

// Stores a sample value at the specified index.
//
// layer: The layer descriptor to set
//
// layerIndex: The index of the sample you want to set.
//
// # Discussion
// 
// The method converts the value to a single precision floating point value.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerArray/setObject:atIndexedSubscript:
func (r MTLRasterizationRateLayerArray) SetObjectAtIndexedSubscript(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint) {
	objc.Send[objc.ID](r.ID, objc.Sel("setObject:atIndexedSubscript:"), layer, layerIndex)
}

// The number of layers in the rate map.
//
// See: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/layercount
func (r MTLRasterizationRateLayerArray) LayerCount() int {
	rv := objc.Send[int](r.ID, objc.Sel("layerCount"))
	return rv
}
func (r MTLRasterizationRateLayerArray) SetLayerCount(value int) {
	objc.Send[struct{}](r.ID, objc.Sel("setLayerCount:"), value)
}

// The rasterization rates for one or more layers in the rate map.
//
// See: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/layers
func (r MTLRasterizationRateLayerArray) Layers() IMTLRasterizationRateLayerArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("layers"))
	return MTLRasterizationRateLayerArrayFromID(objc.ID(rv))
}
func (r MTLRasterizationRateLayerArray) SetLayers(value IMTLRasterizationRateLayerArray) {
	objc.Send[struct{}](r.ID, objc.Sel("setLayers:"), value)
}

