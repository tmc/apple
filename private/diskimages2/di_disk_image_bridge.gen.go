// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIDiskImageBridge] class.
var (
	_DIDiskImageBridgeClass     DIDiskImageBridgeClass
	_DIDiskImageBridgeClassOnce sync.Once
)

func getDIDiskImageBridgeClass() DIDiskImageBridgeClass {
	_DIDiskImageBridgeClassOnce.Do(func() {
		_DIDiskImageBridgeClass = DIDiskImageBridgeClass{class: objc.GetClass("DIDiskImageBridge")}
	})
	return _DIDiskImageBridgeClass
}

// GetDIDiskImageBridgeClass returns the class object for DIDiskImageBridge.
func GetDIDiskImageBridgeClass() DIDiskImageBridgeClass {
	return getDIDiskImageBridgeClass()
}

type DIDiskImageBridgeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIDiskImageBridgeClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIDiskImageBridgeClass) Alloc() DIDiskImageBridge {
	rv := objc.SendIfResponds[DIDiskImageBridge](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIDiskImageBridge.Handles]
//   - [DIDiskImageBridge.SetHandles]
//   - [DIDiskImageBridge.InitWithHandles]
type DIDiskImageBridge struct {
	objectivec.Object
}

// DIDiskImageBridgeFromID constructs a [DIDiskImageBridge] from an objc.ID.
func DIDiskImageBridgeFromID(id objc.ID) DIDiskImageBridge {
	return DIDiskImageBridge{objectivec.Object{ID: id}}
}

// Ensure DIDiskImageBridge implements IDIDiskImageBridge.
var _ IDIDiskImageBridge = DIDiskImageBridge{}

// An interface definition for the [DIDiskImageBridge] class.
//
// # Methods
//
//   - [IDIDiskImageBridge.Handles]
//   - [IDIDiskImageBridge.SetHandles]
//   - [IDIDiskImageBridge.InitWithHandles]
type IDIDiskImageBridge interface {
	objectivec.IObject

	// Topic: Methods

	Handles() foundation.INSArray
	SetHandles(value foundation.INSArray)
	InitWithHandles(handles objectivec.IObject) DIDiskImageBridge
}

// Init initializes the instance.
func (d DIDiskImageBridge) Init() DIDiskImageBridge {
	rv := objc.SendIfResponds[DIDiskImageBridge](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIDiskImageBridge) Autorelease() DIDiskImageBridge {
	rv := objc.SendIfResponds[DIDiskImageBridge](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIDiskImageBridge creates a new DIDiskImageBridge instance.
func NewDIDiskImageBridge() DIDiskImageBridge {
	class := getDIDiskImageBridgeClass()
	rv := objc.SendIfResponds[DIDiskImageBridge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIDiskImageBridgeWithHandles(handles objectivec.IObject) DIDiskImageBridge {
	instance := getDIDiskImageBridgeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHandles:"), handles)
	return DIDiskImageBridgeFromID(rv)
}

func (d DIDiskImageBridge) InitWithHandles(handles objectivec.IObject) DIDiskImageBridge {
	rv := objc.SendIfResponds[DIDiskImageBridge](d.ID, objc.Sel("initWithHandles:"), handles)
	return rv
}

func (d DIDiskImageBridge) Handles() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("handles"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DIDiskImageBridge) SetHandles(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setHandles:"), value)
}
