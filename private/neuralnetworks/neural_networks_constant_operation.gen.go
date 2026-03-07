// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksConstantOperation] class.
var (
	_NeuralNetworksConstantOperationClass     NeuralNetworksConstantOperationClass
	_NeuralNetworksConstantOperationClassOnce sync.Once
)

func getNeuralNetworksConstantOperationClass() NeuralNetworksConstantOperationClass {
	_NeuralNetworksConstantOperationClassOnce.Do(func() {
		_NeuralNetworksConstantOperationClass = NeuralNetworksConstantOperationClass{class: objc.GetClass("NeuralNetworks.ConstantOperation")}
	})
	return _NeuralNetworksConstantOperationClass
}

// GetNeuralNetworksConstantOperationClass returns the class object for NeuralNetworks.ConstantOperation.
func GetNeuralNetworksConstantOperationClass() NeuralNetworksConstantOperationClass {
	return getNeuralNetworksConstantOperationClass()
}

type NeuralNetworksConstantOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksConstantOperationClass) Alloc() NeuralNetworksConstantOperation {
	rv := objc.Send[NeuralNetworksConstantOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConstantOperation
type NeuralNetworksConstantOperation struct {
	objectivec.Object
}

// NeuralNetworksConstantOperationFromID constructs a [NeuralNetworksConstantOperation] from an objc.ID.
func NeuralNetworksConstantOperationFromID(id objc.ID) NeuralNetworksConstantOperation {
	return NeuralNetworksConstantOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksConstantOperation implements INeuralNetworksConstantOperation.
var _ INeuralNetworksConstantOperation = NeuralNetworksConstantOperation{}





// An interface definition for the [NeuralNetworksConstantOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConstantOperation
type INeuralNetworksConstantOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksConstantOperation) Init() NeuralNetworksConstantOperation {
	rv := objc.Send[NeuralNetworksConstantOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksConstantOperation) Autorelease() NeuralNetworksConstantOperation {
	rv := objc.Send[NeuralNetworksConstantOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksConstantOperation creates a new NeuralNetworksConstantOperation instance.
func NewNeuralNetworksConstantOperation() NeuralNetworksConstantOperation {
	class := getNeuralNetworksConstantOperationClass()
	rv := objc.Send[NeuralNetworksConstantOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































