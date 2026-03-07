// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksConcatOperation] class.
var (
	_NeuralNetworksConcatOperationClass     NeuralNetworksConcatOperationClass
	_NeuralNetworksConcatOperationClassOnce sync.Once
)

func getNeuralNetworksConcatOperationClass() NeuralNetworksConcatOperationClass {
	_NeuralNetworksConcatOperationClassOnce.Do(func() {
		_NeuralNetworksConcatOperationClass = NeuralNetworksConcatOperationClass{class: objc.GetClass("NeuralNetworks.ConcatOperation")}
	})
	return _NeuralNetworksConcatOperationClass
}

// GetNeuralNetworksConcatOperationClass returns the class object for NeuralNetworks.ConcatOperation.
func GetNeuralNetworksConcatOperationClass() NeuralNetworksConcatOperationClass {
	return getNeuralNetworksConcatOperationClass()
}

type NeuralNetworksConcatOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksConcatOperationClass) Alloc() NeuralNetworksConcatOperation {
	rv := objc.Send[NeuralNetworksConcatOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConcatOperation
type NeuralNetworksConcatOperation struct {
	objectivec.Object
}

// NeuralNetworksConcatOperationFromID constructs a [NeuralNetworksConcatOperation] from an objc.ID.
func NeuralNetworksConcatOperationFromID(id objc.ID) NeuralNetworksConcatOperation {
	return NeuralNetworksConcatOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksConcatOperation implements INeuralNetworksConcatOperation.
var _ INeuralNetworksConcatOperation = NeuralNetworksConcatOperation{}





// An interface definition for the [NeuralNetworksConcatOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ConcatOperation
type INeuralNetworksConcatOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksConcatOperation) Init() NeuralNetworksConcatOperation {
	rv := objc.Send[NeuralNetworksConcatOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksConcatOperation) Autorelease() NeuralNetworksConcatOperation {
	rv := objc.Send[NeuralNetworksConcatOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksConcatOperation creates a new NeuralNetworksConcatOperation instance.
func NewNeuralNetworksConcatOperation() NeuralNetworksConcatOperation {
	class := getNeuralNetworksConcatOperationClass()
	rv := objc.Send[NeuralNetworksConcatOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































