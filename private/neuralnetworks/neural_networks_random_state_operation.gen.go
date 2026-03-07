// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksRandomStateOperation] class.
var (
	_NeuralNetworksRandomStateOperationClass     NeuralNetworksRandomStateOperationClass
	_NeuralNetworksRandomStateOperationClassOnce sync.Once
)

func getNeuralNetworksRandomStateOperationClass() NeuralNetworksRandomStateOperationClass {
	_NeuralNetworksRandomStateOperationClassOnce.Do(func() {
		_NeuralNetworksRandomStateOperationClass = NeuralNetworksRandomStateOperationClass{class: objc.GetClass("NeuralNetworks.RandomStateOperation")}
	})
	return _NeuralNetworksRandomStateOperationClass
}

// GetNeuralNetworksRandomStateOperationClass returns the class object for NeuralNetworks.RandomStateOperation.
func GetNeuralNetworksRandomStateOperationClass() NeuralNetworksRandomStateOperationClass {
	return getNeuralNetworksRandomStateOperationClass()
}

type NeuralNetworksRandomStateOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksRandomStateOperationClass) Alloc() NeuralNetworksRandomStateOperation {
	rv := objc.Send[NeuralNetworksRandomStateOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RandomStateOperation
type NeuralNetworksRandomStateOperation struct {
	objectivec.Object
}

// NeuralNetworksRandomStateOperationFromID constructs a [NeuralNetworksRandomStateOperation] from an objc.ID.
func NeuralNetworksRandomStateOperationFromID(id objc.ID) NeuralNetworksRandomStateOperation {
	return NeuralNetworksRandomStateOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksRandomStateOperation implements INeuralNetworksRandomStateOperation.
var _ INeuralNetworksRandomStateOperation = NeuralNetworksRandomStateOperation{}





// An interface definition for the [NeuralNetworksRandomStateOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RandomStateOperation
type INeuralNetworksRandomStateOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksRandomStateOperation) Init() NeuralNetworksRandomStateOperation {
	rv := objc.Send[NeuralNetworksRandomStateOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksRandomStateOperation) Autorelease() NeuralNetworksRandomStateOperation {
	rv := objc.Send[NeuralNetworksRandomStateOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksRandomStateOperation creates a new NeuralNetworksRandomStateOperation instance.
func NewNeuralNetworksRandomStateOperation() NeuralNetworksRandomStateOperation {
	class := getNeuralNetworksRandomStateOperationClass()
	rv := objc.Send[NeuralNetworksRandomStateOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































