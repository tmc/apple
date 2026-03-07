// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksLossOperation] class.
var (
	_NeuralNetworksLossOperationClass     NeuralNetworksLossOperationClass
	_NeuralNetworksLossOperationClassOnce sync.Once
)

func getNeuralNetworksLossOperationClass() NeuralNetworksLossOperationClass {
	_NeuralNetworksLossOperationClassOnce.Do(func() {
		_NeuralNetworksLossOperationClass = NeuralNetworksLossOperationClass{class: objc.GetClass("NeuralNetworks.LossOperation")}
	})
	return _NeuralNetworksLossOperationClass
}

// GetNeuralNetworksLossOperationClass returns the class object for NeuralNetworks.LossOperation.
func GetNeuralNetworksLossOperationClass() NeuralNetworksLossOperationClass {
	return getNeuralNetworksLossOperationClass()
}

type NeuralNetworksLossOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLossOperationClass) Alloc() NeuralNetworksLossOperation {
	rv := objc.Send[NeuralNetworksLossOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LossOperation
type NeuralNetworksLossOperation struct {
	NeuralNetworksBaseLossOperation
}

// NeuralNetworksLossOperationFromID constructs a [NeuralNetworksLossOperation] from an objc.ID.
func NeuralNetworksLossOperationFromID(id objc.ID) NeuralNetworksLossOperation {
	return NeuralNetworksLossOperation{NeuralNetworksBaseLossOperation: NeuralNetworksBaseLossOperationFromID(id)}
}
// Ensure NeuralNetworksLossOperation implements INeuralNetworksLossOperation.
var _ INeuralNetworksLossOperation = NeuralNetworksLossOperation{}





// An interface definition for the [NeuralNetworksLossOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LossOperation
type INeuralNetworksLossOperation interface {
	INeuralNetworksBaseLossOperation
}





// Init initializes the instance.
func (n NeuralNetworksLossOperation) Init() NeuralNetworksLossOperation {
	rv := objc.Send[NeuralNetworksLossOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLossOperation) Autorelease() NeuralNetworksLossOperation {
	rv := objc.Send[NeuralNetworksLossOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLossOperation creates a new NeuralNetworksLossOperation instance.
func NewNeuralNetworksLossOperation() NeuralNetworksLossOperation {
	class := getNeuralNetworksLossOperationClass()
	rv := objc.Send[NeuralNetworksLossOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































