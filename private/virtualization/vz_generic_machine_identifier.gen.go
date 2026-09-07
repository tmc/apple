// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGenericMachineIdentifier] class.
var (
	_VZGenericMachineIdentifierClass     VZGenericMachineIdentifierClass
	_VZGenericMachineIdentifierClassOnce sync.Once
)

func getVZGenericMachineIdentifierClass() VZGenericMachineIdentifierClass {
	_VZGenericMachineIdentifierClassOnce.Do(func() {
		_VZGenericMachineIdentifierClass = VZGenericMachineIdentifierClass{class: objc.GetClass("VZGenericMachineIdentifier")}
	})
	return _VZGenericMachineIdentifierClass
}

// GetVZGenericMachineIdentifierClass returns the class object for VZGenericMachineIdentifier.
func GetVZGenericMachineIdentifierClass() VZGenericMachineIdentifierClass {
	return getVZGenericMachineIdentifierClass()
}

type VZGenericMachineIdentifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGenericMachineIdentifierClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGenericMachineIdentifierClass) Alloc() VZGenericMachineIdentifier {
	rv := objc.SendIfResponds[VZGenericMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGenericMachineIdentifier.DebugDescription]
//   - [VZGenericMachineIdentifier.Description]
//   - [VZGenericMachineIdentifier.Hash]
//   - [VZGenericMachineIdentifier.Superclass]
type VZGenericMachineIdentifier struct {
	objectivec.Object
}

// VZGenericMachineIdentifierFromID constructs a [VZGenericMachineIdentifier] from an objc.ID.
func VZGenericMachineIdentifierFromID(id objc.ID) VZGenericMachineIdentifier {
	return VZGenericMachineIdentifier{objectivec.Object{ID: id}}
}

// Ensure VZGenericMachineIdentifier implements IVZGenericMachineIdentifier.
var _ IVZGenericMachineIdentifier = VZGenericMachineIdentifier{}

// An interface definition for the [VZGenericMachineIdentifier] class.
//
// # Methods
//
//   - [IVZGenericMachineIdentifier.DebugDescription]
//   - [IVZGenericMachineIdentifier.Description]
//   - [IVZGenericMachineIdentifier.Hash]
//   - [IVZGenericMachineIdentifier.Superclass]
type IVZGenericMachineIdentifier interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZGenericMachineIdentifier) Init() VZGenericMachineIdentifier {
	rv := objc.SendIfResponds[VZGenericMachineIdentifier](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGenericMachineIdentifier) Autorelease() VZGenericMachineIdentifier {
	rv := objc.SendIfResponds[VZGenericMachineIdentifier](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGenericMachineIdentifier creates a new VZGenericMachineIdentifier instance.
func NewVZGenericMachineIdentifier() VZGenericMachineIdentifier {
	class := getVZGenericMachineIdentifierClass()
	rv := objc.SendIfResponds[VZGenericMachineIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZGenericMachineIdentifier) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZGenericMachineIdentifier) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZGenericMachineIdentifier) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZGenericMachineIdentifier) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
