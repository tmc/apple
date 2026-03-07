// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksCastOperation] class.
var (
	_NeuralNetworksCastOperationClass     NeuralNetworksCastOperationClass
	_NeuralNetworksCastOperationClassOnce sync.Once
)

func getNeuralNetworksCastOperationClass() NeuralNetworksCastOperationClass {
	_NeuralNetworksCastOperationClassOnce.Do(func() {
		_NeuralNetworksCastOperationClass = NeuralNetworksCastOperationClass{class: objc.GetClass("NeuralNetworks.CastOperation")}
	})
	return _NeuralNetworksCastOperationClass
}

// GetNeuralNetworksCastOperationClass returns the class object for NeuralNetworks.CastOperation.
func GetNeuralNetworksCastOperationClass() NeuralNetworksCastOperationClass {
	return getNeuralNetworksCastOperationClass()
}

type NeuralNetworksCastOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksCastOperationClass) Alloc() NeuralNetworksCastOperation {
	rv := objc.Send[NeuralNetworksCastOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.CastOperation
type NeuralNetworksCastOperation struct {
	objectivec.Object
}

// NeuralNetworksCastOperationFromID constructs a [NeuralNetworksCastOperation] from an objc.ID.
func NeuralNetworksCastOperationFromID(id objc.ID) NeuralNetworksCastOperation {
	return NeuralNetworksCastOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksCastOperation implements INeuralNetworksCastOperation.
var _ INeuralNetworksCastOperation = NeuralNetworksCastOperation{}





// An interface definition for the [NeuralNetworksCastOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.CastOperation
type INeuralNetworksCastOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksCastOperation) Init() NeuralNetworksCastOperation {
	rv := objc.Send[NeuralNetworksCastOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksCastOperation) Autorelease() NeuralNetworksCastOperation {
	rv := objc.Send[NeuralNetworksCastOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksCastOperation creates a new NeuralNetworksCastOperation instance.
func NewNeuralNetworksCastOperation() NeuralNetworksCastOperation {
	class := getNeuralNetworksCastOperationClass()
	rv := objc.Send[NeuralNetworksCastOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































