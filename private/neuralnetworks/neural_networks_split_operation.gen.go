// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksSplitOperation] class.
var (
	_NeuralNetworksSplitOperationClass     NeuralNetworksSplitOperationClass
	_NeuralNetworksSplitOperationClassOnce sync.Once
)

func getNeuralNetworksSplitOperationClass() NeuralNetworksSplitOperationClass {
	_NeuralNetworksSplitOperationClassOnce.Do(func() {
		_NeuralNetworksSplitOperationClass = NeuralNetworksSplitOperationClass{class: objc.GetClass("NeuralNetworks.SplitOperation")}
	})
	return _NeuralNetworksSplitOperationClass
}

// GetNeuralNetworksSplitOperationClass returns the class object for NeuralNetworks.SplitOperation.
func GetNeuralNetworksSplitOperationClass() NeuralNetworksSplitOperationClass {
	return getNeuralNetworksSplitOperationClass()
}

type NeuralNetworksSplitOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSplitOperationClass) Alloc() NeuralNetworksSplitOperation {
	rv := objc.Send[NeuralNetworksSplitOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SplitOperation
type NeuralNetworksSplitOperation struct {
	objectivec.Object
}

// NeuralNetworksSplitOperationFromID constructs a [NeuralNetworksSplitOperation] from an objc.ID.
func NeuralNetworksSplitOperationFromID(id objc.ID) NeuralNetworksSplitOperation {
	return NeuralNetworksSplitOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksSplitOperation implements INeuralNetworksSplitOperation.
var _ INeuralNetworksSplitOperation = NeuralNetworksSplitOperation{}





// An interface definition for the [NeuralNetworksSplitOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SplitOperation
type INeuralNetworksSplitOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksSplitOperation) Init() NeuralNetworksSplitOperation {
	rv := objc.Send[NeuralNetworksSplitOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSplitOperation) Autorelease() NeuralNetworksSplitOperation {
	rv := objc.Send[NeuralNetworksSplitOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSplitOperation creates a new NeuralNetworksSplitOperation instance.
func NewNeuralNetworksSplitOperation() NeuralNetworksSplitOperation {
	class := getNeuralNetworksSplitOperationClass()
	rv := objc.Send[NeuralNetworksSplitOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































