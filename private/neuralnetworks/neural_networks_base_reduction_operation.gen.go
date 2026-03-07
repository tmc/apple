// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBaseReductionOperation] class.
var (
	_NeuralNetworksBaseReductionOperationClass     NeuralNetworksBaseReductionOperationClass
	_NeuralNetworksBaseReductionOperationClassOnce sync.Once
)

func getNeuralNetworksBaseReductionOperationClass() NeuralNetworksBaseReductionOperationClass {
	_NeuralNetworksBaseReductionOperationClassOnce.Do(func() {
		_NeuralNetworksBaseReductionOperationClass = NeuralNetworksBaseReductionOperationClass{class: objc.GetClass("NeuralNetworks.BaseReductionOperation")}
	})
	return _NeuralNetworksBaseReductionOperationClass
}

// GetNeuralNetworksBaseReductionOperationClass returns the class object for NeuralNetworks.BaseReductionOperation.
func GetNeuralNetworksBaseReductionOperationClass() NeuralNetworksBaseReductionOperationClass {
	return getNeuralNetworksBaseReductionOperationClass()
}

type NeuralNetworksBaseReductionOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseReductionOperationClass) Alloc() NeuralNetworksBaseReductionOperation {
	rv := objc.Send[NeuralNetworksBaseReductionOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseReductionOperation
type NeuralNetworksBaseReductionOperation struct {
	objectivec.Object
}

// NeuralNetworksBaseReductionOperationFromID constructs a [NeuralNetworksBaseReductionOperation] from an objc.ID.
func NeuralNetworksBaseReductionOperationFromID(id objc.ID) NeuralNetworksBaseReductionOperation {
	return NeuralNetworksBaseReductionOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksBaseReductionOperation implements INeuralNetworksBaseReductionOperation.
var _ INeuralNetworksBaseReductionOperation = NeuralNetworksBaseReductionOperation{}





// An interface definition for the [NeuralNetworksBaseReductionOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseReductionOperation
type INeuralNetworksBaseReductionOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBaseReductionOperation) Init() NeuralNetworksBaseReductionOperation {
	rv := objc.Send[NeuralNetworksBaseReductionOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseReductionOperation) Autorelease() NeuralNetworksBaseReductionOperation {
	rv := objc.Send[NeuralNetworksBaseReductionOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseReductionOperation creates a new NeuralNetworksBaseReductionOperation instance.
func NewNeuralNetworksBaseReductionOperation() NeuralNetworksBaseReductionOperation {
	class := getNeuralNetworksBaseReductionOperationClass()
	rv := objc.Send[NeuralNetworksBaseReductionOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































