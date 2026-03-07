// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksReverseOperation] class.
var (
	_NeuralNetworksReverseOperationClass     NeuralNetworksReverseOperationClass
	_NeuralNetworksReverseOperationClassOnce sync.Once
)

func getNeuralNetworksReverseOperationClass() NeuralNetworksReverseOperationClass {
	_NeuralNetworksReverseOperationClassOnce.Do(func() {
		_NeuralNetworksReverseOperationClass = NeuralNetworksReverseOperationClass{class: objc.GetClass("NeuralNetworks.ReverseOperation")}
	})
	return _NeuralNetworksReverseOperationClass
}

// GetNeuralNetworksReverseOperationClass returns the class object for NeuralNetworks.ReverseOperation.
func GetNeuralNetworksReverseOperationClass() NeuralNetworksReverseOperationClass {
	return getNeuralNetworksReverseOperationClass()
}

type NeuralNetworksReverseOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksReverseOperationClass) Alloc() NeuralNetworksReverseOperation {
	rv := objc.Send[NeuralNetworksReverseOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReverseOperation
type NeuralNetworksReverseOperation struct {
	objectivec.Object
}

// NeuralNetworksReverseOperationFromID constructs a [NeuralNetworksReverseOperation] from an objc.ID.
func NeuralNetworksReverseOperationFromID(id objc.ID) NeuralNetworksReverseOperation {
	return NeuralNetworksReverseOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksReverseOperation implements INeuralNetworksReverseOperation.
var _ INeuralNetworksReverseOperation = NeuralNetworksReverseOperation{}





// An interface definition for the [NeuralNetworksReverseOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ReverseOperation
type INeuralNetworksReverseOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksReverseOperation) Init() NeuralNetworksReverseOperation {
	rv := objc.Send[NeuralNetworksReverseOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksReverseOperation) Autorelease() NeuralNetworksReverseOperation {
	rv := objc.Send[NeuralNetworksReverseOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksReverseOperation creates a new NeuralNetworksReverseOperation instance.
func NewNeuralNetworksReverseOperation() NeuralNetworksReverseOperation {
	class := getNeuralNetworksReverseOperationClass()
	rv := objc.Send[NeuralNetworksReverseOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































