// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksMax2DPoolGradientOperation] class.
var (
	_NeuralNetworksMax2DPoolGradientOperationClass     NeuralNetworksMax2DPoolGradientOperationClass
	_NeuralNetworksMax2DPoolGradientOperationClassOnce sync.Once
)

func getNeuralNetworksMax2DPoolGradientOperationClass() NeuralNetworksMax2DPoolGradientOperationClass {
	_NeuralNetworksMax2DPoolGradientOperationClassOnce.Do(func() {
		_NeuralNetworksMax2DPoolGradientOperationClass = NeuralNetworksMax2DPoolGradientOperationClass{class: objc.GetClass("NeuralNetworks.Max2DPoolGradientOperation")}
	})
	return _NeuralNetworksMax2DPoolGradientOperationClass
}

// GetNeuralNetworksMax2DPoolGradientOperationClass returns the class object for NeuralNetworks.Max2DPoolGradientOperation.
func GetNeuralNetworksMax2DPoolGradientOperationClass() NeuralNetworksMax2DPoolGradientOperationClass {
	return getNeuralNetworksMax2DPoolGradientOperationClass()
}

type NeuralNetworksMax2DPoolGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMax2DPoolGradientOperationClass) Alloc() NeuralNetworksMax2DPoolGradientOperation {
	rv := objc.Send[NeuralNetworksMax2DPoolGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Max2DPoolGradientOperation
type NeuralNetworksMax2DPoolGradientOperation struct {
	NeuralNetworksBase2DPoolOperation
}

// NeuralNetworksMax2DPoolGradientOperationFromID constructs a [NeuralNetworksMax2DPoolGradientOperation] from an objc.ID.
func NeuralNetworksMax2DPoolGradientOperationFromID(id objc.ID) NeuralNetworksMax2DPoolGradientOperation {
	return NeuralNetworksMax2DPoolGradientOperation{NeuralNetworksBase2DPoolOperation: NeuralNetworksBase2DPoolOperationFromID(id)}
}
// Ensure NeuralNetworksMax2DPoolGradientOperation implements INeuralNetworksMax2DPoolGradientOperation.
var _ INeuralNetworksMax2DPoolGradientOperation = NeuralNetworksMax2DPoolGradientOperation{}





// An interface definition for the [NeuralNetworksMax2DPoolGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Max2DPoolGradientOperation
type INeuralNetworksMax2DPoolGradientOperation interface {
	INeuralNetworksBase2DPoolOperation
}





// Init initializes the instance.
func (n NeuralNetworksMax2DPoolGradientOperation) Init() NeuralNetworksMax2DPoolGradientOperation {
	rv := objc.Send[NeuralNetworksMax2DPoolGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMax2DPoolGradientOperation) Autorelease() NeuralNetworksMax2DPoolGradientOperation {
	rv := objc.Send[NeuralNetworksMax2DPoolGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMax2DPoolGradientOperation creates a new NeuralNetworksMax2DPoolGradientOperation instance.
func NewNeuralNetworksMax2DPoolGradientOperation() NeuralNetworksMax2DPoolGradientOperation {
	class := getNeuralNetworksMax2DPoolGradientOperationClass()
	rv := objc.Send[NeuralNetworksMax2DPoolGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































