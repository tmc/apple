// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksParameterInitializer] class.
var (
	_NeuralNetworksParameterInitializerClass     NeuralNetworksParameterInitializerClass
	_NeuralNetworksParameterInitializerClassOnce sync.Once
)

func getNeuralNetworksParameterInitializerClass() NeuralNetworksParameterInitializerClass {
	_NeuralNetworksParameterInitializerClassOnce.Do(func() {
		_NeuralNetworksParameterInitializerClass = NeuralNetworksParameterInitializerClass{class: objc.GetClass("NeuralNetworks.ParameterInitializer")}
	})
	return _NeuralNetworksParameterInitializerClass
}

// GetNeuralNetworksParameterInitializerClass returns the class object for NeuralNetworks.ParameterInitializer.
func GetNeuralNetworksParameterInitializerClass() NeuralNetworksParameterInitializerClass {
	return getNeuralNetworksParameterInitializerClass()
}

type NeuralNetworksParameterInitializerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksParameterInitializerClass) Alloc() NeuralNetworksParameterInitializer {
	rv := objc.Send[NeuralNetworksParameterInitializer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ParameterInitializer
type NeuralNetworksParameterInitializer struct {
	objectivec.Object
}

// NeuralNetworksParameterInitializerFromID constructs a [NeuralNetworksParameterInitializer] from an objc.ID.
func NeuralNetworksParameterInitializerFromID(id objc.ID) NeuralNetworksParameterInitializer {
	return NeuralNetworksParameterInitializer{objectivec.Object{id}}
}
// Ensure NeuralNetworksParameterInitializer implements INeuralNetworksParameterInitializer.
var _ INeuralNetworksParameterInitializer = NeuralNetworksParameterInitializer{}





// An interface definition for the [NeuralNetworksParameterInitializer] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ParameterInitializer
type INeuralNetworksParameterInitializer interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksParameterInitializer) Init() NeuralNetworksParameterInitializer {
	rv := objc.Send[NeuralNetworksParameterInitializer](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksParameterInitializer) Autorelease() NeuralNetworksParameterInitializer {
	rv := objc.Send[NeuralNetworksParameterInitializer](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksParameterInitializer creates a new NeuralNetworksParameterInitializer instance.
func NewNeuralNetworksParameterInitializer() NeuralNetworksParameterInitializer {
	class := getNeuralNetworksParameterInitializerClass()
	rv := objc.Send[NeuralNetworksParameterInitializer](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































