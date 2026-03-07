// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksTensorRepresentation] class.
var (
	_NeuralNetworksTensorRepresentationClass     NeuralNetworksTensorRepresentationClass
	_NeuralNetworksTensorRepresentationClassOnce sync.Once
)

func getNeuralNetworksTensorRepresentationClass() NeuralNetworksTensorRepresentationClass {
	_NeuralNetworksTensorRepresentationClassOnce.Do(func() {
		_NeuralNetworksTensorRepresentationClass = NeuralNetworksTensorRepresentationClass{class: objc.GetClass("NeuralNetworks.TensorRepresentation")}
	})
	return _NeuralNetworksTensorRepresentationClass
}

// GetNeuralNetworksTensorRepresentationClass returns the class object for NeuralNetworks.TensorRepresentation.
func GetNeuralNetworksTensorRepresentationClass() NeuralNetworksTensorRepresentationClass {
	return getNeuralNetworksTensorRepresentationClass()
}

type NeuralNetworksTensorRepresentationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTensorRepresentationClass) Alloc() NeuralNetworksTensorRepresentation {
	rv := objc.Send[NeuralNetworksTensorRepresentation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TensorRepresentation
type NeuralNetworksTensorRepresentation struct {
	objectivec.Object
}

// NeuralNetworksTensorRepresentationFromID constructs a [NeuralNetworksTensorRepresentation] from an objc.ID.
func NeuralNetworksTensorRepresentationFromID(id objc.ID) NeuralNetworksTensorRepresentation {
	return NeuralNetworksTensorRepresentation{objectivec.Object{id}}
}
// Ensure NeuralNetworksTensorRepresentation implements INeuralNetworksTensorRepresentation.
var _ INeuralNetworksTensorRepresentation = NeuralNetworksTensorRepresentation{}





// An interface definition for the [NeuralNetworksTensorRepresentation] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TensorRepresentation
type INeuralNetworksTensorRepresentation interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksTensorRepresentation) Init() NeuralNetworksTensorRepresentation {
	rv := objc.Send[NeuralNetworksTensorRepresentation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTensorRepresentation) Autorelease() NeuralNetworksTensorRepresentation {
	rv := objc.Send[NeuralNetworksTensorRepresentation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTensorRepresentation creates a new NeuralNetworksTensorRepresentation instance.
func NewNeuralNetworksTensorRepresentation() NeuralNetworksTensorRepresentation {
	class := getNeuralNetworksTensorRepresentationClass()
	rv := objc.Send[NeuralNetworksTensorRepresentation](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































