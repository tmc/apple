// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNLossDescriptor] class.
var (
	_MPSCNNLossDescriptorClass     MPSCNNLossDescriptorClass
	_MPSCNNLossDescriptorClassOnce sync.Once
)

func getMPSCNNLossDescriptorClass() MPSCNNLossDescriptorClass {
	_MPSCNNLossDescriptorClassOnce.Do(func() {
		_MPSCNNLossDescriptorClass = MPSCNNLossDescriptorClass{class: objc.GetClass("MPSCNNLossDescriptor")}
	})
	return _MPSCNNLossDescriptorClass
}

// GetMPSCNNLossDescriptorClass returns the class object for MPSCNNLossDescriptor.
func GetMPSCNNLossDescriptorClass() MPSCNNLossDescriptorClass {
	return getMPSCNNLossDescriptorClass()
}

type MPSCNNLossDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLossDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLossDescriptorClass) Alloc() MPSCNNLossDescriptor {
	rv := objc.Send[MPSCNNLossDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies properties used by a loss kernel.
//
// # Instance Properties
//
//   - [MPSCNNLossDescriptor.Delta]
//   - [MPSCNNLossDescriptor.SetDelta]
//   - [MPSCNNLossDescriptor.Epsilon]
//   - [MPSCNNLossDescriptor.SetEpsilon]
//   - [MPSCNNLossDescriptor.LabelSmoothing]
//   - [MPSCNNLossDescriptor.SetLabelSmoothing]
//   - [MPSCNNLossDescriptor.LossType]
//   - [MPSCNNLossDescriptor.SetLossType]
//   - [MPSCNNLossDescriptor.NumberOfClasses]
//   - [MPSCNNLossDescriptor.SetNumberOfClasses]
//   - [MPSCNNLossDescriptor.ReduceAcrossBatch]
//   - [MPSCNNLossDescriptor.SetReduceAcrossBatch]
//   - [MPSCNNLossDescriptor.ReductionType]
//   - [MPSCNNLossDescriptor.SetReductionType]
//   - [MPSCNNLossDescriptor.Weight]
//   - [MPSCNNLossDescriptor.SetWeight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor
type MPSCNNLossDescriptor struct {
	objectivec.Object
}

// MPSCNNLossDescriptorFromID constructs a [MPSCNNLossDescriptor] from an objc.ID.
//
// An object that specifies properties used by a loss kernel.
func MPSCNNLossDescriptorFromID(id objc.ID) MPSCNNLossDescriptor {
	return MPSCNNLossDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSCNNLossDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLossDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSCNNLossDescriptor.Delta]
//   - [IMPSCNNLossDescriptor.SetDelta]
//   - [IMPSCNNLossDescriptor.Epsilon]
//   - [IMPSCNNLossDescriptor.SetEpsilon]
//   - [IMPSCNNLossDescriptor.LabelSmoothing]
//   - [IMPSCNNLossDescriptor.SetLabelSmoothing]
//   - [IMPSCNNLossDescriptor.LossType]
//   - [IMPSCNNLossDescriptor.SetLossType]
//   - [IMPSCNNLossDescriptor.NumberOfClasses]
//   - [IMPSCNNLossDescriptor.SetNumberOfClasses]
//   - [IMPSCNNLossDescriptor.ReduceAcrossBatch]
//   - [IMPSCNNLossDescriptor.SetReduceAcrossBatch]
//   - [IMPSCNNLossDescriptor.ReductionType]
//   - [IMPSCNNLossDescriptor.SetReductionType]
//   - [IMPSCNNLossDescriptor.Weight]
//   - [IMPSCNNLossDescriptor.SetWeight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor
type IMPSCNNLossDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	Delta() float32
	SetDelta(value float32)
	Epsilon() float32
	SetEpsilon(value float32)
	LabelSmoothing() float32
	SetLabelSmoothing(value float32)
	LossType() MPSCNNLossType
	SetLossType(value MPSCNNLossType)
	NumberOfClasses() uint
	SetNumberOfClasses(value uint)
	ReduceAcrossBatch() bool
	SetReduceAcrossBatch(value bool)
	ReductionType() MPSCNNReductionType
	SetReductionType(value MPSCNNReductionType)
	Weight() float32
	SetWeight(value float32)
}

// Init initializes the instance.
func (c MPSCNNLossDescriptor) Init() MPSCNNLossDescriptor {
	rv := objc.Send[MPSCNNLossDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLossDescriptor) Autorelease() MPSCNNLossDescriptor {
	rv := objc.Send[MPSCNNLossDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLossDescriptor creates a new MPSCNNLossDescriptor instance.
func NewMPSCNNLossDescriptor() MPSCNNLossDescriptor {
	class := getMPSCNNLossDescriptorClass()
	rv := objc.Send[MPSCNNLossDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/init(type:reductionType:)
func NewCNNLossDescriptorCnnLossDescriptorWithTypeReductionType(lossType MPSCNNLossType, reductionType MPSCNNReductionType) MPSCNNLossDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNLossDescriptorClass().class), objc.Sel("cnnLossDescriptorWithType:reductionType:"), lossType, reductionType)
	return MPSCNNLossDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/delta
func (c MPSCNNLossDescriptor) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNLossDescriptor) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/epsilon
func (c MPSCNNLossDescriptor) Epsilon() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("epsilon"))
	return rv
}
func (c MPSCNNLossDescriptor) SetEpsilon(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setEpsilon:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/labelSmoothing
func (c MPSCNNLossDescriptor) LabelSmoothing() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("labelSmoothing"))
	return rv
}
func (c MPSCNNLossDescriptor) SetLabelSmoothing(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setLabelSmoothing:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/lossType
func (c MPSCNNLossDescriptor) LossType() MPSCNNLossType {
	rv := objc.Send[MPSCNNLossType](c.ID, objc.Sel("lossType"))
	return MPSCNNLossType(rv)
}
func (c MPSCNNLossDescriptor) SetLossType(value MPSCNNLossType) {
	objc.Send[struct{}](c.ID, objc.Sel("setLossType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/numberOfClasses
func (c MPSCNNLossDescriptor) NumberOfClasses() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("numberOfClasses"))
	return rv
}
func (c MPSCNNLossDescriptor) SetNumberOfClasses(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setNumberOfClasses:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/reduceAcrossBatch
func (c MPSCNNLossDescriptor) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}
func (c MPSCNNLossDescriptor) SetReduceAcrossBatch(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setReduceAcrossBatch:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/reductionType
func (c MPSCNNLossDescriptor) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](c.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}
func (c MPSCNNLossDescriptor) SetReductionType(value MPSCNNReductionType) {
	objc.Send[struct{}](c.ID, objc.Sel("setReductionType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDescriptor/weight
func (c MPSCNNLossDescriptor) Weight() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("weight"))
	return rv
}
func (c MPSCNNLossDescriptor) SetWeight(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setWeight:"), value)
}
