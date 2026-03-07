// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseBinaryElementwiseOperation] class.
var (
	_NeuralNetworksBaseBinaryElementwiseOperationClass     NeuralNetworksBaseBinaryElementwiseOperationClass
	_NeuralNetworksBaseBinaryElementwiseOperationClassOnce sync.Once
)

func getNeuralNetworksBaseBinaryElementwiseOperationClass() NeuralNetworksBaseBinaryElementwiseOperationClass {
	_NeuralNetworksBaseBinaryElementwiseOperationClassOnce.Do(func() {
		_NeuralNetworksBaseBinaryElementwiseOperationClass = NeuralNetworksBaseBinaryElementwiseOperationClass{class: objc.GetClass("NeuralNetworks.BaseBinaryElementwiseOperation")}
	})
	return _NeuralNetworksBaseBinaryElementwiseOperationClass
}

// GetNeuralNetworksBaseBinaryElementwiseOperationClass returns the class object for NeuralNetworks.BaseBinaryElementwiseOperation.
func GetNeuralNetworksBaseBinaryElementwiseOperationClass() NeuralNetworksBaseBinaryElementwiseOperationClass {
	return getNeuralNetworksBaseBinaryElementwiseOperationClass()
}

type NeuralNetworksBaseBinaryElementwiseOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseBinaryElementwiseOperationClass) Alloc() NeuralNetworksBaseBinaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBaseBinaryElementwiseOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseBinaryElementwiseOperation
type NeuralNetworksBaseBinaryElementwiseOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseBinaryElementwiseOperationFromID constructs a [NeuralNetworksBaseBinaryElementwiseOperation] from an objc.ID.
func NeuralNetworksBaseBinaryElementwiseOperationFromID(id objc.ID) NeuralNetworksBaseBinaryElementwiseOperation {
	return NeuralNetworksBaseBinaryElementwiseOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseBinaryElementwiseOperation implements INeuralNetworksBaseBinaryElementwiseOperation.
var _ INeuralNetworksBaseBinaryElementwiseOperation = NeuralNetworksBaseBinaryElementwiseOperation{}





// An interface definition for the [NeuralNetworksBaseBinaryElementwiseOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseBinaryElementwiseOperation
type INeuralNetworksBaseBinaryElementwiseOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseBinaryElementwiseOperation) Init() NeuralNetworksBaseBinaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBaseBinaryElementwiseOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseBinaryElementwiseOperation) Autorelease() NeuralNetworksBaseBinaryElementwiseOperation {
	rv := objc.Send[NeuralNetworksBaseBinaryElementwiseOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseBinaryElementwiseOperation creates a new NeuralNetworksBaseBinaryElementwiseOperation instance.
func NewNeuralNetworksBaseBinaryElementwiseOperation() NeuralNetworksBaseBinaryElementwiseOperation {
	class := getNeuralNetworksBaseBinaryElementwiseOperationClass()
	rv := objc.Send[NeuralNetworksBaseBinaryElementwiseOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































