// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageEuclideanDistanceTransform] class.
var (
	_MPSImageEuclideanDistanceTransformClass     MPSImageEuclideanDistanceTransformClass
	_MPSImageEuclideanDistanceTransformClassOnce sync.Once
)

func getMPSImageEuclideanDistanceTransformClass() MPSImageEuclideanDistanceTransformClass {
	_MPSImageEuclideanDistanceTransformClassOnce.Do(func() {
		_MPSImageEuclideanDistanceTransformClass = MPSImageEuclideanDistanceTransformClass{class: objc.GetClass("MPSImageEuclideanDistanceTransform")}
	})
	return _MPSImageEuclideanDistanceTransformClass
}

// GetMPSImageEuclideanDistanceTransformClass returns the class object for MPSImageEuclideanDistanceTransform.
func GetMPSImageEuclideanDistanceTransformClass() MPSImageEuclideanDistanceTransformClass {
	return getMPSImageEuclideanDistanceTransformClass()
}

type MPSImageEuclideanDistanceTransformClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageEuclideanDistanceTransformClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageEuclideanDistanceTransformClass) Alloc() MPSImageEuclideanDistanceTransform {
	rv := objc.Send[MPSImageEuclideanDistanceTransform](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that performs a Euclidean distance transform on an image.
//
// # Limiting the search for nonzero pixels
//
//   - [MPSImageEuclideanDistanceTransform.SearchLimitRadius]: Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius.
//   - [MPSImageEuclideanDistanceTransform.SetSearchLimitRadius]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform
type MPSImageEuclideanDistanceTransform struct {
	MPSUnaryImageKernel
}

// MPSImageEuclideanDistanceTransformFromID constructs a [MPSImageEuclideanDistanceTransform] from an objc.ID.
//
// A filter that performs a Euclidean distance transform on an image.
func MPSImageEuclideanDistanceTransformFromID(id objc.ID) MPSImageEuclideanDistanceTransform {
	return MPSImageEuclideanDistanceTransform{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageEuclideanDistanceTransform adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageEuclideanDistanceTransform] class.
//
// # Limiting the search for nonzero pixels
//
//   - [IMPSImageEuclideanDistanceTransform.SearchLimitRadius]: Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius.
//   - [IMPSImageEuclideanDistanceTransform.SetSearchLimitRadius]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform
type IMPSImageEuclideanDistanceTransform interface {
	IMPSUnaryImageKernel

	// Topic: Limiting the search for nonzero pixels

	// Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius.
	SearchLimitRadius() float32
	SetSearchLimitRadius(value float32)
}

// Init initializes the instance.
func (i MPSImageEuclideanDistanceTransform) Init() MPSImageEuclideanDistanceTransform {
	rv := objc.Send[MPSImageEuclideanDistanceTransform](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageEuclideanDistanceTransform) Autorelease() MPSImageEuclideanDistanceTransform {
	rv := objc.Send[MPSImageEuclideanDistanceTransform](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageEuclideanDistanceTransform creates a new MPSImageEuclideanDistanceTransform instance.
func NewMPSImageEuclideanDistanceTransform() MPSImageEuclideanDistanceTransform {
	class := getMPSImageEuclideanDistanceTransformClass()
	rv := objc.Send[MPSImageEuclideanDistanceTransform](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageEuclideanDistanceTransformWithCoder(aDecoder foundation.INSCoder) MPSImageEuclideanDistanceTransform {
	instance := getMPSImageEuclideanDistanceTransformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageEuclideanDistanceTransformFromID(rv)
}

// Creates a Euclidean distance transform that uses a specified decoder for
// your data and runs on a specified device.
//
// aDecoder: The decoder for your data.
//
// device: The device that the filter runs on.
//
// # Discussion
//
// Use this initializer to specify the location of your data; otherwise, the
// framework may guess incorrectly.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform/init(coder:device:)
func NewImageEuclideanDistanceTransformWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageEuclideanDistanceTransform {
	instance := getMPSImageEuclideanDistanceTransformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageEuclideanDistanceTransformFromID(rv)
}

// Creates a Euclidean distance transform that runs on a specified device.
//
// device: The device that the filter runs on.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform/init(device:)
func NewImageEuclideanDistanceTransformWithDevice(device metal.MTLDevice) MPSImageEuclideanDistanceTransform {
	instance := getMPSImageEuclideanDistanceTransformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageEuclideanDistanceTransformFromID(rv)
}

// Limits the search in an image from a pixel to the closest nonzero pixel
// within a specified radius.
//
// # Discussion
//
// When the nonzero pixels in an input image are far apart, the search
// algorithm works harder to find the closest nonzero pixel. If you don’t
// need results outside a certain radius, use this property to limit the
// search and improve kernel performance. If the nonzero pixels are outside
// the specified radius, the search result is some number larger that the
// radius.
//
// The default value for this property is [greatestFiniteMagnitude], which
// results in an exact Euclidean distance transform. Typical values for this
// property are `32`, `64`, `96`, and `128`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform/searchLimitRadius
//
// [greatestFiniteMagnitude]: https://developer.apple.com/documentation/Swift/Float/greatestFiniteMagnitude
func (i MPSImageEuclideanDistanceTransform) SearchLimitRadius() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("searchLimitRadius"))
	return rv
}
func (i MPSImageEuclideanDistanceTransform) SetSearchLimitRadius(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setSearchLimitRadius:"), value)
}
