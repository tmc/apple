// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksScatterOperation] class.
var (
	_NeuralNetworksScatterOperationClass     NeuralNetworksScatterOperationClass
	_NeuralNetworksScatterOperationClassOnce sync.Once
)

func getNeuralNetworksScatterOperationClass() NeuralNetworksScatterOperationClass {
	_NeuralNetworksScatterOperationClassOnce.Do(func() {
		_NeuralNetworksScatterOperationClass = NeuralNetworksScatterOperationClass{class: objc.GetClass("NeuralNetworks.ScatterOperation")}
	})
	return _NeuralNetworksScatterOperationClass
}

// GetNeuralNetworksScatterOperationClass returns the class object for NeuralNetworks.ScatterOperation.
func GetNeuralNetworksScatterOperationClass() NeuralNetworksScatterOperationClass {
	return getNeuralNetworksScatterOperationClass()
}

type NeuralNetworksScatterOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksScatterOperationClass) Alloc() NeuralNetworksScatterOperation {
	rv := objc.Send[NeuralNetworksScatterOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ScatterOperation
type NeuralNetworksScatterOperation struct {
	objectivec.Object
}

// NeuralNetworksScatterOperationFromID constructs a [NeuralNetworksScatterOperation] from an objc.ID.
func NeuralNetworksScatterOperationFromID(id objc.ID) NeuralNetworksScatterOperation {
	return NeuralNetworksScatterOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksScatterOperation implements INeuralNetworksScatterOperation.
var _ INeuralNetworksScatterOperation = NeuralNetworksScatterOperation{}





// An interface definition for the [NeuralNetworksScatterOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ScatterOperation
type INeuralNetworksScatterOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksScatterOperation) Init() NeuralNetworksScatterOperation {
	rv := objc.Send[NeuralNetworksScatterOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksScatterOperation) Autorelease() NeuralNetworksScatterOperation {
	rv := objc.Send[NeuralNetworksScatterOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksScatterOperation creates a new NeuralNetworksScatterOperation instance.
func NewNeuralNetworksScatterOperation() NeuralNetworksScatterOperation {
	class := getNeuralNetworksScatterOperationClass()
	rv := objc.Send[NeuralNetworksScatterOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































