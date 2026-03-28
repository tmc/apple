// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBridgedNetworkInterface] class.
var (
	_VZBridgedNetworkInterfaceClass     VZBridgedNetworkInterfaceClass
	_VZBridgedNetworkInterfaceClassOnce sync.Once
)

func getVZBridgedNetworkInterfaceClass() VZBridgedNetworkInterfaceClass {
	_VZBridgedNetworkInterfaceClassOnce.Do(func() {
		_VZBridgedNetworkInterfaceClass = VZBridgedNetworkInterfaceClass{class: objc.GetClass("VZBridgedNetworkInterface")}
	})
	return _VZBridgedNetworkInterfaceClass
}

// GetVZBridgedNetworkInterfaceClass returns the class object for VZBridgedNetworkInterface.
func GetVZBridgedNetworkInterfaceClass() VZBridgedNetworkInterfaceClass {
	return getVZBridgedNetworkInterfaceClass()
}

type VZBridgedNetworkInterfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBridgedNetworkInterfaceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBridgedNetworkInterfaceClass) Alloc() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface
type VZBridgedNetworkInterface struct {
	objectivec.Object
}

// VZBridgedNetworkInterfaceFromID constructs a [VZBridgedNetworkInterface] from an objc.ID.
func VZBridgedNetworkInterfaceFromID(id objc.ID) VZBridgedNetworkInterface {
	return VZBridgedNetworkInterface{objectivec.Object{ID: id}}
}
// Ensure VZBridgedNetworkInterface implements IVZBridgedNetworkInterface.
var _ IVZBridgedNetworkInterface = VZBridgedNetworkInterface{}

// An interface definition for the [VZBridgedNetworkInterface] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface
type IVZBridgedNetworkInterface interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b VZBridgedNetworkInterface) Init() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b VZBridgedNetworkInterface) Autorelease() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBridgedNetworkInterface creates a new VZBridgedNetworkInterface instance.
func NewVZBridgedNetworkInterface() VZBridgedNetworkInterface {
	class := getVZBridgedNetworkInterfaceClass()
	rv := objc.Send[VZBridgedNetworkInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/_interfaceWithIdentifier:error:
func (_VZBridgedNetworkInterfaceClass VZBridgedNetworkInterfaceClass) _interfaceWithIdentifierError(identifier objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_VZBridgedNetworkInterfaceClass.class), objc.Sel("_interfaceWithIdentifier:error:"), identifier, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// InterfaceWithIdentifierError is an exported wrapper for the private method _interfaceWithIdentifierError.
func (_VZBridgedNetworkInterfaceClass VZBridgedNetworkInterfaceClass) InterfaceWithIdentifierError(identifier objectivec.IObject) (objectivec.IObject, error) {
	return _VZBridgedNetworkInterfaceClass._interfaceWithIdentifierError(identifier)
}

