// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacHardwareModel] class.
var (
	_VZMacHardwareModelClass     VZMacHardwareModelClass
	_VZMacHardwareModelClassOnce sync.Once
)

func getVZMacHardwareModelClass() VZMacHardwareModelClass {
	_VZMacHardwareModelClassOnce.Do(func() {
		_VZMacHardwareModelClass = VZMacHardwareModelClass{class: objc.GetClass("VZMacHardwareModel")}
	})
	return _VZMacHardwareModelClass
}

// GetVZMacHardwareModelClass returns the class object for VZMacHardwareModel.
func GetVZMacHardwareModelClass() VZMacHardwareModelClass {
	return getVZMacHardwareModelClass()
}

type VZMacHardwareModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacHardwareModelClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacHardwareModelClass) Alloc() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A specification for the hardware elements and configurations present in a
// particular Mac hardware model.
//
// # Overview
// 
// The Mac hardware model abstracts a set of virtualized hardware elements and
// configurations.
// 
// A version of macOS may only run on certain hardware models. Additionally,
// the host may also only provide certain hardware models based on the version
// of macOS and the underlying hardware.
// 
// The [VZMacHardwareModel.Supported] property allows you to discover if the current host
// supports a particular hardware model.
// 
// Choosing the hardware model starts from a restore image with
// [VZMacOSRestoreImage]. A restore image describes its supported
// configuration requirements through its
// [VZMacHardwareModel.MostFeaturefulSupportedConfiguration] property.
// 
// A configuration requirements object has a corresponding hardware model that
// you can use to configure a VM that meets the requirements. After obtaining
// the hardware model, use the platform configuration’s [VZMacHardwareModel.HardwareModel] to
// configure the Mac platform object and use
// [InitCreatingStorageAtURLHardwareModelOptionsError] to create its auxiliary
// storage.
// 
// After creating the VM, use [VZMacOSInstaller] to install macOS on it.
// 
// If you serialize the VM on disk, preserve the hardware model used for
// installation for subsequent boots. The [VZMacHardwareModel.DataRepresentation] property
// provides a unique binary representation that you serialize to the file
// system. You can recreate the hardware model from the serialized binary
// representation with [VZMacHardwareModel.InitWithDataRepresentation].
//
// # Creating the hardware model
//
//   - [VZMacHardwareModel.InitWithDataRepresentation]: Creates an instance of the hardware model described by the specified data representation.
//
// # Configuring the hardware model
//
//   - [VZMacHardwareModel.DataRepresentation]: Returns the opaque data representation of the hardware model.
//   - [VZMacHardwareModel.Supported]: A Boolean value that indicates whether the host supports this hardware model.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel
type VZMacHardwareModel struct {
	objectivec.Object
}

// VZMacHardwareModelFromID constructs a [VZMacHardwareModel] from an objc.ID.
//
// A specification for the hardware elements and configurations present in a
// particular Mac hardware model.
func VZMacHardwareModelFromID(id objc.ID) VZMacHardwareModel {
	return VZMacHardwareModel{objectivec.Object{ID: id}}
}
// NOTE: VZMacHardwareModel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacHardwareModel] class.
//
// # Creating the hardware model
//
//   - [IVZMacHardwareModel.InitWithDataRepresentation]: Creates an instance of the hardware model described by the specified data representation.
//
// # Configuring the hardware model
//
//   - [IVZMacHardwareModel.DataRepresentation]: Returns the opaque data representation of the hardware model.
//   - [IVZMacHardwareModel.Supported]: A Boolean value that indicates whether the host supports this hardware model.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel
type IVZMacHardwareModel interface {
	objectivec.IObject

	// Topic: Creating the hardware model

	// Creates an instance of the hardware model described by the specified data representation.
	InitWithDataRepresentation(dataRepresentation foundation.INSData) VZMacHardwareModel

	// Topic: Configuring the hardware model

	// Returns the opaque data representation of the hardware model.
	DataRepresentation() foundation.INSData
	// A Boolean value that indicates whether the host supports this hardware model.
	Supported() bool

	// The Mac hardware model.
	HardwareModel() IVZMacHardwareModel
	SetHardwareModel(value IVZMacHardwareModel)
	// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
	MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements
	SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements)
}

// Init initializes the instance.
func (m VZMacHardwareModel) Init() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacHardwareModel) Autorelease() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHardwareModel creates a new VZMacHardwareModel instance.
func NewVZMacHardwareModel() VZMacHardwareModel {
	class := getVZMacHardwareModelClass()
	rv := objc.Send[VZMacHardwareModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an instance of the hardware model described by the specified data
// representation.
//
// dataRepresentation: The opaque data representation of the hardware model.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/init(dataRepresentation:)
func NewMacHardwareModelWithDataRepresentation(dataRepresentation foundation.INSData) VZMacHardwareModel {
	instance := getVZMacHardwareModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	return VZMacHardwareModelFromID(rv)
}

// Creates an instance of the hardware model described by the specified data
// representation.
//
// dataRepresentation: The opaque data representation of the hardware model.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/init(dataRepresentation:)
func (m VZMacHardwareModel) InitWithDataRepresentation(dataRepresentation foundation.INSData) VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](m.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	return rv
}

// Returns the opaque data representation of the hardware model.
//
// # Discussion
// 
// You can use this to recreate the same hardware model with
// [InitWithDataRepresentation].
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/dataRepresentation
func (m VZMacHardwareModel) DataRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the host supports this hardware
// model.
//
// # Discussion
// 
// If this hardware model isn’t supported by the host, the
// [VZVirtualMachineConfiguration] won’t validate.
// 
// The validation error of the [VZVirtualMachineConfiguration] provides more
// information about why the hardware model isn’t supported.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/isSupported
func (m VZMacHardwareModel) Supported() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSupported"))
	return rv
}
// The Mac hardware model.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (m VZMacHardwareModel) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("hardwareModel"))
	return VZMacHardwareModelFromID(objc.ID(rv))
}
func (m VZMacHardwareModel) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[struct{}](m.ID, objc.Sel("setHardwareModel:"), value)
}
// This object represents the most fully featured configuration that’s
// supported by both the current host and by this restore image.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (m VZMacHardwareModel) MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return VZMacOSConfigurationRequirementsFromID(objc.ID(rv))
}
func (m VZMacHardwareModel) SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements) {
	objc.Send[struct{}](m.ID, objc.Sel("setMostFeaturefulSupportedConfiguration:"), value)
}

