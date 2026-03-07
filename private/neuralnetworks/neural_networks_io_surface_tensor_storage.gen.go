// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksIOSurfaceTensorStorage] class.
var (
	_NeuralNetworksIOSurfaceTensorStorageClass     NeuralNetworksIOSurfaceTensorStorageClass
	_NeuralNetworksIOSurfaceTensorStorageClassOnce sync.Once
)

func getNeuralNetworksIOSurfaceTensorStorageClass() NeuralNetworksIOSurfaceTensorStorageClass {
	_NeuralNetworksIOSurfaceTensorStorageClassOnce.Do(func() {
		_NeuralNetworksIOSurfaceTensorStorageClass = NeuralNetworksIOSurfaceTensorStorageClass{class: objc.GetClass("NeuralNetworks.IOSurfaceTensorStorage")}
	})
	return _NeuralNetworksIOSurfaceTensorStorageClass
}

// GetNeuralNetworksIOSurfaceTensorStorageClass returns the class object for NeuralNetworks.IOSurfaceTensorStorage.
func GetNeuralNetworksIOSurfaceTensorStorageClass() NeuralNetworksIOSurfaceTensorStorageClass {
	return getNeuralNetworksIOSurfaceTensorStorageClass()
}

type NeuralNetworksIOSurfaceTensorStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksIOSurfaceTensorStorageClass) Alloc() NeuralNetworksIOSurfaceTensorStorage {
	rv := objc.Send[NeuralNetworksIOSurfaceTensorStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.IOSurfaceTensorStorage
type NeuralNetworksIOSurfaceTensorStorage struct {
	objectivec.Object
}

// NeuralNetworksIOSurfaceTensorStorageFromID constructs a [NeuralNetworksIOSurfaceTensorStorage] from an objc.ID.
func NeuralNetworksIOSurfaceTensorStorageFromID(id objc.ID) NeuralNetworksIOSurfaceTensorStorage {
	return NeuralNetworksIOSurfaceTensorStorage{objectivec.Object{id}}
}
// Ensure NeuralNetworksIOSurfaceTensorStorage implements INeuralNetworksIOSurfaceTensorStorage.
var _ INeuralNetworksIOSurfaceTensorStorage = NeuralNetworksIOSurfaceTensorStorage{}





// An interface definition for the [NeuralNetworksIOSurfaceTensorStorage] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.IOSurfaceTensorStorage
type INeuralNetworksIOSurfaceTensorStorage interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksIOSurfaceTensorStorage) Init() NeuralNetworksIOSurfaceTensorStorage {
	rv := objc.Send[NeuralNetworksIOSurfaceTensorStorage](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksIOSurfaceTensorStorage) Autorelease() NeuralNetworksIOSurfaceTensorStorage {
	rv := objc.Send[NeuralNetworksIOSurfaceTensorStorage](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksIOSurfaceTensorStorage creates a new NeuralNetworksIOSurfaceTensorStorage instance.
func NewNeuralNetworksIOSurfaceTensorStorage() NeuralNetworksIOSurfaceTensorStorage {
	class := getNeuralNetworksIOSurfaceTensorStorageClass()
	rv := objc.Send[NeuralNetworksIOSurfaceTensorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































