// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksTopKGradientOperation] class.
var (
	_NeuralNetworksTopKGradientOperationClass     NeuralNetworksTopKGradientOperationClass
	_NeuralNetworksTopKGradientOperationClassOnce sync.Once
)

func getNeuralNetworksTopKGradientOperationClass() NeuralNetworksTopKGradientOperationClass {
	_NeuralNetworksTopKGradientOperationClassOnce.Do(func() {
		_NeuralNetworksTopKGradientOperationClass = NeuralNetworksTopKGradientOperationClass{class: objc.GetClass("NeuralNetworks.TopKGradientOperation")}
	})
	return _NeuralNetworksTopKGradientOperationClass
}

// GetNeuralNetworksTopKGradientOperationClass returns the class object for NeuralNetworks.TopKGradientOperation.
func GetNeuralNetworksTopKGradientOperationClass() NeuralNetworksTopKGradientOperationClass {
	return getNeuralNetworksTopKGradientOperationClass()
}

type NeuralNetworksTopKGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTopKGradientOperationClass) Alloc() NeuralNetworksTopKGradientOperation {
	rv := objc.Send[NeuralNetworksTopKGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TopKGradientOperation
type NeuralNetworksTopKGradientOperation struct {
	objectivec.Object
}

// NeuralNetworksTopKGradientOperationFromID constructs a [NeuralNetworksTopKGradientOperation] from an objc.ID.
func NeuralNetworksTopKGradientOperationFromID(id objc.ID) NeuralNetworksTopKGradientOperation {
	return NeuralNetworksTopKGradientOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksTopKGradientOperation implements INeuralNetworksTopKGradientOperation.
var _ INeuralNetworksTopKGradientOperation = NeuralNetworksTopKGradientOperation{}





// An interface definition for the [NeuralNetworksTopKGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TopKGradientOperation
type INeuralNetworksTopKGradientOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksTopKGradientOperation) Init() NeuralNetworksTopKGradientOperation {
	rv := objc.Send[NeuralNetworksTopKGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTopKGradientOperation) Autorelease() NeuralNetworksTopKGradientOperation {
	rv := objc.Send[NeuralNetworksTopKGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTopKGradientOperation creates a new NeuralNetworksTopKGradientOperation instance.
func NewNeuralNetworksTopKGradientOperation() NeuralNetworksTopKGradientOperation {
	class := getNeuralNetworksTopKGradientOperationClass()
	rv := objc.Send[NeuralNetworksTopKGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}






































