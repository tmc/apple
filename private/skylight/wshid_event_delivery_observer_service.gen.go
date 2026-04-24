// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSHIDEventDeliveryObserverService] class.
var (
	_WSHIDEventDeliveryObserverServiceClass     WSHIDEventDeliveryObserverServiceClass
	_WSHIDEventDeliveryObserverServiceClassOnce sync.Once
)

func getWSHIDEventDeliveryObserverServiceClass() WSHIDEventDeliveryObserverServiceClass {
	_WSHIDEventDeliveryObserverServiceClassOnce.Do(func() {
		_WSHIDEventDeliveryObserverServiceClass = WSHIDEventDeliveryObserverServiceClass{class: objc.GetClass("WSHIDEventDeliveryObserverService")}
	})
	return _WSHIDEventDeliveryObserverServiceClass
}

// GetWSHIDEventDeliveryObserverServiceClass returns the class object for WSHIDEventDeliveryObserverService.
func GetWSHIDEventDeliveryObserverServiceClass() WSHIDEventDeliveryObserverServiceClass {
	return getWSHIDEventDeliveryObserverServiceClass()
}

type WSHIDEventDeliveryObserverServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSHIDEventDeliveryObserverServiceClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSHIDEventDeliveryObserverServiceClass) Alloc() WSHIDEventDeliveryObserverService {
	rv := objc.Send[WSHIDEventDeliveryObserverService](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSHIDEventDeliveryObserverService.BkObserverService]
//   - [WSHIDEventDeliveryObserverService.InitWithServer]
//
// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverService
type WSHIDEventDeliveryObserverService struct {
	objectivec.Object
}

// WSHIDEventDeliveryObserverServiceFromID constructs a [WSHIDEventDeliveryObserverService] from an objc.ID.
func WSHIDEventDeliveryObserverServiceFromID(id objc.ID) WSHIDEventDeliveryObserverService {
	return WSHIDEventDeliveryObserverService{objectivec.Object{ID: id}}
}

// Ensure WSHIDEventDeliveryObserverService implements IWSHIDEventDeliveryObserverService.
var _ IWSHIDEventDeliveryObserverService = WSHIDEventDeliveryObserverService{}

// An interface definition for the [WSHIDEventDeliveryObserverService] class.
//
// # Methods
//
//   - [IWSHIDEventDeliveryObserverService.BkObserverService]
//   - [IWSHIDEventDeliveryObserverService.InitWithServer]
//
// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverService
type IWSHIDEventDeliveryObserverService interface {
	objectivec.IObject

	// Topic: Methods

	BkObserverService() unsafe.Pointer
	InitWithServer(server objectivec.IObject) WSHIDEventDeliveryObserverService
}

// Init initializes the instance.
func (w WSHIDEventDeliveryObserverService) Init() WSHIDEventDeliveryObserverService {
	rv := objc.Send[WSHIDEventDeliveryObserverService](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSHIDEventDeliveryObserverService) Autorelease() WSHIDEventDeliveryObserverService {
	rv := objc.Send[WSHIDEventDeliveryObserverService](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSHIDEventDeliveryObserverService creates a new WSHIDEventDeliveryObserverService instance.
func NewWSHIDEventDeliveryObserverService() WSHIDEventDeliveryObserverService {
	class := getWSHIDEventDeliveryObserverServiceClass()
	rv := objc.Send[WSHIDEventDeliveryObserverService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverService/initWithServer:
func NewWSHIDEventDeliveryObserverServiceWithServer(server objectivec.IObject) WSHIDEventDeliveryObserverService {
	instance := getWSHIDEventDeliveryObserverServiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServer:"), server)
	return WSHIDEventDeliveryObserverServiceFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverService/initWithServer:
func (w WSHIDEventDeliveryObserverService) InitWithServer(server objectivec.IObject) WSHIDEventDeliveryObserverService {
	rv := objc.Send[WSHIDEventDeliveryObserverService](w.ID, objc.Sel("initWithServer:"), server)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverService/bkObserverService
func (w WSHIDEventDeliveryObserverService) BkObserverService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w.ID, objc.Sel("bkObserverService"))
	return rv
}
