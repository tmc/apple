// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksRandomOperation] class.
var (
	_NeuralNetworksRandomOperationClass     NeuralNetworksRandomOperationClass
	_NeuralNetworksRandomOperationClassOnce sync.Once
)

func getNeuralNetworksRandomOperationClass() NeuralNetworksRandomOperationClass {
	_NeuralNetworksRandomOperationClassOnce.Do(func() {
		_NeuralNetworksRandomOperationClass = NeuralNetworksRandomOperationClass{class: objc.GetClass("NeuralNetworks.RandomOperation")}
	})
	return _NeuralNetworksRandomOperationClass
}

// GetNeuralNetworksRandomOperationClass returns the class object for NeuralNetworks.RandomOperation.
func GetNeuralNetworksRandomOperationClass() NeuralNetworksRandomOperationClass {
	return getNeuralNetworksRandomOperationClass()
}

type NeuralNetworksRandomOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksRandomOperationClass) Alloc() NeuralNetworksRandomOperation {
	rv := objc.Send[NeuralNetworksRandomOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RandomOperation
type NeuralNetworksRandomOperation struct {
	objectivec.Object
}

// NeuralNetworksRandomOperationFromID constructs a [NeuralNetworksRandomOperation] from an objc.ID.
func NeuralNetworksRandomOperationFromID(id objc.ID) NeuralNetworksRandomOperation {
	return NeuralNetworksRandomOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksRandomOperation implements INeuralNetworksRandomOperation.
var _ INeuralNetworksRandomOperation = NeuralNetworksRandomOperation{}





// An interface definition for the [NeuralNetworksRandomOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.RandomOperation
type INeuralNetworksRandomOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksRandomOperation) Init() NeuralNetworksRandomOperation {
	rv := objc.Send[NeuralNetworksRandomOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksRandomOperation) Autorelease() NeuralNetworksRandomOperation {
	rv := objc.Send[NeuralNetworksRandomOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksRandomOperation creates a new NeuralNetworksRandomOperation instance.
func NewNeuralNetworksRandomOperation() NeuralNetworksRandomOperation {
	class := getNeuralNetworksRandomOperationClass()
	rv := objc.Send[NeuralNetworksRandomOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































