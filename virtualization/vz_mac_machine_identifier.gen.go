// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacMachineIdentifier] class.
var (
	_VZMacMachineIdentifierClass     VZMacMachineIdentifierClass
	_VZMacMachineIdentifierClassOnce sync.Once
)

func getVZMacMachineIdentifierClass() VZMacMachineIdentifierClass {
	_VZMacMachineIdentifierClassOnce.Do(func() {
		_VZMacMachineIdentifierClass = VZMacMachineIdentifierClass{class: objc.GetClass("VZMacMachineIdentifier")}
	})
	return _VZMacMachineIdentifierClass
}

// GetVZMacMachineIdentifierClass returns the class object for VZMacMachineIdentifier.
func GetVZMacMachineIdentifierClass() VZMacMachineIdentifierClass {
	return getVZMacMachineIdentifierClass()
}

type VZMacMachineIdentifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacMachineIdentifierClass) Alloc() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A unique identifier for a VM.
//
// # Overview
// 
// This value uniquely identifies a virtual Mac hardware instance. Two VMs
// running concurrently shouldn’t use the same identifier.
// 
// When serializing the VM to disk, you can preserve the identifier in a
// binary representation by serializing the data in the
// [VZMacMachineIdentifier].[VZMacMachineIdentifier.DataRepresentation] property. Conversely, you can
// recreate the identifier with [VZMacMachineIdentifier.InitWithDataRepresentation] from the binary
// representation.
// 
// You can compare the contents of two identifiers with [isEqual(to:)].
//
// [isEqual(to:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isEqual(to:)
//
// # Creating a machine identifier
//
//   - [VZMacMachineIdentifier.InitWithDataRepresentation]: Create a machine identifier described by the specified data representation.
//
// # Machine data representation
//
//   - [VZMacMachineIdentifier.DataRepresentation]: Returns the opaque data representation of the machine identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier
type VZMacMachineIdentifier struct {
	objectivec.Object
}

// VZMacMachineIdentifierFromID constructs a [VZMacMachineIdentifier] from an objc.ID.
//
// A unique identifier for a VM.
func VZMacMachineIdentifierFromID(id objc.ID) VZMacMachineIdentifier {
	return VZMacMachineIdentifier{objectivec.Object{ID: id}}
}
// NOTE: VZMacMachineIdentifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacMachineIdentifier] class.
//
// # Creating a machine identifier
//
//   - [IVZMacMachineIdentifier.InitWithDataRepresentation]: Create a machine identifier described by the specified data representation.
//
// # Machine data representation
//
//   - [IVZMacMachineIdentifier.DataRepresentation]: Returns the opaque data representation of the machine identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier
type IVZMacMachineIdentifier interface {
	objectivec.IObject

	// Topic: Creating a machine identifier

	// Create a machine identifier described by the specified data representation.
	InitWithDataRepresentation(dataRepresentation foundation.INSData) VZMacMachineIdentifier

	// Topic: Machine data representation

	// Returns the opaque data representation of the machine identifier.
	DataRepresentation() foundation.INSData
}

// Init initializes the instance.
func (m VZMacMachineIdentifier) Init() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacMachineIdentifier) Autorelease() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacMachineIdentifier creates a new VZMacMachineIdentifier instance.
func NewVZMacMachineIdentifier() VZMacMachineIdentifier {
	class := getVZMacMachineIdentifierClass()
	rv := objc.Send[VZMacMachineIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create a machine identifier described by the specified data representation.
//
// dataRepresentation: The opaque data representation of the machine identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/init(dataRepresentation:)
func NewMacMachineIdentifierWithDataRepresentation(dataRepresentation foundation.INSData) VZMacMachineIdentifier {
	instance := getVZMacMachineIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	return VZMacMachineIdentifierFromID(rv)
}

// Create a machine identifier described by the specified data representation.
//
// dataRepresentation: The opaque data representation of the machine identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/init(dataRepresentation:)
func (m VZMacMachineIdentifier) InitWithDataRepresentation(dataRepresentation foundation.INSData) VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](m.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	return rv
}

// Returns the opaque data representation of the machine identifier.
//
// # Discussion
// 
// You can use this to recreate the same machine identifier with
// [InitWithDataRepresentation].
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/dataRepresentation
func (m VZMacMachineIdentifier) DataRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}

