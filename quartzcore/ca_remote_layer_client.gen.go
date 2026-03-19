// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CARemoteLayerClient] class.
var (
	_CARemoteLayerClientClass     CARemoteLayerClientClass
	_CARemoteLayerClientClassOnce sync.Once
)

func getCARemoteLayerClientClass() CARemoteLayerClientClass {
	_CARemoteLayerClientClassOnce.Do(func() {
		_CARemoteLayerClientClass = CARemoteLayerClientClass{class: objc.GetClass("CARemoteLayerClient")}
	})
	return _CARemoteLayerClientClass
}

// GetCARemoteLayerClientClass returns the class object for CARemoteLayerClient.
func GetCARemoteLayerClientClass() CARemoteLayerClientClass {
	return getCARemoteLayerClientClass()
}

type CARemoteLayerClientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CARemoteLayerClientClass) Alloc() CARemoteLayerClient {
	rv := objc.Send[CARemoteLayerClient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A legacy class for cross-process rendering.
//
// # Overview
// 
// [CARemoteLaterClient] is a legacy class for cross-process rendering.
// [IOSurfaceCreateMachPort(_:)] and [IOSurfaceCreateXPCObject(_:)], available
// with [IOSurface], offer an improved way to perform cross-process rendering.
//
// [IOSurfaceCreateMachPort(_:)]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateMachPort(_:)
// [IOSurfaceCreateXPCObject(_:)]: https://developer.apple.com/documentation/IOSurface/IOSurfaceCreateXPCObject(_:)
// [IOSurface]: https://developer.apple.com/documentation/IOSurface/IOSurface
//
// # Creating a Client
//
//   - [CARemoteLayerClient.InitWithServerPort]: Creates a layer client from a server port.
//
// # Retrieving Client Properties
//
//   - [CARemoteLayerClient.ClientId]: The ID of the remote layer client.
//   - [CARemoteLayerClient.Layer]: The layer associated with the remote client.
//   - [CARemoteLayerClient.SetLayer]
//
// # Invalidating a Client
//
//   - [CARemoteLayerClient.Invalidate]: Invalidates a remote layer client.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient
type CARemoteLayerClient struct {
	objectivec.Object
}

// CARemoteLayerClientFromID constructs a [CARemoteLayerClient] from an objc.ID.
//
// A legacy class for cross-process rendering.
func CARemoteLayerClientFromID(id objc.ID) CARemoteLayerClient {
	return CARemoteLayerClient{objectivec.Object{ID: id}}
}
// NOTE: CARemoteLayerClient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CARemoteLayerClient] class.
//
// # Creating a Client
//
//   - [ICARemoteLayerClient.InitWithServerPort]: Creates a layer client from a server port.
//
// # Retrieving Client Properties
//
//   - [ICARemoteLayerClient.ClientId]: The ID of the remote layer client.
//   - [ICARemoteLayerClient.Layer]: The layer associated with the remote client.
//   - [ICARemoteLayerClient.SetLayer]
//
// # Invalidating a Client
//
//   - [ICARemoteLayerClient.Invalidate]: Invalidates a remote layer client.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient
type ICARemoteLayerClient interface {
	objectivec.IObject

	// Topic: Creating a Client

	// Creates a layer client from a server port.
	InitWithServerPort(port uint32) CARemoteLayerClient

	// Topic: Retrieving Client Properties

	// The ID of the remote layer client.
	ClientId() uint32
	// The layer associated with the remote client.
	Layer() ICALayer
	SetLayer(value ICALayer)

	// Topic: Invalidating a Client

	// Invalidates a remote layer client.
	Invalidate()
}

// Init initializes the instance.
func (r CARemoteLayerClient) Init() CARemoteLayerClient {
	rv := objc.Send[CARemoteLayerClient](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CARemoteLayerClient) Autorelease() CARemoteLayerClient {
	rv := objc.Send[CARemoteLayerClient](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCARemoteLayerClient creates a new CARemoteLayerClient instance.
func NewCARemoteLayerClient() CARemoteLayerClient {
	class := getCARemoteLayerClientClass()
	rv := objc.Send[CARemoteLayerClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a layer client from a server port.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/init(serverPort:)
func NewRemoteLayerClientWithServerPort(port uint32) CARemoteLayerClient {
	instance := getCARemoteLayerClientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServerPort:"), port)
	return CARemoteLayerClientFromID(rv)
}

// Creates a layer client from a server port.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/init(serverPort:)
func (r CARemoteLayerClient) InitWithServerPort(port uint32) CARemoteLayerClient {
	rv := objc.Send[CARemoteLayerClient](r.ID, objc.Sel("initWithServerPort:"), port)
	return rv
}
// Invalidates a remote layer client.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/invalidate()
func (r CARemoteLayerClient) Invalidate() {
	objc.Send[objc.ID](r.ID, objc.Sel("invalidate"))
}

// The ID of the remote layer client.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/clientId
func (r CARemoteLayerClient) ClientId() uint32 {
	rv := objc.Send[uint32](r.ID, objc.Sel("clientId"))
	return rv
}
// The layer associated with the remote client.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerClient/layer
func (r CARemoteLayerClient) Layer() ICALayer {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("layer"))
	return CALayerFromID(objc.ID(rv))
}
func (r CARemoteLayerClient) SetLayer(value ICALayer) {
	objc.Send[struct{}](r.ID, objc.Sel("setLayer:"), value)
}

