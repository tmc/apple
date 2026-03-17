// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEInMemoryModelCacheManager] class.
var (
	_ANEInMemoryModelCacheManagerClass     ANEInMemoryModelCacheManagerClass
	_ANEInMemoryModelCacheManagerClassOnce sync.Once
)

func getANEInMemoryModelCacheManagerClass() ANEInMemoryModelCacheManagerClass {
	_ANEInMemoryModelCacheManagerClassOnce.Do(func() {
		_ANEInMemoryModelCacheManagerClass = ANEInMemoryModelCacheManagerClass{class: objc.GetClass("_ANEInMemoryModelCacheManager")}
	})
	return _ANEInMemoryModelCacheManagerClass
}

// GetANEInMemoryModelCacheManagerClass returns the class object for _ANEInMemoryModelCacheManager.
func GetANEInMemoryModelCacheManagerClass() ANEInMemoryModelCacheManagerClass {
	return getANEInMemoryModelCacheManagerClass()
}

type ANEInMemoryModelCacheManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEInMemoryModelCacheManagerClass) Alloc() ANEInMemoryModelCacheManager {
	rv := objc.Send[ANEInMemoryModelCacheManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEInMemoryModelCacheManager.URLForBundleID]
//   - [ANEInMemoryModelCacheManager.URLForModelHashBundleID]
//   - [ANEInMemoryModelCacheManager.CacheDir]
//   - [ANEInMemoryModelCacheManager.CachedModelPathMatchingHashCsIdentity]
//   - [ANEInMemoryModelCacheManager.GetDiskSpaceForBundleID]
//   - [ANEInMemoryModelCacheManager.GetDiskSpaceItemizedByBundleIDAndPurge]
//   - [ANEInMemoryModelCacheManager.RemoveAllModelsForBundleID]
//   - [ANEInMemoryModelCacheManager.RemoveStaleModels]
//   - [ANEInMemoryModelCacheManager.ScheduleMaintenanceWithNameDirectoryPaths]
//   - [ANEInMemoryModelCacheManager.ShouldEnforceSizeLimits]
//   - [ANEInMemoryModelCacheManager.InitWithURL]
//   - [ANEInMemoryModelCacheManager.InitWithURLCreateDirectory]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager
type ANEInMemoryModelCacheManager struct {
	objectivec.Object
}

// ANEInMemoryModelCacheManagerFromID constructs a [ANEInMemoryModelCacheManager] from an objc.ID.
func ANEInMemoryModelCacheManagerFromID(id objc.ID) ANEInMemoryModelCacheManager {
	return ANEInMemoryModelCacheManager{objectivec.Object{ID: id}}
}
// Ensure ANEInMemoryModelCacheManager implements IANEInMemoryModelCacheManager.
var _ IANEInMemoryModelCacheManager = ANEInMemoryModelCacheManager{}

// An interface definition for the [ANEInMemoryModelCacheManager] class.
//
// # Methods
//
//   - [IANEInMemoryModelCacheManager.URLForBundleID]
//   - [IANEInMemoryModelCacheManager.URLForModelHashBundleID]
//   - [IANEInMemoryModelCacheManager.CacheDir]
//   - [IANEInMemoryModelCacheManager.CachedModelPathMatchingHashCsIdentity]
//   - [IANEInMemoryModelCacheManager.GetDiskSpaceForBundleID]
//   - [IANEInMemoryModelCacheManager.GetDiskSpaceItemizedByBundleIDAndPurge]
//   - [IANEInMemoryModelCacheManager.RemoveAllModelsForBundleID]
//   - [IANEInMemoryModelCacheManager.RemoveStaleModels]
//   - [IANEInMemoryModelCacheManager.ScheduleMaintenanceWithNameDirectoryPaths]
//   - [IANEInMemoryModelCacheManager.ShouldEnforceSizeLimits]
//   - [IANEInMemoryModelCacheManager.InitWithURL]
//   - [IANEInMemoryModelCacheManager.InitWithURLCreateDirectory]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager
type IANEInMemoryModelCacheManager interface {
	objectivec.IObject

	// Topic: Methods

	URLForBundleID(id objectivec.IObject) objectivec.IObject
	URLForModelHashBundleID(hash objectivec.IObject, id objectivec.IObject) objectivec.IObject
	CacheDir() foundation.INSURL
	CachedModelPathMatchingHashCsIdentity(hash objectivec.IObject, identity objectivec.IObject) objectivec.IObject
	GetDiskSpaceForBundleID(id objectivec.IObject) objectivec.IObject
	GetDiskSpaceItemizedByBundleIDAndPurge(purge bool) objectivec.IObject
	RemoveAllModelsForBundleID(id objectivec.IObject) bool
	RemoveStaleModels() bool
	ScheduleMaintenanceWithNameDirectoryPaths(name objectivec.IObject, paths objectivec.IObject)
	ShouldEnforceSizeLimits() bool
	InitWithURL(url foundation.INSURL) ANEInMemoryModelCacheManager
	InitWithURLCreateDirectory(url foundation.INSURL, directory bool) ANEInMemoryModelCacheManager
}

// Init initializes the instance.
func (a ANEInMemoryModelCacheManager) Init() ANEInMemoryModelCacheManager {
	rv := objc.Send[ANEInMemoryModelCacheManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEInMemoryModelCacheManager) Autorelease() ANEInMemoryModelCacheManager {
	rv := objc.Send[ANEInMemoryModelCacheManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEInMemoryModelCacheManager creates a new ANEInMemoryModelCacheManager instance.
func NewANEInMemoryModelCacheManager() ANEInMemoryModelCacheManager {
	class := getANEInMemoryModelCacheManagerClass()
	rv := objc.Send[ANEInMemoryModelCacheManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/initWithURL:
func NewANEInMemoryModelCacheManagerWithURL(url foundation.INSURL) ANEInMemoryModelCacheManager {
	instance := getANEInMemoryModelCacheManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return ANEInMemoryModelCacheManagerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/initWithURL:createDirectory:
func NewANEInMemoryModelCacheManagerWithURLCreateDirectory(url foundation.INSURL, directory bool) ANEInMemoryModelCacheManager {
	instance := getANEInMemoryModelCacheManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:createDirectory:"), url, directory)
	return ANEInMemoryModelCacheManagerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/URLForBundleID:
func (a ANEInMemoryModelCacheManager) URLForBundleID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForBundleID:"), id)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/URLForModelHash:bundleID:
func (a ANEInMemoryModelCacheManager) URLForModelHashBundleID(hash objectivec.IObject, id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModelHash:bundleID:"), hash, id)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/cachedModelPathMatchingHash:csIdentity:
func (a ANEInMemoryModelCacheManager) CachedModelPathMatchingHashCsIdentity(hash objectivec.IObject, identity objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cachedModelPathMatchingHash:csIdentity:"), hash, identity)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/getDiskSpaceForBundleID:
func (a ANEInMemoryModelCacheManager) GetDiskSpaceForBundleID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getDiskSpaceForBundleID:"), id)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/getDiskSpaceItemizedByBundleIDAndPurge:
func (a ANEInMemoryModelCacheManager) GetDiskSpaceItemizedByBundleIDAndPurge(purge bool) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getDiskSpaceItemizedByBundleIDAndPurge:"), purge)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/removeAllModelsForBundleID:
func (a ANEInMemoryModelCacheManager) RemoveAllModelsForBundleID(id objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removeAllModelsForBundleID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/removeStaleModels
func (a ANEInMemoryModelCacheManager) RemoveStaleModels() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removeStaleModels"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/scheduleMaintenanceWithName:directoryPaths:
func (a ANEInMemoryModelCacheManager) ScheduleMaintenanceWithNameDirectoryPaths(name objectivec.IObject, paths objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleMaintenanceWithName:directoryPaths:"), name, paths)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/shouldEnforceSizeLimits
func (a ANEInMemoryModelCacheManager) ShouldEnforceSizeLimits() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldEnforceSizeLimits"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/initWithURL:
func (a ANEInMemoryModelCacheManager) InitWithURL(url foundation.INSURL) ANEInMemoryModelCacheManager {
	rv := objc.Send[ANEInMemoryModelCacheManager](a.ID, objc.Sel("initWithURL:"), url)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/initWithURL:createDirectory:
func (a ANEInMemoryModelCacheManager) InitWithURLCreateDirectory(url foundation.INSURL, directory bool) ANEInMemoryModelCacheManager {
	rv := objc.Send[ANEInMemoryModelCacheManager](a.ID, objc.Sel("initWithURL:createDirectory:"), url, directory)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/notRecentlyUsedSecondsThreshold
func (_ANEInMemoryModelCacheManagerClass ANEInMemoryModelCacheManagerClass) NotRecentlyUsedSecondsThreshold() uint64 {
	rv := objc.Send[uint64](objc.ID(_ANEInMemoryModelCacheManagerClass.class), objc.Sel("notRecentlyUsedSecondsThreshold"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/removeFilesFromDirectory:notAccessedInSeconds:
func (_ANEInMemoryModelCacheManagerClass ANEInMemoryModelCacheManagerClass) RemoveFilesFromDirectoryNotAccessedInSeconds(directory objectivec.IObject, seconds float64) bool {
	rv := objc.Send[bool](objc.ID(_ANEInMemoryModelCacheManagerClass.class), objc.Sel("removeFilesFromDirectory:notAccessedInSeconds:"), directory, seconds)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/removeStaleModelsAtPath:
func (_ANEInMemoryModelCacheManagerClass ANEInMemoryModelCacheManagerClass) RemoveStaleModelsAtPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEInMemoryModelCacheManagerClass.class), objc.Sel("removeStaleModelsAtPath:"), path)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelCacheManager/cacheDir
func (a ANEInMemoryModelCacheManager) CacheDir() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cacheDir"))
	return foundation.NSURLFromID(objc.ID(rv))
}

