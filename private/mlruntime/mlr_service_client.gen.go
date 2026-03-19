// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRServiceClient] class.
var (
	_MLRServiceClientClass     MLRServiceClientClass
	_MLRServiceClientClassOnce sync.Once
)

func getMLRServiceClientClass() MLRServiceClientClass {
	_MLRServiceClientClassOnce.Do(func() {
		_MLRServiceClientClass = MLRServiceClientClass{class: objc.GetClass("MLRServiceClient")}
	})
	return _MLRServiceClientClass
}

// GetMLRServiceClientClass returns the class object for MLRServiceClient.
func GetMLRServiceClientClass() MLRServiceClientClass {
	return getMLRServiceClientClass()
}

type MLRServiceClientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRServiceClientClass) Alloc() MLRServiceClient {
	rv := objc.Send[MLRServiceClient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRServiceClient.DonateJSONResultIdentifierCompletion]
//   - [MLRServiceClient.PerformOnRemoteObjectWithBlockErrorHandler]
//   - [MLRServiceClient.PerformOnRemoteObjectWithBlockIsSynchronousErrorHandler]
//   - [MLRServiceClient.PerformSynchronouslyOnRemoteObjectWithBlockErrorHandler]
//   - [MLRServiceClient.InitWithConnection]
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient
type MLRServiceClient struct {
	objectivec.Object
}

// MLRServiceClientFromID constructs a [MLRServiceClient] from an objc.ID.
func MLRServiceClientFromID(id objc.ID) MLRServiceClient {
	return MLRServiceClient{objectivec.Object{ID: id}}
}
// Ensure MLRServiceClient implements IMLRServiceClient.
var _ IMLRServiceClient = MLRServiceClient{}

// An interface definition for the [MLRServiceClient] class.
//
// # Methods
//
//   - [IMLRServiceClient.DonateJSONResultIdentifierCompletion]
//   - [IMLRServiceClient.PerformOnRemoteObjectWithBlockErrorHandler]
//   - [IMLRServiceClient.PerformOnRemoteObjectWithBlockIsSynchronousErrorHandler]
//   - [IMLRServiceClient.PerformSynchronouslyOnRemoteObjectWithBlockErrorHandler]
//   - [IMLRServiceClient.InitWithConnection]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient
type IMLRServiceClient interface {
	objectivec.IObject

	// Topic: Methods

	DonateJSONResultIdentifierCompletion(jSONResult objectivec.IObject, identifier objectivec.IObject, completion ErrorHandler)
	PerformOnRemoteObjectWithBlockErrorHandler(block objectivec.IObject, handler ErrorHandler)
	PerformOnRemoteObjectWithBlockIsSynchronousErrorHandler(block objectivec.IObject, synchronous bool, handler ErrorHandler)
	PerformSynchronouslyOnRemoteObjectWithBlockErrorHandler(block objectivec.IObject, handler ErrorHandler)
	InitWithConnection(connection objectivec.IObject) MLRServiceClient
}

// Init initializes the instance.
func (r MLRServiceClient) Init() MLRServiceClient {
	rv := objc.Send[MLRServiceClient](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRServiceClient) Autorelease() MLRServiceClient {
	rv := objc.Send[MLRServiceClient](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRServiceClient creates a new MLRServiceClient instance.
func NewMLRServiceClient() MLRServiceClient {
	class := getMLRServiceClientClass()
	rv := objc.Send[MLRServiceClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient/initWithConnection:
func NewRServiceClientWithConnection(connection objectivec.IObject) MLRServiceClient {
	instance := getMLRServiceClientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConnection:"), connection)
	return MLRServiceClientFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient/donateJSONResult:identifier:completion:
func (r MLRServiceClient) DonateJSONResultIdentifierCompletion(jSONResult objectivec.IObject, identifier objectivec.IObject, completion ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completion)
	defer _cleanup2()
	objc.Send[objc.ID](r.ID, objc.Sel("donateJSONResult:identifier:completion:"), jSONResult, identifier, _block2)
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient/performOnRemoteObjectWithBlock:errorHandler:
func (r MLRServiceClient) PerformOnRemoteObjectWithBlockErrorHandler(block objectivec.IObject, handler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](r.ID, objc.Sel("performOnRemoteObjectWithBlock:errorHandler:"), block, _block1)
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient/performOnRemoteObjectWithBlock:isSynchronous:errorHandler:
func (r MLRServiceClient) PerformOnRemoteObjectWithBlockIsSynchronousErrorHandler(block objectivec.IObject, synchronous bool, handler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(handler)
	defer _cleanup2()
	objc.Send[objc.ID](r.ID, objc.Sel("performOnRemoteObjectWithBlock:isSynchronous:errorHandler:"), block, synchronous, _block2)
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient/performSynchronouslyOnRemoteObjectWithBlock:errorHandler:
func (r MLRServiceClient) PerformSynchronouslyOnRemoteObjectWithBlockErrorHandler(block objectivec.IObject, handler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](r.ID, objc.Sel("performSynchronouslyOnRemoteObjectWithBlock:errorHandler:"), block, _block1)
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient/initWithConnection:
func (r MLRServiceClient) InitWithConnection(connection objectivec.IObject) MLRServiceClient {
	rv := objc.Send[MLRServiceClient](r.ID, objc.Sel("initWithConnection:"), connection)
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRServiceClient/sharedConnection
func (_MLRServiceClientClass MLRServiceClientClass) SharedConnection() MLRServiceClient {
	rv := objc.Send[objc.ID](objc.ID(_MLRServiceClientClass.class), objc.Sel("sharedConnection"))
	return MLRServiceClientFromID(rv)
}

