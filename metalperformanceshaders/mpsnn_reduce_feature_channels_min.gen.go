// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsMin] class.
var (
	_MPSNNReduceFeatureChannelsMinClass     MPSNNReduceFeatureChannelsMinClass
	_MPSNNReduceFeatureChannelsMinClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsMinClass() MPSNNReduceFeatureChannelsMinClass {
	_MPSNNReduceFeatureChannelsMinClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsMinClass = MPSNNReduceFeatureChannelsMinClass{class: objc.GetClass("MPSNNReduceFeatureChannelsMin")}
	})
	return _MPSNNReduceFeatureChannelsMinClass
}

// GetMPSNNReduceFeatureChannelsMinClass returns the class object for MPSNNReduceFeatureChannelsMin.
func GetMPSNNReduceFeatureChannelsMinClass() MPSNNReduceFeatureChannelsMinClass {
	return getMPSNNReduceFeatureChannelsMinClass()
}

type MPSNNReduceFeatureChannelsMinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsMinClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsMinClass) Alloc() MPSNNReduceFeatureChannelsMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the minimum value for each feature channel
// in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMin
type MPSNNReduceFeatureChannelsMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsMinFromID constructs a [MPSNNReduceFeatureChannelsMin] from an objc.ID.
//
// A reduction filter that returns the minimum value for each feature channel
// in an image.
func MPSNNReduceFeatureChannelsMinFromID(id objc.ID) MPSNNReduceFeatureChannelsMin {
	return MPSNNReduceFeatureChannelsMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsMin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsMin] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMin
type IMPSNNReduceFeatureChannelsMin interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsMin) Init() MPSNNReduceFeatureChannelsMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsMin) Autorelease() MPSNNReduceFeatureChannelsMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsMin creates a new MPSNNReduceFeatureChannelsMin instance.
func NewMPSNNReduceFeatureChannelsMin() MPSNNReduceFeatureChannelsMin {
	class := getMPSNNReduceFeatureChannelsMinClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsMinWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsMin {
	instance := getMPSNNReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMin/init(coder:device:)
func NewReduceFeatureChannelsMinWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsMin {
	instance := getMPSNNReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMin/init(device:)
func NewReduceFeatureChannelsMinWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsMin {
	instance := getMPSNNReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsMinFromID(rv)
}
