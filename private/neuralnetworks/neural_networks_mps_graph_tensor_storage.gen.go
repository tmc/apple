// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMPSGraphTensorStorage] class.
var (
	_NeuralNetworksMPSGraphTensorStorageClass     NeuralNetworksMPSGraphTensorStorageClass
	_NeuralNetworksMPSGraphTensorStorageClassOnce sync.Once
)

func getNeuralNetworksMPSGraphTensorStorageClass() NeuralNetworksMPSGraphTensorStorageClass {
	_NeuralNetworksMPSGraphTensorStorageClassOnce.Do(func() {
		_NeuralNetworksMPSGraphTensorStorageClass = NeuralNetworksMPSGraphTensorStorageClass{class: objc.GetClass("NeuralNetworks.MPSGraphTensorStorage")}
	})
	return _NeuralNetworksMPSGraphTensorStorageClass
}

// GetNeuralNetworksMPSGraphTensorStorageClass returns the class object for NeuralNetworks.MPSGraphTensorStorage.
func GetNeuralNetworksMPSGraphTensorStorageClass() NeuralNetworksMPSGraphTensorStorageClass {
	return getNeuralNetworksMPSGraphTensorStorageClass()
}

type NeuralNetworksMPSGraphTensorStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMPSGraphTensorStorageClass) Alloc() NeuralNetworksMPSGraphTensorStorage {
	rv := objc.Send[NeuralNetworksMPSGraphTensorStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSGraphTensorStorage
type NeuralNetworksMPSGraphTensorStorage struct {
	objectivec.Object
}

// NeuralNetworksMPSGraphTensorStorageFromID constructs a [NeuralNetworksMPSGraphTensorStorage] from an objc.ID.
func NeuralNetworksMPSGraphTensorStorageFromID(id objc.ID) NeuralNetworksMPSGraphTensorStorage {
	return NeuralNetworksMPSGraphTensorStorage{objectivec.Object{id}}
}
// Ensure NeuralNetworksMPSGraphTensorStorage implements INeuralNetworksMPSGraphTensorStorage.
var _ INeuralNetworksMPSGraphTensorStorage = NeuralNetworksMPSGraphTensorStorage{}





// An interface definition for the [NeuralNetworksMPSGraphTensorStorage] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSGraphTensorStorage
type INeuralNetworksMPSGraphTensorStorage interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMPSGraphTensorStorage) Init() NeuralNetworksMPSGraphTensorStorage {
	rv := objc.Send[NeuralNetworksMPSGraphTensorStorage](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMPSGraphTensorStorage) Autorelease() NeuralNetworksMPSGraphTensorStorage {
	rv := objc.Send[NeuralNetworksMPSGraphTensorStorage](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMPSGraphTensorStorage creates a new NeuralNetworksMPSGraphTensorStorage instance.
func NewNeuralNetworksMPSGraphTensorStorage() NeuralNetworksMPSGraphTensorStorage {
	class := getNeuralNetworksMPSGraphTensorStorageClass()
	rv := objc.Send[NeuralNetworksMPSGraphTensorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































