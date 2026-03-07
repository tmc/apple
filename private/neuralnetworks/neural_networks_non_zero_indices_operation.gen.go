// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksNonZeroIndicesOperation] class.
var (
	_NeuralNetworksNonZeroIndicesOperationClass     NeuralNetworksNonZeroIndicesOperationClass
	_NeuralNetworksNonZeroIndicesOperationClassOnce sync.Once
)

func getNeuralNetworksNonZeroIndicesOperationClass() NeuralNetworksNonZeroIndicesOperationClass {
	_NeuralNetworksNonZeroIndicesOperationClassOnce.Do(func() {
		_NeuralNetworksNonZeroIndicesOperationClass = NeuralNetworksNonZeroIndicesOperationClass{class: objc.GetClass("NeuralNetworks.NonZeroIndicesOperation")}
	})
	return _NeuralNetworksNonZeroIndicesOperationClass
}

// GetNeuralNetworksNonZeroIndicesOperationClass returns the class object for NeuralNetworks.NonZeroIndicesOperation.
func GetNeuralNetworksNonZeroIndicesOperationClass() NeuralNetworksNonZeroIndicesOperationClass {
	return getNeuralNetworksNonZeroIndicesOperationClass()
}

type NeuralNetworksNonZeroIndicesOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksNonZeroIndicesOperationClass) Alloc() NeuralNetworksNonZeroIndicesOperation {
	rv := objc.Send[NeuralNetworksNonZeroIndicesOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.NonZeroIndicesOperation
type NeuralNetworksNonZeroIndicesOperation struct {
	objectivec.Object
}

// NeuralNetworksNonZeroIndicesOperationFromID constructs a [NeuralNetworksNonZeroIndicesOperation] from an objc.ID.
func NeuralNetworksNonZeroIndicesOperationFromID(id objc.ID) NeuralNetworksNonZeroIndicesOperation {
	return NeuralNetworksNonZeroIndicesOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksNonZeroIndicesOperation implements INeuralNetworksNonZeroIndicesOperation.
var _ INeuralNetworksNonZeroIndicesOperation = NeuralNetworksNonZeroIndicesOperation{}





// An interface definition for the [NeuralNetworksNonZeroIndicesOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.NonZeroIndicesOperation
type INeuralNetworksNonZeroIndicesOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksNonZeroIndicesOperation) Init() NeuralNetworksNonZeroIndicesOperation {
	rv := objc.Send[NeuralNetworksNonZeroIndicesOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksNonZeroIndicesOperation) Autorelease() NeuralNetworksNonZeroIndicesOperation {
	rv := objc.Send[NeuralNetworksNonZeroIndicesOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksNonZeroIndicesOperation creates a new NeuralNetworksNonZeroIndicesOperation instance.
func NewNeuralNetworksNonZeroIndicesOperation() NeuralNetworksNonZeroIndicesOperation {
	class := getNeuralNetworksNonZeroIndicesOperationClass()
	rv := objc.Send[NeuralNetworksNonZeroIndicesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































