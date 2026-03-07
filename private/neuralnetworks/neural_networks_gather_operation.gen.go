// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksGatherOperation] class.
var (
	_NeuralNetworksGatherOperationClass     NeuralNetworksGatherOperationClass
	_NeuralNetworksGatherOperationClassOnce sync.Once
)

func getNeuralNetworksGatherOperationClass() NeuralNetworksGatherOperationClass {
	_NeuralNetworksGatherOperationClassOnce.Do(func() {
		_NeuralNetworksGatherOperationClass = NeuralNetworksGatherOperationClass{class: objc.GetClass("NeuralNetworks.GatherOperation")}
	})
	return _NeuralNetworksGatherOperationClass
}

// GetNeuralNetworksGatherOperationClass returns the class object for NeuralNetworks.GatherOperation.
func GetNeuralNetworksGatherOperationClass() NeuralNetworksGatherOperationClass {
	return getNeuralNetworksGatherOperationClass()
}

type NeuralNetworksGatherOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksGatherOperationClass) Alloc() NeuralNetworksGatherOperation {
	rv := objc.Send[NeuralNetworksGatherOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.GatherOperation
type NeuralNetworksGatherOperation struct {
	objectivec.Object
}

// NeuralNetworksGatherOperationFromID constructs a [NeuralNetworksGatherOperation] from an objc.ID.
func NeuralNetworksGatherOperationFromID(id objc.ID) NeuralNetworksGatherOperation {
	return NeuralNetworksGatherOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksGatherOperation implements INeuralNetworksGatherOperation.
var _ INeuralNetworksGatherOperation = NeuralNetworksGatherOperation{}





// An interface definition for the [NeuralNetworksGatherOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.GatherOperation
type INeuralNetworksGatherOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksGatherOperation) Init() NeuralNetworksGatherOperation {
	rv := objc.Send[NeuralNetworksGatherOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksGatherOperation) Autorelease() NeuralNetworksGatherOperation {
	rv := objc.Send[NeuralNetworksGatherOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksGatherOperation creates a new NeuralNetworksGatherOperation instance.
func NewNeuralNetworksGatherOperation() NeuralNetworksGatherOperation {
	class := getNeuralNetworksGatherOperationClass()
	rv := objc.Send[NeuralNetworksGatherOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































