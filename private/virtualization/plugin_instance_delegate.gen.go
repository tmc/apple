// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PluginInstanceDelegate] class.
var (
	_PluginInstanceDelegateClass     PluginInstanceDelegateClass
	_PluginInstanceDelegateClassOnce sync.Once
)

func getPluginInstanceDelegateClass() PluginInstanceDelegateClass {
	_PluginInstanceDelegateClassOnce.Do(func() {
		_PluginInstanceDelegateClass = PluginInstanceDelegateClass{class: objc.GetClass("PluginInstanceDelegate")}
	})
	return _PluginInstanceDelegateClass
}

// GetPluginInstanceDelegateClass returns the class object for PluginInstanceDelegate.
func GetPluginInstanceDelegateClass() PluginInstanceDelegateClass {
	return getPluginInstanceDelegateClass()
}

type PluginInstanceDelegateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PluginInstanceDelegateClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PluginInstanceDelegateClass) Alloc() PluginInstanceDelegate {
	rv := objc.Send[PluginInstanceDelegate](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PluginInstanceDelegate.HandleConnectionError]
//   - [PluginInstanceDelegate.InvalidateConnection]
//   - [PluginInstanceDelegate.DebugDescription]
//   - [PluginInstanceDelegate.Description]
//   - [PluginInstanceDelegate.Hash]
//   - [PluginInstanceDelegate.Superclass]
type PluginInstanceDelegate struct {
	objectivec.Object
}

// PluginInstanceDelegateFromID constructs a [PluginInstanceDelegate] from an objc.ID.
func PluginInstanceDelegateFromID(id objc.ID) PluginInstanceDelegate {
	return PluginInstanceDelegate{objectivec.Object{ID: id}}
}

// Ensure PluginInstanceDelegate implements IPluginInstanceDelegate.
var _ IPluginInstanceDelegate = PluginInstanceDelegate{}

// An interface definition for the [PluginInstanceDelegate] class.
//
// # Methods
//
//   - [IPluginInstanceDelegate.HandleConnectionError]
//   - [IPluginInstanceDelegate.InvalidateConnection]
//   - [IPluginInstanceDelegate.DebugDescription]
//   - [IPluginInstanceDelegate.Description]
//   - [IPluginInstanceDelegate.Hash]
//   - [IPluginInstanceDelegate.Superclass]
type IPluginInstanceDelegate interface {
	objectivec.IObject

	// Topic: Methods

	HandleConnectionError(error_ objectivec.IObject)
	InvalidateConnection()
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (p PluginInstanceDelegate) Init() PluginInstanceDelegate {
	rv := objc.Send[PluginInstanceDelegate](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PluginInstanceDelegate) Autorelease() PluginInstanceDelegate {
	rv := objc.Send[PluginInstanceDelegate](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPluginInstanceDelegate creates a new PluginInstanceDelegate instance.
func NewPluginInstanceDelegate() PluginInstanceDelegate {
	class := getPluginInstanceDelegateClass()
	rv := objc.Send[PluginInstanceDelegate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PluginInstanceDelegate) HandleConnectionError(error_ objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("handleConnectionError:"), error_)
}
func (p PluginInstanceDelegate) InvalidateConnection() {
	objc.Send[objc.ID](p.ID, objc.Sel("invalidateConnection"))
}

func (p PluginInstanceDelegate) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (p PluginInstanceDelegate) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (p PluginInstanceDelegate) Hash() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("hash"))
	return rv
}
func (p PluginInstanceDelegate) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](p.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
