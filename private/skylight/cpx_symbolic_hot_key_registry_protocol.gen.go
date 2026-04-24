// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXSymbolicHotKeyRegistry protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXSymbolicHotKeyRegistry
type CPXSymbolicHotKeyRegistry interface {
	objectivec.IObject

	// GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSymbolicHotKeyRegistry/getSymbolicHotKeyValue:outTrigger:outKeyChar:outVirtualKey:outModifiers:
	GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers(value uint32, trigger unsafe.Pointer, char unsafe.Pointer, key unsafe.Pointer, modifiers unsafe.Pointer) int

	// RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSymbolicHotKeyRegistry/registerSymbolicHotKeyConnection:hotKeyID:symbolicHotKey:option:callbackFunc:
	RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc(connection unsafe.Pointer, id uint64, key uint32, option uint32, func_ VoidHandler) int

	// UnregisterHotKeyConnectionHotKeyID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXSymbolicHotKeyRegistry/unregisterHotKeyConnection:hotKeyID:
	UnregisterHotKeyConnectionHotKeyID(connection unsafe.Pointer, id uint64) int
}

// CPXSymbolicHotKeyRegistryObject wraps an existing Objective-C object that conforms to the CPXSymbolicHotKeyRegistry protocol.
type CPXSymbolicHotKeyRegistryObject struct {
	objectivec.Object
}

func (o CPXSymbolicHotKeyRegistryObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXSymbolicHotKeyRegistryObjectFromID constructs a [CPXSymbolicHotKeyRegistryObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXSymbolicHotKeyRegistryObjectFromID(id objc.ID) CPXSymbolicHotKeyRegistryObject {
	return CPXSymbolicHotKeyRegistryObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSymbolicHotKeyRegistry/getSymbolicHotKeyValue:outTrigger:outKeyChar:outVirtualKey:outModifiers:
func (o CPXSymbolicHotKeyRegistryObject) GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers(value uint32, trigger unsafe.Pointer, char unsafe.Pointer, key unsafe.Pointer, modifiers unsafe.Pointer) int {
	rv := objc.Send[int](o.ID, objc.Sel("getSymbolicHotKeyValue:outTrigger:outKeyChar:outVirtualKey:outModifiers:"), value, trigger, char, key, modifiers)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSymbolicHotKeyRegistry/registerSymbolicHotKeyConnection:hotKeyID:symbolicHotKey:option:callbackFunc:
func (o CPXSymbolicHotKeyRegistryObject) RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc(connection unsafe.Pointer, id uint64, key uint32, option uint32, func_ VoidHandler) int {
	rv := objc.Send[int](o.ID, objc.Sel("registerSymbolicHotKeyConnection:hotKeyID:symbolicHotKey:option:callbackFunc:"), connection, id, key, option, func_)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSymbolicHotKeyRegistry/unregisterHotKeyConnection:hotKeyID:
func (o CPXSymbolicHotKeyRegistryObject) UnregisterHotKeyConnectionHotKeyID(connection unsafe.Pointer, id uint64) int {
	rv := objc.Send[int](o.ID, objc.Sel("unregisterHotKeyConnection:hotKeyID:"), connection, id)
	return rv
}
