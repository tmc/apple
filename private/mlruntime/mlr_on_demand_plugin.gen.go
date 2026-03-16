// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLROnDemandPlugin] class.
var (
	_MLROnDemandPluginClass     MLROnDemandPluginClass
	_MLROnDemandPluginClassOnce sync.Once
)

func getMLROnDemandPluginClass() MLROnDemandPluginClass {
	_MLROnDemandPluginClassOnce.Do(func() {
		_MLROnDemandPluginClass = MLROnDemandPluginClass{class: objc.GetClass("MLROnDemandPlugin")}
	})
	return _MLROnDemandPluginClass
}

// GetMLROnDemandPluginClass returns the class object for MLROnDemandPlugin.
func GetMLROnDemandPluginClass() MLROnDemandPluginClass {
	return getMLROnDemandPluginClass()
}

type MLROnDemandPluginClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLROnDemandPluginClass) Alloc() MLROnDemandPlugin {
	rv := objc.Send[MLROnDemandPlugin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandPlugin
type MLROnDemandPlugin struct {
	objectivec.Object
}

// MLROnDemandPluginFromID constructs a [MLROnDemandPlugin] from an objc.ID.
func MLROnDemandPluginFromID(id objc.ID) MLROnDemandPlugin {
	return MLROnDemandPlugin{objectivec.Object{id}}
}
// Ensure MLROnDemandPlugin implements IMLROnDemandPlugin.
var _ IMLROnDemandPlugin = MLROnDemandPlugin{}

// An interface definition for the [MLROnDemandPlugin] class.
//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandPlugin
type IMLROnDemandPlugin interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r MLROnDemandPlugin) Init() MLROnDemandPlugin {
	rv := objc.Send[MLROnDemandPlugin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLROnDemandPlugin) Autorelease() MLROnDemandPlugin {
	rv := objc.Send[MLROnDemandPlugin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLROnDemandPlugin creates a new MLROnDemandPlugin instance.
func NewMLROnDemandPlugin() MLROnDemandPlugin {
	class := getMLROnDemandPluginClass()
	rv := objc.Send[MLROnDemandPlugin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandPlugin/onDemandPluginIDs
func (_MLROnDemandPluginClass MLROnDemandPluginClass) OnDemandPluginIDs() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLROnDemandPluginClass.class), objc.Sel("onDemandPluginIDs"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandPlugin/performTask:forPluginID:completionHandler:
func (_MLROnDemandPluginClass MLROnDemandPluginClass) PerformTaskForPluginIDCompletionHandler(task objectivec.IObject, id objectivec.IObject, handler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(handler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_MLROnDemandPluginClass.class), objc.Sel("performTask:forPluginID:completionHandler:"), task, id, _block2)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandPlugin/synchronouslyPerformTask:forPluginID:error:
func (_MLROnDemandPluginClass MLROnDemandPluginClass) SynchronouslyPerformTaskForPluginIDError(task objectivec.IObject, id objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLROnDemandPluginClass.class), objc.Sel("synchronouslyPerformTask:forPluginID:error:"), task, id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

