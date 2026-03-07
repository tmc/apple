// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksExecutionContext] class.
var (
	_NeuralNetworksExecutionContextClass     NeuralNetworksExecutionContextClass
	_NeuralNetworksExecutionContextClassOnce sync.Once
)

func getNeuralNetworksExecutionContextClass() NeuralNetworksExecutionContextClass {
	_NeuralNetworksExecutionContextClassOnce.Do(func() {
		_NeuralNetworksExecutionContextClass = NeuralNetworksExecutionContextClass{class: objc.GetClass("NeuralNetworks.ExecutionContext")}
	})
	return _NeuralNetworksExecutionContextClass
}

// GetNeuralNetworksExecutionContextClass returns the class object for NeuralNetworks.ExecutionContext.
func GetNeuralNetworksExecutionContextClass() NeuralNetworksExecutionContextClass {
	return getNeuralNetworksExecutionContextClass()
}

type NeuralNetworksExecutionContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksExecutionContextClass) Alloc() NeuralNetworksExecutionContext {
	rv := objc.Send[NeuralNetworksExecutionContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ExecutionContext
type NeuralNetworksExecutionContext struct {
	objectivec.Object
}

// NeuralNetworksExecutionContextFromID constructs a [NeuralNetworksExecutionContext] from an objc.ID.
func NeuralNetworksExecutionContextFromID(id objc.ID) NeuralNetworksExecutionContext {
	return NeuralNetworksExecutionContext{objectivec.Object{id}}
}
// Ensure NeuralNetworksExecutionContext implements INeuralNetworksExecutionContext.
var _ INeuralNetworksExecutionContext = NeuralNetworksExecutionContext{}





// An interface definition for the [NeuralNetworksExecutionContext] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ExecutionContext
type INeuralNetworksExecutionContext interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksExecutionContext) Init() NeuralNetworksExecutionContext {
	rv := objc.Send[NeuralNetworksExecutionContext](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksExecutionContext) Autorelease() NeuralNetworksExecutionContext {
	rv := objc.Send[NeuralNetworksExecutionContext](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksExecutionContext creates a new NeuralNetworksExecutionContext instance.
func NewNeuralNetworksExecutionContext() NeuralNetworksExecutionContext {
	class := getNeuralNetworksExecutionContextClass()
	rv := objc.Send[NeuralNetworksExecutionContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































