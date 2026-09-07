// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BKHitTestDispatchingShim] class.
var (
	_BKHitTestDispatchingShimClass     BKHitTestDispatchingShimClass
	_BKHitTestDispatchingShimClassOnce sync.Once
)

func getBKHitTestDispatchingShimClass() BKHitTestDispatchingShimClass {
	_BKHitTestDispatchingShimClassOnce.Do(func() {
		_BKHitTestDispatchingShimClass = BKHitTestDispatchingShimClass{class: objc.GetClass("_BKHitTestDispatchingShim")}
	})
	return _BKHitTestDispatchingShimClass
}

// GetBKHitTestDispatchingShimClass returns the class object for _BKHitTestDispatchingShim.
func GetBKHitTestDispatchingShimClass() BKHitTestDispatchingShimClass {
	return getBKHitTestDispatchingShimClass()
}

type BKHitTestDispatchingShimClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BKHitTestDispatchingShimClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BKHitTestDispatchingShimClass) Alloc() BKHitTestDispatchingShim {
	rv := objc.SendIfResponds[BKHitTestDispatchingShim](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BKHitTestDispatchingShim.SendEventForTargetIDToClientConnectionIdentifier]
//   - [BKHitTestDispatchingShim.SendEventToClientTaskNameToTargetID]
type BKHitTestDispatchingShim struct {
	objectivec.Object
}

// BKHitTestDispatchingShimFromID constructs a [BKHitTestDispatchingShim] from an objc.ID.
func BKHitTestDispatchingShimFromID(id objc.ID) BKHitTestDispatchingShim {
	return BKHitTestDispatchingShim{objectivec.Object{ID: id}}
}

// Ensure BKHitTestDispatchingShim implements IBKHitTestDispatchingShim.
var _ IBKHitTestDispatchingShim = BKHitTestDispatchingShim{}

// An interface definition for the [BKHitTestDispatchingShim] class.
//
// # Methods
//
//   - [IBKHitTestDispatchingShim.SendEventForTargetIDToClientConnectionIdentifier]
//   - [IBKHitTestDispatchingShim.SendEventToClientTaskNameToTargetID]
type IBKHitTestDispatchingShim interface {
	objectivec.IObject

	// Topic: Methods

	SendEventForTargetIDToClientConnectionIdentifier(event uintptr, id unsafe.Pointer, identifier uint64)
	SendEventToClientTaskNameToTargetID(event uintptr, name uint32, id unsafe.Pointer)
}

// Init initializes the instance.
func (b BKHitTestDispatchingShim) Init() BKHitTestDispatchingShim {
	rv := objc.SendIfResponds[BKHitTestDispatchingShim](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BKHitTestDispatchingShim) Autorelease() BKHitTestDispatchingShim {
	rv := objc.SendIfResponds[BKHitTestDispatchingShim](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBKHitTestDispatchingShim creates a new BKHitTestDispatchingShim instance.
func NewBKHitTestDispatchingShim() BKHitTestDispatchingShim {
	class := getBKHitTestDispatchingShimClass()
	rv := objc.SendIfResponds[BKHitTestDispatchingShim](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (b BKHitTestDispatchingShim) SendEventForTargetIDToClientConnectionIdentifier(event uintptr, id unsafe.Pointer, identifier uint64) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("sendEvent:forTargetID:toClientConnectionIdentifier:"), event, id, identifier)
}
func (b BKHitTestDispatchingShim) SendEventToClientTaskNameToTargetID(event uintptr, name uint32, id unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("sendEvent:toClientTaskName:toTargetID:"), event, name, id)
}
