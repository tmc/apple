// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksEspressoBuffer] class.
var (
	_NeuralNetworksEspressoBufferClass     NeuralNetworksEspressoBufferClass
	_NeuralNetworksEspressoBufferClassOnce sync.Once
)

func getNeuralNetworksEspressoBufferClass() NeuralNetworksEspressoBufferClass {
	_NeuralNetworksEspressoBufferClassOnce.Do(func() {
		_NeuralNetworksEspressoBufferClass = NeuralNetworksEspressoBufferClass{class: objc.GetClass("NeuralNetworks.EspressoBuffer")}
	})
	return _NeuralNetworksEspressoBufferClass
}

// GetNeuralNetworksEspressoBufferClass returns the class object for NeuralNetworks.EspressoBuffer.
func GetNeuralNetworksEspressoBufferClass() NeuralNetworksEspressoBufferClass {
	return getNeuralNetworksEspressoBufferClass()
}

type NeuralNetworksEspressoBufferClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksEspressoBufferClass) Alloc() NeuralNetworksEspressoBuffer {
	rv := objc.Send[NeuralNetworksEspressoBuffer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.EspressoBuffer
type NeuralNetworksEspressoBuffer struct {
	objectivec.Object
}

// NeuralNetworksEspressoBufferFromID constructs a [NeuralNetworksEspressoBuffer] from an objc.ID.
func NeuralNetworksEspressoBufferFromID(id objc.ID) NeuralNetworksEspressoBuffer {
	return NeuralNetworksEspressoBuffer{objectivec.Object{id}}
}
// Ensure NeuralNetworksEspressoBuffer implements INeuralNetworksEspressoBuffer.
var _ INeuralNetworksEspressoBuffer = NeuralNetworksEspressoBuffer{}





// An interface definition for the [NeuralNetworksEspressoBuffer] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.EspressoBuffer
type INeuralNetworksEspressoBuffer interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksEspressoBuffer) Init() NeuralNetworksEspressoBuffer {
	rv := objc.Send[NeuralNetworksEspressoBuffer](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksEspressoBuffer) Autorelease() NeuralNetworksEspressoBuffer {
	rv := objc.Send[NeuralNetworksEspressoBuffer](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksEspressoBuffer creates a new NeuralNetworksEspressoBuffer instance.
func NewNeuralNetworksEspressoBuffer() NeuralNetworksEspressoBuffer {
	class := getNeuralNetworksEspressoBufferClass()
	rv := objc.Send[NeuralNetworksEspressoBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































