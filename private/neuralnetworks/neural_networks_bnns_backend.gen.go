// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksBNNSBackend] class.
var (
	_NeuralNetworksBNNSBackendClass     NeuralNetworksBNNSBackendClass
	_NeuralNetworksBNNSBackendClassOnce sync.Once
)

func getNeuralNetworksBNNSBackendClass() NeuralNetworksBNNSBackendClass {
	_NeuralNetworksBNNSBackendClassOnce.Do(func() {
		_NeuralNetworksBNNSBackendClass = NeuralNetworksBNNSBackendClass{class: objc.GetClass("NeuralNetworks.BNNSBackend")}
	})
	return _NeuralNetworksBNNSBackendClass
}

// GetNeuralNetworksBNNSBackendClass returns the class object for NeuralNetworks.BNNSBackend.
func GetNeuralNetworksBNNSBackendClass() NeuralNetworksBNNSBackendClass {
	return getNeuralNetworksBNNSBackendClass()
}

type NeuralNetworksBNNSBackendClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksBNNSBackendClass) Alloc() NeuralNetworksBNNSBackend {
	rv := objc.Send[NeuralNetworksBNNSBackend](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BNNSBackend
type NeuralNetworksBNNSBackend struct {
	objectivec.Object
}

// NeuralNetworksBNNSBackendFromID constructs a [NeuralNetworksBNNSBackend] from an objc.ID.
func NeuralNetworksBNNSBackendFromID(id objc.ID) NeuralNetworksBNNSBackend {
	return NeuralNetworksBNNSBackend{objectivec.Object{id}}
}
// Ensure NeuralNetworksBNNSBackend implements INeuralNetworksBNNSBackend.
var _ INeuralNetworksBNNSBackend = NeuralNetworksBNNSBackend{}





// An interface definition for the [NeuralNetworksBNNSBackend] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.BNNSBackend
type INeuralNetworksBNNSBackend interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksBNNSBackend) Init() NeuralNetworksBNNSBackend {
	rv := objc.Send[NeuralNetworksBNNSBackend](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksBNNSBackend) Autorelease() NeuralNetworksBNNSBackend {
	rv := objc.Send[NeuralNetworksBNNSBackend](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksBNNSBackend creates a new NeuralNetworksBNNSBackend instance.
func NewNeuralNetworksBNNSBackend() NeuralNetworksBNNSBackend {
	class := getNeuralNetworksBNNSBackendClass()
	rv := objc.Send[NeuralNetworksBNNSBackend](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































