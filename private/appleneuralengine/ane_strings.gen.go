// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEStrings] class.
var (
	_ANEStringsClass     ANEStringsClass
	_ANEStringsClassOnce sync.Once
)

func getANEStringsClass() ANEStringsClass {
	_ANEStringsClassOnce.Do(func() {
		_ANEStringsClass = ANEStringsClass{class: objc.GetClass("_ANEStrings")}
	})
	return _ANEStringsClass
}

// GetANEStringsClass returns the class object for _ANEStrings.
func GetANEStringsClass() ANEStringsClass {
	return getANEStringsClass()
}

type ANEStringsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEStringsClass) Alloc() ANEStrings {
	rv := objc.Send[ANEStrings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings
type ANEStrings struct {
	objectivec.Object
}

// ANEStringsFromID constructs a [ANEStrings] from an objc.ID.
func ANEStringsFromID(id objc.ID) ANEStrings {
	return ANEStrings{objectivec.Object{id}}
}
// Ensure ANEStrings implements IANEStrings.
var _ IANEStrings = ANEStrings{}





// An interface definition for the [ANEStrings] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings
type IANEStrings interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANEStrings) Init() ANEStrings {
	rv := objc.Send[ANEStrings](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEStrings) Autorelease() ANEStrings {
	rv := objc.Send[ANEStrings](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEStrings creates a new ANEStrings instance.
func NewANEStrings() ANEStrings {
	class := getANEStringsClass()
	rv := objc.Send[ANEStrings](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/adapterWeightsAccessEntitlement
func (_ANEStringsClass ANEStringsClass) AdapterWeightsAccessEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("adapterWeightsAccessEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/adapterWeightsAccessEntitlementBypassBootArg
func (_ANEStringsClass ANEStringsClass) AdapterWeightsAccessEntitlementBypassBootArg() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("adapterWeightsAccessEntitlementBypassBootArg"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/aggressivePowerSavingEntitlement
func (_ANEStringsClass ANEStringsClass) AggressivePowerSavingEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("aggressivePowerSavingEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/binExtension
func (_ANEStringsClass ANEStringsClass) BinExtension() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("binExtension"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/buildSpecificModelDataVaultDirectory
func (_ANEStringsClass ANEStringsClass) BuildSpecificModelDataVaultDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("buildSpecificModelDataVaultDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/buildSpecificUserModelDataVaultDirectory
func (_ANEStringsClass ANEStringsClass) BuildSpecificUserModelDataVaultDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("buildSpecificUserModelDataVaultDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/cloneDirectory
func (_ANEStringsClass ANEStringsClass) CloneDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("cloneDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/compilerServiceAccessEntitlement
func (_ANEStringsClass ANEStringsClass) CompilerServiceAccessEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("compilerServiceAccessEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/coreAnalyticsANEUsageDefaultReportedClient
func (_ANEStringsClass ANEStringsClass) CoreAnalyticsANEUsageDefaultReportedClient() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("coreAnalyticsANEUsageDefaultReportedClient"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/coreAnalyticsANEUsageKeyGroup
func (_ANEStringsClass ANEStringsClass) CoreAnalyticsANEUsageKeyGroup() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("coreAnalyticsANEUsageKeyGroup"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/dataVaultStorageClass
func (_ANEStringsClass ANEStringsClass) DataVaultStorageClass() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("dataVaultStorageClass"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/defaultANECIRFileName
func (_ANEStringsClass ANEStringsClass) DefaultANECIRFileName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("defaultANECIRFileName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/defaultANECIROptionsFileName
func (_ANEStringsClass ANEStringsClass) DefaultANECIROptionsFileName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("defaultANECIROptionsFileName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/defaultCompilerOptionsFilename
func (_ANEStringsClass ANEStringsClass) DefaultCompilerOptionsFilename() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("defaultCompilerOptionsFilename"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/defaultLLIRBundleName
func (_ANEStringsClass ANEStringsClass) DefaultLLIRBundleName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("defaultLLIRBundleName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/defaultMILFileName
func (_ANEStringsClass ANEStringsClass) DefaultMILFileName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("defaultMILFileName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/defaultMLIRFileName
func (_ANEStringsClass ANEStringsClass) DefaultMLIRFileName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("defaultMLIRFileName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/defaultWeightFileName
func (_ANEStringsClass ANEStringsClass) DefaultWeightFileName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("defaultWeightFileName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/errorDomainCompiler
func (_ANEStringsClass ANEStringsClass) ErrorDomainCompiler() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("errorDomainCompiler"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/errorDomainEspresso
func (_ANEStringsClass ANEStringsClass) ErrorDomainEspresso() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("errorDomainEspresso"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/errorDomainGeneric
func (_ANEStringsClass ANEStringsClass) ErrorDomainGeneric() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("errorDomainGeneric"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/errorDomainVirtIO
func (_ANEStringsClass ANEStringsClass) ErrorDomainVirtIO() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("errorDomainVirtIO"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/hwxExtension
func (_ANEStringsClass ANEStringsClass) HwxExtension() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("hwxExtension"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/inMemoryModelCacheName
func (_ANEStringsClass ANEStringsClass) InMemoryModelCacheName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("inMemoryModelCacheName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/internalLibraryPath
func (_ANEStringsClass ANEStringsClass) InternalLibraryPath() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("internalLibraryPath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/launchIOKitEvent
func (_ANEStringsClass ANEStringsClass) LaunchIOKitEvent() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("launchIOKitEvent"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/launchUserIOKitEvent
func (_ANEStringsClass ANEStringsClass) LaunchUserIOKitEvent() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("launchUserIOKitEvent"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/llirBundleExtension
func (_ANEStringsClass ANEStringsClass) LlirBundleExtension() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("llirBundleExtension"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/machServiceName
func (_ANEStringsClass ANEStringsClass) MachServiceName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("machServiceName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/machServiceNamePrivate
func (_ANEStringsClass ANEStringsClass) MachServiceNamePrivate() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("machServiceNamePrivate"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/memoryUnwireAccessEntitlement
func (_ANEStringsClass ANEStringsClass) MemoryUnwireAccessEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("memoryUnwireAccessEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/modelAssetsCacheName
func (_ANEStringsClass ANEStringsClass) ModelAssetsCacheName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("modelAssetsCacheName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/modelBinaryName
func (_ANEStringsClass ANEStringsClass) ModelBinaryName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("modelBinaryName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/modelCacheRetainName
func (_ANEStringsClass ANEStringsClass) ModelCacheRetainName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("modelCacheRetainName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/modelDataVaultDirectory
func (_ANEStringsClass ANEStringsClass) ModelDataVaultDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("modelDataVaultDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/modelPurgeInAllPartitionsEntitlement
func (_ANEStringsClass ANEStringsClass) ModelPurgeInAllPartitionsEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("modelPurgeInAllPartitionsEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/modelSourceStoreName
func (_ANEStringsClass ANEStringsClass) ModelSourceStoreName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("modelSourceStoreName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/noSandboxExtension
func (_ANEStringsClass ANEStringsClass) NoSandboxExtension() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("noSandboxExtension"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/ppsCategoryForANE
func (_ANEStringsClass ANEStringsClass) PpsCategoryForANE() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("ppsCategoryForANE"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/ppsSubsystemForANE
func (_ANEStringsClass ANEStringsClass) PpsSubsystemForANE() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("ppsSubsystemForANE"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_applicationDir
func (_ANEStringsClass ANEStringsClass) Pps_applicationDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_applicationDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_catalogDir
func (_ANEStringsClass ANEStringsClass) Pps_catalogDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_catalogDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_defaultSystemPathDir
func (_ANEStringsClass ANEStringsClass) Pps_defaultSystemPathDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_defaultSystemPathDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_defaultUserPathPrefix
func (_ANEStringsClass ANEStringsClass) Pps_defaultUserPathPrefix() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_defaultUserPathPrefix"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_frameworkDir
func (_ANEStringsClass ANEStringsClass) Pps_frameworkDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_frameworkDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_internalDir
func (_ANEStringsClass ANEStringsClass) Pps_internalDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_internalDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_privateFrameworkDir
func (_ANEStringsClass ANEStringsClass) Pps_privateFrameworkDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_privateFrameworkDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_tmpDir
func (_ANEStringsClass ANEStringsClass) Pps_tmpDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_tmpDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/pps_varDir
func (_ANEStringsClass ANEStringsClass) Pps_varDir() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("pps_varDir"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/processModelShareAccessEntitlement
func (_ANEStringsClass ANEStringsClass) ProcessModelShareAccessEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("processModelShareAccessEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/restrictedAccessEntitlement
func (_ANEStringsClass ANEStringsClass) RestrictedAccessEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("restrictedAccessEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/secondaryANECompilerServiceAccessEntitlement
func (_ANEStringsClass ANEStringsClass) SecondaryANECompilerServiceAccessEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("secondaryANECompilerServiceAccessEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/storageMaintainerAccessEntitlement
func (_ANEStringsClass ANEStringsClass) StorageMaintainerAccessEntitlement() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("storageMaintainerAccessEntitlement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/systemLibraryPath
func (_ANEStringsClass ANEStringsClass) SystemLibraryPath() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("systemLibraryPath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/systemModelsCacheDirectory
func (_ANEStringsClass ANEStringsClass) SystemModelsCacheDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("systemModelsCacheDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/tempDirectory
func (_ANEStringsClass ANEStringsClass) TempDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("tempDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_ThreeSixtyModelName
func (_ANEStringsClass ANEStringsClass) Testing_ThreeSixtyModelName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_ThreeSixtyModelName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_cacheDirectory
func (_ANEStringsClass ANEStringsClass) Testing_cacheDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_cacheDirectory"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_cacheDirectoryWithSuffix:
func (_ANEStringsClass ANEStringsClass) Testing_cacheDirectoryWithSuffix(suffix objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_cacheDirectoryWithSuffix:"), suffix)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_cacheDirectoryWithSuffix:buildVersion:
func (_ANEStringsClass ANEStringsClass) Testing_cacheDirectoryWithSuffixBuildVersion(suffix objectivec.IObject, version objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_cacheDirectoryWithSuffix:buildVersion:"), suffix, version)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_cloneDirectory:
func (_ANEStringsClass ANEStringsClass) Testing_cloneDirectory(directory objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_cloneDirectory:"), directory)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_dataVaultStorageClass
func (_ANEStringsClass ANEStringsClass) Testing_dataVaultStorageClass() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_dataVaultStorageClass"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_defaultMLIRModelName
func (_ANEStringsClass ANEStringsClass) Testing_defaultMLIRModelName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_defaultMLIRModelName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_encryptedModelNames
func (_ANEStringsClass ANEStringsClass) Testing_encryptedModelNames() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_encryptedModelNames"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_external_modelPath
func (_ANEStringsClass ANEStringsClass) Testing_external_modelPath() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_external_modelPath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_external_precompiledModelPath
func (_ANEStringsClass ANEStringsClass) Testing_external_precompiledModelPath() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_external_precompiledModelPath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_inputDirectory
func (_ANEStringsClass ANEStringsClass) Testing_inputDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_inputDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_modelDirectory
func (_ANEStringsClass ANEStringsClass) Testing_modelDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_modelDirectory"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_modelDirectory:
func (_ANEStringsClass ANEStringsClass) Testing_modelDirectoryWithDirectory(directory objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_modelDirectory:"), directory)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_modelNames
func (_ANEStringsClass ANEStringsClass) Testing_modelNames() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_modelNames"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_tempDirectory:
func (_ANEStringsClass ANEStringsClass) Testing_tempDirectory(directory objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_tempDirectory:"), directory)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_userCacheDirectory
func (_ANEStringsClass ANEStringsClass) Testing_userCacheDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_userCacheDirectory"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_userCacheDirectoryWithSuffix:
func (_ANEStringsClass ANEStringsClass) Testing_userCacheDirectoryWithSuffix(suffix objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_userCacheDirectoryWithSuffix:"), suffix)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_userCacheDirectoryWithSuffix:buildVersion:
func (_ANEStringsClass ANEStringsClass) Testing_userCacheDirectoryWithSuffixBuildVersion(suffix objectivec.IObject, version objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_userCacheDirectoryWithSuffix:buildVersion:"), suffix, version)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_userCloneDirectory:
func (_ANEStringsClass ANEStringsClass) Testing_userCloneDirectory(directory objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_userCloneDirectory:"), directory)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_userTempDirectory:
func (_ANEStringsClass ANEStringsClass) Testing_userTempDirectory(directory objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_userTempDirectory:"), directory)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/testing_zeroModelName
func (_ANEStringsClass ANEStringsClass) Testing_zeroModelName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("testing_zeroModelName"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/trimmedModelPath:trimmedPath:
func (_ANEStringsClass ANEStringsClass) TrimmedModelPathTrimmedPath(path objectivec.IObject, path2 []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEStringsClass.class), objc.Sel("trimmedModelPath:trimmedPath:"), path, objectivec.IObjectSliceToNSArray(path2))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/userCloneDirectory
func (_ANEStringsClass ANEStringsClass) UserCloneDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("userCloneDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/userMachServiceName
func (_ANEStringsClass ANEStringsClass) UserMachServiceName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("userMachServiceName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/userModelDataVaultDirectory
func (_ANEStringsClass ANEStringsClass) UserModelDataVaultDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("userModelDataVaultDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/userTempDirectory
func (_ANEStringsClass ANEStringsClass) UserTempDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("userTempDirectory"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/vm_allowPrecompiledBinaryBootArg
func (_ANEStringsClass ANEStringsClass) Vm_allowPrecompiledBinaryBootArg() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("vm_allowPrecompiledBinaryBootArg"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/vm_debugDumpBootArg
func (_ANEStringsClass ANEStringsClass) Vm_debugDumpBootArg() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("vm_debugDumpBootArg"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/vm_forceValidationOnGuestBootArg
func (_ANEStringsClass ANEStringsClass) Vm_forceValidationOnGuestBootArg() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("vm_forceValidationOnGuestBootArg"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStrings/vm_tmpBaseDirectory
func (_ANEStringsClass ANEStringsClass) Vm_tmpBaseDirectory() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStringsClass.class), objc.Sel("vm_tmpBaseDirectory"))
	return objectivec.Object{ID: rv}
}






















