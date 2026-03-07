// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseLSTMLayerOperation] class.
var (
	_NeuralNetworksBaseLSTMLayerOperationClass     NeuralNetworksBaseLSTMLayerOperationClass
	_NeuralNetworksBaseLSTMLayerOperationClassOnce sync.Once
)

func getNeuralNetworksBaseLSTMLayerOperationClass() NeuralNetworksBaseLSTMLayerOperationClass {
	_NeuralNetworksBaseLSTMLayerOperationClassOnce.Do(func() {
		_NeuralNetworksBaseLSTMLayerOperationClass = NeuralNetworksBaseLSTMLayerOperationClass{class: objc.GetClass("NeuralNetworks.BaseLSTMLayerOperation")}
	})
	return _NeuralNetworksBaseLSTMLayerOperationClass
}

// GetNeuralNetworksBaseLSTMLayerOperationClass returns the class object for NeuralNetworks.BaseLSTMLayerOperation.
func GetNeuralNetworksBaseLSTMLayerOperationClass() NeuralNetworksBaseLSTMLayerOperationClass {
	return getNeuralNetworksBaseLSTMLayerOperationClass()
}

type NeuralNetworksBaseLSTMLayerOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseLSTMLayerOperationClass) Alloc() NeuralNetworksBaseLSTMLayerOperation {
	rv := objc.Send[NeuralNetworksBaseLSTMLayerOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseLSTMLayerOperation
type NeuralNetworksBaseLSTMLayerOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseLSTMLayerOperationFromID constructs a [NeuralNetworksBaseLSTMLayerOperation] from an objc.ID.
func NeuralNetworksBaseLSTMLayerOperationFromID(id objc.ID) NeuralNetworksBaseLSTMLayerOperation {
	return NeuralNetworksBaseLSTMLayerOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseLSTMLayerOperation implements INeuralNetworksBaseLSTMLayerOperation.
var _ INeuralNetworksBaseLSTMLayerOperation = NeuralNetworksBaseLSTMLayerOperation{}





// An interface definition for the [NeuralNetworksBaseLSTMLayerOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseLSTMLayerOperation
type INeuralNetworksBaseLSTMLayerOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseLSTMLayerOperation) Init() NeuralNetworksBaseLSTMLayerOperation {
	rv := objc.Send[NeuralNetworksBaseLSTMLayerOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseLSTMLayerOperation) Autorelease() NeuralNetworksBaseLSTMLayerOperation {
	rv := objc.Send[NeuralNetworksBaseLSTMLayerOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseLSTMLayerOperation creates a new NeuralNetworksBaseLSTMLayerOperation instance.
func NewNeuralNetworksBaseLSTMLayerOperation() NeuralNetworksBaseLSTMLayerOperation {
	class := getNeuralNetworksBaseLSTMLayerOperationClass()
	rv := objc.Send[NeuralNetworksBaseLSTMLayerOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































