// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksCumulativeOperation] class.
var (
	_NeuralNetworksCumulativeOperationClass     NeuralNetworksCumulativeOperationClass
	_NeuralNetworksCumulativeOperationClassOnce sync.Once
)

func getNeuralNetworksCumulativeOperationClass() NeuralNetworksCumulativeOperationClass {
	_NeuralNetworksCumulativeOperationClassOnce.Do(func() {
		_NeuralNetworksCumulativeOperationClass = NeuralNetworksCumulativeOperationClass{class: objc.GetClass("NeuralNetworks.CumulativeOperation")}
	})
	return _NeuralNetworksCumulativeOperationClass
}

// GetNeuralNetworksCumulativeOperationClass returns the class object for NeuralNetworks.CumulativeOperation.
func GetNeuralNetworksCumulativeOperationClass() NeuralNetworksCumulativeOperationClass {
	return getNeuralNetworksCumulativeOperationClass()
}

type NeuralNetworksCumulativeOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksCumulativeOperationClass) Alloc() NeuralNetworksCumulativeOperation {
	rv := objc.Send[NeuralNetworksCumulativeOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.CumulativeOperation
type NeuralNetworksCumulativeOperation struct {
	objectivec.Object
}

// NeuralNetworksCumulativeOperationFromID constructs a [NeuralNetworksCumulativeOperation] from an objc.ID.
func NeuralNetworksCumulativeOperationFromID(id objc.ID) NeuralNetworksCumulativeOperation {
	return NeuralNetworksCumulativeOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksCumulativeOperation implements INeuralNetworksCumulativeOperation.
var _ INeuralNetworksCumulativeOperation = NeuralNetworksCumulativeOperation{}





// An interface definition for the [NeuralNetworksCumulativeOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.CumulativeOperation
type INeuralNetworksCumulativeOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksCumulativeOperation) Init() NeuralNetworksCumulativeOperation {
	rv := objc.Send[NeuralNetworksCumulativeOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksCumulativeOperation) Autorelease() NeuralNetworksCumulativeOperation {
	rv := objc.Send[NeuralNetworksCumulativeOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksCumulativeOperation creates a new NeuralNetworksCumulativeOperation instance.
func NewNeuralNetworksCumulativeOperation() NeuralNetworksCumulativeOperation {
	class := getNeuralNetworksCumulativeOperationClass()
	rv := objc.Send[NeuralNetworksCumulativeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































