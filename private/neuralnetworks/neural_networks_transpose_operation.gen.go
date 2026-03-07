// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksTransposeOperation] class.
var (
	_NeuralNetworksTransposeOperationClass     NeuralNetworksTransposeOperationClass
	_NeuralNetworksTransposeOperationClassOnce sync.Once
)

func getNeuralNetworksTransposeOperationClass() NeuralNetworksTransposeOperationClass {
	_NeuralNetworksTransposeOperationClassOnce.Do(func() {
		_NeuralNetworksTransposeOperationClass = NeuralNetworksTransposeOperationClass{class: objc.GetClass("NeuralNetworks.TransposeOperation")}
	})
	return _NeuralNetworksTransposeOperationClass
}

// GetNeuralNetworksTransposeOperationClass returns the class object for NeuralNetworks.TransposeOperation.
func GetNeuralNetworksTransposeOperationClass() NeuralNetworksTransposeOperationClass {
	return getNeuralNetworksTransposeOperationClass()
}

type NeuralNetworksTransposeOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTransposeOperationClass) Alloc() NeuralNetworksTransposeOperation {
	rv := objc.Send[NeuralNetworksTransposeOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TransposeOperation
type NeuralNetworksTransposeOperation struct {
	objectivec.Object
}

// NeuralNetworksTransposeOperationFromID constructs a [NeuralNetworksTransposeOperation] from an objc.ID.
func NeuralNetworksTransposeOperationFromID(id objc.ID) NeuralNetworksTransposeOperation {
	return NeuralNetworksTransposeOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksTransposeOperation implements INeuralNetworksTransposeOperation.
var _ INeuralNetworksTransposeOperation = NeuralNetworksTransposeOperation{}





// An interface definition for the [NeuralNetworksTransposeOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TransposeOperation
type INeuralNetworksTransposeOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksTransposeOperation) Init() NeuralNetworksTransposeOperation {
	rv := objc.Send[NeuralNetworksTransposeOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTransposeOperation) Autorelease() NeuralNetworksTransposeOperation {
	rv := objc.Send[NeuralNetworksTransposeOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTransposeOperation creates a new NeuralNetworksTransposeOperation instance.
func NewNeuralNetworksTransposeOperation() NeuralNetworksTransposeOperation {
	class := getNeuralNetworksTransposeOperationClass()
	rv := objc.Send[NeuralNetworksTransposeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































