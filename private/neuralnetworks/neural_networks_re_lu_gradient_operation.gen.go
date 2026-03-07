// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksReLUGradientOperation] class.
var (
	_NeuralNetworksReLUGradientOperationClass     NeuralNetworksReLUGradientOperationClass
	_NeuralNetworksReLUGradientOperationClassOnce sync.Once
)

func getNeuralNetworksReLUGradientOperationClass() NeuralNetworksReLUGradientOperationClass {
	_NeuralNetworksReLUGradientOperationClassOnce.Do(func() {
		_NeuralNetworksReLUGradientOperationClass = NeuralNetworksReLUGradientOperationClass{class: objc.GetClass("NeuralNetworks.ReLUGradientOperation")}
	})
	return _NeuralNetworksReLUGradientOperationClass
}

// GetNeuralNetworksReLUGradientOperationClass returns the class object for NeuralNetworks.ReLUGradientOperation.
func GetNeuralNetworksReLUGradientOperationClass() NeuralNetworksReLUGradientOperationClass {
	return getNeuralNetworksReLUGradientOperationClass()
}

type NeuralNetworksReLUGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksReLUGradientOperationClass) Alloc() NeuralNetworksReLUGradientOperation {
	rv := objc.Send[NeuralNetworksReLUGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReLUGradientOperation
type NeuralNetworksReLUGradientOperation struct {
	NeuralNetworksBaseUnaryElementwiseOperation
}

// NeuralNetworksReLUGradientOperationFromID constructs a [NeuralNetworksReLUGradientOperation] from an objc.ID.
func NeuralNetworksReLUGradientOperationFromID(id objc.ID) NeuralNetworksReLUGradientOperation {
	return NeuralNetworksReLUGradientOperation{NeuralNetworksBaseUnaryElementwiseOperation: NeuralNetworksBaseUnaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksReLUGradientOperation implements INeuralNetworksReLUGradientOperation.
var _ INeuralNetworksReLUGradientOperation = NeuralNetworksReLUGradientOperation{}





// An interface definition for the [NeuralNetworksReLUGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReLUGradientOperation
type INeuralNetworksReLUGradientOperation interface {
	INeuralNetworksBaseUnaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksReLUGradientOperation) Init() NeuralNetworksReLUGradientOperation {
	rv := objc.Send[NeuralNetworksReLUGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksReLUGradientOperation) Autorelease() NeuralNetworksReLUGradientOperation {
	rv := objc.Send[NeuralNetworksReLUGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksReLUGradientOperation creates a new NeuralNetworksReLUGradientOperation instance.
func NewNeuralNetworksReLUGradientOperation() NeuralNetworksReLUGradientOperation {
	class := getNeuralNetworksReLUGradientOperationClass()
	rv := objc.Send[NeuralNetworksReLUGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































