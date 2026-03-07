// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksViewingTensorStorage] class.
var (
	_NeuralNetworksViewingTensorStorageClass     NeuralNetworksViewingTensorStorageClass
	_NeuralNetworksViewingTensorStorageClassOnce sync.Once
)

func getNeuralNetworksViewingTensorStorageClass() NeuralNetworksViewingTensorStorageClass {
	_NeuralNetworksViewingTensorStorageClassOnce.Do(func() {
		_NeuralNetworksViewingTensorStorageClass = NeuralNetworksViewingTensorStorageClass{class: objc.GetClass("NeuralNetworks.ViewingTensorStorage")}
	})
	return _NeuralNetworksViewingTensorStorageClass
}

// GetNeuralNetworksViewingTensorStorageClass returns the class object for NeuralNetworks.ViewingTensorStorage.
func GetNeuralNetworksViewingTensorStorageClass() NeuralNetworksViewingTensorStorageClass {
	return getNeuralNetworksViewingTensorStorageClass()
}

type NeuralNetworksViewingTensorStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksViewingTensorStorageClass) Alloc() NeuralNetworksViewingTensorStorage {
	rv := objc.Send[NeuralNetworksViewingTensorStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ViewingTensorStorage
type NeuralNetworksViewingTensorStorage struct {
	objectivec.Object
}

// NeuralNetworksViewingTensorStorageFromID constructs a [NeuralNetworksViewingTensorStorage] from an objc.ID.
func NeuralNetworksViewingTensorStorageFromID(id objc.ID) NeuralNetworksViewingTensorStorage {
	return NeuralNetworksViewingTensorStorage{objectivec.Object{id}}
}
// Ensure NeuralNetworksViewingTensorStorage implements INeuralNetworksViewingTensorStorage.
var _ INeuralNetworksViewingTensorStorage = NeuralNetworksViewingTensorStorage{}





// An interface definition for the [NeuralNetworksViewingTensorStorage] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ViewingTensorStorage
type INeuralNetworksViewingTensorStorage interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksViewingTensorStorage) Init() NeuralNetworksViewingTensorStorage {
	rv := objc.Send[NeuralNetworksViewingTensorStorage](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksViewingTensorStorage) Autorelease() NeuralNetworksViewingTensorStorage {
	rv := objc.Send[NeuralNetworksViewingTensorStorage](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksViewingTensorStorage creates a new NeuralNetworksViewingTensorStorage instance.
func NewNeuralNetworksViewingTensorStorage() NeuralNetworksViewingTensorStorage {
	class := getNeuralNetworksViewingTensorStorageClass()
	rv := objc.Send[NeuralNetworksViewingTensorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































