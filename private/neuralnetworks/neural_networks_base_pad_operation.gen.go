// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBasePadOperation] class.
var (
	_NeuralNetworksBasePadOperationClass     NeuralNetworksBasePadOperationClass
	_NeuralNetworksBasePadOperationClassOnce sync.Once
)

func getNeuralNetworksBasePadOperationClass() NeuralNetworksBasePadOperationClass {
	_NeuralNetworksBasePadOperationClassOnce.Do(func() {
		_NeuralNetworksBasePadOperationClass = NeuralNetworksBasePadOperationClass{class: objc.GetClass("NeuralNetworks.BasePadOperation")}
	})
	return _NeuralNetworksBasePadOperationClass
}

// GetNeuralNetworksBasePadOperationClass returns the class object for NeuralNetworks.BasePadOperation.
func GetNeuralNetworksBasePadOperationClass() NeuralNetworksBasePadOperationClass {
	return getNeuralNetworksBasePadOperationClass()
}

type NeuralNetworksBasePadOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBasePadOperationClass) Alloc() NeuralNetworksBasePadOperation {
	rv := objc.Send[NeuralNetworksBasePadOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BasePadOperation
type NeuralNetworksBasePadOperation struct {
	objectivec.Object
}

// NeuralNetworksBasePadOperationFromID constructs a [NeuralNetworksBasePadOperation] from an objc.ID.
func NeuralNetworksBasePadOperationFromID(id objc.ID) NeuralNetworksBasePadOperation {
	return NeuralNetworksBasePadOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBasePadOperation implements INeuralNetworksBasePadOperation.
var _ INeuralNetworksBasePadOperation = NeuralNetworksBasePadOperation{}





// An interface definition for the [NeuralNetworksBasePadOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BasePadOperation
type INeuralNetworksBasePadOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBasePadOperation) Init() NeuralNetworksBasePadOperation {
	rv := objc.Send[NeuralNetworksBasePadOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBasePadOperation) Autorelease() NeuralNetworksBasePadOperation {
	rv := objc.Send[NeuralNetworksBasePadOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBasePadOperation creates a new NeuralNetworksBasePadOperation instance.
func NewNeuralNetworksBasePadOperation() NeuralNetworksBasePadOperation {
	class := getNeuralNetworksBasePadOperationClass()
	rv := objc.Send[NeuralNetworksBasePadOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































