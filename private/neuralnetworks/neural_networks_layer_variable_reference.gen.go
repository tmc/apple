// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksLayerVariableReference] class.
var (
	_NeuralNetworksLayerVariableReferenceClass     NeuralNetworksLayerVariableReferenceClass
	_NeuralNetworksLayerVariableReferenceClassOnce sync.Once
)

func getNeuralNetworksLayerVariableReferenceClass() NeuralNetworksLayerVariableReferenceClass {
	_NeuralNetworksLayerVariableReferenceClassOnce.Do(func() {
		_NeuralNetworksLayerVariableReferenceClass = NeuralNetworksLayerVariableReferenceClass{class: objc.GetClass("NeuralNetworks.LayerVariableReference")}
	})
	return _NeuralNetworksLayerVariableReferenceClass
}

// GetNeuralNetworksLayerVariableReferenceClass returns the class object for NeuralNetworks.LayerVariableReference.
func GetNeuralNetworksLayerVariableReferenceClass() NeuralNetworksLayerVariableReferenceClass {
	return getNeuralNetworksLayerVariableReferenceClass()
}

type NeuralNetworksLayerVariableReferenceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLayerVariableReferenceClass) Alloc() NeuralNetworksLayerVariableReference {
	rv := objc.Send[NeuralNetworksLayerVariableReference](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LayerVariableReference
type NeuralNetworksLayerVariableReference struct {
	objectivec.Object
}

// NeuralNetworksLayerVariableReferenceFromID constructs a [NeuralNetworksLayerVariableReference] from an objc.ID.
func NeuralNetworksLayerVariableReferenceFromID(id objc.ID) NeuralNetworksLayerVariableReference {
	return NeuralNetworksLayerVariableReference{objectivec.Object{id}}
}
// Ensure NeuralNetworksLayerVariableReference implements INeuralNetworksLayerVariableReference.
var _ INeuralNetworksLayerVariableReference = NeuralNetworksLayerVariableReference{}





// An interface definition for the [NeuralNetworksLayerVariableReference] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LayerVariableReference
type INeuralNetworksLayerVariableReference interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksLayerVariableReference) Init() NeuralNetworksLayerVariableReference {
	rv := objc.Send[NeuralNetworksLayerVariableReference](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLayerVariableReference) Autorelease() NeuralNetworksLayerVariableReference {
	rv := objc.Send[NeuralNetworksLayerVariableReference](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLayerVariableReference creates a new NeuralNetworksLayerVariableReference instance.
func NewNeuralNetworksLayerVariableReference() NeuralNetworksLayerVariableReference {
	class := getNeuralNetworksLayerVariableReferenceClass()
	rv := objc.Send[NeuralNetworksLayerVariableReference](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































