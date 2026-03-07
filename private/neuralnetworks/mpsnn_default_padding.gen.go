// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNDefaultPadding] class.
var (
	_MPSNNDefaultPaddingClass     MPSNNDefaultPaddingClass
	_MPSNNDefaultPaddingClassOnce sync.Once
)

func getMPSNNDefaultPaddingClass() MPSNNDefaultPaddingClass {
	_MPSNNDefaultPaddingClassOnce.Do(func() {
		_MPSNNDefaultPaddingClass = MPSNNDefaultPaddingClass{class: objc.GetClass("MPSNNDefaultPadding")}
	})
	return _MPSNNDefaultPaddingClass
}

// GetMPSNNDefaultPaddingClass returns the class object for MPSNNDefaultPadding.
func GetMPSNNDefaultPaddingClass() MPSNNDefaultPaddingClass {
	return getMPSNNDefaultPaddingClass()
}

type MPSNNDefaultPaddingClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNDefaultPaddingClass) Alloc() MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNDefaultPadding.EncodeWithCoder]
//   - [MPSNNDefaultPadding.Label]
//   - [MPSNNDefaultPadding.PaddingMethod]
//   - [MPSNNDefaultPadding.InitWithCoder]
//   - [MPSNNDefaultPadding.InitWithPaddingMethod]
//   - [MPSNNDefaultPadding.DebugDescription]
//   - [MPSNNDefaultPadding.Description]
//   - [MPSNNDefaultPadding.Hash]
//   - [MPSNNDefaultPadding.Superclass]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding
type MPSNNDefaultPadding struct {
	objectivec.Object
}

// MPSNNDefaultPaddingFromID constructs a [MPSNNDefaultPadding] from an objc.ID.
func MPSNNDefaultPaddingFromID(id objc.ID) MPSNNDefaultPadding {
	return MPSNNDefaultPadding{objectivec.Object{id}}
}
// Ensure MPSNNDefaultPadding implements IMPSNNDefaultPadding.
var _ IMPSNNDefaultPadding = MPSNNDefaultPadding{}





// An interface definition for the [MPSNNDefaultPadding] class.
//
// # Methods
//
//   - [IMPSNNDefaultPadding.EncodeWithCoder]
//   - [IMPSNNDefaultPadding.Label]
//   - [IMPSNNDefaultPadding.PaddingMethod]
//   - [IMPSNNDefaultPadding.InitWithCoder]
//   - [IMPSNNDefaultPadding.InitWithPaddingMethod]
//   - [IMPSNNDefaultPadding.DebugDescription]
//   - [IMPSNNDefaultPadding.Description]
//   - [IMPSNNDefaultPadding.Hash]
//   - [IMPSNNDefaultPadding.Superclass]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding
type IMPSNNDefaultPadding interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder objectivec.IObject)
	Label() objectivec.IObject
	PaddingMethod() uint64
	InitWithCoder(coder objectivec.IObject) MPSNNDefaultPadding
	InitWithPaddingMethod(method uint64) MPSNNDefaultPadding
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}





// Init initializes the instance.
func (d MPSNNDefaultPadding) Init() MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MPSNNDefaultPadding) Autorelease() MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNDefaultPadding creates a new MPSNNDefaultPadding instance.
func NewMPSNNDefaultPadding() MPSNNDefaultPadding {
	class := getMPSNNDefaultPaddingClass()
	rv := objc.Send[MPSNNDefaultPadding](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/initWithCoder:
func NewDefaultPaddingWithCoder(coder objectivec.IObject) MPSNNDefaultPadding {
	instance := getMPSNNDefaultPaddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MPSNNDefaultPaddingFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/initWithPaddingMethod:
func NewDefaultPaddingWithPaddingMethod(method uint64) MPSNNDefaultPadding {
	instance := getMPSNNDefaultPaddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPaddingMethod:"), method)
	return MPSNNDefaultPaddingFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/encodeWithCoder:
func (d MPSNNDefaultPadding) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/label
func (d MPSNNDefaultPadding) Label() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("label"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/paddingMethod
func (d MPSNNDefaultPadding) PaddingMethod() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("paddingMethod"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/initWithCoder:
func (d MPSNNDefaultPadding) InitWithCoder(coder objectivec.IObject) MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/initWithPaddingMethod:
func (d MPSNNDefaultPadding) InitWithPaddingMethod(method uint64) MPSNNDefaultPadding {
	rv := objc.Send[MPSNNDefaultPadding](d.ID, objc.Sel("initWithPaddingMethod:"), method)
	return rv
}





// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/paddingForTensorflowAveragePooling
func (_MPSNNDefaultPaddingClass MPSNNDefaultPaddingClass) PaddingForTensorflowAveragePooling() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNDefaultPaddingClass.class), objc.Sel("paddingForTensorflowAveragePooling"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/paddingForTensorflowAveragePoolingValidOnly
func (_MPSNNDefaultPaddingClass MPSNNDefaultPaddingClass) PaddingForTensorflowAveragePoolingValidOnly() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNDefaultPaddingClass.class), objc.Sel("paddingForTensorflowAveragePoolingValidOnly"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/paddingWithMethod:
func (_MPSNNDefaultPaddingClass MPSNNDefaultPaddingClass) PaddingWithMethod(method uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNDefaultPaddingClass.class), objc.Sel("paddingWithMethod:"), method)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/supportsSecureCoding
func (_MPSNNDefaultPaddingClass MPSNNDefaultPaddingClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MPSNNDefaultPaddingClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/zeroPaddingWithTopAmount:bottomAmount:leftAmount:rightAmount:
func (_MPSNNDefaultPaddingClass MPSNNDefaultPaddingClass) ZeroPaddingWithTopAmountBottomAmountLeftAmountRightAmount(amount uint64, amount2 uint64, amount3 uint64, amount4 uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNDefaultPaddingClass.class), objc.Sel("zeroPaddingWithTopAmount:bottomAmount:leftAmount:rightAmount:"), amount, amount2, amount3, amount4)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/debugDescription
func (d MPSNNDefaultPadding) DebugDescription() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/description
func (d MPSNNDefaultPadding) Description() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/hash
func (d MPSNNDefaultPadding) Hash() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("hash"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/superclass
func (d MPSNNDefaultPadding) Superclass() objc.Class {
	rv := objc.Send[objc.Class](d.ID, objc.Sel("superclass"))
	return rv
}

















