// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacAuxiliaryStorage] class.
var (
	_VZMacAuxiliaryStorageClass     VZMacAuxiliaryStorageClass
	_VZMacAuxiliaryStorageClassOnce sync.Once
)

func getVZMacAuxiliaryStorageClass() VZMacAuxiliaryStorageClass {
	_VZMacAuxiliaryStorageClassOnce.Do(func() {
		_VZMacAuxiliaryStorageClass = VZMacAuxiliaryStorageClass{class: objc.GetClass("VZMacAuxiliaryStorage")}
	})
	return _VZMacAuxiliaryStorageClass
}

// GetVZMacAuxiliaryStorageClass returns the class object for VZMacAuxiliaryStorage.
func GetVZMacAuxiliaryStorageClass() VZMacAuxiliaryStorageClass {
	return getVZMacAuxiliaryStorageClass()
}

type VZMacAuxiliaryStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacAuxiliaryStorageClass) Alloc() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that contains information the boot loader needs for booting macOS
// as a guest operating system.
//
// # Overview
// 
// The Mac auxiliary storage contains data used by the boot loader and the
// guest operating system. It’s necessary to boot a macOS guest OS.
// 
// When creating a new VM, use
// [VZMacAuxiliaryStorage.InitCreatingStorageAtURLHardwareModelOptionsError] to create a default
// initialized auxiliary storage.
// 
// The hardware model you use when creating the new auxiliary storage depends
// on the restore image that you’ll use for installation. From the restore
// image, use [VZMacAuxiliaryStorage.MostFeaturefulSupportedConfiguration] to get a supported
// configuration. A configuration has a [VZMacHardwareModel] associated with
// it.
// 
// After initializing the new auxiliary storage, set it on
// [VZMacPlatformConfiguration].[VZMacAuxiliaryStorage.AuxiliaryStorage].
// 
// The hardware model in [VZMacPlatformConfiguration].[VZMacAuxiliaryStorage.HardwareModel] must be
// identical to the one used to create the empty auxiliary storage., otherwise
// the behavior isn’t defined.
// 
// When installing macOS, the [VZMacOSInstaller] lays out data on the
// auxiliary storage. After installation, the macOS guest uses the auxiliary
// storage for every subsequent boot.
// 
// When moving or performing a backup of a VM, you must move or copy the file
// containing the auxiliary storage along with the main disk image.
// 
// To boot a VM created with [VZMacOSInstaller], use [init(contentsOfURL:)] to
// set up the auxiliary storage from the existing file used during
// installation.
// 
// When using an existing file, the hardware model of the
// [VZMacPlatformConfiguration].[VZMacAuxiliaryStorage.HardwareModel] must match the hardware model
// used when creating the original file.
//
// [init(contentsOfURL:)]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(contentsOfURL:)
//
// # Creating the auxiliary storage
//
//   - [VZMacAuxiliaryStorage.InitWithURL]: Initializes an auxiliary storage object with data from the location at the URL you provide.
//   - [VZMacAuxiliaryStorage.InitCreatingStorageAtURLHardwareModelOptionsError]: Creates an initialized Mac auxiliary storage instance that describes a specific hardware model at a URL you specify.
//
// # Configuring the auxiliary storage location
//
//   - [VZMacAuxiliaryStorage.URL]: The URL of the auxiliary storage on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage
type VZMacAuxiliaryStorage struct {
	objectivec.Object
}

// VZMacAuxiliaryStorageFromID constructs a [VZMacAuxiliaryStorage] from an objc.ID.
//
// An object that contains information the boot loader needs for booting macOS
// as a guest operating system.
func VZMacAuxiliaryStorageFromID(id objc.ID) VZMacAuxiliaryStorage {
	return VZMacAuxiliaryStorage{objectivec.Object{id}}
}
// NOTE: VZMacAuxiliaryStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMacAuxiliaryStorage] class.
//
// # Creating the auxiliary storage
//
//   - [IVZMacAuxiliaryStorage.InitWithURL]: Initializes an auxiliary storage object with data from the location at the URL you provide.
//   - [IVZMacAuxiliaryStorage.InitCreatingStorageAtURLHardwareModelOptionsError]: Creates an initialized Mac auxiliary storage instance that describes a specific hardware model at a URL you specify.
//
// # Configuring the auxiliary storage location
//
//   - [IVZMacAuxiliaryStorage.URL]: The URL of the auxiliary storage on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage
type IVZMacAuxiliaryStorage interface {
	objectivec.IObject

	// Topic: Creating the auxiliary storage

	// Initializes an auxiliary storage object with data from the location at the URL you provide.
	InitWithURL(URL foundation.INSURL) VZMacAuxiliaryStorage
	// Creates an initialized Mac auxiliary storage instance that describes a specific hardware model at a URL you specify.
	InitCreatingStorageAtURLHardwareModelOptionsError(URL foundation.INSURL, hardwareModel IVZMacHardwareModel, options VZMacAuxiliaryStorageInitializationOptions) (VZMacAuxiliaryStorage, error)

	// Topic: Configuring the auxiliary storage location

	// The URL of the auxiliary storage on the local file system.
	URL() foundation.INSURL

	// The Mac auxiliary storage.
	AuxiliaryStorage() IVZMacAuxiliaryStorage
	SetAuxiliaryStorage(value IVZMacAuxiliaryStorage)
	// The Mac hardware model.
	HardwareModel() IVZMacHardwareModel
	SetHardwareModel(value IVZMacHardwareModel)
	// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
	MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements
	SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements)
}





// Init initializes the instance.
func (m VZMacAuxiliaryStorage) Init() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacAuxiliaryStorage) Autorelease() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacAuxiliaryStorage creates a new VZMacAuxiliaryStorage instance.
func NewVZMacAuxiliaryStorage() VZMacAuxiliaryStorage {
	class := getVZMacAuxiliaryStorageClass()
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates an initialized Mac auxiliary storage instance that describes a
// specific hardware model at a URL you specify.
//
// URL: The [URL] to write the auxiliary storage to on the local file system.
//
// hardwareModel: The [VZMacHardwareModel] model to use. The auxiliary storage can have
// different layouts for different hardware models.
//
// options: Initialization options from the available
// [VZMacAuxiliaryStorage.InitializationOptions].
// //
// [VZMacAuxiliaryStorage.InitializationOptions]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/InitializationOptions
//
// # Return Value
// 
// Returns a newly initialized [VZMacAuxiliaryStorage] object on success or
// `nil` if there was an error. On failure, `error` contains the [NSError]
// that describes reason for the failure.
//
// # Discussion
// 
// Use this method to create a new auxiliary storage object that describes a
// specific hardware model. To restore data from a previously saved existing
// auxiliary storage object, use [init(contentsOfURL:)].
//
// [init(contentsOfURL:)]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(contentsOfURL:)
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(creatingStorageAt:hardwareModel:options:)
func NewMacAuxiliaryStorageCreatingStorageAtURLHardwareModelOptionsError(URL foundation.INSURL, hardwareModel IVZMacHardwareModel, options VZMacAuxiliaryStorageInitializationOptions) (VZMacAuxiliaryStorage, error) {
	var errorPtr objc.ID
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initCreatingStorageAtURL:hardwareModel:options:error:"), URL, hardwareModel, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZMacAuxiliaryStorageFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VZMacAuxiliaryStorageFromID(rv), nil
}


//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(contentsOf:)
func NewMacAuxiliaryStorageWithContentsOfURL(URL foundation.INSURL) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), URL)
	return VZMacAuxiliaryStorageFromID(rv)
}


// Initializes an auxiliary storage object with data from the location at the
// URL you provide.
//
// URL: The URL of the auxiliary storage on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(url:)
func NewMacAuxiliaryStorageWithURL(URL foundation.INSURL) VZMacAuxiliaryStorage {
	instance := getVZMacAuxiliaryStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return VZMacAuxiliaryStorageFromID(rv)
}







// Initializes an auxiliary storage object with data from the location at the
// URL you provide.
//
// URL: The URL of the auxiliary storage on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(url:)
func (m VZMacAuxiliaryStorage) InitWithURL(URL foundation.INSURL) VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](m.ID, objc.Sel("initWithURL:"), URL)
	return rv
}

// Creates an initialized Mac auxiliary storage instance that describes a
// specific hardware model at a URL you specify.
//
// URL: The [URL] to write the auxiliary storage to on the local file system.
//
// hardwareModel: The [VZMacHardwareModel] model to use. The auxiliary storage can have
// different layouts for different hardware models.
//
// options: Initialization options from the available
// [VZMacAuxiliaryStorage.InitializationOptions].
// //
// [VZMacAuxiliaryStorage.InitializationOptions]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/InitializationOptions
//
// # Return Value
// 
// Returns a newly initialized [VZMacAuxiliaryStorage] object on success or
// `nil` if there was an error. On failure, `error` contains the [NSError]
// that describes reason for the failure.
//
// # Discussion
// 
// Use this method to create a new auxiliary storage object that describes a
// specific hardware model. To restore data from a previously saved existing
// auxiliary storage object, use [init(contentsOfURL:)].
//
// [init(contentsOfURL:)]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(contentsOfURL:)
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/init(creatingStorageAt:hardwareModel:options:)
func (m VZMacAuxiliaryStorage) InitCreatingStorageAtURLHardwareModelOptionsError(URL foundation.INSURL, hardwareModel IVZMacHardwareModel, options VZMacAuxiliaryStorageInitializationOptions) (VZMacAuxiliaryStorage, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initCreatingStorageAtURL:hardwareModel:options:error:"), URL, hardwareModel, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZMacAuxiliaryStorage{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZMacAuxiliaryStorageFromID(rv), nil

}












// The URL of the auxiliary storage on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/url
func (m VZMacAuxiliaryStorage) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}



// The Mac auxiliary storage.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/auxiliarystorage
func (m VZMacAuxiliaryStorage) AuxiliaryStorage() IVZMacAuxiliaryStorage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("auxiliaryStorage"))
	return VZMacAuxiliaryStorageFromID(objc.ID(rv))
}
func (m VZMacAuxiliaryStorage) SetAuxiliaryStorage(value IVZMacAuxiliaryStorage) {
	objc.Send[struct{}](m.ID, objc.Sel("setAuxiliaryStorage:"), value)
}



// The Mac hardware model.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (m VZMacAuxiliaryStorage) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("hardwareModel"))
	return VZMacHardwareModelFromID(objc.ID(rv))
}
func (m VZMacAuxiliaryStorage) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[struct{}](m.ID, objc.Sel("setHardwareModel:"), value)
}



// This object represents the most fully featured configuration that’s
// supported by both the current host and by this restore image.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (m VZMacAuxiliaryStorage) MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return VZMacOSConfigurationRequirementsFromID(objc.ID(rv))
}
func (m VZMacAuxiliaryStorage) SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements) {
	objc.Send[struct{}](m.ID, objc.Sel("setMostFeaturefulSupportedConfiguration:"), value)
}























