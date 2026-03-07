// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEModelCacheManager] class.
var (
	_ANEModelCacheManagerClass     ANEModelCacheManagerClass
	_ANEModelCacheManagerClassOnce sync.Once
)

func getANEModelCacheManagerClass() ANEModelCacheManagerClass {
	_ANEModelCacheManagerClassOnce.Do(func() {
		_ANEModelCacheManagerClass = ANEModelCacheManagerClass{class: objc.GetClass("_ANEModelCacheManager")}
	})
	return _ANEModelCacheManagerClass
}

// GetANEModelCacheManagerClass returns the class object for _ANEModelCacheManager.
func GetANEModelCacheManagerClass() ANEModelCacheManagerClass {
	return getANEModelCacheManagerClass()
}

type ANEModelCacheManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEModelCacheManagerClass) Alloc() ANEModelCacheManager {
	rv := objc.Send[ANEModelCacheManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [ANEModelCacheManager.URLForBundleID]
//   - [ANEModelCacheManager.URLForModelBundleID]
//   - [ANEModelCacheManager.URLForModelBundleIDAotCacheUrlIdentifier]
//   - [ANEModelCacheManager.URLForModelBundleIDForAllSegments]
//   - [ANEModelCacheManager.URLForModelBundleIDForAllSegmentsAotCacheUrlIdentifier]
//   - [ANEModelCacheManager.URLForModelBundleIDUseSourceURL]
//   - [ANEModelCacheManager.URLForModelBundleIDUseSourceURLAotCacheUrlIdentifier]
//   - [ANEModelCacheManager.URLForModelBundleIDUseSourceURLForAllSegmentsAotCacheUrlIdentifier]
//   - [ANEModelCacheManager.CacheDir]
//   - [ANEModelCacheManager.CacheURLIdentifierForModelUseSourceURLWithReply]
//   - [ANEModelCacheManager.CachedModelAllSegmentsPathForCsIdentity]
//   - [ANEModelCacheManager.CachedModelPathForCsIdentity]
//   - [ANEModelCacheManager.CachedModelPathForCsIdentityUseSourceURL]
//   - [ANEModelCacheManager.CachedModelRetainNameForCsIdentity]
//   - [ANEModelCacheManager.CachedSourceModelStoreNameForCsIdentity]
//   - [ANEModelCacheManager.FilePathForModelBundleID]
//   - [ANEModelCacheManager.GarbageCollectDanglingModels]
//   - [ANEModelCacheManager.GetDiskSpaceForBundleID]
//   - [ANEModelCacheManager.GetDiskSpaceItemizedByBundleIDAndPurge]
//   - [ANEModelCacheManager.GetModelBinaryPathFromURLIdentifierBundleID]
//   - [ANEModelCacheManager.RemoveAllModelsForBundleID]
//   - [ANEModelCacheManager.ScanAllPartitionsForModelCsIdentityExpunge]
//   - [ANEModelCacheManager.ScheduleMaintenanceWithNameDirectoryPaths]
//   - [ANEModelCacheManager.ShouldEnforceSizeLimits]
//   - [ANEModelCacheManager.StartDanglingModelGC]
//   - [ANEModelCacheManager.InitWithURL]
//   - [ANEModelCacheManager.InitWithURLCreateDirectory]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager
type ANEModelCacheManager struct {
	objectivec.Object
}

// ANEModelCacheManagerFromID constructs a [ANEModelCacheManager] from an objc.ID.
func ANEModelCacheManagerFromID(id objc.ID) ANEModelCacheManager {
	return ANEModelCacheManager{objectivec.Object{id}}
}
// Ensure ANEModelCacheManager implements IANEModelCacheManager.
var _ IANEModelCacheManager = ANEModelCacheManager{}





// An interface definition for the [ANEModelCacheManager] class.
//
// # Methods
//
//   - [IANEModelCacheManager.URLForBundleID]
//   - [IANEModelCacheManager.URLForModelBundleID]
//   - [IANEModelCacheManager.URLForModelBundleIDAotCacheUrlIdentifier]
//   - [IANEModelCacheManager.URLForModelBundleIDForAllSegments]
//   - [IANEModelCacheManager.URLForModelBundleIDForAllSegmentsAotCacheUrlIdentifier]
//   - [IANEModelCacheManager.URLForModelBundleIDUseSourceURL]
//   - [IANEModelCacheManager.URLForModelBundleIDUseSourceURLAotCacheUrlIdentifier]
//   - [IANEModelCacheManager.URLForModelBundleIDUseSourceURLForAllSegmentsAotCacheUrlIdentifier]
//   - [IANEModelCacheManager.CacheDir]
//   - [IANEModelCacheManager.CacheURLIdentifierForModelUseSourceURLWithReply]
//   - [IANEModelCacheManager.CachedModelAllSegmentsPathForCsIdentity]
//   - [IANEModelCacheManager.CachedModelPathForCsIdentity]
//   - [IANEModelCacheManager.CachedModelPathForCsIdentityUseSourceURL]
//   - [IANEModelCacheManager.CachedModelRetainNameForCsIdentity]
//   - [IANEModelCacheManager.CachedSourceModelStoreNameForCsIdentity]
//   - [IANEModelCacheManager.FilePathForModelBundleID]
//   - [IANEModelCacheManager.GarbageCollectDanglingModels]
//   - [IANEModelCacheManager.GetDiskSpaceForBundleID]
//   - [IANEModelCacheManager.GetDiskSpaceItemizedByBundleIDAndPurge]
//   - [IANEModelCacheManager.GetModelBinaryPathFromURLIdentifierBundleID]
//   - [IANEModelCacheManager.RemoveAllModelsForBundleID]
//   - [IANEModelCacheManager.ScanAllPartitionsForModelCsIdentityExpunge]
//   - [IANEModelCacheManager.ScheduleMaintenanceWithNameDirectoryPaths]
//   - [IANEModelCacheManager.ShouldEnforceSizeLimits]
//   - [IANEModelCacheManager.StartDanglingModelGC]
//   - [IANEModelCacheManager.InitWithURL]
//   - [IANEModelCacheManager.InitWithURLCreateDirectory]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager
type IANEModelCacheManager interface {
	objectivec.IObject

	// Topic: Methods

	URLForBundleID(id objectivec.IObject) objectivec.IObject
	URLForModelBundleID(model objectivec.IObject, id objectivec.IObject) objectivec.IObject
	URLForModelBundleIDAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, identifier objectivec.IObject) objectivec.IObject
	URLForModelBundleIDForAllSegments(model objectivec.IObject, id objectivec.IObject, segments bool) objectivec.IObject
	URLForModelBundleIDForAllSegmentsAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, segments bool, identifier objectivec.IObject) objectivec.IObject
	URLForModelBundleIDUseSourceURL(model objectivec.IObject, id objectivec.IObject, url bool) objectivec.IObject
	URLForModelBundleIDUseSourceURLAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, url bool, identifier objectivec.IObject) objectivec.IObject
	URLForModelBundleIDUseSourceURLForAllSegmentsAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, url bool, segments bool, identifier objectivec.IObject) objectivec.IObject
	CacheDir() foundation.INSURL
	CacheURLIdentifierForModelUseSourceURLWithReply(model objectivec.IObject, url bool, reply VoidHandler) bool
	CachedModelAllSegmentsPathForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject
	CachedModelPathForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject
	CachedModelPathForCsIdentityUseSourceURL(for_ objectivec.IObject, identity objectivec.IObject, url bool) objectivec.IObject
	CachedModelRetainNameForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject
	CachedSourceModelStoreNameForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject
	FilePathForModelBundleID(model objectivec.IObject, id objectivec.IObject) objectivec.IObject
	GarbageCollectDanglingModels() bool
	GetDiskSpaceForBundleID(id objectivec.IObject) objectivec.IObject
	GetDiskSpaceItemizedByBundleIDAndPurge(purge bool) objectivec.IObject
	GetModelBinaryPathFromURLIdentifierBundleID(uRLIdentifier objectivec.IObject, id objectivec.IObject) objectivec.IObject
	RemoveAllModelsForBundleID(id objectivec.IObject) bool
	ScanAllPartitionsForModelCsIdentityExpunge(model objectivec.IObject, identity objectivec.IObject, expunge bool) bool
	ScheduleMaintenanceWithNameDirectoryPaths(name objectivec.IObject, paths objectivec.IObject)
	ShouldEnforceSizeLimits() bool
	StartDanglingModelGC()
	InitWithURL(url foundation.INSURL) ANEModelCacheManager
	InitWithURLCreateDirectory(url foundation.INSURL, directory bool) ANEModelCacheManager
}





// Init initializes the instance.
func (a ANEModelCacheManager) Init() ANEModelCacheManager {
	rv := objc.Send[ANEModelCacheManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEModelCacheManager) Autorelease() ANEModelCacheManager {
	rv := objc.Send[ANEModelCacheManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEModelCacheManager creates a new ANEModelCacheManager instance.
func NewANEModelCacheManager() ANEModelCacheManager {
	class := getANEModelCacheManagerClass()
	rv := objc.Send[ANEModelCacheManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/initWithURL:
func NewANEModelCacheManagerWithURL(url foundation.INSURL) ANEModelCacheManager {
	instance := getANEModelCacheManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return ANEModelCacheManagerFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/initWithURL:createDirectory:
func NewANEModelCacheManagerWithURLCreateDirectory(url foundation.INSURL, directory bool) ANEModelCacheManager {
	instance := getANEModelCacheManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:createDirectory:"), url, directory)
	return ANEModelCacheManagerFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForBundleID:
func (a ANEModelCacheManager) URLForBundleID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForBundleID:"), id)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForModel:bundleID:
func (a ANEModelCacheManager) URLForModelBundleID(model objectivec.IObject, id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModel:bundleID:"), model, id)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForModel:bundleID:aotCacheUrlIdentifier:
func (a ANEModelCacheManager) URLForModelBundleIDAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModel:bundleID:aotCacheUrlIdentifier:"), model, id, identifier)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForModel:bundleID:forAllSegments:
func (a ANEModelCacheManager) URLForModelBundleIDForAllSegments(model objectivec.IObject, id objectivec.IObject, segments bool) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModel:bundleID:forAllSegments:"), model, id, segments)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForModel:bundleID:forAllSegments:aotCacheUrlIdentifier:
func (a ANEModelCacheManager) URLForModelBundleIDForAllSegmentsAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, segments bool, identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModel:bundleID:forAllSegments:aotCacheUrlIdentifier:"), model, id, segments, identifier)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForModel:bundleID:useSourceURL:
func (a ANEModelCacheManager) URLForModelBundleIDUseSourceURL(model objectivec.IObject, id objectivec.IObject, url bool) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModel:bundleID:useSourceURL:"), model, id, url)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForModel:bundleID:useSourceURL:aotCacheUrlIdentifier:
func (a ANEModelCacheManager) URLForModelBundleIDUseSourceURLAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, url bool, identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModel:bundleID:useSourceURL:aotCacheUrlIdentifier:"), model, id, url, identifier)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/URLForModel:bundleID:useSourceURL:forAllSegments:aotCacheUrlIdentifier:
func (a ANEModelCacheManager) URLForModelBundleIDUseSourceURLForAllSegmentsAotCacheUrlIdentifier(model objectivec.IObject, id objectivec.IObject, url bool, segments bool, identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLForModel:bundleID:useSourceURL:forAllSegments:aotCacheUrlIdentifier:"), model, id, url, segments, identifier)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cacheURLIdentifierForModel:useSourceURL:withReply:
func (a ANEModelCacheManager) CacheURLIdentifierForModelUseSourceURLWithReply(model objectivec.IObject, url bool, reply VoidHandler) bool {
		_block2, _cleanup2 := NewVoidBlock(reply)
	defer _cleanup2()
		rv := objc.Send[bool](a.ID, objc.Sel("cacheURLIdentifierForModel:useSourceURL:withReply:"), model, url, _block2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cachedModelAllSegmentsPathFor:csIdentity:
func (a ANEModelCacheManager) CachedModelAllSegmentsPathForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cachedModelAllSegmentsPathFor:csIdentity:"), for_, identity)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cachedModelPathFor:csIdentity:
func (a ANEModelCacheManager) CachedModelPathForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cachedModelPathFor:csIdentity:"), for_, identity)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cachedModelPathFor:csIdentity:useSourceURL:
func (a ANEModelCacheManager) CachedModelPathForCsIdentityUseSourceURL(for_ objectivec.IObject, identity objectivec.IObject, url bool) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cachedModelPathFor:csIdentity:useSourceURL:"), for_, identity, url)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cachedModelRetainNameFor:csIdentity:
func (a ANEModelCacheManager) CachedModelRetainNameForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cachedModelRetainNameFor:csIdentity:"), for_, identity)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cachedSourceModelStoreNameFor:csIdentity:
func (a ANEModelCacheManager) CachedSourceModelStoreNameForCsIdentity(for_ objectivec.IObject, identity objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cachedSourceModelStoreNameFor:csIdentity:"), for_, identity)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/filePathForModel:bundleID:
func (a ANEModelCacheManager) FilePathForModelBundleID(model objectivec.IObject, id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("filePathForModel:bundleID:"), model, id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/garbageCollectDanglingModels
func (a ANEModelCacheManager) GarbageCollectDanglingModels() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("garbageCollectDanglingModels"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/getDiskSpaceForBundleID:
func (a ANEModelCacheManager) GetDiskSpaceForBundleID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getDiskSpaceForBundleID:"), id)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/getDiskSpaceItemizedByBundleIDAndPurge:
func (a ANEModelCacheManager) GetDiskSpaceItemizedByBundleIDAndPurge(purge bool) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getDiskSpaceItemizedByBundleIDAndPurge:"), purge)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/getModelBinaryPathFromURLIdentifier:bundleID:
func (a ANEModelCacheManager) GetModelBinaryPathFromURLIdentifierBundleID(uRLIdentifier objectivec.IObject, id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getModelBinaryPathFromURLIdentifier:bundleID:"), uRLIdentifier, id)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/removeAllModelsForBundleID:
func (a ANEModelCacheManager) RemoveAllModelsForBundleID(id objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removeAllModelsForBundleID:"), id)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/scanAllPartitionsForModel:csIdentity:expunge:
func (a ANEModelCacheManager) ScanAllPartitionsForModelCsIdentityExpunge(model objectivec.IObject, identity objectivec.IObject, expunge bool) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("scanAllPartitionsForModel:csIdentity:expunge:"), model, identity, expunge)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/scheduleMaintenanceWithName:directoryPaths:
func (a ANEModelCacheManager) ScheduleMaintenanceWithNameDirectoryPaths(name objectivec.IObject, paths objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleMaintenanceWithName:directoryPaths:"), name, paths)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/shouldEnforceSizeLimits
func (a ANEModelCacheManager) ShouldEnforceSizeLimits() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldEnforceSizeLimits"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/startDanglingModelGC
func (a ANEModelCacheManager) StartDanglingModelGC() {
	objc.Send[objc.ID](a.ID, objc.Sel("startDanglingModelGC"))
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/initWithURL:
func (a ANEModelCacheManager) InitWithURL(url foundation.INSURL) ANEModelCacheManager {
	rv := objc.Send[ANEModelCacheManager](a.ID, objc.Sel("initWithURL:"), url)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/initWithURL:createDirectory:
func (a ANEModelCacheManager) InitWithURLCreateDirectory(url foundation.INSURL, directory bool) ANEModelCacheManager {
	rv := objc.Send[ANEModelCacheManager](a.ID, objc.Sel("initWithURL:createDirectory:"), url, directory)
	return rv
}





//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cachedModelRetainNameFor:
func (_ANEModelCacheManagerClass ANEModelCacheManagerClass) CachedModelRetainNameFor(for_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelCacheManagerClass.class), objc.Sel("cachedModelRetainNameFor:"), for_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cachedSourceModelStoreNameFor:
func (_ANEModelCacheManagerClass ANEModelCacheManagerClass) CachedSourceModelStoreNameFor(for_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelCacheManagerClass.class), objc.Sel("cachedSourceModelStoreNameFor:"), for_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/createModelCacheRetain:
func (_ANEModelCacheManagerClass ANEModelCacheManagerClass) CreateModelCacheRetain(retain objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEModelCacheManagerClass.class), objc.Sel("createModelCacheRetain:"), retain)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/isSystemModelPath:
func (_ANEModelCacheManagerClass ANEModelCacheManagerClass) IsSystemModelPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEModelCacheManagerClass.class), objc.Sel("isSystemModelPath:"), path)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/removeIfStaleBinary:forModelPath:
func (_ANEModelCacheManagerClass ANEModelCacheManagerClass) RemoveIfStaleBinaryForModelPath(binary objectivec.IObject, path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEModelCacheManagerClass.class), objc.Sel("removeIfStaleBinary:forModelPath:"), binary, path)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/saveSourceModelPath:outputModelDirectory:
func (_ANEModelCacheManagerClass ANEModelCacheManagerClass) SaveSourceModelPathOutputModelDirectory(path objectivec.IObject, directory objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEModelCacheManagerClass.class), objc.Sel("saveSourceModelPath:outputModelDirectory:"), path, directory)
	return rv
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModelCacheManager/cacheDir
func (a ANEModelCacheManager) CacheDir() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cacheDir"))
	return foundation.NSURLFromID(objc.ID(rv))
}












// CacheURLIdentifierForModelUseSourceURLWithReplySync is a synchronous wrapper around [ANEModelCacheManager.CacheURLIdentifierForModelUseSourceURLWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEModelCacheManager) CacheURLIdentifierForModelUseSourceURLWithReplySync(ctx context.Context, model objectivec.IObject, url bool) error {
	done := make(chan struct{}, 1)
	a.CacheURLIdentifierForModelUseSourceURLWithReply(model, url, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






