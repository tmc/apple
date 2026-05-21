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
	rv := objc.Send[VZGenericMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a unique identifier for a virtual machine.
//
// # Overview
//
// Use the data representation in
// [VZGenericMachineIdentifier.DataRepresentation] to save the VM’s
// identifier. To restore a previously saved identifier use
// [VZGenericMachineIdentifier.InitWithDataRepresentation].
//
// # Creating a Machine Identifier
//
//   - [VZGenericMachineIdentifier.InitWithDataRepresentation]: Creates a new unique identifier for a VM with the provided data.
//
// # Getting Information About the Machine Identifier
//
//   - [VZGenericMachineIdentifier.DataRepresentation]: An opaque data representation of the VM’s identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier
type VZGenericMachineIdentifier struct {
	objectivec.Object
}

// VZGenericMachineIdentifierFromID constructs a [VZGenericMachineIdentifier] from an objc.ID.
//
// An object that represents a unique identifier for a virtual machine.
func VZGenericMachineIdentifierFromID(id objc.ID) VZGenericMachineIdentifier {
	return VZGenericMachineIdentifier{objectivec.Object{ID: id}}
}

// NOTE: VZGenericMachineIdentifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZGenericMachineIdentifier] class.
//
// # Creating a Machine Identifier
//
//   - [IVZGenericMachineIdentifier.InitWithDataRepresentation]: Creates a new unique identifier for a VM with the provided data.
//
// # Getting Information About the Machine Identifier
//
//   - [IVZGenericMachineIdentifier.DataRepresentation]: An opaque data representation of the VM’s identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier
type IVZGenericMachineIdentifier interface {
	objectivec.IObject

	// Topic: Creating a Machine Identifier

	// Creates a new unique identifier for a VM with the provided data.
	InitWithDataRepresentation(dataRepresentation foundation.NSData) VZGenericMachineIdentifier

	// Topic: Getting Information About the Machine Identifier

	// An opaque data representation of the VM’s identifier.
	DataRepresentation() foundation.NSData
}

// Init initializes the instance.
func (g VZGenericMachineIdentifier) Init() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGenericMachineIdentifier) Autorelease() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGenericMachineIdentifier creates a new VZGenericMachineIdentifier instance.
func NewVZGenericMachineIdentifier() VZGenericMachineIdentifier {
	class := getVZGenericMachineIdentifierClass()
	rv := objc.Send[VZGenericMachineIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new unique identifier for a VM with the provided data.
//
// dataRepresentation: A data object that describes the machine identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier/init(dataRepresentation:)
func NewGenericMachineIdentifierWithDataRepresentation(dataRepresentation foundation.NSData) VZGenericMachineIdentifier {
	instance := getVZGenericMachineIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	return VZGenericMachineIdentifierFromID(rv)
}

// Creates a new unique identifier for a VM with the provided data.
//
// dataRepresentation: A data object that describes the machine identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier/init(dataRepresentation:)
func (g VZGenericMachineIdentifier) InitWithDataRepresentation(dataRepresentation foundation.NSData) VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](g.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	return rv
}

// An opaque data representation of the VM’s identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier/dataRepresentation
func (g VZGenericMachineIdentifier) DataRepresentation() foundation.NSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dataRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}
