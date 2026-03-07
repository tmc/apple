// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksTrackedOperation] class.
var (
	_NeuralNetworksTrackedOperationClass     NeuralNetworksTrackedOperationClass
	_NeuralNetworksTrackedOperationClassOnce sync.Once
)

func getNeuralNetworksTrackedOperationClass() NeuralNetworksTrackedOperationClass {
	_NeuralNetworksTrackedOperationClassOnce.Do(func() {
		_NeuralNetworksTrackedOperationClass = NeuralNetworksTrackedOperationClass{class: objc.GetClass("NeuralNetworks.TrackedOperation")}
	})
	return _NeuralNetworksTrackedOperationClass
}

// GetNeuralNetworksTrackedOperationClass returns the class object for NeuralNetworks.TrackedOperation.
func GetNeuralNetworksTrackedOperationClass() NeuralNetworksTrackedOperationClass {
	return getNeuralNetworksTrackedOperationClass()
}

type NeuralNetworksTrackedOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTrackedOperationClass) Alloc() NeuralNetworksTrackedOperation {
	rv := objc.Send[NeuralNetworksTrackedOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TrackedOperation
type NeuralNetworksTrackedOperation struct {
	NeuralNetworksPassthroughOperation
}

// NeuralNetworksTrackedOperationFromID constructs a [NeuralNetworksTrackedOperation] from an objc.ID.
func NeuralNetworksTrackedOperationFromID(id objc.ID) NeuralNetworksTrackedOperation {
	return NeuralNetworksTrackedOperation{NeuralNetworksPassthroughOperation: NeuralNetworksPassthroughOperationFromID(id)}
}
// Ensure NeuralNetworksTrackedOperation implements INeuralNetworksTrackedOperation.
var _ INeuralNetworksTrackedOperation = NeuralNetworksTrackedOperation{}





// An interface definition for the [NeuralNetworksTrackedOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TrackedOperation
type INeuralNetworksTrackedOperation interface {
	INeuralNetworksPassthroughOperation
}





// Init initializes the instance.
func (n NeuralNetworksTrackedOperation) Init() NeuralNetworksTrackedOperation {
	rv := objc.Send[NeuralNetworksTrackedOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTrackedOperation) Autorelease() NeuralNetworksTrackedOperation {
	rv := objc.Send[NeuralNetworksTrackedOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTrackedOperation creates a new NeuralNetworksTrackedOperation instance.
func NewNeuralNetworksTrackedOperation() NeuralNetworksTrackedOperation {
	class := getNeuralNetworksTrackedOperationClass()
	rv := objc.Send[NeuralNetworksTrackedOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































