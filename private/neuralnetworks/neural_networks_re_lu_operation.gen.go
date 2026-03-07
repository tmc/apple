// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksReLUOperation] class.
var (
	_NeuralNetworksReLUOperationClass     NeuralNetworksReLUOperationClass
	_NeuralNetworksReLUOperationClassOnce sync.Once
)

func getNeuralNetworksReLUOperationClass() NeuralNetworksReLUOperationClass {
	_NeuralNetworksReLUOperationClassOnce.Do(func() {
		_NeuralNetworksReLUOperationClass = NeuralNetworksReLUOperationClass{class: objc.GetClass("NeuralNetworks.ReLUOperation")}
	})
	return _NeuralNetworksReLUOperationClass
}

// GetNeuralNetworksReLUOperationClass returns the class object for NeuralNetworks.ReLUOperation.
func GetNeuralNetworksReLUOperationClass() NeuralNetworksReLUOperationClass {
	return getNeuralNetworksReLUOperationClass()
}

type NeuralNetworksReLUOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksReLUOperationClass) Alloc() NeuralNetworksReLUOperation {
	rv := objc.Send[NeuralNetworksReLUOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReLUOperation
type NeuralNetworksReLUOperation struct {
	NeuralNetworksUnaryElementwiseOperation
}

// NeuralNetworksReLUOperationFromID constructs a [NeuralNetworksReLUOperation] from an objc.ID.
func NeuralNetworksReLUOperationFromID(id objc.ID) NeuralNetworksReLUOperation {
	return NeuralNetworksReLUOperation{NeuralNetworksUnaryElementwiseOperation: NeuralNetworksUnaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksReLUOperation implements INeuralNetworksReLUOperation.
var _ INeuralNetworksReLUOperation = NeuralNetworksReLUOperation{}





// An interface definition for the [NeuralNetworksReLUOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReLUOperation
type INeuralNetworksReLUOperation interface {
	INeuralNetworksUnaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksReLUOperation) Init() NeuralNetworksReLUOperation {
	rv := objc.Send[NeuralNetworksReLUOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksReLUOperation) Autorelease() NeuralNetworksReLUOperation {
	rv := objc.Send[NeuralNetworksReLUOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksReLUOperation creates a new NeuralNetworksReLUOperation instance.
func NewNeuralNetworksReLUOperation() NeuralNetworksReLUOperation {
	class := getNeuralNetworksReLUOperationClass()
	rv := objc.Send[NeuralNetworksReLUOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































