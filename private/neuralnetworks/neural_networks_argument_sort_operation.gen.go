// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksArgumentSortOperation] class.
var (
	_NeuralNetworksArgumentSortOperationClass     NeuralNetworksArgumentSortOperationClass
	_NeuralNetworksArgumentSortOperationClassOnce sync.Once
)

func getNeuralNetworksArgumentSortOperationClass() NeuralNetworksArgumentSortOperationClass {
	_NeuralNetworksArgumentSortOperationClassOnce.Do(func() {
		_NeuralNetworksArgumentSortOperationClass = NeuralNetworksArgumentSortOperationClass{class: objc.GetClass("NeuralNetworks.ArgumentSortOperation")}
	})
	return _NeuralNetworksArgumentSortOperationClass
}

// GetNeuralNetworksArgumentSortOperationClass returns the class object for NeuralNetworks.ArgumentSortOperation.
func GetNeuralNetworksArgumentSortOperationClass() NeuralNetworksArgumentSortOperationClass {
	return getNeuralNetworksArgumentSortOperationClass()
}

type NeuralNetworksArgumentSortOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksArgumentSortOperationClass) Alloc() NeuralNetworksArgumentSortOperation {
	rv := objc.Send[NeuralNetworksArgumentSortOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ArgumentSortOperation
type NeuralNetworksArgumentSortOperation struct {
	objectivec.Object
}

// NeuralNetworksArgumentSortOperationFromID constructs a [NeuralNetworksArgumentSortOperation] from an objc.ID.
func NeuralNetworksArgumentSortOperationFromID(id objc.ID) NeuralNetworksArgumentSortOperation {
	return NeuralNetworksArgumentSortOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksArgumentSortOperation implements INeuralNetworksArgumentSortOperation.
var _ INeuralNetworksArgumentSortOperation = NeuralNetworksArgumentSortOperation{}





// An interface definition for the [NeuralNetworksArgumentSortOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ArgumentSortOperation
type INeuralNetworksArgumentSortOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksArgumentSortOperation) Init() NeuralNetworksArgumentSortOperation {
	rv := objc.Send[NeuralNetworksArgumentSortOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksArgumentSortOperation) Autorelease() NeuralNetworksArgumentSortOperation {
	rv := objc.Send[NeuralNetworksArgumentSortOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksArgumentSortOperation creates a new NeuralNetworksArgumentSortOperation instance.
func NewNeuralNetworksArgumentSortOperation() NeuralNetworksArgumentSortOperation {
	class := getNeuralNetworksArgumentSortOperationClass()
	rv := objc.Send[NeuralNetworksArgumentSortOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































