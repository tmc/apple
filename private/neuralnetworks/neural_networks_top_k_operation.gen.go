// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksTopKOperation] class.
var (
	_NeuralNetworksTopKOperationClass     NeuralNetworksTopKOperationClass
	_NeuralNetworksTopKOperationClassOnce sync.Once
)

func getNeuralNetworksTopKOperationClass() NeuralNetworksTopKOperationClass {
	_NeuralNetworksTopKOperationClassOnce.Do(func() {
		_NeuralNetworksTopKOperationClass = NeuralNetworksTopKOperationClass{class: objc.GetClass("NeuralNetworks.TopKOperation")}
	})
	return _NeuralNetworksTopKOperationClass
}

// GetNeuralNetworksTopKOperationClass returns the class object for NeuralNetworks.TopKOperation.
func GetNeuralNetworksTopKOperationClass() NeuralNetworksTopKOperationClass {
	return getNeuralNetworksTopKOperationClass()
}

type NeuralNetworksTopKOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTopKOperationClass) Alloc() NeuralNetworksTopKOperation {
	rv := objc.Send[NeuralNetworksTopKOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TopKOperation
type NeuralNetworksTopKOperation struct {
	objectivec.Object
}

// NeuralNetworksTopKOperationFromID constructs a [NeuralNetworksTopKOperation] from an objc.ID.
func NeuralNetworksTopKOperationFromID(id objc.ID) NeuralNetworksTopKOperation {
	return NeuralNetworksTopKOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksTopKOperation implements INeuralNetworksTopKOperation.
var _ INeuralNetworksTopKOperation = NeuralNetworksTopKOperation{}





// An interface definition for the [NeuralNetworksTopKOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TopKOperation
type INeuralNetworksTopKOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksTopKOperation) Init() NeuralNetworksTopKOperation {
	rv := objc.Send[NeuralNetworksTopKOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTopKOperation) Autorelease() NeuralNetworksTopKOperation {
	rv := objc.Send[NeuralNetworksTopKOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTopKOperation creates a new NeuralNetworksTopKOperation instance.
func NewNeuralNetworksTopKOperation() NeuralNetworksTopKOperation {
	class := getNeuralNetworksTopKOperationClass()
	rv := objc.Send[NeuralNetworksTopKOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































