// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksPadGradientOperation] class.
var (
	_NeuralNetworksPadGradientOperationClass     NeuralNetworksPadGradientOperationClass
	_NeuralNetworksPadGradientOperationClassOnce sync.Once
)

func getNeuralNetworksPadGradientOperationClass() NeuralNetworksPadGradientOperationClass {
	_NeuralNetworksPadGradientOperationClassOnce.Do(func() {
		_NeuralNetworksPadGradientOperationClass = NeuralNetworksPadGradientOperationClass{class: objc.GetClass("NeuralNetworks.PadGradientOperation")}
	})
	return _NeuralNetworksPadGradientOperationClass
}

// GetNeuralNetworksPadGradientOperationClass returns the class object for NeuralNetworks.PadGradientOperation.
func GetNeuralNetworksPadGradientOperationClass() NeuralNetworksPadGradientOperationClass {
	return getNeuralNetworksPadGradientOperationClass()
}

type NeuralNetworksPadGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksPadGradientOperationClass) Alloc() NeuralNetworksPadGradientOperation {
	rv := objc.Send[NeuralNetworksPadGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PadGradientOperation
type NeuralNetworksPadGradientOperation struct {
	NeuralNetworksBasePadOperation
}

// NeuralNetworksPadGradientOperationFromID constructs a [NeuralNetworksPadGradientOperation] from an objc.ID.
func NeuralNetworksPadGradientOperationFromID(id objc.ID) NeuralNetworksPadGradientOperation {
	return NeuralNetworksPadGradientOperation{NeuralNetworksBasePadOperation: NeuralNetworksBasePadOperationFromID(id)}
}
// Ensure NeuralNetworksPadGradientOperation implements INeuralNetworksPadGradientOperation.
var _ INeuralNetworksPadGradientOperation = NeuralNetworksPadGradientOperation{}





// An interface definition for the [NeuralNetworksPadGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PadGradientOperation
type INeuralNetworksPadGradientOperation interface {
	INeuralNetworksBasePadOperation
}





// Init initializes the instance.
func (n NeuralNetworksPadGradientOperation) Init() NeuralNetworksPadGradientOperation {
	rv := objc.Send[NeuralNetworksPadGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksPadGradientOperation) Autorelease() NeuralNetworksPadGradientOperation {
	rv := objc.Send[NeuralNetworksPadGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksPadGradientOperation creates a new NeuralNetworksPadGradientOperation instance.
func NewNeuralNetworksPadGradientOperation() NeuralNetworksPadGradientOperation {
	class := getNeuralNetworksPadGradientOperationClass()
	rv := objc.Send[NeuralNetworksPadGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































