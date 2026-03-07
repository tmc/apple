// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksCustomOperation] class.
var (
	_NeuralNetworksCustomOperationClass     NeuralNetworksCustomOperationClass
	_NeuralNetworksCustomOperationClassOnce sync.Once
)

func getNeuralNetworksCustomOperationClass() NeuralNetworksCustomOperationClass {
	_NeuralNetworksCustomOperationClassOnce.Do(func() {
		_NeuralNetworksCustomOperationClass = NeuralNetworksCustomOperationClass{class: objc.GetClass("NeuralNetworks.CustomOperation")}
	})
	return _NeuralNetworksCustomOperationClass
}

// GetNeuralNetworksCustomOperationClass returns the class object for NeuralNetworks.CustomOperation.
func GetNeuralNetworksCustomOperationClass() NeuralNetworksCustomOperationClass {
	return getNeuralNetworksCustomOperationClass()
}

type NeuralNetworksCustomOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksCustomOperationClass) Alloc() NeuralNetworksCustomOperation {
	rv := objc.Send[NeuralNetworksCustomOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.CustomOperation
type NeuralNetworksCustomOperation struct {
	objectivec.Object
}

// NeuralNetworksCustomOperationFromID constructs a [NeuralNetworksCustomOperation] from an objc.ID.
func NeuralNetworksCustomOperationFromID(id objc.ID) NeuralNetworksCustomOperation {
	return NeuralNetworksCustomOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksCustomOperation implements INeuralNetworksCustomOperation.
var _ INeuralNetworksCustomOperation = NeuralNetworksCustomOperation{}





// An interface definition for the [NeuralNetworksCustomOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.CustomOperation
type INeuralNetworksCustomOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksCustomOperation) Init() NeuralNetworksCustomOperation {
	rv := objc.Send[NeuralNetworksCustomOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksCustomOperation) Autorelease() NeuralNetworksCustomOperation {
	rv := objc.Send[NeuralNetworksCustomOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksCustomOperation creates a new NeuralNetworksCustomOperation instance.
func NewNeuralNetworksCustomOperation() NeuralNetworksCustomOperation {
	class := getNeuralNetworksCustomOperationClass()
	rv := objc.Send[NeuralNetworksCustomOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































