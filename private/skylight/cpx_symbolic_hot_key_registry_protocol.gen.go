// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXSymbolicHotKeyRegistry protocol.
type CPXSymbolicHotKeyRegistry interface {
	objectivec.IObject

	// GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers protocol.
	GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers(value uint32, trigger *uint32, char *uint16, key *uint16, modifiers *uint32) int

	// RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc protocol.
	RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc(connection *CGXConnection, id uint64, key uint32, option uint32, func_ VoidHandler) int

	// UnregisterHotKeyConnectionHotKeyID protocol.
	UnregisterHotKeyConnectionHotKeyID(connection *CGXConnection, id uint64) int
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

func (o CPXSymbolicHotKeyRegistryObject) GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers(value uint32, trigger *uint32, char *uint16, key *uint16, modifiers *uint32) int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("getSymbolicHotKeyValue:outTrigger:outKeyChar:outVirtualKey:outModifiers:"), value, unsafe.Pointer(trigger), unsafe.Pointer(char), unsafe.Pointer(key), unsafe.Pointer(modifiers))
	return rv
}
func (o CPXSymbolicHotKeyRegistryObject) RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc(connection *CGXConnection, id uint64, key uint32, option uint32, func_ VoidHandler) int {
	_block4, _cleanup4 := NewVoidBlock(func_)
	defer _cleanup4()
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("registerSymbolicHotKeyConnection:hotKeyID:symbolicHotKey:option:callbackFunc:"), connection, id, key, option, objc.ID(_block4))
	return rv
}
func (o CPXSymbolicHotKeyRegistryObject) UnregisterHotKeyConnectionHotKeyID(connection *CGXConnection, id uint64) int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("unregisterHotKeyConnection:hotKeyID:"), unsafe.Pointer(connection), id)
	return rv
}
