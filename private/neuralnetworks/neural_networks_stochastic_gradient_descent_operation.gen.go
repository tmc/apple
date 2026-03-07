// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksStochasticGradientDescentOperation] class.
var (
	_NeuralNetworksStochasticGradientDescentOperationClass     NeuralNetworksStochasticGradientDescentOperationClass
	_NeuralNetworksStochasticGradientDescentOperationClassOnce sync.Once
)

func getNeuralNetworksStochasticGradientDescentOperationClass() NeuralNetworksStochasticGradientDescentOperationClass {
	_NeuralNetworksStochasticGradientDescentOperationClassOnce.Do(func() {
		_NeuralNetworksStochasticGradientDescentOperationClass = NeuralNetworksStochasticGradientDescentOperationClass{class: objc.GetClass("NeuralNetworks.StochasticGradientDescentOperation")}
	})
	return _NeuralNetworksStochasticGradientDescentOperationClass
}

// GetNeuralNetworksStochasticGradientDescentOperationClass returns the class object for NeuralNetworks.StochasticGradientDescentOperation.
func GetNeuralNetworksStochasticGradientDescentOperationClass() NeuralNetworksStochasticGradientDescentOperationClass {
	return getNeuralNetworksStochasticGradientDescentOperationClass()
}

type NeuralNetworksStochasticGradientDescentOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksStochasticGradientDescentOperationClass) Alloc() NeuralNetworksStochasticGradientDescentOperation {
	rv := objc.Send[NeuralNetworksStochasticGradientDescentOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.StochasticGradientDescentOperation
type NeuralNetworksStochasticGradientDescentOperation struct {
	NeuralNetworksBaseOptimizerOperation
}

// NeuralNetworksStochasticGradientDescentOperationFromID constructs a [NeuralNetworksStochasticGradientDescentOperation] from an objc.ID.
func NeuralNetworksStochasticGradientDescentOperationFromID(id objc.ID) NeuralNetworksStochasticGradientDescentOperation {
	return NeuralNetworksStochasticGradientDescentOperation{NeuralNetworksBaseOptimizerOperation: NeuralNetworksBaseOptimizerOperationFromID(id)}
}
// Ensure NeuralNetworksStochasticGradientDescentOperation implements INeuralNetworksStochasticGradientDescentOperation.
var _ INeuralNetworksStochasticGradientDescentOperation = NeuralNetworksStochasticGradientDescentOperation{}





// An interface definition for the [NeuralNetworksStochasticGradientDescentOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.StochasticGradientDescentOperation
type INeuralNetworksStochasticGradientDescentOperation interface {
	INeuralNetworksBaseOptimizerOperation
}





// Init initializes the instance.
func (n NeuralNetworksStochasticGradientDescentOperation) Init() NeuralNetworksStochasticGradientDescentOperation {
	rv := objc.Send[NeuralNetworksStochasticGradientDescentOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksStochasticGradientDescentOperation) Autorelease() NeuralNetworksStochasticGradientDescentOperation {
	rv := objc.Send[NeuralNetworksStochasticGradientDescentOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksStochasticGradientDescentOperation creates a new NeuralNetworksStochasticGradientDescentOperation instance.
func NewNeuralNetworksStochasticGradientDescentOperation() NeuralNetworksStochasticGradientDescentOperation {
	class := getNeuralNetworksStochasticGradientDescentOperationClass()
	rv := objc.Send[NeuralNetworksStochasticGradientDescentOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































