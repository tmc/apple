// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRSandboxExtensionRequest] class.
var (
	_MLRSandboxExtensionRequestClass     MLRSandboxExtensionRequestClass
	_MLRSandboxExtensionRequestClassOnce sync.Once
)

func getMLRSandboxExtensionRequestClass() MLRSandboxExtensionRequestClass {
	_MLRSandboxExtensionRequestClassOnce.Do(func() {
		_MLRSandboxExtensionRequestClass = MLRSandboxExtensionRequestClass{class: objc.GetClass("MLRSandboxExtensionRequest")}
	})
	return _MLRSandboxExtensionRequestClass
}

// GetMLRSandboxExtensionRequestClass returns the class object for MLRSandboxExtensionRequest.
func GetMLRSandboxExtensionRequestClass() MLRSandboxExtensionRequestClass {
	return getMLRSandboxExtensionRequestClass()
}

type MLRSandboxExtensionRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRSandboxExtensionRequestClass) Alloc() MLRSandboxExtensionRequest {
	rv := objc.Send[MLRSandboxExtensionRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRSandboxExtensionRequest.URLs]
//   - [MLRSandboxExtensionRequest.RequireWrite]
//   - [MLRSandboxExtensionRequest.InitWithURLsRequireWrite]
// See: https://developer.apple.com/documentation/MLRuntime/MLRSandboxExtensionRequest
type MLRSandboxExtensionRequest struct {
	objectivec.Object
}

// MLRSandboxExtensionRequestFromID constructs a [MLRSandboxExtensionRequest] from an objc.ID.
func MLRSandboxExtensionRequestFromID(id objc.ID) MLRSandboxExtensionRequest {
	return MLRSandboxExtensionRequest{objectivec.Object{id}}
}
// Ensure MLRSandboxExtensionRequest implements IMLRSandboxExtensionRequest.
var _ IMLRSandboxExtensionRequest = MLRSandboxExtensionRequest{}

// An interface definition for the [MLRSandboxExtensionRequest] class.
//
// # Methods
//
//   - [IMLRSandboxExtensionRequest.URLs]
//   - [IMLRSandboxExtensionRequest.RequireWrite]
//   - [IMLRSandboxExtensionRequest.InitWithURLsRequireWrite]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRSandboxExtensionRequest
type IMLRSandboxExtensionRequest interface {
	objectivec.IObject

	// Topic: Methods

	URLs() foundation.INSArray
	RequireWrite() bool
	InitWithURLsRequireWrite(uRLs objectivec.IObject, write bool) MLRSandboxExtensionRequest
}

// Init initializes the instance.
func (r MLRSandboxExtensionRequest) Init() MLRSandboxExtensionRequest {
	rv := objc.Send[MLRSandboxExtensionRequest](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRSandboxExtensionRequest) Autorelease() MLRSandboxExtensionRequest {
	rv := objc.Send[MLRSandboxExtensionRequest](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRSandboxExtensionRequest creates a new MLRSandboxExtensionRequest instance.
func NewMLRSandboxExtensionRequest() MLRSandboxExtensionRequest {
	class := getMLRSandboxExtensionRequestClass()
	rv := objc.Send[MLRSandboxExtensionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRSandboxExtensionRequest/initWithURLs:requireWrite:
func NewRSandboxExtensionRequestWithURLsRequireWrite(uRLs objectivec.IObject, write bool) MLRSandboxExtensionRequest {
	instance := getMLRSandboxExtensionRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURLs:requireWrite:"), uRLs, write)
	return MLRSandboxExtensionRequestFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRSandboxExtensionRequest/initWithURLs:requireWrite:
func (r MLRSandboxExtensionRequest) InitWithURLsRequireWrite(uRLs objectivec.IObject, write bool) MLRSandboxExtensionRequest {
	rv := objc.Send[MLRSandboxExtensionRequest](r.ID, objc.Sel("initWithURLs:requireWrite:"), uRLs, write)
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRSandboxExtensionRequest/URLs
func (r MLRSandboxExtensionRequest) URLs() foundation.INSArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("URLs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRSandboxExtensionRequest/requireWrite
func (r MLRSandboxExtensionRequest) RequireWrite() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("requireWrite"))
	return rv
}

