// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKTouchLifecycleObserver protocol.
type BKTouchLifecycleObserver interface {
	objectivec.IObject

	// HitTestRegionsDidChange protocol.
	HitTestRegionsDidChange(change objectivec.IObject)

	// TouchDidAlwaysRouteToTargetIDClientTaskName protocol.
	TouchDidAlwaysRouteToTargetIDClientTaskName(touch uint32, id unsafe.Pointer, name uint32)

	// TouchDidHitTestToDestinationHostingChainIndex protocol.
	TouchDidHitTestToDestinationHostingChainIndex(touch uint32, destination objectivec.IObject, index int64)

	// TouchDidMoveToPointEventMaskZMaxZ protocol.
	TouchDidMoveToPointEventMaskZMaxZ(touch uint32, point corefoundation.CGPoint, mask uint32, z float64, z2 float64)

	// TouchPathIndexDownAtPointEventMaskTransducerType protocol.
	TouchPathIndexDownAtPointEventMaskTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, type_ uint32)

	// TouchPathIndexRangeInAtPointEventMaskZMaxZTransducerType protocol.
	TouchPathIndexRangeInAtPointEventMaskZMaxZTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, z float64, z2 float64, type_ uint32)

	// TouchPathIndexRangeOutAtPoint protocol.
	TouchPathIndexRangeOutAtPoint(touch uint32, index int64, point corefoundation.CGPoint)

	// TouchPathIndexUpAtPointDetached protocol.
	TouchPathIndexUpAtPointDetached(touch uint32, index int64, point corefoundation.CGPoint, detached bool)

	// TouchDidDetachDestinations protocol.
	TouchDidDetachDestinations(detach uint32, destinations objectivec.IObject)

	// TouchDidFinishProcessingTouchCollection protocol.
	TouchDidFinishProcessingTouchCollection()

	// TouchDidHIDCancel protocol.
	TouchDidHIDCancel(hIDCancel uint32)

	// TouchDidSoftCancel protocol.
	TouchDidSoftCancel(cancel uint32)

	// TouchDidTransferDestination protocol.
	TouchDidTransferDestination(transfer uint32, destination objectivec.IObject)

	// TouchWillStartProcessingTouchCollection protocol.
	TouchWillStartProcessingTouchCollection()
}

// BKTouchLifecycleObserverObject wraps an existing Objective-C object that conforms to the BKTouchLifecycleObserver protocol.
type BKTouchLifecycleObserverObject struct {
	objectivec.Object
}

func (o BKTouchLifecycleObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKTouchLifecycleObserverObjectFromID constructs a [BKTouchLifecycleObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKTouchLifecycleObserverObjectFromID(id objc.ID) BKTouchLifecycleObserverObject {
	return BKTouchLifecycleObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o BKTouchLifecycleObserverObject) HitTestRegionsDidChange(change objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("hitTestRegionsDidChange:"), change)
}
func (o BKTouchLifecycleObserverObject) TouchDidAlwaysRouteToTargetIDClientTaskName(touch uint32, id unsafe.Pointer, name uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touch:didAlwaysRouteToTargetID:clientTaskName:"), touch, id, name)
}
func (o BKTouchLifecycleObserverObject) TouchDidHitTestToDestinationHostingChainIndex(touch uint32, destination objectivec.IObject, index int64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touch:didHitTestToDestination:hostingChainIndex:"), touch, destination, index)
}
func (o BKTouchLifecycleObserverObject) TouchDidMoveToPointEventMaskZMaxZ(touch uint32, point corefoundation.CGPoint, mask uint32, z float64, z2 float64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touch:didMoveToPoint:eventMask:z:maxZ:"), touch, point, mask, z, z2)
}
func (o BKTouchLifecycleObserverObject) TouchPathIndexDownAtPointEventMaskTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, type_ uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touch:pathIndex:downAtPoint:eventMask:transducerType:"), touch, index, point, mask, type_)
}
func (o BKTouchLifecycleObserverObject) TouchPathIndexRangeInAtPointEventMaskZMaxZTransducerType(touch uint32, index int64, point corefoundation.CGPoint, mask uint32, z float64, z2 float64, type_ uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touch:pathIndex:rangeInAtPoint:eventMask:z:maxZ:transducerType:"), touch, index, point, mask, z, z2, type_)
}
func (o BKTouchLifecycleObserverObject) TouchPathIndexRangeOutAtPoint(touch uint32, index int64, point corefoundation.CGPoint) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touch:pathIndex:rangeOutAtPoint:"), touch, index, point)
}
func (o BKTouchLifecycleObserverObject) TouchPathIndexUpAtPointDetached(touch uint32, index int64, point corefoundation.CGPoint, detached bool) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touch:pathIndex:upAtPoint:detached:"), touch, index, point, detached)
}
func (o BKTouchLifecycleObserverObject) TouchDidDetachDestinations(detach uint32, destinations objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touchDidDetach:destinations:"), detach, destinations)
}
func (o BKTouchLifecycleObserverObject) TouchDidFinishProcessingTouchCollection() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touchDidFinishProcessingTouchCollection"))
}
func (o BKTouchLifecycleObserverObject) TouchDidHIDCancel(hIDCancel uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touchDidHIDCancel:"), hIDCancel)
}
func (o BKTouchLifecycleObserverObject) TouchDidSoftCancel(cancel uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touchDidSoftCancel:"), cancel)
}
func (o BKTouchLifecycleObserverObject) TouchDidTransferDestination(transfer uint32, destination objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touchDidTransfer:destination:"), transfer, destination)
}
func (o BKTouchLifecycleObserverObject) TouchWillStartProcessingTouchCollection() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("touchWillStartProcessingTouchCollection"))
}
