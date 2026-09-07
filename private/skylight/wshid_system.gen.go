// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSHIDSystem] class.
var (
	_WSHIDSystemClass     WSHIDSystemClass
	_WSHIDSystemClassOnce sync.Once
)

func getWSHIDSystemClass() WSHIDSystemClass {
	_WSHIDSystemClassOnce.Do(func() {
		_WSHIDSystemClass = WSHIDSystemClass{class: objc.GetClass("WSHIDSystem")}
	})
	return _WSHIDSystemClass
}

// GetWSHIDSystemClass returns the class object for WSHIDSystem.
func GetWSHIDSystemClass() WSHIDSystemClass {
	return getWSHIDSystemClass()
}

type WSHIDSystemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSHIDSystemClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSHIDSystemClass) Alloc() WSHIDSystem {
	rv := objc.SendIfResponds[WSHIDSystem](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSHIDSystem.IOHIDServicesMatching]
//   - [WSHIDSystem._init]
//   - [WSHIDSystem.HoistOnThread]
//   - [WSHIDSystem.HoistOnThreadForSessionID]
//   - [WSHIDSystem.RegisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon]
//   - [WSHIDSystem.SenderCache]
//   - [WSHIDSystem.UnregisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon]
//   - [WSHIDSystem.DebugDescription]
//   - [WSHIDSystem.Description]
//   - [WSHIDSystem.Hash]
//   - [WSHIDSystem.Superclass]
type WSHIDSystem struct {
	objectivec.Object
}

// WSHIDSystemFromID constructs a [WSHIDSystem] from an objc.ID.
func WSHIDSystemFromID(id objc.ID) WSHIDSystem {
	return WSHIDSystem{objectivec.Object{ID: id}}
}

// Ensure WSHIDSystem implements IWSHIDSystem.
var _ IWSHIDSystem = WSHIDSystem{}

// An interface definition for the [WSHIDSystem] class.
//
// # Methods
//
//   - [IWSHIDSystem.IOHIDServicesMatching]
//   - [IWSHIDSystem._init]
//   - [IWSHIDSystem.HoistOnThread]
//   - [IWSHIDSystem.HoistOnThreadForSessionID]
//   - [IWSHIDSystem.RegisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon]
//   - [IWSHIDSystem.SenderCache]
//   - [IWSHIDSystem.UnregisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon]
//   - [IWSHIDSystem.DebugDescription]
//   - [IWSHIDSystem.Description]
//   - [IWSHIDSystem.Hash]
//   - [IWSHIDSystem.Superclass]
type IWSHIDSystem interface {
	objectivec.IObject

	// Topic: Methods

	IOHIDServicesMatching(matching objectivec.IObject) objectivec.IObject
	_init() objectivec.IObject
	HoistOnThread(thread VoidHandler)
	HoistOnThreadForSessionID(thread VoidHandler, id uint32)
	RegisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer)
	SenderCache() unsafe.Pointer
	UnregisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSHIDSystem) Init() WSHIDSystem {
	rv := objc.SendIfResponds[WSHIDSystem](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSHIDSystem) Autorelease() WSHIDSystem {
	rv := objc.SendIfResponds[WSHIDSystem](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSHIDSystem creates a new WSHIDSystem instance.
func NewWSHIDSystem() WSHIDSystem {
	class := getWSHIDSystemClass()
	rv := objc.SendIfResponds[WSHIDSystem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSHIDSystem) IOHIDServicesMatching(matching objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("IOHIDServicesMatching:"), matching)
	return objectivec.Object{ID: rv}
}
func (w WSHIDSystem) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (w WSHIDSystem) HoistOnThread(thread VoidHandler) {
	_block0, _ := NewVoidBlock(thread)
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("hoistOnThread:"), _block0)
}
func (w WSHIDSystem) HoistOnThreadForSessionID(thread VoidHandler, id uint32) {
	_block0, _ := NewVoidBlock(thread)
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("hoistOnThread:forSessionID:"), _block0, id)
}
func (w WSHIDSystem) RegisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("registerIOHIDServicesCallback:matchingDictionary:target:refCon:"), _block0, dictionary, target, con)
}
func (w WSHIDSystem) UnregisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("unregisterIOHIDServicesCallback:matchingDictionary:target:refCon:"), _block0, dictionary, target, con)
}

func (_WSHIDSystemClass WSHIDSystemClass) SharedInstance() WSHIDSystem {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSHIDSystemClass.class), objc.Sel("sharedInstance"))
	return WSHIDSystemFromID(rv)
}

func (w WSHIDSystem) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDSystem) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDSystem) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSHIDSystem) SenderCache() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("senderCache"))
	return rv
}
func (w WSHIDSystem) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// HoistOnThreadSync is a synchronous wrapper around [WSHIDSystem.HoistOnThread].
// It blocks until the completion handler fires or the context is cancelled.
func (w WSHIDSystem) HoistOnThreadSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	w.HoistOnThread(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
