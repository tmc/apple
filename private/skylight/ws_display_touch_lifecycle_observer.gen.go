// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSDisplayTouchLifecycleObserver] class.
var (
	_WSDisplayTouchLifecycleObserverClass     WSDisplayTouchLifecycleObserverClass
	_WSDisplayTouchLifecycleObserverClassOnce sync.Once
)

func getWSDisplayTouchLifecycleObserverClass() WSDisplayTouchLifecycleObserverClass {
	_WSDisplayTouchLifecycleObserverClassOnce.Do(func() {
		_WSDisplayTouchLifecycleObserverClass = WSDisplayTouchLifecycleObserverClass{class: objc.GetClass("WSDisplayTouchLifecycleObserver")}
	})
	return _WSDisplayTouchLifecycleObserverClass
}

// GetWSDisplayTouchLifecycleObserverClass returns the class object for WSDisplayTouchLifecycleObserver.
func GetWSDisplayTouchLifecycleObserverClass() WSDisplayTouchLifecycleObserverClass {
	return getWSDisplayTouchLifecycleObserverClass()
}

type WSDisplayTouchLifecycleObserverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSDisplayTouchLifecycleObserverClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSDisplayTouchLifecycleObserverClass) Alloc() WSDisplayTouchLifecycleObserver {
	rv := objc.SendIfResponds[WSDisplayTouchLifecycleObserver](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSDisplayTouchLifecycleObserver.DisplayUUID]
//   - [WSDisplayTouchLifecycleObserver.SetDisplayUUID]
//   - [WSDisplayTouchLifecycleObserver.HitTestRegionsDidChange]
//   - [WSDisplayTouchLifecycleObserver.TouchDidAlwaysRouteToTargetIDClientTaskName]
//   - [WSDisplayTouchLifecycleObserver.TouchDidHitTestToDestinationHostingChainIndex]
//   - [WSDisplayTouchLifecycleObserver.TouchDidMoveToPointEventMaskZMaxZ]
//   - [WSDisplayTouchLifecycleObserver.TouchPathIndexDownAtPointEventMaskTransducerType]
//   - [WSDisplayTouchLifecycleObserver.TouchPathIndexRangeInAtPointEventMaskZMaxZTransducerType]
//   - [WSDisplayTouchLifecycleObserver.TouchPathIndexRangeOutAtPoint]
//   - [WSDisplayTouchLifecycleObserver.TouchPathIndexUpAtPointDetached]
//   - [WSDisplayTouchLifecycleObserver.TouchDidDetachDestinations]
//   - [WSDisplayTouchLifecycleObserver.TouchDidFinishProcessingTouchCollection]
//   - [WSDisplayTouchLifecycleObserver.TouchDidHIDCancel]
//   - [WSDisplayTouchLifecycleObserver.TouchDidSoftCancel]
//   - [WSDisplayTouchLifecycleObserver.TouchDidTransferDestination]
//   - [WSDisplayTouchLifecycleObserver.TouchWillStartProcessingTouchCollection]
//   - [WSDisplayTouchLifecycleObserver.InitWithDisplayUUID]
//   - [WSDisplayTouchLifecycleObserver.DebugDescription]
//   - [WSDisplayTouchLifecycleObserver.Description]
//   - [WSDisplayTouchLifecycleObserver.Hash]
//   - [WSDisplayTouchLifecycleObserver.Superclass]
type WSDisplayTouchLifecycleObserver struct {
	objectivec.Object
}

// WSDisplayTouchLifecycleObserverFromID constructs a [WSDisplayTouchLifecycleObserver] from an objc.ID.
func WSDisplayTouchLifecycleObserverFromID(id objc.ID) WSDisplayTouchLifecycleObserver {
	return WSDisplayTouchLifecycleObserver{objectivec.Object{ID: id}}
}

// Ensure WSDisplayTouchLifecycleObserver implements IWSDisplayTouchLifecycleObserver.
var _ IWSDisplayTouchLifecycleObserver = WSDisplayTouchLifecycleObserver{}

// An interface definition for the [WSDisplayTouchLifecycleObserver] class.
//
// # Methods
//
//   - [IWSDisplayTouchLifecycleObserver.DisplayUUID]
//   - [IWSDisplayTouchLifecycleObserver.SetDisplayUUID]
//   - [IWSDisplayTouchLifecycleObserver.HitTestRegionsDidChange]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidAlwaysRouteToTargetIDClientTaskName]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidHitTestToDestinationHostingChainIndex]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidMoveToPointEventMaskZMaxZ]
//   - [IWSDisplayTouchLifecycleObserver.TouchPathIndexDownAtPointEventMaskTransducerType]
//   - [IWSDisplayTouchLifecycleObserver.TouchPathIndexRangeInAtPointEventMaskZMaxZTransducerType]
//   - [IWSDisplayTouchLifecycleObserver.TouchPathIndexRangeOutAtPoint]
//   - [IWSDisplayTouchLifecycleObserver.TouchPathIndexUpAtPointDetached]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidDetachDestinations]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidFinishProcessingTouchCollection]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidHIDCancel]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidSoftCancel]
//   - [IWSDisplayTouchLifecycleObserver.TouchDidTransferDestination]
//   - [IWSDisplayTouchLifecycleObserver.TouchWillStartProcessingTouchCollection]
//   - [IWSDisplayTouchLifecycleObserver.InitWithDisplayUUID]
//   - [IWSDisplayTouchLifecycleObserver.DebugDescription]
//   - [IWSDisplayTouchLifecycleObserver.Description]
//   - [IWSDisplayTouchLifecycleObserver.Hash]
//   - [IWSDisplayTouchLifecycleObserver.Superclass]
type IWSDisplayTouchLifecycleObserver interface {
	objectivec.IObject

	// Topic: Methods

	DisplayUUID() string
	SetDisplayUUID(value string)
	HitTestRegionsDidChange(change objectivec.IObject)
	TouchDidAlwaysRouteToTargetIDClientTaskName(touch uint32, id unsafe.Pointer, name uint32)
	TouchDidHitTestToDestinationHostingChainIndex(touch uint32, destination objectivec.IObject, index int64)
	TouchDidMoveToPointEventMaskZMaxZ(touch uint32, point corefoundation.CGPoint, mask uint32, z float64, z2 float64)
	TouchPathIndexDownAtPointEventMaskTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, type_ uint32)
	TouchPathIndexRangeInAtPointEventMaskZMaxZTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, z float64, z2 float64, type_ uint32)
	TouchPathIndexRangeOutAtPoint(touch uint32, index int64, point corefoundation.CGPoint)
	TouchPathIndexUpAtPointDetached(touch uint32, index int64, point corefoundation.CGPoint, detached bool)
	TouchDidDetachDestinations(detach uint32, destinations objectivec.IObject)
	TouchDidFinishProcessingTouchCollection()
	TouchDidHIDCancel(hIDCancel uint32)
	TouchDidSoftCancel(cancel uint32)
	TouchDidTransferDestination(transfer uint32, destination objectivec.IObject)
	TouchWillStartProcessingTouchCollection()
	InitWithDisplayUUID(uuid objectivec.IObject) WSDisplayTouchLifecycleObserver
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSDisplayTouchLifecycleObserver) Init() WSDisplayTouchLifecycleObserver {
	rv := objc.SendIfResponds[WSDisplayTouchLifecycleObserver](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSDisplayTouchLifecycleObserver) Autorelease() WSDisplayTouchLifecycleObserver {
	rv := objc.SendIfResponds[WSDisplayTouchLifecycleObserver](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSDisplayTouchLifecycleObserver creates a new WSDisplayTouchLifecycleObserver instance.
func NewWSDisplayTouchLifecycleObserver() WSDisplayTouchLifecycleObserver {
	class := getWSDisplayTouchLifecycleObserverClass()
	rv := objc.SendIfResponds[WSDisplayTouchLifecycleObserver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSDisplayTouchLifecycleObserverWithDisplayUUID(uuid objectivec.IObject) WSDisplayTouchLifecycleObserver {
	instance := getWSDisplayTouchLifecycleObserverClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplayUUID:"), uuid)
	return WSDisplayTouchLifecycleObserverFromID(rv)
}

func (w WSDisplayTouchLifecycleObserver) HitTestRegionsDidChange(change objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("hitTestRegionsDidChange:"), change)
}
func (w WSDisplayTouchLifecycleObserver) TouchDidAlwaysRouteToTargetIDClientTaskName(touch uint32, id unsafe.Pointer, name uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touch:didAlwaysRouteToTargetID:clientTaskName:"), touch, id, name)
}
func (w WSDisplayTouchLifecycleObserver) TouchDidHitTestToDestinationHostingChainIndex(touch uint32, destination objectivec.IObject, index int64) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touch:didHitTestToDestination:hostingChainIndex:"), touch, destination, index)
}
func (w WSDisplayTouchLifecycleObserver) TouchDidMoveToPointEventMaskZMaxZ(touch uint32, point corefoundation.CGPoint, mask uint32, z float64, z2 float64) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touch:didMoveToPoint:eventMask:z:maxZ:"), touch, point, mask, z, z2)
}
func (w WSDisplayTouchLifecycleObserver) TouchPathIndexDownAtPointEventMaskTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, type_ uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touch:pathIndex:downAtPoint:eventMask:transducerType:"), touch, index, point, mask, type_)
}
func (w WSDisplayTouchLifecycleObserver) TouchPathIndexRangeInAtPointEventMaskZMaxZTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, z float64, z2 float64, type_ uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touch:pathIndex:rangeInAtPoint:eventMask:z:maxZ:transducerType:"), touch, index, point, mask, z, z2, type_)
}
func (w WSDisplayTouchLifecycleObserver) TouchPathIndexRangeOutAtPoint(touch uint32, index int64, point corefoundation.CGPoint) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touch:pathIndex:rangeOutAtPoint:"), touch, index, point)
}
func (w WSDisplayTouchLifecycleObserver) TouchPathIndexUpAtPointDetached(touch uint32, index int64, point corefoundation.CGPoint, detached bool) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touch:pathIndex:upAtPoint:detached:"), touch, index, point, detached)
}
func (w WSDisplayTouchLifecycleObserver) TouchDidDetachDestinations(detach uint32, destinations objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touchDidDetach:destinations:"), detach, destinations)
}
func (w WSDisplayTouchLifecycleObserver) TouchDidFinishProcessingTouchCollection() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touchDidFinishProcessingTouchCollection"))
}
func (w WSDisplayTouchLifecycleObserver) TouchDidHIDCancel(hIDCancel uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touchDidHIDCancel:"), hIDCancel)
}
func (w WSDisplayTouchLifecycleObserver) TouchDidSoftCancel(cancel uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touchDidSoftCancel:"), cancel)
}
func (w WSDisplayTouchLifecycleObserver) TouchDidTransferDestination(transfer uint32, destination objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touchDidTransfer:destination:"), transfer, destination)
}
func (w WSDisplayTouchLifecycleObserver) TouchWillStartProcessingTouchCollection() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touchWillStartProcessingTouchCollection"))
}
func (w WSDisplayTouchLifecycleObserver) InitWithDisplayUUID(uuid objectivec.IObject) WSDisplayTouchLifecycleObserver {
	rv := objc.SendIfResponds[WSDisplayTouchLifecycleObserver](w.ID, objc.Sel("initWithDisplayUUID:"), uuid)
	return rv
}

func (w WSDisplayTouchLifecycleObserver) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSDisplayTouchLifecycleObserver) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSDisplayTouchLifecycleObserver) DisplayUUID() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("displayUUID"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSDisplayTouchLifecycleObserver) SetDisplayUUID(value string) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setDisplayUUID:"), objc.String(value))
}
func (w WSDisplayTouchLifecycleObserver) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSDisplayTouchLifecycleObserver) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
