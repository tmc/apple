// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksMax2DPoolOperation] class.
var (
	_NeuralNetworksMax2DPoolOperationClass     NeuralNetworksMax2DPoolOperationClass
	_NeuralNetworksMax2DPoolOperationClassOnce sync.Once
)

func getNeuralNetworksMax2DPoolOperationClass() NeuralNetworksMax2DPoolOperationClass {
	_NeuralNetworksMax2DPoolOperationClassOnce.Do(func() {
		_NeuralNetworksMax2DPoolOperationClass = NeuralNetworksMax2DPoolOperationClass{class: objc.GetClass("NeuralNetworks.Max2DPoolOperation")}
	})
	return _NeuralNetworksMax2DPoolOperationClass
}

// GetNeuralNetworksMax2DPoolOperationClass returns the class object for NeuralNetworks.Max2DPoolOperation.
func GetNeuralNetworksMax2DPoolOperationClass() NeuralNetworksMax2DPoolOperationClass {
	return getNeuralNetworksMax2DPoolOperationClass()
}

type NeuralNetworksMax2DPoolOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMax2DPoolOperationClass) Alloc() NeuralNetworksMax2DPoolOperation {
	rv := objc.Send[NeuralNetworksMax2DPoolOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Max2DPoolOperation
type NeuralNetworksMax2DPoolOperation struct {
	NeuralNetworksBase2DPoolOperation
}

// NeuralNetworksMax2DPoolOperationFromID constructs a [NeuralNetworksMax2DPoolOperation] from an objc.ID.
func NeuralNetworksMax2DPoolOperationFromID(id objc.ID) NeuralNetworksMax2DPoolOperation {
	return NeuralNetworksMax2DPoolOperation{NeuralNetworksBase2DPoolOperation: NeuralNetworksBase2DPoolOperationFromID(id)}
}
// Ensure NeuralNetworksMax2DPoolOperation implements INeuralNetworksMax2DPoolOperation.
var _ INeuralNetworksMax2DPoolOperation = NeuralNetworksMax2DPoolOperation{}





// An interface definition for the [NeuralNetworksMax2DPoolOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.Max2DPoolOperation
type INeuralNetworksMax2DPoolOperation interface {
	INeuralNetworksBase2DPoolOperation
}





// Init initializes the instance.
func (n NeuralNetworksMax2DPoolOperation) Init() NeuralNetworksMax2DPoolOperation {
	rv := objc.Send[NeuralNetworksMax2DPoolOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMax2DPoolOperation) Autorelease() NeuralNetworksMax2DPoolOperation {
	rv := objc.Send[NeuralNetworksMax2DPoolOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMax2DPoolOperation creates a new NeuralNetworksMax2DPoolOperation instance.
func NewNeuralNetworksMax2DPoolOperation() NeuralNetworksMax2DPoolOperation {
	class := getNeuralNetworksMax2DPoolOperationClass()
	rv := objc.Send[NeuralNetworksMax2DPoolOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































