// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksLSTMLayerOperationBackwardContext] class.
var (
	_NeuralNetworksLSTMLayerOperationBackwardContextClass     NeuralNetworksLSTMLayerOperationBackwardContextClass
	_NeuralNetworksLSTMLayerOperationBackwardContextClassOnce sync.Once
)

func getNeuralNetworksLSTMLayerOperationBackwardContextClass() NeuralNetworksLSTMLayerOperationBackwardContextClass {
	_NeuralNetworksLSTMLayerOperationBackwardContextClassOnce.Do(func() {
		_NeuralNetworksLSTMLayerOperationBackwardContextClass = NeuralNetworksLSTMLayerOperationBackwardContextClass{class: objc.GetClass("NeuralNetworks.LSTMLayerOperationBackwardContext")}
	})
	return _NeuralNetworksLSTMLayerOperationBackwardContextClass
}

// GetNeuralNetworksLSTMLayerOperationBackwardContextClass returns the class object for NeuralNetworks.LSTMLayerOperationBackwardContext.
func GetNeuralNetworksLSTMLayerOperationBackwardContextClass() NeuralNetworksLSTMLayerOperationBackwardContextClass {
	return getNeuralNetworksLSTMLayerOperationBackwardContextClass()
}

type NeuralNetworksLSTMLayerOperationBackwardContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLSTMLayerOperationBackwardContextClass) Alloc() NeuralNetworksLSTMLayerOperationBackwardContext {
	rv := objc.Send[NeuralNetworksLSTMLayerOperationBackwardContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LSTMLayerOperationBackwardContext
type NeuralNetworksLSTMLayerOperationBackwardContext struct {
	objectivec.Object
}

// NeuralNetworksLSTMLayerOperationBackwardContextFromID constructs a [NeuralNetworksLSTMLayerOperationBackwardContext] from an objc.ID.
func NeuralNetworksLSTMLayerOperationBackwardContextFromID(id objc.ID) NeuralNetworksLSTMLayerOperationBackwardContext {
	return NeuralNetworksLSTMLayerOperationBackwardContext{objectivec.Object{id}}
}
// Ensure NeuralNetworksLSTMLayerOperationBackwardContext implements INeuralNetworksLSTMLayerOperationBackwardContext.
var _ INeuralNetworksLSTMLayerOperationBackwardContext = NeuralNetworksLSTMLayerOperationBackwardContext{}





// An interface definition for the [NeuralNetworksLSTMLayerOperationBackwardContext] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LSTMLayerOperationBackwardContext
type INeuralNetworksLSTMLayerOperationBackwardContext interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksLSTMLayerOperationBackwardContext) Init() NeuralNetworksLSTMLayerOperationBackwardContext {
	rv := objc.Send[NeuralNetworksLSTMLayerOperationBackwardContext](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLSTMLayerOperationBackwardContext) Autorelease() NeuralNetworksLSTMLayerOperationBackwardContext {
	rv := objc.Send[NeuralNetworksLSTMLayerOperationBackwardContext](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLSTMLayerOperationBackwardContext creates a new NeuralNetworksLSTMLayerOperationBackwardContext instance.
func NewNeuralNetworksLSTMLayerOperationBackwardContext() NeuralNetworksLSTMLayerOperationBackwardContext {
	class := getNeuralNetworksLSTMLayerOperationBackwardContextClass()
	rv := objc.Send[NeuralNetworksLSTMLayerOperationBackwardContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































