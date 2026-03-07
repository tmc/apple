// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksAverage2DPoolGradientOperation] class.
var (
	_NeuralNetworksAverage2DPoolGradientOperationClass     NeuralNetworksAverage2DPoolGradientOperationClass
	_NeuralNetworksAverage2DPoolGradientOperationClassOnce sync.Once
)

func getNeuralNetworksAverage2DPoolGradientOperationClass() NeuralNetworksAverage2DPoolGradientOperationClass {
	_NeuralNetworksAverage2DPoolGradientOperationClassOnce.Do(func() {
		_NeuralNetworksAverage2DPoolGradientOperationClass = NeuralNetworksAverage2DPoolGradientOperationClass{class: objc.GetClass("NeuralNetworks.Average2DPoolGradientOperation")}
	})
	return _NeuralNetworksAverage2DPoolGradientOperationClass
}

// GetNeuralNetworksAverage2DPoolGradientOperationClass returns the class object for NeuralNetworks.Average2DPoolGradientOperation.
func GetNeuralNetworksAverage2DPoolGradientOperationClass() NeuralNetworksAverage2DPoolGradientOperationClass {
	return getNeuralNetworksAverage2DPoolGradientOperationClass()
}

type NeuralNetworksAverage2DPoolGradientOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksAverage2DPoolGradientOperationClass) Alloc() NeuralNetworksAverage2DPoolGradientOperation {
	rv := objc.Send[NeuralNetworksAverage2DPoolGradientOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Average2DPoolGradientOperation
type NeuralNetworksAverage2DPoolGradientOperation struct {
	NeuralNetworksBase2DPoolOperation
}

// NeuralNetworksAverage2DPoolGradientOperationFromID constructs a [NeuralNetworksAverage2DPoolGradientOperation] from an objc.ID.
func NeuralNetworksAverage2DPoolGradientOperationFromID(id objc.ID) NeuralNetworksAverage2DPoolGradientOperation {
	return NeuralNetworksAverage2DPoolGradientOperation{NeuralNetworksBase2DPoolOperation: NeuralNetworksBase2DPoolOperationFromID(id)}
}
// Ensure NeuralNetworksAverage2DPoolGradientOperation implements INeuralNetworksAverage2DPoolGradientOperation.
var _ INeuralNetworksAverage2DPoolGradientOperation = NeuralNetworksAverage2DPoolGradientOperation{}





// An interface definition for the [NeuralNetworksAverage2DPoolGradientOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Average2DPoolGradientOperation
type INeuralNetworksAverage2DPoolGradientOperation interface {
	INeuralNetworksBase2DPoolOperation
}





// Init initializes the instance.
func (n NeuralNetworksAverage2DPoolGradientOperation) Init() NeuralNetworksAverage2DPoolGradientOperation {
	rv := objc.Send[NeuralNetworksAverage2DPoolGradientOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksAverage2DPoolGradientOperation) Autorelease() NeuralNetworksAverage2DPoolGradientOperation {
	rv := objc.Send[NeuralNetworksAverage2DPoolGradientOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksAverage2DPoolGradientOperation creates a new NeuralNetworksAverage2DPoolGradientOperation instance.
func NewNeuralNetworksAverage2DPoolGradientOperation() NeuralNetworksAverage2DPoolGradientOperation {
	class := getNeuralNetworksAverage2DPoolGradientOperationClass()
	rv := objc.Send[NeuralNetworksAverage2DPoolGradientOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































