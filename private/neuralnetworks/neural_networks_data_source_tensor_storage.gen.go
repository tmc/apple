// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksDataSourceTensorStorage] class.
var (
	_NeuralNetworksDataSourceTensorStorageClass     NeuralNetworksDataSourceTensorStorageClass
	_NeuralNetworksDataSourceTensorStorageClassOnce sync.Once
)

func getNeuralNetworksDataSourceTensorStorageClass() NeuralNetworksDataSourceTensorStorageClass {
	_NeuralNetworksDataSourceTensorStorageClassOnce.Do(func() {
		_NeuralNetworksDataSourceTensorStorageClass = NeuralNetworksDataSourceTensorStorageClass{class: objc.GetClass("NeuralNetworks.DataSourceTensorStorage")}
	})
	return _NeuralNetworksDataSourceTensorStorageClass
}

// GetNeuralNetworksDataSourceTensorStorageClass returns the class object for NeuralNetworks.DataSourceTensorStorage.
func GetNeuralNetworksDataSourceTensorStorageClass() NeuralNetworksDataSourceTensorStorageClass {
	return getNeuralNetworksDataSourceTensorStorageClass()
}

type NeuralNetworksDataSourceTensorStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksDataSourceTensorStorageClass) Alloc() NeuralNetworksDataSourceTensorStorage {
	rv := objc.Send[NeuralNetworksDataSourceTensorStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.DataSourceTensorStorage
type NeuralNetworksDataSourceTensorStorage struct {
	objectivec.Object
}

// NeuralNetworksDataSourceTensorStorageFromID constructs a [NeuralNetworksDataSourceTensorStorage] from an objc.ID.
func NeuralNetworksDataSourceTensorStorageFromID(id objc.ID) NeuralNetworksDataSourceTensorStorage {
	return NeuralNetworksDataSourceTensorStorage{objectivec.Object{id}}
}
// Ensure NeuralNetworksDataSourceTensorStorage implements INeuralNetworksDataSourceTensorStorage.
var _ INeuralNetworksDataSourceTensorStorage = NeuralNetworksDataSourceTensorStorage{}





// An interface definition for the [NeuralNetworksDataSourceTensorStorage] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.DataSourceTensorStorage
type INeuralNetworksDataSourceTensorStorage interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksDataSourceTensorStorage) Init() NeuralNetworksDataSourceTensorStorage {
	rv := objc.Send[NeuralNetworksDataSourceTensorStorage](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksDataSourceTensorStorage) Autorelease() NeuralNetworksDataSourceTensorStorage {
	rv := objc.Send[NeuralNetworksDataSourceTensorStorage](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksDataSourceTensorStorage creates a new NeuralNetworksDataSourceTensorStorage instance.
func NewNeuralNetworksDataSourceTensorStorage() NeuralNetworksDataSourceTensorStorage {
	class := getNeuralNetworksDataSourceTensorStorageClass()
	rv := objc.Send[NeuralNetworksDataSourceTensorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































