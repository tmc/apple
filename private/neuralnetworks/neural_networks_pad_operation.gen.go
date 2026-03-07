// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksPadOperation] class.
var (
	_NeuralNetworksPadOperationClass     NeuralNetworksPadOperationClass
	_NeuralNetworksPadOperationClassOnce sync.Once
)

func getNeuralNetworksPadOperationClass() NeuralNetworksPadOperationClass {
	_NeuralNetworksPadOperationClassOnce.Do(func() {
		_NeuralNetworksPadOperationClass = NeuralNetworksPadOperationClass{class: objc.GetClass("NeuralNetworks.PadOperation")}
	})
	return _NeuralNetworksPadOperationClass
}

// GetNeuralNetworksPadOperationClass returns the class object for NeuralNetworks.PadOperation.
func GetNeuralNetworksPadOperationClass() NeuralNetworksPadOperationClass {
	return getNeuralNetworksPadOperationClass()
}

type NeuralNetworksPadOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksPadOperationClass) Alloc() NeuralNetworksPadOperation {
	rv := objc.Send[NeuralNetworksPadOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PadOperation
type NeuralNetworksPadOperation struct {
	NeuralNetworksBasePadOperation
}

// NeuralNetworksPadOperationFromID constructs a [NeuralNetworksPadOperation] from an objc.ID.
func NeuralNetworksPadOperationFromID(id objc.ID) NeuralNetworksPadOperation {
	return NeuralNetworksPadOperation{NeuralNetworksBasePadOperation: NeuralNetworksBasePadOperationFromID(id)}
}
// Ensure NeuralNetworksPadOperation implements INeuralNetworksPadOperation.
var _ INeuralNetworksPadOperation = NeuralNetworksPadOperation{}





// An interface definition for the [NeuralNetworksPadOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PadOperation
type INeuralNetworksPadOperation interface {
	INeuralNetworksBasePadOperation
}





// Init initializes the instance.
func (n NeuralNetworksPadOperation) Init() NeuralNetworksPadOperation {
	rv := objc.Send[NeuralNetworksPadOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksPadOperation) Autorelease() NeuralNetworksPadOperation {
	rv := objc.Send[NeuralNetworksPadOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksPadOperation creates a new NeuralNetworksPadOperation instance.
func NewNeuralNetworksPadOperation() NeuralNetworksPadOperation {
	class := getNeuralNetworksPadOperationClass()
	rv := objc.Send[NeuralNetworksPadOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































