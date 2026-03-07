// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NeuralNetworksBaseForwardNormalizeOperation] class.
var (
	_NeuralNetworksBaseForwardNormalizeOperationClass     NeuralNetworksBaseForwardNormalizeOperationClass
	_NeuralNetworksBaseForwardNormalizeOperationClassOnce sync.Once
)

func getNeuralNetworksBaseForwardNormalizeOperationClass() NeuralNetworksBaseForwardNormalizeOperationClass {
	_NeuralNetworksBaseForwardNormalizeOperationClassOnce.Do(func() {
		_NeuralNetworksBaseForwardNormalizeOperationClass = NeuralNetworksBaseForwardNormalizeOperationClass{class: objc.GetClass("NeuralNetworks.BaseForwardNormalizeOperation")}
	})
	return _NeuralNetworksBaseForwardNormalizeOperationClass
}

// GetNeuralNetworksBaseForwardNormalizeOperationClass returns the class object for NeuralNetworks.BaseForwardNormalizeOperation.
func GetNeuralNetworksBaseForwardNormalizeOperationClass() NeuralNetworksBaseForwardNormalizeOperationClass {
	return getNeuralNetworksBaseForwardNormalizeOperationClass()
}

type NeuralNetworksBaseForwardNormalizeOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBaseForwardNormalizeOperationClass) Alloc() NeuralNetworksBaseForwardNormalizeOperation {
	rv := objc.Send[NeuralNetworksBaseForwardNormalizeOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseForwardNormalizeOperation
type NeuralNetworksBaseForwardNormalizeOperation struct {
	NeuralNetworksBaseNormalizeOperation
}

// NeuralNetworksBaseForwardNormalizeOperationFromID constructs a [NeuralNetworksBaseForwardNormalizeOperation] from an objc.ID.
func NeuralNetworksBaseForwardNormalizeOperationFromID(id objc.ID) NeuralNetworksBaseForwardNormalizeOperation {
	return NeuralNetworksBaseForwardNormalizeOperation{NeuralNetworksBaseNormalizeOperation: NeuralNetworksBaseNormalizeOperationFromID(id)}
}
// Ensure NeuralNetworksBaseForwardNormalizeOperation implements INeuralNetworksBaseForwardNormalizeOperation.
var _ INeuralNetworksBaseForwardNormalizeOperation = NeuralNetworksBaseForwardNormalizeOperation{}





// An interface definition for the [NeuralNetworksBaseForwardNormalizeOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BaseForwardNormalizeOperation
type INeuralNetworksBaseForwardNormalizeOperation interface {
	INeuralNetworksBaseNormalizeOperation
}





// Init initializes the instance.
func (n NeuralNetworksBaseForwardNormalizeOperation) Init() NeuralNetworksBaseForwardNormalizeOperation {
	rv := objc.Send[NeuralNetworksBaseForwardNormalizeOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBaseForwardNormalizeOperation) Autorelease() NeuralNetworksBaseForwardNormalizeOperation {
	rv := objc.Send[NeuralNetworksBaseForwardNormalizeOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBaseForwardNormalizeOperation creates a new NeuralNetworksBaseForwardNormalizeOperation instance.
func NewNeuralNetworksBaseForwardNormalizeOperation() NeuralNetworksBaseForwardNormalizeOperation {
	class := getNeuralNetworksBaseForwardNormalizeOperationClass()
	rv := objc.Send[NeuralNetworksBaseForwardNormalizeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































