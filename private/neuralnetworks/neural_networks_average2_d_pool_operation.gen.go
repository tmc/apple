// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksAverage2DPoolOperation] class.
var (
	_NeuralNetworksAverage2DPoolOperationClass     NeuralNetworksAverage2DPoolOperationClass
	_NeuralNetworksAverage2DPoolOperationClassOnce sync.Once
)

func getNeuralNetworksAverage2DPoolOperationClass() NeuralNetworksAverage2DPoolOperationClass {
	_NeuralNetworksAverage2DPoolOperationClassOnce.Do(func() {
		_NeuralNetworksAverage2DPoolOperationClass = NeuralNetworksAverage2DPoolOperationClass{class: objc.GetClass("NeuralNetworks.Average2DPoolOperation")}
	})
	return _NeuralNetworksAverage2DPoolOperationClass
}

// GetNeuralNetworksAverage2DPoolOperationClass returns the class object for NeuralNetworks.Average2DPoolOperation.
func GetNeuralNetworksAverage2DPoolOperationClass() NeuralNetworksAverage2DPoolOperationClass {
	return getNeuralNetworksAverage2DPoolOperationClass()
}

type NeuralNetworksAverage2DPoolOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksAverage2DPoolOperationClass) Alloc() NeuralNetworksAverage2DPoolOperation {
	rv := objc.Send[NeuralNetworksAverage2DPoolOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Average2DPoolOperation
type NeuralNetworksAverage2DPoolOperation struct {
	NeuralNetworksBase2DPoolOperation
}

// NeuralNetworksAverage2DPoolOperationFromID constructs a [NeuralNetworksAverage2DPoolOperation] from an objc.ID.
func NeuralNetworksAverage2DPoolOperationFromID(id objc.ID) NeuralNetworksAverage2DPoolOperation {
	return NeuralNetworksAverage2DPoolOperation{NeuralNetworksBase2DPoolOperation: NeuralNetworksBase2DPoolOperationFromID(id)}
}
// Ensure NeuralNetworksAverage2DPoolOperation implements INeuralNetworksAverage2DPoolOperation.
var _ INeuralNetworksAverage2DPoolOperation = NeuralNetworksAverage2DPoolOperation{}





// An interface definition for the [NeuralNetworksAverage2DPoolOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Average2DPoolOperation
type INeuralNetworksAverage2DPoolOperation interface {
	INeuralNetworksBase2DPoolOperation
}





// Init initializes the instance.
func (n NeuralNetworksAverage2DPoolOperation) Init() NeuralNetworksAverage2DPoolOperation {
	rv := objc.Send[NeuralNetworksAverage2DPoolOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksAverage2DPoolOperation) Autorelease() NeuralNetworksAverage2DPoolOperation {
	rv := objc.Send[NeuralNetworksAverage2DPoolOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksAverage2DPoolOperation creates a new NeuralNetworksAverage2DPoolOperation instance.
func NewNeuralNetworksAverage2DPoolOperation() NeuralNetworksAverage2DPoolOperation {
	class := getNeuralNetworksAverage2DPoolOperationClass()
	rv := objc.Send[NeuralNetworksAverage2DPoolOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































