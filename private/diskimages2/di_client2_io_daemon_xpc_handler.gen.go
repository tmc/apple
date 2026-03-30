// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIClient2IODaemonXPCHandler] class.
var (
	_DIClient2IODaemonXPCHandlerClass     DIClient2IODaemonXPCHandlerClass
	_DIClient2IODaemonXPCHandlerClassOnce sync.Once
)

func getDIClient2IODaemonXPCHandlerClass() DIClient2IODaemonXPCHandlerClass {
	_DIClient2IODaemonXPCHandlerClassOnce.Do(func() {
		_DIClient2IODaemonXPCHandlerClass = DIClient2IODaemonXPCHandlerClass{class: objc.GetClass("DIClient2IODaemonXPCHandler")}
	})
	return _DIClient2IODaemonXPCHandlerClass
}

// GetDIClient2IODaemonXPCHandlerClass returns the class object for DIClient2IODaemonXPCHandler.
func GetDIClient2IODaemonXPCHandlerClass() DIClient2IODaemonXPCHandlerClass {
	return getDIClient2IODaemonXPCHandlerClass()
}

type DIClient2IODaemonXPCHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIClient2IODaemonXPCHandlerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIClient2IODaemonXPCHandlerClass) Alloc() DIClient2IODaemonXPCHandler {
	rv := objc.Send[DIClient2IODaemonXPCHandler](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIClient2IODaemonXPCHandler.AddToRefCountWithError]
//   - [DIClient2IODaemonXPCHandler.XpcListenerEndpoint]
//   - [DIClient2IODaemonXPCHandler.SetXpcListenerEndpoint]
//   - [DIClient2IODaemonXPCHandler.InitWithEndpoint]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIClient2IODaemonXPCHandler
type DIClient2IODaemonXPCHandler struct {
	DIBaseXPCHandler
}

// DIClient2IODaemonXPCHandlerFromID constructs a [DIClient2IODaemonXPCHandler] from an objc.ID.
func DIClient2IODaemonXPCHandlerFromID(id objc.ID) DIClient2IODaemonXPCHandler {
	return DIClient2IODaemonXPCHandler{DIBaseXPCHandler: DIBaseXPCHandlerFromID(id)}
}

// Ensure DIClient2IODaemonXPCHandler implements IDIClient2IODaemonXPCHandler.
var _ IDIClient2IODaemonXPCHandler = DIClient2IODaemonXPCHandler{}

// An interface definition for the [DIClient2IODaemonXPCHandler] class.
//
// # Methods
//
//   - [IDIClient2IODaemonXPCHandler.AddToRefCountWithError]
//   - [IDIClient2IODaemonXPCHandler.XpcListenerEndpoint]
//   - [IDIClient2IODaemonXPCHandler.SetXpcListenerEndpoint]
//   - [IDIClient2IODaemonXPCHandler.InitWithEndpoint]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIClient2IODaemonXPCHandler
type IDIClient2IODaemonXPCHandler interface {
	IDIBaseXPCHandler

	// Topic: Methods

	AddToRefCountWithError() (bool, error)
	XpcListenerEndpoint() foundation.NSXPCListenerEndpoint
	SetXpcListenerEndpoint(value foundation.NSXPCListenerEndpoint)
	InitWithEndpoint(endpoint objectivec.IObject) DIClient2IODaemonXPCHandler
}

// Init initializes the instance.
func (d DIClient2IODaemonXPCHandler) Init() DIClient2IODaemonXPCHandler {
	rv := objc.Send[DIClient2IODaemonXPCHandler](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIClient2IODaemonXPCHandler) Autorelease() DIClient2IODaemonXPCHandler {
	rv := objc.Send[DIClient2IODaemonXPCHandler](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIClient2IODaemonXPCHandler creates a new DIClient2IODaemonXPCHandler instance.
func NewDIClient2IODaemonXPCHandler() DIClient2IODaemonXPCHandler {
	class := getDIClient2IODaemonXPCHandlerClass()
	rv := objc.Send[DIClient2IODaemonXPCHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2IODaemonXPCHandler/initWithEndpoint:
func NewDIClient2IODaemonXPCHandlerWithEndpoint(endpoint objectivec.IObject) DIClient2IODaemonXPCHandler {
	instance := getDIClient2IODaemonXPCHandlerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEndpoint:"), endpoint)
	return DIClient2IODaemonXPCHandlerFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2IODaemonXPCHandler/addToRefCountWithError:
func (d DIClient2IODaemonXPCHandler) AddToRefCountWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addToRefCountWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addToRefCountWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2IODaemonXPCHandler/initWithEndpoint:
func (d DIClient2IODaemonXPCHandler) InitWithEndpoint(endpoint objectivec.IObject) DIClient2IODaemonXPCHandler {
	rv := objc.Send[DIClient2IODaemonXPCHandler](d.ID, objc.Sel("initWithEndpoint:"), endpoint)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2IODaemonXPCHandler/xpcListenerEndpoint
func (d DIClient2IODaemonXPCHandler) XpcListenerEndpoint() foundation.NSXPCListenerEndpoint {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("xpcListenerEndpoint"))
	return foundation.NSXPCListenerEndpointFromID(objc.ID(rv))
}
func (d DIClient2IODaemonXPCHandler) SetXpcListenerEndpoint(value foundation.NSXPCListenerEndpoint) {
	objc.Send[struct{}](d.ID, objc.Sel("setXpcListenerEndpoint:"), value)
}
