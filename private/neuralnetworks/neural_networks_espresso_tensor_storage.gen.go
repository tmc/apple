// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksEspressoTensorStorage] class.
var (
	_NeuralNetworksEspressoTensorStorageClass     NeuralNetworksEspressoTensorStorageClass
	_NeuralNetworksEspressoTensorStorageClassOnce sync.Once
)

func getNeuralNetworksEspressoTensorStorageClass() NeuralNetworksEspressoTensorStorageClass {
	_NeuralNetworksEspressoTensorStorageClassOnce.Do(func() {
		_NeuralNetworksEspressoTensorStorageClass = NeuralNetworksEspressoTensorStorageClass{class: objc.GetClass("NeuralNetworks.EspressoTensorStorage")}
	})
	return _NeuralNetworksEspressoTensorStorageClass
}

// GetNeuralNetworksEspressoTensorStorageClass returns the class object for NeuralNetworks.EspressoTensorStorage.
func GetNeuralNetworksEspressoTensorStorageClass() NeuralNetworksEspressoTensorStorageClass {
	return getNeuralNetworksEspressoTensorStorageClass()
}

type NeuralNetworksEspressoTensorStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksEspressoTensorStorageClass) Alloc() NeuralNetworksEspressoTensorStorage {
	rv := objc.Send[NeuralNetworksEspressoTensorStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.EspressoTensorStorage
type NeuralNetworksEspressoTensorStorage struct {
	objectivec.Object
}

// NeuralNetworksEspressoTensorStorageFromID constructs a [NeuralNetworksEspressoTensorStorage] from an objc.ID.
func NeuralNetworksEspressoTensorStorageFromID(id objc.ID) NeuralNetworksEspressoTensorStorage {
	return NeuralNetworksEspressoTensorStorage{objectivec.Object{id}}
}
// Ensure NeuralNetworksEspressoTensorStorage implements INeuralNetworksEspressoTensorStorage.
var _ INeuralNetworksEspressoTensorStorage = NeuralNetworksEspressoTensorStorage{}





// An interface definition for the [NeuralNetworksEspressoTensorStorage] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.EspressoTensorStorage
type INeuralNetworksEspressoTensorStorage interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksEspressoTensorStorage) Init() NeuralNetworksEspressoTensorStorage {
	rv := objc.Send[NeuralNetworksEspressoTensorStorage](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksEspressoTensorStorage) Autorelease() NeuralNetworksEspressoTensorStorage {
	rv := objc.Send[NeuralNetworksEspressoTensorStorage](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksEspressoTensorStorage creates a new NeuralNetworksEspressoTensorStorage instance.
func NewNeuralNetworksEspressoTensorStorage() NeuralNetworksEspressoTensorStorage {
	class := getNeuralNetworksEspressoTensorStorageClass()
	rv := objc.Send[NeuralNetworksEspressoTensorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































