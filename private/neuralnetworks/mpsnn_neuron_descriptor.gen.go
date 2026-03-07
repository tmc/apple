// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNNeuronDescriptorClass) Alloc() MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNNeuronDescriptor.A]
//   - [MPSNNNeuronDescriptor.SetA]
//   - [MPSNNNeuronDescriptor.B]
//   - [MPSNNNeuronDescriptor.SetB]
//   - [MPSNNNeuronDescriptor.C]
//   - [MPSNNNeuronDescriptor.SetC]
//   - [MPSNNNeuronDescriptor.Data]
//   - [MPSNNNeuronDescriptor.SetData]
//   - [MPSNNNeuronDescriptor.EncodeWithCoder]
//   - [MPSNNNeuronDescriptor.InitializeWithPReLUWithDataNoCopy]
//   - [MPSNNNeuronDescriptor.InitializeWithTypeABC]
//   - [MPSNNNeuronDescriptor.NeuronInfo]
//   - [MPSNNNeuronDescriptor.NeuronType]
//   - [MPSNNNeuronDescriptor.SetNeuronType]
//   - [MPSNNNeuronDescriptor.InitWithCoder]
//   - [MPSNNNeuronDescriptor.InitWithPReLUWithDataNoCopy]
//   - [MPSNNNeuronDescriptor.InitWithTypeABC]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor
type MPSNNNeuronDescriptor struct {
	objectivec.Object
}

// MPSNNNeuronDescriptorFromID constructs a [MPSNNNeuronDescriptor] from an objc.ID.
func MPSNNNeuronDescriptorFromID(id objc.ID) MPSNNNeuronDescriptor {
	return MPSNNNeuronDescriptor{objectivec.Object{id}}
}
// Ensure MPSNNNeuronDescriptor implements IMPSNNNeuronDescriptor.
var _ IMPSNNNeuronDescriptor = MPSNNNeuronDescriptor{}





// An interface definition for the [MPSNNNeuronDescriptor] class.
//
// # Methods
//
//   - [IMPSNNNeuronDescriptor.A]
//   - [IMPSNNNeuronDescriptor.SetA]
//   - [IMPSNNNeuronDescriptor.B]
//   - [IMPSNNNeuronDescriptor.SetB]
//   - [IMPSNNNeuronDescriptor.C]
//   - [IMPSNNNeuronDescriptor.SetC]
//   - [IMPSNNNeuronDescriptor.Data]
//   - [IMPSNNNeuronDescriptor.SetData]
//   - [IMPSNNNeuronDescriptor.EncodeWithCoder]
//   - [IMPSNNNeuronDescriptor.InitializeWithPReLUWithDataNoCopy]
//   - [IMPSNNNeuronDescriptor.InitializeWithTypeABC]
//   - [IMPSNNNeuronDescriptor.NeuronInfo]
//   - [IMPSNNNeuronDescriptor.NeuronType]
//   - [IMPSNNNeuronDescriptor.SetNeuronType]
//   - [IMPSNNNeuronDescriptor.InitWithCoder]
//   - [IMPSNNNeuronDescriptor.InitWithPReLUWithDataNoCopy]
//   - [IMPSNNNeuronDescriptor.InitWithTypeABC]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor
type IMPSNNNeuronDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	A() float32
	SetA(value float32)
	B() float32
	SetB(value float32)
	C() float32
	SetC(value float32)
	Data() foundation.INSData
	SetData(value foundation.INSData)
	EncodeWithCoder(coder objectivec.IObject)
	InitializeWithPReLUWithDataNoCopy(data objectivec.IObject, copy_ bool) MPSNNNeuronDescriptor
	InitializeWithTypeABC(type_ int, a float32, b float32, c float32) MPSNNNeuronDescriptor
	NeuronInfo() objectivec.IObject
	NeuronType() int
	SetNeuronType(value int)
	InitWithCoder(coder objectivec.IObject) MPSNNNeuronDescriptor
	InitWithPReLUWithDataNoCopy(data objectivec.IObject, copy_ bool) MPSNNNeuronDescriptor
	InitWithTypeABC(type_ int, a float32, b float32, c float32) MPSNNNeuronDescriptor
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initWithCoder:
func NewNeuronDescriptorWithCoder(coder objectivec.IObject) MPSNNNeuronDescriptor {
	instance := getMPSNNNeuronDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MPSNNNeuronDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initWithPReLUWithData:noCopy:
func NewNeuronDescriptorWithPReLUWithDataNoCopy(data objectivec.IObject, copy_ bool) MPSNNNeuronDescriptor {
	instance := getMPSNNNeuronDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPReLUWithData:noCopy:"), data, copy_)
	return MPSNNNeuronDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initWithType:a:b:c:
func NewNeuronDescriptorWithTypeABC(type_ int, a float32, b float32, c float32) MPSNNNeuronDescriptor {
	instance := getMPSNNNeuronDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:a:b:c:"), type_, a, b, c)
	return MPSNNNeuronDescriptorFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/encodeWithCoder:
func (n MPSNNNeuronDescriptor) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initializeWithPReLUWithData:noCopy:
func (n MPSNNNeuronDescriptor) InitializeWithPReLUWithDataNoCopy(data objectivec.IObject, copy_ bool) MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](n.ID, objc.Sel("initializeWithPReLUWithData:noCopy:"), data, copy_)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initializeWithType:a:b:c:
func (n MPSNNNeuronDescriptor) InitializeWithTypeABC(type_ int, a float32, b float32, c float32) MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](n.ID, objc.Sel("initializeWithType:a:b:c:"), type_, a, b, c)
	return rv
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/neuronInfo
func (n MPSNNNeuronDescriptor) NeuronInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("neuronInfo"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initWithCoder:
func (n MPSNNNeuronDescriptor) InitWithCoder(coder objectivec.IObject) MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initWithPReLUWithData:noCopy:
func (n MPSNNNeuronDescriptor) InitWithPReLUWithDataNoCopy(data objectivec.IObject, copy_ bool) MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](n.ID, objc.Sel("initWithPReLUWithData:noCopy:"), data, copy_)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/initWithType:a:b:c:
func (n MPSNNNeuronDescriptor) InitWithTypeABC(type_ int, a float32, b float32, c float32) MPSNNNeuronDescriptor {
	rv := objc.Send[MPSNNNeuronDescriptor](n.ID, objc.Sel("initWithType:a:b:c:"), type_, a, b, c)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/cnnNeuronDescriptorWithType:
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithType(type_ int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:"), type_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/cnnNeuronDescriptorWithType:a:
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithTypeA(type_ int, a float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:a:"), type_, a)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/cnnNeuronDescriptorWithType:a:b:
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithTypeAB(type_ int, a float32, b float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:a:b:"), type_, a, b)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/cnnNeuronDescriptorWithType:a:b:c:
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronDescriptorWithTypeABC(type_ int, a float32, b float32, c float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronDescriptorWithType:a:b:c:"), type_, a, b, c)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/cnnNeuronPReLUDescriptorWithData:noCopy:
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) CnnNeuronPReLUDescriptorWithDataNoCopy(data objectivec.IObject, copy_ bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("cnnNeuronPReLUDescriptorWithData:noCopy:"), data, copy_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/supportsSecureCoding
func (_MPSNNNeuronDescriptorClass MPSNNNeuronDescriptorClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MPSNNNeuronDescriptorClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/a
func (n MPSNNNeuronDescriptor) A() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("a"))
	return rv
}
func (n MPSNNNeuronDescriptor) SetA(value float32) {
	objc.Send[struct{}](n.ID, objc.Sel("setA:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/b
func (n MPSNNNeuronDescriptor) B() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("b"))
	return rv
}
func (n MPSNNNeuronDescriptor) SetB(value float32) {
	objc.Send[struct{}](n.ID, objc.Sel("setB:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/c
func (n MPSNNNeuronDescriptor) C() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("c"))
	return rv
}
func (n MPSNNNeuronDescriptor) SetC(value float32) {
	objc.Send[struct{}](n.ID, objc.Sel("setC:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/data
func (n MPSNNNeuronDescriptor) Data() foundation.INSData {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (n MPSNNNeuronDescriptor) SetData(value foundation.INSData) {
	objc.Send[struct{}](n.ID, objc.Sel("setData:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNNeuronDescriptor/neuronType
func (n MPSNNNeuronDescriptor) NeuronType() int {
	rv := objc.Send[int](n.ID, objc.Sel("neuronType"))
	return rv
}
func (n MPSNNNeuronDescriptor) SetNeuronType(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setNeuronType:"), value)
}

















