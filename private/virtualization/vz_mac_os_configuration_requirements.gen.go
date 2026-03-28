// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacOSConfigurationRequirements] class.
var (
	_VZMacOSConfigurationRequirementsClass     VZMacOSConfigurationRequirementsClass
	_VZMacOSConfigurationRequirementsClassOnce sync.Once
)

func getVZMacOSConfigurationRequirementsClass() VZMacOSConfigurationRequirementsClass {
	_VZMacOSConfigurationRequirementsClassOnce.Do(func() {
		_VZMacOSConfigurationRequirementsClass = VZMacOSConfigurationRequirementsClass{class: objc.GetClass("VZMacOSConfigurationRequirements")}
	})
	return _VZMacOSConfigurationRequirementsClass
}

// GetVZMacOSConfigurationRequirementsClass returns the class object for VZMacOSConfigurationRequirements.
func GetVZMacOSConfigurationRequirementsClass() VZMacOSConfigurationRequirementsClass {
	return getVZMacOSConfigurationRequirementsClass()
}

type VZMacOSConfigurationRequirementsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacOSConfigurationRequirementsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSConfigurationRequirementsClass) Alloc() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacOSConfigurationRequirements._variants]
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements
type VZMacOSConfigurationRequirements struct {
	objectivec.Object
}

// VZMacOSConfigurationRequirementsFromID constructs a [VZMacOSConfigurationRequirements] from an objc.ID.
func VZMacOSConfigurationRequirementsFromID(id objc.ID) VZMacOSConfigurationRequirements {
	return VZMacOSConfigurationRequirements{objectivec.Object{ID: id}}
}
// Ensure VZMacOSConfigurationRequirements implements IVZMacOSConfigurationRequirements.
var _ IVZMacOSConfigurationRequirements = VZMacOSConfigurationRequirements{}

// An interface definition for the [VZMacOSConfigurationRequirements] class.
//
// # Methods
//
//   - [IVZMacOSConfigurationRequirements._variants]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements
type IVZMacOSConfigurationRequirements interface {
	objectivec.IObject

	// Topic: Methods

	_variants() objectivec.IObject
}

// Init initializes the instance.
func (m VZMacOSConfigurationRequirements) Init() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSConfigurationRequirements) Autorelease() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSConfigurationRequirements creates a new VZMacOSConfigurationRequirements instance.
func NewVZMacOSConfigurationRequirements() VZMacOSConfigurationRequirements {
	class := getVZMacOSConfigurationRequirementsClass()
	rv := objc.Send[VZMacOSConfigurationRequirements](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/_variants
func (m VZMacOSConfigurationRequirements) _variants() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_variants"))
	return objectivec.Object{ID: rv}
}

// Variants is an exported wrapper for the private method _variants.
func (m VZMacOSConfigurationRequirements) Variants() objectivec.IObject {
	return m._variants()
}

