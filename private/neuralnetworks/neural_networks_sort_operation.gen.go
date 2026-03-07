// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksSortOperation] class.
var (
	_NeuralNetworksSortOperationClass     NeuralNetworksSortOperationClass
	_NeuralNetworksSortOperationClassOnce sync.Once
)

func getNeuralNetworksSortOperationClass() NeuralNetworksSortOperationClass {
	_NeuralNetworksSortOperationClassOnce.Do(func() {
		_NeuralNetworksSortOperationClass = NeuralNetworksSortOperationClass{class: objc.GetClass("NeuralNetworks.SortOperation")}
	})
	return _NeuralNetworksSortOperationClass
}

// GetNeuralNetworksSortOperationClass returns the class object for NeuralNetworks.SortOperation.
func GetNeuralNetworksSortOperationClass() NeuralNetworksSortOperationClass {
	return getNeuralNetworksSortOperationClass()
}

type NeuralNetworksSortOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSortOperationClass) Alloc() NeuralNetworksSortOperation {
	rv := objc.Send[NeuralNetworksSortOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SortOperation
type NeuralNetworksSortOperation struct {
	objectivec.Object
}

// NeuralNetworksSortOperationFromID constructs a [NeuralNetworksSortOperation] from an objc.ID.
func NeuralNetworksSortOperationFromID(id objc.ID) NeuralNetworksSortOperation {
	return NeuralNetworksSortOperation{objectivec.Object{id}}
}
// Ensure NeuralNetworksSortOperation implements INeuralNetworksSortOperation.
var _ INeuralNetworksSortOperation = NeuralNetworksSortOperation{}





// An interface definition for the [NeuralNetworksSortOperation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SortOperation
type INeuralNetworksSortOperation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksSortOperation) Init() NeuralNetworksSortOperation {
	rv := objc.Send[NeuralNetworksSortOperation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSortOperation) Autorelease() NeuralNetworksSortOperation {
	rv := objc.Send[NeuralNetworksSortOperation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSortOperation creates a new NeuralNetworksSortOperation instance.
func NewNeuralNetworksSortOperation() NeuralNetworksSortOperation {
	class := getNeuralNetworksSortOperationClass()
	rv := objc.Send[NeuralNetworksSortOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































