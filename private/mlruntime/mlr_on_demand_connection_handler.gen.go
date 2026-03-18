// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLROnDemandConnectionHandler] class.
var (
	_MLROnDemandConnectionHandlerClass     MLROnDemandConnectionHandlerClass
	_MLROnDemandConnectionHandlerClassOnce sync.Once
)

func getMLROnDemandConnectionHandlerClass() MLROnDemandConnectionHandlerClass {
	_MLROnDemandConnectionHandlerClassOnce.Do(func() {
		_MLROnDemandConnectionHandlerClass = MLROnDemandConnectionHandlerClass{class: objc.GetClass("MLROnDemandConnectionHandler")}
	})
	return _MLROnDemandConnectionHandlerClass
}

// GetMLROnDemandConnectionHandlerClass returns the class object for MLROnDemandConnectionHandler.
func GetMLROnDemandConnectionHandlerClass() MLROnDemandConnectionHandlerClass {
	return getMLROnDemandConnectionHandlerClass()
}

type MLROnDemandConnectionHandlerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLROnDemandConnectionHandlerClass) Alloc() MLROnDemandConnectionHandler {
	rv := objc.Send[MLROnDemandConnectionHandler](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLROnDemandConnectionHandler.PerformTaskCompletionHandler]
//   - [MLROnDemandConnectionHandler.PerformTaskInternalCompletionHandler]
//   - [MLROnDemandConnectionHandler.PluginPrincipal]
//   - [MLROnDemandConnectionHandler.SetPluginPrincipal]
//   - [MLROnDemandConnectionHandler.PrincipalObject]
//   - [MLROnDemandConnectionHandler.ShouldAcceptXPCConnection]
//   - [MLROnDemandConnectionHandler.XpcConnection]
//   - [MLROnDemandConnectionHandler.SetXpcConnection]
//   - [MLROnDemandConnectionHandler.InitWithPrincipalObject]
//   - [MLROnDemandConnectionHandler.DebugDescription]
//   - [MLROnDemandConnectionHandler.Description]
//   - [MLROnDemandConnectionHandler.Hash]
//   - [MLROnDemandConnectionHandler.Superclass]
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler
type MLROnDemandConnectionHandler struct {
	objectivec.Object
}

// MLROnDemandConnectionHandlerFromID constructs a [MLROnDemandConnectionHandler] from an objc.ID.
func MLROnDemandConnectionHandlerFromID(id objc.ID) MLROnDemandConnectionHandler {
	return MLROnDemandConnectionHandler{objectivec.Object{ID: id}}
}
// Ensure MLROnDemandConnectionHandler implements IMLROnDemandConnectionHandler.
var _ IMLROnDemandConnectionHandler = MLROnDemandConnectionHandler{}

// An interface definition for the [MLROnDemandConnectionHandler] class.
//
// # Methods
//
//   - [IMLROnDemandConnectionHandler.PerformTaskCompletionHandler]
//   - [IMLROnDemandConnectionHandler.PerformTaskInternalCompletionHandler]
//   - [IMLROnDemandConnectionHandler.PluginPrincipal]
//   - [IMLROnDemandConnectionHandler.SetPluginPrincipal]
//   - [IMLROnDemandConnectionHandler.PrincipalObject]
//   - [IMLROnDemandConnectionHandler.ShouldAcceptXPCConnection]
//   - [IMLROnDemandConnectionHandler.XpcConnection]
//   - [IMLROnDemandConnectionHandler.SetXpcConnection]
//   - [IMLROnDemandConnectionHandler.InitWithPrincipalObject]
//   - [IMLROnDemandConnectionHandler.DebugDescription]
//   - [IMLROnDemandConnectionHandler.Description]
//   - [IMLROnDemandConnectionHandler.Hash]
//   - [IMLROnDemandConnectionHandler.Superclass]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler
type IMLROnDemandConnectionHandler interface {
	objectivec.IObject

	// Topic: Methods

	PerformTaskCompletionHandler(task objectivec.IObject, handler ErrorHandler)
	PerformTaskInternalCompletionHandler(internal objectivec.IObject, handler ErrorHandler)
	PluginPrincipal() objectivec.IObject
	SetPluginPrincipal(value objectivec.IObject)
	PrincipalObject() objectivec.IObject
	ShouldAcceptXPCConnection(xPCConnection objectivec.IObject) bool
	XpcConnection() foundation.NSXPCConnection
	SetXpcConnection(value foundation.NSXPCConnection)
	InitWithPrincipalObject(object objectivec.IObject) MLROnDemandConnectionHandler
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (r MLROnDemandConnectionHandler) Init() MLROnDemandConnectionHandler {
	rv := objc.Send[MLROnDemandConnectionHandler](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLROnDemandConnectionHandler) Autorelease() MLROnDemandConnectionHandler {
	rv := objc.Send[MLROnDemandConnectionHandler](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLROnDemandConnectionHandler creates a new MLROnDemandConnectionHandler instance.
func NewMLROnDemandConnectionHandler() MLROnDemandConnectionHandler {
	class := getMLROnDemandConnectionHandlerClass()
	rv := objc.Send[MLROnDemandConnectionHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/initWithPrincipalObject:
func NewROnDemandConnectionHandlerWithPrincipalObject(object objectivec.IObject) MLROnDemandConnectionHandler {
	instance := getMLROnDemandConnectionHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPrincipalObject:"), object)
	return MLROnDemandConnectionHandlerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/performTask:completionHandler:
func (r MLROnDemandConnectionHandler) PerformTaskCompletionHandler(task objectivec.IObject, handler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](r.ID, objc.Sel("performTask:completionHandler:"), task, _block1)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/performTaskInternal:completionHandler:
func (r MLROnDemandConnectionHandler) PerformTaskInternalCompletionHandler(internal objectivec.IObject, handler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](r.ID, objc.Sel("performTaskInternal:completionHandler:"), internal, _block1)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/shouldAcceptXPCConnection:
func (r MLROnDemandConnectionHandler) ShouldAcceptXPCConnection(xPCConnection objectivec.IObject) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("shouldAcceptXPCConnection:"), xPCConnection)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/initWithPrincipalObject:
func (r MLROnDemandConnectionHandler) InitWithPrincipalObject(object objectivec.IObject) MLROnDemandConnectionHandler {
	rv := objc.Send[MLROnDemandConnectionHandler](r.ID, objc.Sel("initWithPrincipalObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/debugDescription
func (r MLROnDemandConnectionHandler) DebugDescription() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/description
func (r MLROnDemandConnectionHandler) Description() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/hash
func (r MLROnDemandConnectionHandler) Hash() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/pluginPrincipal
func (r MLROnDemandConnectionHandler) PluginPrincipal() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("pluginPrincipal"))
	return objectivec.Object{ID: rv}
}
func (r MLROnDemandConnectionHandler) SetPluginPrincipal(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setPluginPrincipal:"), value)
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/principalObject
func (r MLROnDemandConnectionHandler) PrincipalObject() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("principalObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/superclass
func (r MLROnDemandConnectionHandler) Superclass() objc.Class {
	rv := objc.Send[objc.Class](r.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLROnDemandConnectionHandler/xpcConnection
func (r MLROnDemandConnectionHandler) XpcConnection() foundation.NSXPCConnection {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("xpcConnection"))
	return foundation.NSXPCConnectionFromID(objc.ID(rv))
}
func (r MLROnDemandConnectionHandler) SetXpcConnection(value foundation.NSXPCConnection) {
	objc.Send[struct{}](r.ID, objc.Sel("setXpcConnection:"), value)
}

