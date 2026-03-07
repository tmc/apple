// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRExtensionRemoteContext] class.
var (
	_MLRExtensionRemoteContextClass     MLRExtensionRemoteContextClass
	_MLRExtensionRemoteContextClassOnce sync.Once
)

func getMLRExtensionRemoteContextClass() MLRExtensionRemoteContextClass {
	_MLRExtensionRemoteContextClassOnce.Do(func() {
		_MLRExtensionRemoteContextClass = MLRExtensionRemoteContextClass{class: objc.GetClass("MLRExtensionRemoteContext")}
	})
	return _MLRExtensionRemoteContextClass
}

// GetMLRExtensionRemoteContextClass returns the class object for MLRExtensionRemoteContext.
func GetMLRExtensionRemoteContextClass() MLRExtensionRemoteContextClass {
	return getMLRExtensionRemoteContextClass()
}

type MLRExtensionRemoteContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRExtensionRemoteContextClass) Alloc() MLRExtensionRemoteContext {
	rv := objc.Send[MLRExtensionRemoteContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MLRExtensionRemoteContext.InitWithInputItemsListenerEndpointContextUUID]
//   - [MLRExtensionRemoteContext.InitWithInputItemsListenerEndpointContextUUIDPlugin]
//   - [MLRExtensionRemoteContext.DebugDescription]
//   - [MLRExtensionRemoteContext.Description]
//   - [MLRExtensionRemoteContext.Hash]
//   - [MLRExtensionRemoteContext.Superclass]
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext
type MLRExtensionRemoteContext struct {
	foundation.NSExtensionContext
}

// MLRExtensionRemoteContextFromID constructs a [MLRExtensionRemoteContext] from an objc.ID.
func MLRExtensionRemoteContextFromID(id objc.ID) MLRExtensionRemoteContext {
	return MLRExtensionRemoteContext{NSExtensionContext: foundation.NSExtensionContextFromID(id)}
}
// Ensure MLRExtensionRemoteContext implements IMLRExtensionRemoteContext.
var _ IMLRExtensionRemoteContext = MLRExtensionRemoteContext{}





// An interface definition for the [MLRExtensionRemoteContext] class.
//
// # Methods
//
//   - [IMLRExtensionRemoteContext.InitWithInputItemsListenerEndpointContextUUID]
//   - [IMLRExtensionRemoteContext.InitWithInputItemsListenerEndpointContextUUIDPlugin]
//   - [IMLRExtensionRemoteContext.DebugDescription]
//   - [IMLRExtensionRemoteContext.Description]
//   - [IMLRExtensionRemoteContext.Hash]
//   - [IMLRExtensionRemoteContext.Superclass]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext
type IMLRExtensionRemoteContext interface {
	foundation.INSExtensionContext

	// Topic: Methods

	InitWithInputItemsListenerEndpointContextUUID(items objectivec.IObject, endpoint objectivec.IObject, uuid objectivec.IObject) MLRExtensionRemoteContext
	InitWithInputItemsListenerEndpointContextUUIDPlugin(items objectivec.IObject, endpoint objectivec.IObject, uuid objectivec.IObject, plugin objectivec.IObject) MLRExtensionRemoteContext
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}





// Init initializes the instance.
func (r MLRExtensionRemoteContext) Init() MLRExtensionRemoteContext {
	rv := objc.Send[MLRExtensionRemoteContext](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRExtensionRemoteContext) Autorelease() MLRExtensionRemoteContext {
	rv := objc.Send[MLRExtensionRemoteContext](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRExtensionRemoteContext creates a new MLRExtensionRemoteContext instance.
func NewMLRExtensionRemoteContext() MLRExtensionRemoteContext {
	class := getMLRExtensionRemoteContextClass()
	rv := objc.Send[MLRExtensionRemoteContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/initWithInputItems:listenerEndpoint:contextUUID:
func NewRExtensionRemoteContextWithInputItemsListenerEndpointContextUUID(items objectivec.IObject, endpoint objectivec.IObject, uuid objectivec.IObject) MLRExtensionRemoteContext {
	instance := getMLRExtensionRemoteContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputItems:listenerEndpoint:contextUUID:"), items, endpoint, uuid)
	return MLRExtensionRemoteContextFromID(rv)
}


//
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/initWithInputItems:listenerEndpoint:contextUUID:plugin:
func NewRExtensionRemoteContextWithInputItemsListenerEndpointContextUUIDPlugin(items objectivec.IObject, endpoint objectivec.IObject, uuid objectivec.IObject, plugin objectivec.IObject) MLRExtensionRemoteContext {
	instance := getMLRExtensionRemoteContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputItems:listenerEndpoint:contextUUID:plugin:"), items, endpoint, uuid, plugin)
	return MLRExtensionRemoteContextFromID(rv)
}







//
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/initWithInputItems:listenerEndpoint:contextUUID:
func (r MLRExtensionRemoteContext) InitWithInputItemsListenerEndpointContextUUID(items objectivec.IObject, endpoint objectivec.IObject, uuid objectivec.IObject) MLRExtensionRemoteContext {
	rv := objc.Send[MLRExtensionRemoteContext](r.ID, objc.Sel("initWithInputItems:listenerEndpoint:contextUUID:"), items, endpoint, uuid)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/initWithInputItems:listenerEndpoint:contextUUID:plugin:
func (r MLRExtensionRemoteContext) InitWithInputItemsListenerEndpointContextUUIDPlugin(items objectivec.IObject, endpoint objectivec.IObject, uuid objectivec.IObject, plugin objectivec.IObject) MLRExtensionRemoteContext {
	rv := objc.Send[MLRExtensionRemoteContext](r.ID, objc.Sel("initWithInputItems:listenerEndpoint:contextUUID:plugin:"), items, endpoint, uuid, plugin)
	return rv
}












// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/debugDescription
func (r MLRExtensionRemoteContext) DebugDescription() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/description
func (r MLRExtensionRemoteContext) Description() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/hash
func (r MLRExtensionRemoteContext) Hash() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("hash"))
	return rv
}



// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionRemoteContext/superclass
func (r MLRExtensionRemoteContext) Superclass() objc.Class {
	rv := objc.Send[objc.Class](r.ID, objc.Sel("superclass"))
	return rv
}

















