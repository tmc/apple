// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMILEspressoBackend] class.
var (
	_NeuralNetworksMILEspressoBackendClass     NeuralNetworksMILEspressoBackendClass
	_NeuralNetworksMILEspressoBackendClassOnce sync.Once
)

func getNeuralNetworksMILEspressoBackendClass() NeuralNetworksMILEspressoBackendClass {
	_NeuralNetworksMILEspressoBackendClassOnce.Do(func() {
		_NeuralNetworksMILEspressoBackendClass = NeuralNetworksMILEspressoBackendClass{class: objc.GetClass("NeuralNetworks.MILEspressoBackend")}
	})
	return _NeuralNetworksMILEspressoBackendClass
}

// GetNeuralNetworksMILEspressoBackendClass returns the class object for NeuralNetworks.MILEspressoBackend.
func GetNeuralNetworksMILEspressoBackendClass() NeuralNetworksMILEspressoBackendClass {
	return getNeuralNetworksMILEspressoBackendClass()
}

type NeuralNetworksMILEspressoBackendClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMILEspressoBackendClass) Alloc() NeuralNetworksMILEspressoBackend {
	rv := objc.Send[NeuralNetworksMILEspressoBackend](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILEspressoBackend
type NeuralNetworksMILEspressoBackend struct {
	objectivec.Object
}

// NeuralNetworksMILEspressoBackendFromID constructs a [NeuralNetworksMILEspressoBackend] from an objc.ID.
func NeuralNetworksMILEspressoBackendFromID(id objc.ID) NeuralNetworksMILEspressoBackend {
	return NeuralNetworksMILEspressoBackend{objectivec.Object{id}}
}
// Ensure NeuralNetworksMILEspressoBackend implements INeuralNetworksMILEspressoBackend.
var _ INeuralNetworksMILEspressoBackend = NeuralNetworksMILEspressoBackend{}





// An interface definition for the [NeuralNetworksMILEspressoBackend] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILEspressoBackend
type INeuralNetworksMILEspressoBackend interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMILEspressoBackend) Init() NeuralNetworksMILEspressoBackend {
	rv := objc.Send[NeuralNetworksMILEspressoBackend](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMILEspressoBackend) Autorelease() NeuralNetworksMILEspressoBackend {
	rv := objc.Send[NeuralNetworksMILEspressoBackend](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMILEspressoBackend creates a new NeuralNetworksMILEspressoBackend instance.
func NewNeuralNetworksMILEspressoBackend() NeuralNetworksMILEspressoBackend {
	class := getNeuralNetworksMILEspressoBackendClass()
	rv := objc.Send[NeuralNetworksMILEspressoBackend](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































