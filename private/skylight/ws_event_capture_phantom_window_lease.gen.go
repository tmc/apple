// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventCapturePhantomWindowLease] class.
var (
	_WSEventCapturePhantomWindowLeaseClass     WSEventCapturePhantomWindowLeaseClass
	_WSEventCapturePhantomWindowLeaseClassOnce sync.Once
)

func getWSEventCapturePhantomWindowLeaseClass() WSEventCapturePhantomWindowLeaseClass {
	_WSEventCapturePhantomWindowLeaseClassOnce.Do(func() {
		_WSEventCapturePhantomWindowLeaseClass = WSEventCapturePhantomWindowLeaseClass{class: objc.GetClass("WSEventCapturePhantomWindowLease")}
	})
	return _WSEventCapturePhantomWindowLeaseClass
}

// GetWSEventCapturePhantomWindowLeaseClass returns the class object for WSEventCapturePhantomWindowLease.
func GetWSEventCapturePhantomWindowLeaseClass() WSEventCapturePhantomWindowLeaseClass {
	return getWSEventCapturePhantomWindowLeaseClass()
}

type WSEventCapturePhantomWindowLeaseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventCapturePhantomWindowLeaseClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventCapturePhantomWindowLeaseClass) Alloc() WSEventCapturePhantomWindowLease {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowLease](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSEventCapturePhantomWindowLease.Invalidate]
//   - [WSEventCapturePhantomWindowLease.Underlying]
//   - [WSEventCapturePhantomWindowLease.SetUnderlying]
//   - [WSEventCapturePhantomWindowLease.WindowID]
//   - [WSEventCapturePhantomWindowLease.InitWithWindowIDUnderlying]
//   - [WSEventCapturePhantomWindowLease.DebugDescription]
//   - [WSEventCapturePhantomWindowLease.Description]
//   - [WSEventCapturePhantomWindowLease.Hash]
//   - [WSEventCapturePhantomWindowLease.Superclass]
type WSEventCapturePhantomWindowLease struct {
	objectivec.Object
}

// WSEventCapturePhantomWindowLeaseFromID constructs a [WSEventCapturePhantomWindowLease] from an objc.ID.
func WSEventCapturePhantomWindowLeaseFromID(id objc.ID) WSEventCapturePhantomWindowLease {
	return WSEventCapturePhantomWindowLease{objectivec.Object{ID: id}}
}

// Ensure WSEventCapturePhantomWindowLease implements IWSEventCapturePhantomWindowLease.
var _ IWSEventCapturePhantomWindowLease = WSEventCapturePhantomWindowLease{}

// An interface definition for the [WSEventCapturePhantomWindowLease] class.
//
// # Methods
//
//   - [IWSEventCapturePhantomWindowLease.Invalidate]
//   - [IWSEventCapturePhantomWindowLease.Underlying]
//   - [IWSEventCapturePhantomWindowLease.SetUnderlying]
//   - [IWSEventCapturePhantomWindowLease.WindowID]
//   - [IWSEventCapturePhantomWindowLease.InitWithWindowIDUnderlying]
//   - [IWSEventCapturePhantomWindowLease.DebugDescription]
//   - [IWSEventCapturePhantomWindowLease.Description]
//   - [IWSEventCapturePhantomWindowLease.Hash]
//   - [IWSEventCapturePhantomWindowLease.Superclass]
type IWSEventCapturePhantomWindowLease interface {
	objectivec.IObject

	// Topic: Methods

	Invalidate()
	Underlying() unsafe.Pointer
	SetUnderlying(value unsafe.Pointer)
	WindowID() uint32
	InitWithWindowIDUnderlying(id uint32, underlying objectivec.IObject) WSEventCapturePhantomWindowLease
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSEventCapturePhantomWindowLease) Init() WSEventCapturePhantomWindowLease {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowLease](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventCapturePhantomWindowLease) Autorelease() WSEventCapturePhantomWindowLease {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowLease](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventCapturePhantomWindowLease creates a new WSEventCapturePhantomWindowLease instance.
func NewWSEventCapturePhantomWindowLease() WSEventCapturePhantomWindowLease {
	class := getWSEventCapturePhantomWindowLeaseClass()
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowLease](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSEventCapturePhantomWindowLeaseWithWindowIDUnderlying(id uint32, underlying objectivec.IObject) WSEventCapturePhantomWindowLease {
	instance := getWSEventCapturePhantomWindowLeaseClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithWindowID:underlying:"), id, underlying)
	return WSEventCapturePhantomWindowLeaseFromID(rv)
}

func (w WSEventCapturePhantomWindowLease) Invalidate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("invalidate"))
}
func (w WSEventCapturePhantomWindowLease) InitWithWindowIDUnderlying(id uint32, underlying objectivec.IObject) WSEventCapturePhantomWindowLease {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowLease](w.ID, objc.Sel("initWithWindowID:underlying:"), id, underlying)
	return rv
}

func (w WSEventCapturePhantomWindowLease) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventCapturePhantomWindowLease) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventCapturePhantomWindowLease) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSEventCapturePhantomWindowLease) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (w WSEventCapturePhantomWindowLease) Underlying() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("underlying"))
	return rv
}
func (w WSEventCapturePhantomWindowLease) SetUnderlying(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setUnderlying:"), value)
}
func (w WSEventCapturePhantomWindowLease) WindowID() uint32 {
	rv := objc.SendIfResponds[uint32](w.ID, objc.Sel("windowID"))
	return rv
}
