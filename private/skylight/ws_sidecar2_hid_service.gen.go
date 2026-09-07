// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSSidecar2HIDService] class.
var (
	_WSSidecar2HIDServiceClass     WSSidecar2HIDServiceClass
	_WSSidecar2HIDServiceClassOnce sync.Once
)

func getWSSidecar2HIDServiceClass() WSSidecar2HIDServiceClass {
	_WSSidecar2HIDServiceClassOnce.Do(func() {
		_WSSidecar2HIDServiceClass = WSSidecar2HIDServiceClass{class: objc.GetClass("WSSidecar2HIDService")}
	})
	return _WSSidecar2HIDServiceClass
}

// GetWSSidecar2HIDServiceClass returns the class object for WSSidecar2HIDService.
func GetWSSidecar2HIDServiceClass() WSSidecar2HIDServiceClass {
	return getWSSidecar2HIDServiceClass()
}

type WSSidecar2HIDServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSidecar2HIDServiceClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSidecar2HIDServiceClass) Alloc() WSSidecar2HIDService {
	rv := objc.SendIfResponds[WSSidecar2HIDService](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSSidecar2HIDService.DisplayUUID]
//   - [WSSidecar2HIDService.InitWithIOHIDServiceRefDisplayUUID]
type WSSidecar2HIDService struct {
	objectivec.Object
}

// WSSidecar2HIDServiceFromID constructs a [WSSidecar2HIDService] from an objc.ID.
func WSSidecar2HIDServiceFromID(id objc.ID) WSSidecar2HIDService {
	return WSSidecar2HIDService{objectivec.Object{ID: id}}
}

// NOTE: WSSidecar2HIDService embeds objectivec.Object because the parent type is
// unavailable, but IWSSidecar2HIDService embeds IBKIOHIDService, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [WSSidecar2HIDService] class.
//
// # Methods
//
//   - [IWSSidecar2HIDService.DisplayUUID]
//   - [IWSSidecar2HIDService.InitWithIOHIDServiceRefDisplayUUID]
type IWSSidecar2HIDService interface {
	IBKIOHIDService

	// Topic: Methods

	DisplayUUID() objectivec.IObject
	InitWithIOHIDServiceRefDisplayUUID(ref uintptr, uuid objectivec.IObject) WSSidecar2HIDService
}

// Init initializes the instance.
func (w WSSidecar2HIDService) Init() WSSidecar2HIDService {
	rv := objc.SendIfResponds[WSSidecar2HIDService](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSidecar2HIDService) Autorelease() WSSidecar2HIDService {
	rv := objc.SendIfResponds[WSSidecar2HIDService](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSidecar2HIDService creates a new WSSidecar2HIDService instance.
func NewWSSidecar2HIDService() WSSidecar2HIDService {
	class := getWSSidecar2HIDServiceClass()
	rv := objc.SendIfResponds[WSSidecar2HIDService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSSidecar2HIDServiceWithIOHIDServiceRefDisplayUUID(ref uintptr, uuid objectivec.IObject) WSSidecar2HIDService {
	instance := getWSSidecar2HIDServiceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIOHIDServiceRef:displayUUID:"), ref, uuid)
	return WSSidecar2HIDServiceFromID(rv)
}

func (w WSSidecar2HIDService) DisplayUUID() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("displayUUID"))
	return objectivec.Object{ID: rv}
}
func (w WSSidecar2HIDService) InitWithIOHIDServiceRefDisplayUUID(ref uintptr, uuid objectivec.IObject) WSSidecar2HIDService {
	rv := objc.SendIfResponds[WSSidecar2HIDService](w.ID, objc.Sel("initWithIOHIDServiceRef:displayUUID:"), ref, uuid)
	return rv
}
