// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksLSTMLayerGradientOperation] class.
var (
	_NeuralNetworksLSTMLayerGradientOperationClass     NeuralNetworksLSTMLayerGradientOperationClass
	_NeuralNetworksLSTMLayerGradientOperationClassOnce sync.Once
)

func getNeuralNetworksLSTMLayerGradientOperationClass() NeuralNetworksLSTMLayerGradientOperationClass {
	_NeuralNetworksLSTMLayerGradientOperationClassOnce.Do(func() {
		_NeuralNetworksLSTMLayerGradientOperationClass = NeuralNetworksLSTMLayerGradientOperationClass{class: objc.GetClass("NeuralNetworks.LSTMLayerGradientOperation")}
	})
	return _NeuralNetworksLSTMLayerGradientOperationClass
}

// GetNeuralNetworksLSTMLayerGradientOperationClass returns the class object for NeuralNetworks.LSTMLayerGradientOperation.
func GetNeuralNetworksLSTMLayerGradientOperationClass() NeuralNetworksLSTMLayerGradientOperationClass {
	return getNeuralNetworksLSTMLayerGradientOperationClass()
}

type NeuralNetworksLSTMLayerGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLSTMLayerGradientOperationClass) Alloc() NeuralNetworksLSTMLayerGradientOperation {
	rv := objc.Send[NeuralNetworksLSTMLayerGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LSTMLayerGradientOperation
type NeuralNetworksLSTMLayerGradientOperation struct {
	NeuralNetworksBaseLSTMLayerOperation
}

// NeuralNetworksLSTMLayerGradientOperationFromID constructs a [NeuralNetworksLSTMLayerGradientOperation] from an objc.ID.
func NeuralNetworksLSTMLayerGradientOperationFromID(id objc.ID) NeuralNetworksLSTMLayerGradientOperation {
	return NeuralNetworksLSTMLayerGradientOperation{NeuralNetworksBaseLSTMLayerOperation: NeuralNetworksBaseLSTMLayerOperationFromID(id)}
}
// Ensure NeuralNetworksLSTMLayerGradientOperation implements INeuralNetworksLSTMLayerGradientOperation.
var _ INeuralNetworksLSTMLayerGradientOperation = NeuralNetworksLSTMLayerGradientOperation{}





// An interface definition for the [NeuralNetworksLSTMLayerGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LSTMLayerGradientOperation
type INeuralNetworksLSTMLayerGradientOperation interface {
	INeuralNetworksBaseLSTMLayerOperation
}





// Init initializes the instance.
func (n NeuralNetworksLSTMLayerGradientOperation) Init() NeuralNetworksLSTMLayerGradientOperation {
	rv := objc.Send[NeuralNetworksLSTMLayerGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLSTMLayerGradientOperation) Autorelease() NeuralNetworksLSTMLayerGradientOperation {
	rv := objc.Send[NeuralNetworksLSTMLayerGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLSTMLayerGradientOperation creates a new NeuralNetworksLSTMLayerGradientOperation instance.
func NewNeuralNetworksLSTMLayerGradientOperation() NeuralNetworksLSTMLayerGradientOperation {
	class := getNeuralNetworksLSTMLayerGradientOperationClass()
	rv := objc.Send[NeuralNetworksLSTMLayerGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































