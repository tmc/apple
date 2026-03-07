// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksSharedMTLBufferTensorStorage] class.
var (
	_NeuralNetworksSharedMTLBufferTensorStorageClass     NeuralNetworksSharedMTLBufferTensorStorageClass
	_NeuralNetworksSharedMTLBufferTensorStorageClassOnce sync.Once
)

func getNeuralNetworksSharedMTLBufferTensorStorageClass() NeuralNetworksSharedMTLBufferTensorStorageClass {
	_NeuralNetworksSharedMTLBufferTensorStorageClassOnce.Do(func() {
		_NeuralNetworksSharedMTLBufferTensorStorageClass = NeuralNetworksSharedMTLBufferTensorStorageClass{class: objc.GetClass("NeuralNetworks.SharedMTLBufferTensorStorage")}
	})
	return _NeuralNetworksSharedMTLBufferTensorStorageClass
}

// GetNeuralNetworksSharedMTLBufferTensorStorageClass returns the class object for NeuralNetworks.SharedMTLBufferTensorStorage.
func GetNeuralNetworksSharedMTLBufferTensorStorageClass() NeuralNetworksSharedMTLBufferTensorStorageClass {
	return getNeuralNetworksSharedMTLBufferTensorStorageClass()
}

type NeuralNetworksSharedMTLBufferTensorStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksSharedMTLBufferTensorStorageClass) Alloc() NeuralNetworksSharedMTLBufferTensorStorage {
	rv := objc.Send[NeuralNetworksSharedMTLBufferTensorStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SharedMTLBufferTensorStorage
type NeuralNetworksSharedMTLBufferTensorStorage struct {
	objectivec.Object
}

// NeuralNetworksSharedMTLBufferTensorStorageFromID constructs a [NeuralNetworksSharedMTLBufferTensorStorage] from an objc.ID.
func NeuralNetworksSharedMTLBufferTensorStorageFromID(id objc.ID) NeuralNetworksSharedMTLBufferTensorStorage {
	return NeuralNetworksSharedMTLBufferTensorStorage{objectivec.Object{id}}
}
// Ensure NeuralNetworksSharedMTLBufferTensorStorage implements INeuralNetworksSharedMTLBufferTensorStorage.
var _ INeuralNetworksSharedMTLBufferTensorStorage = NeuralNetworksSharedMTLBufferTensorStorage{}





// An interface definition for the [NeuralNetworksSharedMTLBufferTensorStorage] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.SharedMTLBufferTensorStorage
type INeuralNetworksSharedMTLBufferTensorStorage interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksSharedMTLBufferTensorStorage) Init() NeuralNetworksSharedMTLBufferTensorStorage {
	rv := objc.Send[NeuralNetworksSharedMTLBufferTensorStorage](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksSharedMTLBufferTensorStorage) Autorelease() NeuralNetworksSharedMTLBufferTensorStorage {
	rv := objc.Send[NeuralNetworksSharedMTLBufferTensorStorage](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksSharedMTLBufferTensorStorage creates a new NeuralNetworksSharedMTLBufferTensorStorage instance.
func NewNeuralNetworksSharedMTLBufferTensorStorage() NeuralNetworksSharedMTLBufferTensorStorage {
	class := getNeuralNetworksSharedMTLBufferTensorStorageClass()
	rv := objc.Send[NeuralNetworksSharedMTLBufferTensorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































