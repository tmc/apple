// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBroadcastOperation] class.
var (
	_NeuralNetworksBroadcastOperationClass     NeuralNetworksBroadcastOperationClass
	_NeuralNetworksBroadcastOperationClassOnce sync.Once
)

func getNeuralNetworksBroadcastOperationClass() NeuralNetworksBroadcastOperationClass {
	_NeuralNetworksBroadcastOperationClassOnce.Do(func() {
		_NeuralNetworksBroadcastOperationClass = NeuralNetworksBroadcastOperationClass{class: objc.GetClass("NeuralNetworks.BroadcastOperation")}
	})
	return _NeuralNetworksBroadcastOperationClass
}

// GetNeuralNetworksBroadcastOperationClass returns the class object for NeuralNetworks.BroadcastOperation.
func GetNeuralNetworksBroadcastOperationClass() NeuralNetworksBroadcastOperationClass {
	return getNeuralNetworksBroadcastOperationClass()
}

type NeuralNetworksBroadcastOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBroadcastOperationClass) Alloc() NeuralNetworksBroadcastOperation {
	rv := objc.Send[NeuralNetworksBroadcastOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BroadcastOperation
type NeuralNetworksBroadcastOperation struct {
	objectivec.Object
}

// NeuralNetworksBroadcastOperationFromID constructs a [NeuralNetworksBroadcastOperation] from an objc.ID.
func NeuralNetworksBroadcastOperationFromID(id objc.ID) NeuralNetworksBroadcastOperation {
	return NeuralNetworksBroadcastOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBroadcastOperation implements INeuralNetworksBroadcastOperation.
var _ INeuralNetworksBroadcastOperation = NeuralNetworksBroadcastOperation{}





// An interface definition for the [NeuralNetworksBroadcastOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BroadcastOperation
type INeuralNetworksBroadcastOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBroadcastOperation) Init() NeuralNetworksBroadcastOperation {
	rv := objc.Send[NeuralNetworksBroadcastOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBroadcastOperation) Autorelease() NeuralNetworksBroadcastOperation {
	rv := objc.Send[NeuralNetworksBroadcastOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBroadcastOperation creates a new NeuralNetworksBroadcastOperation instance.
func NewNeuralNetworksBroadcastOperation() NeuralNetworksBroadcastOperation {
	class := getNeuralNetworksBroadcastOperationClass()
	rv := objc.Send[NeuralNetworksBroadcastOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































