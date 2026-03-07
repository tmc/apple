// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksLSTMLayerOperation] class.
var (
	_NeuralNetworksLSTMLayerOperationClass     NeuralNetworksLSTMLayerOperationClass
	_NeuralNetworksLSTMLayerOperationClassOnce sync.Once
)

func getNeuralNetworksLSTMLayerOperationClass() NeuralNetworksLSTMLayerOperationClass {
	_NeuralNetworksLSTMLayerOperationClassOnce.Do(func() {
		_NeuralNetworksLSTMLayerOperationClass = NeuralNetworksLSTMLayerOperationClass{class: objc.GetClass("NeuralNetworks.LSTMLayerOperation")}
	})
	return _NeuralNetworksLSTMLayerOperationClass
}

// GetNeuralNetworksLSTMLayerOperationClass returns the class object for NeuralNetworks.LSTMLayerOperation.
func GetNeuralNetworksLSTMLayerOperationClass() NeuralNetworksLSTMLayerOperationClass {
	return getNeuralNetworksLSTMLayerOperationClass()
}

type NeuralNetworksLSTMLayerOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLSTMLayerOperationClass) Alloc() NeuralNetworksLSTMLayerOperation {
	rv := objc.Send[NeuralNetworksLSTMLayerOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LSTMLayerOperation
type NeuralNetworksLSTMLayerOperation struct {
	NeuralNetworksBaseLSTMLayerOperation
}

// NeuralNetworksLSTMLayerOperationFromID constructs a [NeuralNetworksLSTMLayerOperation] from an objc.ID.
func NeuralNetworksLSTMLayerOperationFromID(id objc.ID) NeuralNetworksLSTMLayerOperation {
	return NeuralNetworksLSTMLayerOperation{NeuralNetworksBaseLSTMLayerOperation: NeuralNetworksBaseLSTMLayerOperationFromID(id)}
}
// Ensure NeuralNetworksLSTMLayerOperation implements INeuralNetworksLSTMLayerOperation.
var _ INeuralNetworksLSTMLayerOperation = NeuralNetworksLSTMLayerOperation{}





// An interface definition for the [NeuralNetworksLSTMLayerOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LSTMLayerOperation
type INeuralNetworksLSTMLayerOperation interface {
	INeuralNetworksBaseLSTMLayerOperation
}





// Init initializes the instance.
func (n NeuralNetworksLSTMLayerOperation) Init() NeuralNetworksLSTMLayerOperation {
	rv := objc.Send[NeuralNetworksLSTMLayerOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLSTMLayerOperation) Autorelease() NeuralNetworksLSTMLayerOperation {
	rv := objc.Send[NeuralNetworksLSTMLayerOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLSTMLayerOperation creates a new NeuralNetworksLSTMLayerOperation instance.
func NewNeuralNetworksLSTMLayerOperation() NeuralNetworksLSTMLayerOperation {
	class := getNeuralNetworksLSTMLayerOperationClass()
	rv := objc.Send[NeuralNetworksLSTMLayerOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































