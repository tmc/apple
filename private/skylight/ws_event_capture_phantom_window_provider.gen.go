// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventCapturePhantomWindowProvider] class.
var (
	_WSEventCapturePhantomWindowProviderClass     WSEventCapturePhantomWindowProviderClass
	_WSEventCapturePhantomWindowProviderClassOnce sync.Once
)

func getWSEventCapturePhantomWindowProviderClass() WSEventCapturePhantomWindowProviderClass {
	_WSEventCapturePhantomWindowProviderClassOnce.Do(func() {
		_WSEventCapturePhantomWindowProviderClass = WSEventCapturePhantomWindowProviderClass{class: objc.GetClass("WSEventCapturePhantomWindowProvider")}
	})
	return _WSEventCapturePhantomWindowProviderClass
}

// GetWSEventCapturePhantomWindowProviderClass returns the class object for WSEventCapturePhantomWindowProvider.
func GetWSEventCapturePhantomWindowProviderClass() WSEventCapturePhantomWindowProviderClass {
	return getWSEventCapturePhantomWindowProviderClass()
}

type WSEventCapturePhantomWindowProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventCapturePhantomWindowProviderClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventCapturePhantomWindowProviderClass) Alloc() WSEventCapturePhantomWindowProvider {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowProvider](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSEventCapturePhantomWindowProvider.CreatePhantomWindowForConnection]
//   - [WSEventCapturePhantomWindowProvider.ReleasePhantomWindowID]
//   - [WSEventCapturePhantomWindowProvider.StartGestureLifetimeTrackingForPhantomWindowIDConnectionManager]
//   - [WSEventCapturePhantomWindowProvider.DebugDescription]
//   - [WSEventCapturePhantomWindowProvider.Description]
//   - [WSEventCapturePhantomWindowProvider.Hash]
//   - [WSEventCapturePhantomWindowProvider.Superclass]
type WSEventCapturePhantomWindowProvider struct {
	objectivec.Object
}

// WSEventCapturePhantomWindowProviderFromID constructs a [WSEventCapturePhantomWindowProvider] from an objc.ID.
func WSEventCapturePhantomWindowProviderFromID(id objc.ID) WSEventCapturePhantomWindowProvider {
	return WSEventCapturePhantomWindowProvider{objectivec.Object{ID: id}}
}

// Ensure WSEventCapturePhantomWindowProvider implements IWSEventCapturePhantomWindowProvider.
var _ IWSEventCapturePhantomWindowProvider = WSEventCapturePhantomWindowProvider{}

// An interface definition for the [WSEventCapturePhantomWindowProvider] class.
//
// # Methods
//
//   - [IWSEventCapturePhantomWindowProvider.CreatePhantomWindowForConnection]
//   - [IWSEventCapturePhantomWindowProvider.ReleasePhantomWindowID]
//   - [IWSEventCapturePhantomWindowProvider.StartGestureLifetimeTrackingForPhantomWindowIDConnectionManager]
//   - [IWSEventCapturePhantomWindowProvider.DebugDescription]
//   - [IWSEventCapturePhantomWindowProvider.Description]
//   - [IWSEventCapturePhantomWindowProvider.Hash]
//   - [IWSEventCapturePhantomWindowProvider.Superclass]
type IWSEventCapturePhantomWindowProvider interface {
	objectivec.IObject

	// Topic: Methods

	CreatePhantomWindowForConnection(connection *CGXConnection) uint32
	ReleasePhantomWindowID(id uint32)
	StartGestureLifetimeTrackingForPhantomWindowIDConnectionManager(id uint32, connection *CGXConnection, manager objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSEventCapturePhantomWindowProvider) Init() WSEventCapturePhantomWindowProvider {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowProvider](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventCapturePhantomWindowProvider) Autorelease() WSEventCapturePhantomWindowProvider {
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowProvider](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventCapturePhantomWindowProvider creates a new WSEventCapturePhantomWindowProvider instance.
func NewWSEventCapturePhantomWindowProvider() WSEventCapturePhantomWindowProvider {
	class := getWSEventCapturePhantomWindowProviderClass()
	rv := objc.SendIfResponds[WSEventCapturePhantomWindowProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSEventCapturePhantomWindowProvider) CreatePhantomWindowForConnection(connection *CGXConnection) uint32 {
	rv := objc.SendIfResponds[uint32](w.ID, objc.Sel("createPhantomWindowForConnection:"), unsafe.Pointer(connection))
	return rv
}
func (w WSEventCapturePhantomWindowProvider) ReleasePhantomWindowID(id uint32) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("releasePhantomWindowID:"), id)
}
func (w WSEventCapturePhantomWindowProvider) StartGestureLifetimeTrackingForPhantomWindowIDConnectionManager(id uint32, connection *CGXConnection, manager objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("startGestureLifetimeTrackingForPhantomWindowID:connection:manager:"), id, unsafe.Pointer(connection), manager)
}

func (w WSEventCapturePhantomWindowProvider) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventCapturePhantomWindowProvider) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventCapturePhantomWindowProvider) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSEventCapturePhantomWindowProvider) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
