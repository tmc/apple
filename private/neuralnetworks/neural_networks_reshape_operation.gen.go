// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksReshapeOperation] class.
var (
	_NeuralNetworksReshapeOperationClass     NeuralNetworksReshapeOperationClass
	_NeuralNetworksReshapeOperationClassOnce sync.Once
)

func getNeuralNetworksReshapeOperationClass() NeuralNetworksReshapeOperationClass {
	_NeuralNetworksReshapeOperationClassOnce.Do(func() {
		_NeuralNetworksReshapeOperationClass = NeuralNetworksReshapeOperationClass{class: objc.GetClass("NeuralNetworks.ReshapeOperation")}
	})
	return _NeuralNetworksReshapeOperationClass
}

// GetNeuralNetworksReshapeOperationClass returns the class object for NeuralNetworks.ReshapeOperation.
func GetNeuralNetworksReshapeOperationClass() NeuralNetworksReshapeOperationClass {
	return getNeuralNetworksReshapeOperationClass()
}

type NeuralNetworksReshapeOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksReshapeOperationClass) Alloc() NeuralNetworksReshapeOperation {
	rv := objc.Send[NeuralNetworksReshapeOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReshapeOperation
type NeuralNetworksReshapeOperation struct {
	objectivec.Object
}

// NeuralNetworksReshapeOperationFromID constructs a [NeuralNetworksReshapeOperation] from an objc.ID.
func NeuralNetworksReshapeOperationFromID(id objc.ID) NeuralNetworksReshapeOperation {
	return NeuralNetworksReshapeOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksReshapeOperation implements INeuralNetworksReshapeOperation.
var _ INeuralNetworksReshapeOperation = NeuralNetworksReshapeOperation{}





// An interface definition for the [NeuralNetworksReshapeOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReshapeOperation
type INeuralNetworksReshapeOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksReshapeOperation) Init() NeuralNetworksReshapeOperation {
	rv := objc.Send[NeuralNetworksReshapeOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksReshapeOperation) Autorelease() NeuralNetworksReshapeOperation {
	rv := objc.Send[NeuralNetworksReshapeOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksReshapeOperation creates a new NeuralNetworksReshapeOperation instance.
func NewNeuralNetworksReshapeOperation() NeuralNetworksReshapeOperation {
	class := getNeuralNetworksReshapeOperationClass()
	rv := objc.Send[NeuralNetworksReshapeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































