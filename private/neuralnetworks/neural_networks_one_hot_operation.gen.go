// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksOneHotOperation] class.
var (
	_NeuralNetworksOneHotOperationClass     NeuralNetworksOneHotOperationClass
	_NeuralNetworksOneHotOperationClassOnce sync.Once
)

func getNeuralNetworksOneHotOperationClass() NeuralNetworksOneHotOperationClass {
	_NeuralNetworksOneHotOperationClassOnce.Do(func() {
		_NeuralNetworksOneHotOperationClass = NeuralNetworksOneHotOperationClass{class: objc.GetClass("NeuralNetworks.OneHotOperation")}
	})
	return _NeuralNetworksOneHotOperationClass
}

// GetNeuralNetworksOneHotOperationClass returns the class object for NeuralNetworks.OneHotOperation.
func GetNeuralNetworksOneHotOperationClass() NeuralNetworksOneHotOperationClass {
	return getNeuralNetworksOneHotOperationClass()
}

type NeuralNetworksOneHotOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksOneHotOperationClass) Alloc() NeuralNetworksOneHotOperation {
	rv := objc.Send[NeuralNetworksOneHotOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.OneHotOperation
type NeuralNetworksOneHotOperation struct {
	objectivec.Object
}

// NeuralNetworksOneHotOperationFromID constructs a [NeuralNetworksOneHotOperation] from an objc.ID.
func NeuralNetworksOneHotOperationFromID(id objc.ID) NeuralNetworksOneHotOperation {
	return NeuralNetworksOneHotOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksOneHotOperation implements INeuralNetworksOneHotOperation.
var _ INeuralNetworksOneHotOperation = NeuralNetworksOneHotOperation{}





// An interface definition for the [NeuralNetworksOneHotOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.OneHotOperation
type INeuralNetworksOneHotOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksOneHotOperation) Init() NeuralNetworksOneHotOperation {
	rv := objc.Send[NeuralNetworksOneHotOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksOneHotOperation) Autorelease() NeuralNetworksOneHotOperation {
	rv := objc.Send[NeuralNetworksOneHotOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksOneHotOperation creates a new NeuralNetworksOneHotOperation instance.
func NewNeuralNetworksOneHotOperation() NeuralNetworksOneHotOperation {
	class := getNeuralNetworksOneHotOperationClass()
	rv := objc.Send[NeuralNetworksOneHotOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































