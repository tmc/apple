// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CARemoteLayerServer] class.
var (
	_CARemoteLayerServerClass     CARemoteLayerServerClass
	_CARemoteLayerServerClassOnce sync.Once
)

func getCARemoteLayerServerClass() CARemoteLayerServerClass {
	_CARemoteLayerServerClassOnce.Do(func() {
		_CARemoteLayerServerClass = CARemoteLayerServerClass{class: objc.GetClass("CARemoteLayerServer")}
	})
	return _CARemoteLayerServerClass
}

// GetCARemoteLayerServerClass returns the class object for CARemoteLayerServer.
func GetCARemoteLayerServerClass() CARemoteLayerServerClass {
	return getCARemoteLayerServerClass()
}

type CARemoteLayerServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CARemoteLayerServerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CARemoteLayerServerClass) Alloc() CARemoteLayerServer {
	rv := objc.Send[CARemoteLayerServer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A legacy class for cross-process rendering.
//
// # Overview
// 
// [CARemoteLaterServer] is a legacy class for cross-process rendering.
// [IOSurfaceCreateMachPort(_:)] and [IOSurfaceCreateXPCObject(_:)], available
// with [IOSurface], offer an improved way to perform cross-process rendering.
//
// [IOSurfaceCreateMachPort(_:)]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateMachPort(_:)
// [IOSurfaceCreateXPCObject(_:)]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateXPCObject(_:)
// [IOSurface]: https://developer.apple.com/documentation/IOSurface/IOSurface
//
// # Creating a Server
//
//   - [CARemoteLayerServer.ServerPort]: The port number of the server.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer
type CARemoteLayerServer struct {
	objectivec.Object
}

// CARemoteLayerServerFromID constructs a [CARemoteLayerServer] from an objc.ID.
//
// A legacy class for cross-process rendering.
func CARemoteLayerServerFromID(id objc.ID) CARemoteLayerServer {
	return CARemoteLayerServer{objectivec.Object{ID: id}}
}
// NOTE: CARemoteLayerServer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CARemoteLayerServer] class.
//
// # Creating a Server
//
//   - [ICARemoteLayerServer.ServerPort]: The port number of the server.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer
type ICARemoteLayerServer interface {
	objectivec.IObject

	// Topic: Creating a Server

	// The port number of the server.
	ServerPort() uint32
}

// Init initializes the instance.
func (r CARemoteLayerServer) Init() CARemoteLayerServer {
	rv := objc.Send[CARemoteLayerServer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CARemoteLayerServer) Autorelease() CARemoteLayerServer {
	rv := objc.Send[CARemoteLayerServer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCARemoteLayerServer creates a new CARemoteLayerServer instance.
func NewCARemoteLayerServer() CARemoteLayerServer {
	class := getCARemoteLayerServerClass()
	rv := objc.Send[CARemoteLayerServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the (singleton) instance of the shared remote layer server.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer/shared()
func (_CARemoteLayerServerClass CARemoteLayerServerClass) SharedServer() CARemoteLayerServer {
	rv := objc.Send[objc.ID](objc.ID(_CARemoteLayerServerClass.class), objc.Sel("sharedServer"))
	return CARemoteLayerServerFromID(rv)
}

// The port number of the server.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer/serverPort
func (r CARemoteLayerServer) ServerPort() uint32 {
	rv := objc.Send[uint32](r.ID, objc.Sel("serverPort"))
	return rv
}

