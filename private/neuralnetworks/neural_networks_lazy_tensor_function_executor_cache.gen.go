// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksLazyTensorFunctionExecutorCache] class.
var (
	_NeuralNetworksLazyTensorFunctionExecutorCacheClass     NeuralNetworksLazyTensorFunctionExecutorCacheClass
	_NeuralNetworksLazyTensorFunctionExecutorCacheClassOnce sync.Once
)

func getNeuralNetworksLazyTensorFunctionExecutorCacheClass() NeuralNetworksLazyTensorFunctionExecutorCacheClass {
	_NeuralNetworksLazyTensorFunctionExecutorCacheClassOnce.Do(func() {
		_NeuralNetworksLazyTensorFunctionExecutorCacheClass = NeuralNetworksLazyTensorFunctionExecutorCacheClass{class: objc.GetClass("NeuralNetworks.LazyTensorFunctionExecutorCache")}
	})
	return _NeuralNetworksLazyTensorFunctionExecutorCacheClass
}

// GetNeuralNetworksLazyTensorFunctionExecutorCacheClass returns the class object for NeuralNetworks.LazyTensorFunctionExecutorCache.
func GetNeuralNetworksLazyTensorFunctionExecutorCacheClass() NeuralNetworksLazyTensorFunctionExecutorCacheClass {
	return getNeuralNetworksLazyTensorFunctionExecutorCacheClass()
}

type NeuralNetworksLazyTensorFunctionExecutorCacheClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLazyTensorFunctionExecutorCacheClass) Alloc() NeuralNetworksLazyTensorFunctionExecutorCache {
	rv := objc.Send[NeuralNetworksLazyTensorFunctionExecutorCache](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LazyTensorFunctionExecutorCache
type NeuralNetworksLazyTensorFunctionExecutorCache struct {
	objectivec.Object
}

// NeuralNetworksLazyTensorFunctionExecutorCacheFromID constructs a [NeuralNetworksLazyTensorFunctionExecutorCache] from an objc.ID.
func NeuralNetworksLazyTensorFunctionExecutorCacheFromID(id objc.ID) NeuralNetworksLazyTensorFunctionExecutorCache {
	return NeuralNetworksLazyTensorFunctionExecutorCache{objectivec.Object{id}}
}
// Ensure NeuralNetworksLazyTensorFunctionExecutorCache implements INeuralNetworksLazyTensorFunctionExecutorCache.
var _ INeuralNetworksLazyTensorFunctionExecutorCache = NeuralNetworksLazyTensorFunctionExecutorCache{}





// An interface definition for the [NeuralNetworksLazyTensorFunctionExecutorCache] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LazyTensorFunctionExecutorCache
type INeuralNetworksLazyTensorFunctionExecutorCache interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksLazyTensorFunctionExecutorCache) Init() NeuralNetworksLazyTensorFunctionExecutorCache {
	rv := objc.Send[NeuralNetworksLazyTensorFunctionExecutorCache](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLazyTensorFunctionExecutorCache) Autorelease() NeuralNetworksLazyTensorFunctionExecutorCache {
	rv := objc.Send[NeuralNetworksLazyTensorFunctionExecutorCache](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLazyTensorFunctionExecutorCache creates a new NeuralNetworksLazyTensorFunctionExecutorCache instance.
func NewNeuralNetworksLazyTensorFunctionExecutorCache() NeuralNetworksLazyTensorFunctionExecutorCache {
	class := getNeuralNetworksLazyTensorFunctionExecutorCacheClass()
	rv := objc.Send[NeuralNetworksLazyTensorFunctionExecutorCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































