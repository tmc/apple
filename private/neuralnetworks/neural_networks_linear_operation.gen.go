// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksLinearOperation] class.
var (
	_NeuralNetworksLinearOperationClass     NeuralNetworksLinearOperationClass
	_NeuralNetworksLinearOperationClassOnce sync.Once
)

func getNeuralNetworksLinearOperationClass() NeuralNetworksLinearOperationClass {
	_NeuralNetworksLinearOperationClassOnce.Do(func() {
		_NeuralNetworksLinearOperationClass = NeuralNetworksLinearOperationClass{class: objc.GetClass("NeuralNetworks.LinearOperation")}
	})
	return _NeuralNetworksLinearOperationClass
}

// GetNeuralNetworksLinearOperationClass returns the class object for NeuralNetworks.LinearOperation.
func GetNeuralNetworksLinearOperationClass() NeuralNetworksLinearOperationClass {
	return getNeuralNetworksLinearOperationClass()
}

type NeuralNetworksLinearOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLinearOperationClass) Alloc() NeuralNetworksLinearOperation {
	rv := objc.Send[NeuralNetworksLinearOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LinearOperation
type NeuralNetworksLinearOperation struct {
	objectivec.Object
}

// NeuralNetworksLinearOperationFromID constructs a [NeuralNetworksLinearOperation] from an objc.ID.
func NeuralNetworksLinearOperationFromID(id objc.ID) NeuralNetworksLinearOperation {
	return NeuralNetworksLinearOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksLinearOperation implements INeuralNetworksLinearOperation.
var _ INeuralNetworksLinearOperation = NeuralNetworksLinearOperation{}





// An interface definition for the [NeuralNetworksLinearOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LinearOperation
type INeuralNetworksLinearOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksLinearOperation) Init() NeuralNetworksLinearOperation {
	rv := objc.Send[NeuralNetworksLinearOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLinearOperation) Autorelease() NeuralNetworksLinearOperation {
	rv := objc.Send[NeuralNetworksLinearOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLinearOperation creates a new NeuralNetworksLinearOperation instance.
func NewNeuralNetworksLinearOperation() NeuralNetworksLinearOperation {
	class := getNeuralNetworksLinearOperationClass()
	rv := objc.Send[NeuralNetworksLinearOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































