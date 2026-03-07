// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksPassthroughOperation] class.
var (
	_NeuralNetworksPassthroughOperationClass     NeuralNetworksPassthroughOperationClass
	_NeuralNetworksPassthroughOperationClassOnce sync.Once
)

func getNeuralNetworksPassthroughOperationClass() NeuralNetworksPassthroughOperationClass {
	_NeuralNetworksPassthroughOperationClassOnce.Do(func() {
		_NeuralNetworksPassthroughOperationClass = NeuralNetworksPassthroughOperationClass{class: objc.GetClass("NeuralNetworks.PassthroughOperation")}
	})
	return _NeuralNetworksPassthroughOperationClass
}

// GetNeuralNetworksPassthroughOperationClass returns the class object for NeuralNetworks.PassthroughOperation.
func GetNeuralNetworksPassthroughOperationClass() NeuralNetworksPassthroughOperationClass {
	return getNeuralNetworksPassthroughOperationClass()
}

type NeuralNetworksPassthroughOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksPassthroughOperationClass) Alloc() NeuralNetworksPassthroughOperation {
	rv := objc.Send[NeuralNetworksPassthroughOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PassthroughOperation
type NeuralNetworksPassthroughOperation struct {
	objectivec.Object
}

// NeuralNetworksPassthroughOperationFromID constructs a [NeuralNetworksPassthroughOperation] from an objc.ID.
func NeuralNetworksPassthroughOperationFromID(id objc.ID) NeuralNetworksPassthroughOperation {
	return NeuralNetworksPassthroughOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksPassthroughOperation implements INeuralNetworksPassthroughOperation.
var _ INeuralNetworksPassthroughOperation = NeuralNetworksPassthroughOperation{}





// An interface definition for the [NeuralNetworksPassthroughOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PassthroughOperation
type INeuralNetworksPassthroughOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksPassthroughOperation) Init() NeuralNetworksPassthroughOperation {
	rv := objc.Send[NeuralNetworksPassthroughOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksPassthroughOperation) Autorelease() NeuralNetworksPassthroughOperation {
	rv := objc.Send[NeuralNetworksPassthroughOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksPassthroughOperation creates a new NeuralNetworksPassthroughOperation instance.
func NewNeuralNetworksPassthroughOperation() NeuralNetworksPassthroughOperation {
	class := getNeuralNetworksPassthroughOperationClass()
	rv := objc.Send[NeuralNetworksPassthroughOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































