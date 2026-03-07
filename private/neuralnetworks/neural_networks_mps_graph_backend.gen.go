// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMPSGraphBackend] class.
var (
	_NeuralNetworksMPSGraphBackendClass     NeuralNetworksMPSGraphBackendClass
	_NeuralNetworksMPSGraphBackendClassOnce sync.Once
)

func getNeuralNetworksMPSGraphBackendClass() NeuralNetworksMPSGraphBackendClass {
	_NeuralNetworksMPSGraphBackendClassOnce.Do(func() {
		_NeuralNetworksMPSGraphBackendClass = NeuralNetworksMPSGraphBackendClass{class: objc.GetClass("NeuralNetworks.MPSGraphBackend")}
	})
	return _NeuralNetworksMPSGraphBackendClass
}

// GetNeuralNetworksMPSGraphBackendClass returns the class object for NeuralNetworks.MPSGraphBackend.
func GetNeuralNetworksMPSGraphBackendClass() NeuralNetworksMPSGraphBackendClass {
	return getNeuralNetworksMPSGraphBackendClass()
}

type NeuralNetworksMPSGraphBackendClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMPSGraphBackendClass) Alloc() NeuralNetworksMPSGraphBackend {
	rv := objc.Send[NeuralNetworksMPSGraphBackend](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSGraphBackend
type NeuralNetworksMPSGraphBackend struct {
	objectivec.Object
}

// NeuralNetworksMPSGraphBackendFromID constructs a [NeuralNetworksMPSGraphBackend] from an objc.ID.
func NeuralNetworksMPSGraphBackendFromID(id objc.ID) NeuralNetworksMPSGraphBackend {
	return NeuralNetworksMPSGraphBackend{objectivec.Object{id}}
}
// Ensure NeuralNetworksMPSGraphBackend implements INeuralNetworksMPSGraphBackend.
var _ INeuralNetworksMPSGraphBackend = NeuralNetworksMPSGraphBackend{}





// An interface definition for the [NeuralNetworksMPSGraphBackend] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MPSGraphBackend
type INeuralNetworksMPSGraphBackend interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMPSGraphBackend) Init() NeuralNetworksMPSGraphBackend {
	rv := objc.Send[NeuralNetworksMPSGraphBackend](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMPSGraphBackend) Autorelease() NeuralNetworksMPSGraphBackend {
	rv := objc.Send[NeuralNetworksMPSGraphBackend](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMPSGraphBackend creates a new NeuralNetworksMPSGraphBackend instance.
func NewNeuralNetworksMPSGraphBackend() NeuralNetworksMPSGraphBackend {
	class := getNeuralNetworksMPSGraphBackendClass()
	rv := objc.Send[NeuralNetworksMPSGraphBackend](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































