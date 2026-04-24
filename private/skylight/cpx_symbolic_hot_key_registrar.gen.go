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

// The class instance for the [CPXSymbolicHotKeyRegistrar] class.
var (
	_CPXSymbolicHotKeyRegistrarClass     CPXSymbolicHotKeyRegistrarClass
	_CPXSymbolicHotKeyRegistrarClassOnce sync.Once
)

func getCPXSymbolicHotKeyRegistrarClass() CPXSymbolicHotKeyRegistrarClass {
	_CPXSymbolicHotKeyRegistrarClassOnce.Do(func() {
		_CPXSymbolicHotKeyRegistrarClass = CPXSymbolicHotKeyRegistrarClass{class: objc.GetClass("_CPXSymbolicHotKeyRegistrar")}
	})
	return _CPXSymbolicHotKeyRegistrarClass
}

// GetCPXSymbolicHotKeyRegistrarClass returns the class object for _CPXSymbolicHotKeyRegistrar.
func GetCPXSymbolicHotKeyRegistrarClass() CPXSymbolicHotKeyRegistrarClass {
	return getCPXSymbolicHotKeyRegistrarClass()
}

type CPXSymbolicHotKeyRegistrarClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXSymbolicHotKeyRegistrarClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXSymbolicHotKeyRegistrarClass) Alloc() CPXSymbolicHotKeyRegistrar {
	rv := objc.Send[CPXSymbolicHotKeyRegistrar](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXSymbolicHotKeyRegistrar.GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers]
//   - [CPXSymbolicHotKeyRegistrar.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc]
//   - [CPXSymbolicHotKeyRegistrar.UnregisterHotKeyConnectionHotKeyID]
//   - [CPXSymbolicHotKeyRegistrar.DebugDescription]
//   - [CPXSymbolicHotKeyRegistrar.Description]
//   - [CPXSymbolicHotKeyRegistrar.Hash]
//   - [CPXSymbolicHotKeyRegistrar.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar
type CPXSymbolicHotKeyRegistrar struct {
	objectivec.Object
}

// CPXSymbolicHotKeyRegistrarFromID constructs a [CPXSymbolicHotKeyRegistrar] from an objc.ID.
func CPXSymbolicHotKeyRegistrarFromID(id objc.ID) CPXSymbolicHotKeyRegistrar {
	return CPXSymbolicHotKeyRegistrar{objectivec.Object{ID: id}}
}

// Ensure CPXSymbolicHotKeyRegistrar implements ICPXSymbolicHotKeyRegistrar.
var _ ICPXSymbolicHotKeyRegistrar = CPXSymbolicHotKeyRegistrar{}

// An interface definition for the [CPXSymbolicHotKeyRegistrar] class.
//
// # Methods
//
//   - [ICPXSymbolicHotKeyRegistrar.GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers]
//   - [ICPXSymbolicHotKeyRegistrar.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc]
//   - [ICPXSymbolicHotKeyRegistrar.UnregisterHotKeyConnectionHotKeyID]
//   - [ICPXSymbolicHotKeyRegistrar.DebugDescription]
//   - [ICPXSymbolicHotKeyRegistrar.Description]
//   - [ICPXSymbolicHotKeyRegistrar.Hash]
//   - [ICPXSymbolicHotKeyRegistrar.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar
type ICPXSymbolicHotKeyRegistrar interface {
	objectivec.IObject

	// Topic: Methods

	GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers(value uint32, trigger unsafe.Pointer, char unsafe.Pointer, key unsafe.Pointer, modifiers unsafe.Pointer) int
	RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc(connection unsafe.Pointer, id uint64, key uint32, option uint32, func_ VoidHandler) int
	UnregisterHotKeyConnectionHotKeyID(connection unsafe.Pointer, id uint64) int
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXSymbolicHotKeyRegistrar) Init() CPXSymbolicHotKeyRegistrar {
	rv := objc.Send[CPXSymbolicHotKeyRegistrar](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXSymbolicHotKeyRegistrar) Autorelease() CPXSymbolicHotKeyRegistrar {
	rv := objc.Send[CPXSymbolicHotKeyRegistrar](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXSymbolicHotKeyRegistrar creates a new CPXSymbolicHotKeyRegistrar instance.
func NewCPXSymbolicHotKeyRegistrar() CPXSymbolicHotKeyRegistrar {
	class := getCPXSymbolicHotKeyRegistrarClass()
	rv := objc.Send[CPXSymbolicHotKeyRegistrar](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar/getSymbolicHotKeyValue:outTrigger:outKeyChar:outVirtualKey:outModifiers:
func (c CPXSymbolicHotKeyRegistrar) GetSymbolicHotKeyValueOutTriggerOutKeyCharOutVirtualKeyOutModifiers(value uint32, trigger unsafe.Pointer, char unsafe.Pointer, key unsafe.Pointer, modifiers unsafe.Pointer) int {
	rv := objc.Send[int](c.ID, objc.Sel("getSymbolicHotKeyValue:outTrigger:outKeyChar:outVirtualKey:outModifiers:"), value, trigger, char, key, modifiers)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar/registerSymbolicHotKeyConnection:hotKeyID:symbolicHotKey:option:callbackFunc:
func (c CPXSymbolicHotKeyRegistrar) RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc(connection unsafe.Pointer, id uint64, key uint32, option uint32, func_ VoidHandler) int {
	_block4, _ := NewVoidBlock(func_)
	rv := objc.Send[int](c.ID, objc.Sel("registerSymbolicHotKeyConnection:hotKeyID:symbolicHotKey:option:callbackFunc:"), connection, id, key, option, _block4)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar/unregisterHotKeyConnection:hotKeyID:
func (c CPXSymbolicHotKeyRegistrar) UnregisterHotKeyConnectionHotKeyID(connection unsafe.Pointer, id uint64) int {
	rv := objc.Send[int](c.ID, objc.Sel("unregisterHotKeyConnection:hotKeyID:"), connection, id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar/debugDescription
func (c CPXSymbolicHotKeyRegistrar) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar/description
func (c CPXSymbolicHotKeyRegistrar) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar/hash
func (c CPXSymbolicHotKeyRegistrar) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXSymbolicHotKeyRegistrar/superclass
func (c CPXSymbolicHotKeyRegistrar) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

// RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFuncSync is a synchronous wrapper around [CPXSymbolicHotKeyRegistrar.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc].
// It blocks until the completion handler fires or the context is cancelled.
func (c CPXSymbolicHotKeyRegistrar) RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFuncSync(ctx context.Context, connection unsafe.Pointer, id uint64, key uint32, option uint32) error {
	done := make(chan struct{}, 1)
	c.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc(connection, id, key, option, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
