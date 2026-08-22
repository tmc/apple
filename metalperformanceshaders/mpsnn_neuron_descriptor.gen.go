// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNNeuronDescriptor] class.
var (
	_MPSNNNeuronDescriptorClass     MPSNNNeuronDescriptorClass
	_MPSNNNeuronDescriptorClassOnce sync.Once
)

func getMPSNNNeuronDescriptorClass() MPSNNNeuronDescriptorClass {
	_MPSNNNeuronDescriptorClassOnce.Do(func() {
		_MPSNNNeuronDescriptorClass = MPSNNNeuronDescriptorClass{class: objc.GetClass("MPSNNNeuronDescriptor")}
	})
	return _MPSNNNeuronDescriptorClass
}

// GetMPSNNNeuronDescriptorClass returns the class object for MPSNNNeuronDescriptor.
func GetMPSNNNeuronDescriptorClass() MPSNNNeuronDescriptorClass {
	return getMPSNNNeuronDescriptorClass()
}

type MPSNNNeuronDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNNeuronDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNNeuronDescriptorClass) Alloc() MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies properties used by a neuron kernel.
//
// # Instance Properties
//
//   - [MPSNNNeuronDescriptor.A]
//   - [MPSNNNeuronDescriptor.SetA]
//   - [MPSNNNeuronDescriptor.B]
//   - [MPSNNNeuronDescriptor.SetB]
//   - [MPSNNNeuronDescriptor.C]
//   - [MPSNNNeuronDescriptor.SetC]
//   - [MPSNNNeuronDescriptor.Data]
//   - [MPSNNNeuronDescriptor.SetData]
//   - [MPSNNNeuronDescriptor.NeuronType]
//   - [MPSNNNeuronDescriptor.SetNeuronType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor
type MPSNNNeuronDescriptor struct {
	objectivec.Object
}

// MPSNNNeuronDescriptorFromID constructs a [MPSNNNeuronDescriptor] from an objc.ID.
//
// An object that specifies properties used by a neuron kernel.
func MPSNNNeuronDescriptorFromID(id objc.ID) MPSNNNeuronDescriptor {
	return MPSNNNeuronDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSNNNeuronDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNNeuronDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSNNNeuronDescriptor.A]
//   - [IMPSNNNeuronDescriptor.SetA]
//   - [IMPSNNNeuronDescriptor.B]
//   - [IMPSNNNeuronDescriptor.SetB]
//   - [IMPSNNNeuronDescriptor.C]
//   - [IMPSNNNeuronDescriptor.SetC]
//   - [IMPSNNNeuronDescriptor.Data]
//   - [IMPSNNNeuronDescriptor.SetData]
//   - [IMPSNNNeuronDescriptor.NeuronType]
//   - [IMPSNNNeuronDescriptor.SetNeuronType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor
type IMPSNNNeuronDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	A() float32
	SetA(value float32)
	B() float32
	SetB(value float32)
	C() float32
	SetC(value float32)
	Data() foundation.NSData
	SetData(value foundation.NSData)
	NeuronType() MPSCNNNeuronType
	SetNeuronType(value MPSCNNNeuronType)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (n MPSNNNeuronDescriptor) Init() MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNNNeuronDescriptor) Autorelease() MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNNeuronDescriptor creates a new MPSNNNeuronDescriptor instance.
func NewMPSNNNeuronDescriptor() MPSNNNeuronDescriptor {
	class := getMPSNNNeuronDescriptorClass()
	rv := objc.Send[MPSNNNeuronDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (n MPSNNNeuronDescriptor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/cnnNeuronDescriptor(with:)
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithType(neuronType MPSCNNNeuronType) MPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:"), neuronType)
	return MPSNNNeuronDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/cnnNeuronDescriptor(with:a:)
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithTypeA(neuronType MPSCNNNeuronType, a float32) MPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:a:"), neuronType, a)
	return MPSNNNeuronDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/cnnNeuronDescriptor(with:a:b:)
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithTypeAB(neuronType MPSCNNNeuronType, a float32, b float32) MPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:a:b:"), neuronType, a, b)
	return MPSNNNeuronDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/cnnNeuronDescriptor(with:a:b:c:)
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithTypeABC(neuronType MPSCNNNeuronType, a float32, b float32, c float32) MPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:a:b:c:"), neuronType, a, b, c)
	return MPSNNNeuronDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/cnnNeuronPReLUDescriptor(with:noCopy:)
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronPReLUDescriptorWithDataNoCopy(data foundation.NSData, noCopy bool) MPSNNNeuronDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronPReLUDescriptorWithData:noCopy:"), data, noCopy)
	return MPSNNNeuronDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/a
func (n MPSNNNeuronDescriptor) A() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("a"))
	return rv
}
func (n MPSNNNeuronDescriptor) SetA(value float32) {
	objc.Send[struct{}](n.ID, objc.Sel("setA:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/b
func (n MPSNNNeuronDescriptor) B() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("b"))
	return rv
}
func (n MPSNNNeuronDescriptor) SetB(value float32) {
	objc.Send[struct{}](n.ID, objc.Sel("setB:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/c
func (n MPSNNNeuronDescriptor) C() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("c"))
	return rv
}
func (n MPSNNNeuronDescriptor) SetC(value float32) {
	objc.Send[struct{}](n.ID, objc.Sel("setC:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/data
func (n MPSNNNeuronDescriptor) Data() foundation.NSData {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (n MPSNNNeuronDescriptor) SetData(value foundation.NSData) {
	objc.Send[struct{}](n.ID, objc.Sel("setData:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor/neuronType
func (n MPSNNNeuronDescriptor) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](n.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}
func (n MPSNNNeuronDescriptor) SetNeuronType(value MPSCNNNeuronType) {
	objc.Send[struct{}](n.ID, objc.Sel("setNeuronType:"), value)
}
