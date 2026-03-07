// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksNativeTensorStorage] class.
var (
	_NeuralNetworksNativeTensorStorageClass     NeuralNetworksNativeTensorStorageClass
	_NeuralNetworksNativeTensorStorageClassOnce sync.Once
)

func getNeuralNetworksNativeTensorStorageClass() NeuralNetworksNativeTensorStorageClass {
	_NeuralNetworksNativeTensorStorageClassOnce.Do(func() {
		_NeuralNetworksNativeTensorStorageClass = NeuralNetworksNativeTensorStorageClass{class: objc.GetClass("NeuralNetworks.NativeTensorStorage")}
	})
	return _NeuralNetworksNativeTensorStorageClass
}

// GetNeuralNetworksNativeTensorStorageClass returns the class object for NeuralNetworks.NativeTensorStorage.
func GetNeuralNetworksNativeTensorStorageClass() NeuralNetworksNativeTensorStorageClass {
	return getNeuralNetworksNativeTensorStorageClass()
}

type NeuralNetworksNativeTensorStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksNativeTensorStorageClass) Alloc() NeuralNetworksNativeTensorStorage {
	rv := objc.Send[NeuralNetworksNativeTensorStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.NativeTensorStorage
type NeuralNetworksNativeTensorStorage struct {
	objectivec.Object
}

// NeuralNetworksNativeTensorStorageFromID constructs a [NeuralNetworksNativeTensorStorage] from an objc.ID.
func NeuralNetworksNativeTensorStorageFromID(id objc.ID) NeuralNetworksNativeTensorStorage {
	return NeuralNetworksNativeTensorStorage{objectivec.Object{id}}
}
// Ensure NeuralNetworksNativeTensorStorage implements INeuralNetworksNativeTensorStorage.
var _ INeuralNetworksNativeTensorStorage = NeuralNetworksNativeTensorStorage{}





// An interface definition for the [NeuralNetworksNativeTensorStorage] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.NativeTensorStorage
type INeuralNetworksNativeTensorStorage interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksNativeTensorStorage) Init() NeuralNetworksNativeTensorStorage {
	rv := objc.Send[NeuralNetworksNativeTensorStorage](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksNativeTensorStorage) Autorelease() NeuralNetworksNativeTensorStorage {
	rv := objc.Send[NeuralNetworksNativeTensorStorage](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksNativeTensorStorage creates a new NeuralNetworksNativeTensorStorage instance.
func NewNeuralNetworksNativeTensorStorage() NeuralNetworksNativeTensorStorage {
	class := getNeuralNetworksNativeTensorStorageClass()
	rv := objc.Send[NeuralNetworksNativeTensorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































