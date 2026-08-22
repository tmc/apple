// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCardSlotManager] class.
var (
	_TKSmartCardSlotManagerClass     TKSmartCardSlotManagerClass
	_TKSmartCardSlotManagerClassOnce sync.Once
)

func getTKSmartCardSlotManagerClass() TKSmartCardSlotManagerClass {
	_TKSmartCardSlotManagerClassOnce.Do(func() {
		_TKSmartCardSlotManagerClass = TKSmartCardSlotManagerClass{class: objc.GetClass("TKSmartCardSlotManager")}
	})
	return _TKSmartCardSlotManagerClass
}

// GetTKSmartCardSlotManagerClass returns the class object for TKSmartCardSlotManager.
func GetTKSmartCardSlotManagerClass() TKSmartCardSlotManagerClass {
	return getTKSmartCardSlotManagerClass()
}

type TKSmartCardSlotManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardSlotManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardSlotManagerClass) Alloc() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// An interface to all available smart card reader slots.
//
// # Overview
//
// Get a list of all known smart card reader slots in the system using the
// [TKSmartCardSlotManager.SlotNames] property, and access individual slots by
// name using the [TKSmartCardSlotManager.GetSlotWithNameReply] method.
//
// # Accessing Smart Card Slots
//
//   - [TKSmartCardSlotManager.SlotNames]: A list of identifiers for all the Smart Card reader slots available to the system.
//   - [TKSmartCardSlotManager.GetSlotWithNameReply]: Asynchronously calls a block with a Smart Card reader slot for a specified name.
//   - [TKSmartCardSlotManager.SlotNamed]: Returns the Smart Card slot with a given name.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager
type TKSmartCardSlotManager struct {
	objectivec.Object
}

// TKSmartCardSlotManagerFromID constructs a [TKSmartCardSlotManager] from an objc.ID.
//
// An interface to all available smart card reader slots.
func TKSmartCardSlotManagerFromID(id objc.ID) TKSmartCardSlotManager {
	return TKSmartCardSlotManager{objectivec.Object{ID: id}}
}

// NOTE: TKSmartCardSlotManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardSlotManager] class.
//
// # Accessing Smart Card Slots
//
//   - [ITKSmartCardSlotManager.SlotNames]: A list of identifiers for all the Smart Card reader slots available to the system.
//   - [ITKSmartCardSlotManager.GetSlotWithNameReply]: Asynchronously calls a block with a Smart Card reader slot for a specified name.
//   - [ITKSmartCardSlotManager.SlotNamed]: Returns the Smart Card slot with a given name.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager
type ITKSmartCardSlotManager interface {
	objectivec.IObject

	// Topic: Accessing Smart Card Slots

	// A list of identifiers for all the Smart Card reader slots available to the system.
	SlotNames() []string
	// Asynchronously calls a block with a Smart Card reader slot for a specified name.
	GetSlotWithNameReply(name string, reply TKSmartCardSlotHandler)
	// Returns the Smart Card slot with a given name.
	SlotNamed(name string) ITKSmartCardSlot
}

// Init initializes the instance.
func (t TKSmartCardSlotManager) Init() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardSlotManager) Autorelease() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlotManager creates a new TKSmartCardSlotManager instance.
func NewTKSmartCardSlotManager() TKSmartCardSlotManager {
	class := getTKSmartCardSlotManagerClass()
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Asynchronously calls a block with a Smart Card reader slot for a specified
// name.
//
// name: The name of the Smart Card reader slot.
//
// reply: slot: The Smart Card reader slot corresponding to the specified name. If no
// slot exists with that name, this argument is `nil`.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/getSlot(withName:reply:)
func (t TKSmartCardSlotManager) GetSlotWithNameReply(name string, reply TKSmartCardSlotHandler) {
	_block1, _ := NewTKSmartCardSlotBlock(reply)
	objc.Send[objc.ID](t.ID, objc.Sel("getSlotWithName:reply:"), objc.String(name), _block1)
}

// Returns the Smart Card slot with a given name.
//
// # Return Value
//
// The slot with the specified name, or `nil` if no slot with that name
// exists.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/slotNamed(_:)
func (t TKSmartCardSlotManager) SlotNamed(name string) ITKSmartCardSlot {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("slotNamed:"), objc.String(name))
	return TKSmartCardSlotFromID(rv)
}

// A list of identifiers for all the Smart Card reader slots available to the
// system.
//
// # Discussion
//
// Use Key-Value Observing on this property to be notified for changes to
// available Smart Card reader slots. For more information, see [Key-Value
// Observing Programming Guide].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/slotNames
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (t TKSmartCardSlotManager) SlotNames() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("slotNames"))
	return objc.ConvertSliceToStrings(rv)
}

// The shared singleton Smart Card reader slot manager.
//
// # Discussion
//
// This method returns `nil` unless the [com.apple.security.smartcard]
// entitlement is enabled.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/default
//
// [com.apple.security.smartcard]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.security.smartcard
func (_TKSmartCardSlotManagerClass TKSmartCardSlotManagerClass) DefaultManager() TKSmartCardSlotManager {
	rv := objc.Send[objc.ID](objc.ID(_TKSmartCardSlotManagerClass.class), objc.Sel("defaultManager"))
	return TKSmartCardSlotManagerFromID(objc.ID(rv))
}

// GetSlotWithNameReplySync is a synchronous wrapper around [TKSmartCardSlotManager.GetSlotWithNameReply].
// It blocks until the completion handler fires or the context is cancelled.
func (t TKSmartCardSlotManager) GetSlotWithNameReplySync(ctx context.Context, name string) (*TKSmartCardSlot, error) {
	done := make(chan *TKSmartCardSlot, 1)
	t.GetSlotWithNameReply(name, func(val *TKSmartCardSlot) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
