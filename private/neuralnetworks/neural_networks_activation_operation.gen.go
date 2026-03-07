// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksActivationOperation] class.
var (
	_NeuralNetworksActivationOperationClass     NeuralNetworksActivationOperationClass
	_NeuralNetworksActivationOperationClassOnce sync.Once
)

func getNeuralNetworksActivationOperationClass() NeuralNetworksActivationOperationClass {
	_NeuralNetworksActivationOperationClassOnce.Do(func() {
		_NeuralNetworksActivationOperationClass = NeuralNetworksActivationOperationClass{class: objc.GetClass("NeuralNetworks.ActivationOperation")}
	})
	return _NeuralNetworksActivationOperationClass
}

// GetNeuralNetworksActivationOperationClass returns the class object for NeuralNetworks.ActivationOperation.
func GetNeuralNetworksActivationOperationClass() NeuralNetworksActivationOperationClass {
	return getNeuralNetworksActivationOperationClass()
}

type NeuralNetworksActivationOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksActivationOperationClass) Alloc() NeuralNetworksActivationOperation {
	rv := objc.Send[NeuralNetworksActivationOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ActivationOperation
type NeuralNetworksActivationOperation struct {
	NeuralNetworksUnaryElementwiseOperation
}

// NeuralNetworksActivationOperationFromID constructs a [NeuralNetworksActivationOperation] from an objc.ID.
func NeuralNetworksActivationOperationFromID(id objc.ID) NeuralNetworksActivationOperation {
	return NeuralNetworksActivationOperation{NeuralNetworksUnaryElementwiseOperation: NeuralNetworksUnaryElementwiseOperationFromID(id)}
}
// Ensure NeuralNetworksActivationOperation implements INeuralNetworksActivationOperation.
var _ INeuralNetworksActivationOperation = NeuralNetworksActivationOperation{}





// An interface definition for the [NeuralNetworksActivationOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ActivationOperation
type INeuralNetworksActivationOperation interface {
	INeuralNetworksUnaryElementwiseOperation
}





// Init initializes the instance.
func (n NeuralNetworksActivationOperation) Init() NeuralNetworksActivationOperation {
	rv := objc.Send[NeuralNetworksActivationOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksActivationOperation) Autorelease() NeuralNetworksActivationOperation {
	rv := objc.Send[NeuralNetworksActivationOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksActivationOperation creates a new NeuralNetworksActivationOperation instance.
func NewNeuralNetworksActivationOperation() NeuralNetworksActivationOperation {
	class := getNeuralNetworksActivationOperationClass()
	rv := objc.Send[NeuralNetworksActivationOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































