// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRasterizationRateSampleArray] class.
var (
	_MTLRasterizationRateSampleArrayClass     MTLRasterizationRateSampleArrayClass
	_MTLRasterizationRateSampleArrayClassOnce sync.Once
)

func getMTLRasterizationRateSampleArrayClass() MTLRasterizationRateSampleArrayClass {
	_MTLRasterizationRateSampleArrayClassOnce.Do(func() {
		_MTLRasterizationRateSampleArrayClass = MTLRasterizationRateSampleArrayClass{class: objc.GetClass("MTLRasterizationRateSampleArray")}
	})
	return _MTLRasterizationRateSampleArrayClass
}

// GetMTLRasterizationRateSampleArrayClass returns the class object for MTLRasterizationRateSampleArray.
func GetMTLRasterizationRateSampleArrayClass() MTLRasterizationRateSampleArrayClass {
	return getMTLRasterizationRateSampleArrayClass()
}

type MTLRasterizationRateSampleArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRasterizationRateSampleArrayClass) Alloc() MTLRasterizationRateSampleArray {
	rv := objc.Send[MTLRasterizationRateSampleArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array instance that contains rasterization rates.
//
// # Overview
// 
// The [MTLRasterizationRateSampleArray.Horizontal] and [MTLRasterizationRateSampleArray.Vertical] properties of an
// [MTLRasterizationRateLayerDescriptor] point to
// [MTLRasterizationRateSampleArray] instances that contains rasterization
// rates for the layer map. You can use array subscript syntax to access the
// samples. [MTLRasterizationRateSampleArray] instances perform bounds
// checking on any memory operations you make to their sample data.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray
type MTLRasterizationRateSampleArray struct {
	objectivec.Object
}

// MTLRasterizationRateSampleArrayFromID constructs a [MTLRasterizationRateSampleArray] from an objc.ID.
//
// An array instance that contains rasterization rates.
func MTLRasterizationRateSampleArrayFromID(id objc.ID) MTLRasterizationRateSampleArray {
	return MTLRasterizationRateSampleArray{objectivec.Object{ID: id}}
}
// NOTE: MTLRasterizationRateSampleArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRasterizationRateSampleArray] class.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray
type IMTLRasterizationRateSampleArray interface {
	objectivec.IObject

	// The horizontal rasterization rates for the layer map’s rows.
	Horizontal() IMTLRasterizationRateSampleArray
	SetHorizontal(value IMTLRasterizationRateSampleArray)
	// The maximum number of rows and columns in the layer map.
	MaxSampleCount() MTLSize
	SetMaxSampleCount(value MTLSize)
	// The number of rows and columns in the layer map.
	SampleCount() MTLSize
	SetSampleCount(value MTLSize)
	// The vertical rasterization rates for the layer map’s rows.
	Vertical() IMTLRasterizationRateSampleArray
	SetVertical(value IMTLRasterizationRateSampleArray)
	// Retrieves the sample value at the specified index.
	ObjectAtIndexedSubscript(index uint) foundation.NSNumber
	// Stores a sample value at the specified index.
	SetObjectAtIndexedSubscript(value foundation.NSNumber, index uint)
}

// Init initializes the instance.
func (r MTLRasterizationRateSampleArray) Init() MTLRasterizationRateSampleArray {
	rv := objc.Send[MTLRasterizationRateSampleArray](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRasterizationRateSampleArray) Autorelease() MTLRasterizationRateSampleArray {
	rv := objc.Send[MTLRasterizationRateSampleArray](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRasterizationRateSampleArray creates a new MTLRasterizationRateSampleArray instance.
func NewMTLRasterizationRateSampleArray() MTLRasterizationRateSampleArray {
	class := getMTLRasterizationRateSampleArrayClass()
	rv := objc.Send[MTLRasterizationRateSampleArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the sample value at the specified index.
//
// index: The index of the element to retrieve.
//
// # Return Value
// 
// An [NSNumber] object. It contains the value of the sample at the specified
// index or `0` if the index you specified is out of bounds.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray/objectAtIndexedSubscript:
func (r MTLRasterizationRateSampleArray) ObjectAtIndexedSubscript(index uint) foundation.NSNumber {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return foundation.NSNumberFromID(rv)
}
// Stores a sample value at the specified index.
//
// value: The new value to set.
//
// index: The index of the element you want to set.
//
// # Discussion
// 
// The method converts the value to a single precision floating-point value.
// If the index you specified is out of bounds, this method does nothing.
//
// See: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray/setObject:atIndexedSubscript:
func (r MTLRasterizationRateSampleArray) SetObjectAtIndexedSubscript(value foundation.NSNumber, index uint) {
	objc.Send[objc.ID](r.ID, objc.Sel("setObject:atIndexedSubscript:"), value, index)
}

// The horizontal rasterization rates for the layer map’s rows.
//
// See: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/horizontal
func (r MTLRasterizationRateSampleArray) Horizontal() IMTLRasterizationRateSampleArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("horizontal"))
	return MTLRasterizationRateSampleArrayFromID(objc.ID(rv))
}
func (r MTLRasterizationRateSampleArray) SetHorizontal(value IMTLRasterizationRateSampleArray) {
	objc.Send[struct{}](r.ID, objc.Sel("setHorizontal:"), value)
}
// The maximum number of rows and columns in the layer map.
//
// See: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/maxsamplecount
func (r MTLRasterizationRateSampleArray) MaxSampleCount() MTLSize {
	rv := objc.Send[MTLSize](r.ID, objc.Sel("maxSampleCount"))
	return MTLSize(rv)
}
func (r MTLRasterizationRateSampleArray) SetMaxSampleCount(value MTLSize) {
	objc.Send[struct{}](r.ID, objc.Sel("setMaxSampleCount:"), value)
}
// The number of rows and columns in the layer map.
//
// See: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/samplecount
func (r MTLRasterizationRateSampleArray) SampleCount() MTLSize {
	rv := objc.Send[MTLSize](r.ID, objc.Sel("sampleCount"))
	return MTLSize(rv)
}
func (r MTLRasterizationRateSampleArray) SetSampleCount(value MTLSize) {
	objc.Send[struct{}](r.ID, objc.Sel("setSampleCount:"), value)
}
// The vertical rasterization rates for the layer map’s rows.
//
// See: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/vertical
func (r MTLRasterizationRateSampleArray) Vertical() IMTLRasterizationRateSampleArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("vertical"))
	return MTLRasterizationRateSampleArrayFromID(objc.ID(rv))
}
func (r MTLRasterizationRateSampleArray) SetVertical(value IMTLRasterizationRateSampleArray) {
	objc.Send[struct{}](r.ID, objc.Sel("setVertical:"), value)
}

