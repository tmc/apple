// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEStorageHelper] class.
var (
	_ANEStorageHelperClass     ANEStorageHelperClass
	_ANEStorageHelperClassOnce sync.Once
)

func getANEStorageHelperClass() ANEStorageHelperClass {
	_ANEStorageHelperClassOnce.Do(func() {
		_ANEStorageHelperClass = ANEStorageHelperClass{class: objc.GetClass("_ANEStorageHelper")}
	})
	return _ANEStorageHelperClass
}

// GetANEStorageHelperClass returns the class object for _ANEStorageHelper.
func GetANEStorageHelperClass() ANEStorageHelperClass {
	return getANEStorageHelperClass()
}

type ANEStorageHelperClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEStorageHelperClass) Alloc() ANEStorageHelper {
	rv := objc.Send[ANEStorageHelper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper
type ANEStorageHelper struct {
	objectivec.Object
}

// ANEStorageHelperFromID constructs a [ANEStorageHelper] from an objc.ID.
func ANEStorageHelperFromID(id objc.ID) ANEStorageHelper {
	return ANEStorageHelper{objectivec.Object{ID: id}}
}
// Ensure ANEStorageHelper implements IANEStorageHelper.
var _ IANEStorageHelper = ANEStorageHelper{}

// An interface definition for the [ANEStorageHelper] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper
type IANEStorageHelper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEStorageHelper) Init() ANEStorageHelper {
	rv := objc.Send[ANEStorageHelper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEStorageHelper) Autorelease() ANEStorageHelper {
	rv := objc.Send[ANEStorageHelper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEStorageHelper creates a new ANEStorageHelper instance.
func NewANEStorageHelper() ANEStorageHelper {
	class := getANEStorageHelperClass()
	rv := objc.Send[ANEStorageHelper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/addSubdirectoryDetails:directoryPath:size:
func (_ANEStorageHelperClass ANEStorageHelperClass) AddSubdirectoryDetailsDirectoryPathSize(details objectivec.IObject, path objectivec.IObject, size uint64) {
	objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("addSubdirectoryDetails:directoryPath:size:"), details, path, size)
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/createModelCacheDictionary
func (_ANEStorageHelperClass ANEStorageHelperClass) CreateModelCacheDictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("createModelCacheDictionary"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/enableApfsPurging
func (_ANEStorageHelperClass ANEStorageHelperClass) EnableApfsPurging() bool {
	rv := objc.Send[bool](objc.ID(_ANEStorageHelperClass.class), objc.Sel("enableApfsPurging"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/garbageCollectDanglingModelsAtPath:
func (_ANEStorageHelperClass ANEStorageHelperClass) GarbageCollectDanglingModelsAtPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEStorageHelperClass.class), objc.Sel("garbageCollectDanglingModelsAtPath:"), path)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/getAccessTimeForFilePath:
func (_ANEStorageHelperClass ANEStorageHelperClass) GetAccessTimeForFilePath(path objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("getAccessTimeForFilePath:"), path)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/markPathAndDirectParentPurgeable:error:
func (_ANEStorageHelperClass ANEStorageHelperClass) MarkPathAndDirectParentPurgeableError(purgeable objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_ANEStorageHelperClass.class), objc.Sel("markPathAndDirectParentPurgeable:error:"), purgeable, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("markPathAndDirectParentPurgeable:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/memoryMapModelAtPath:isPrecompiled:modelAttributes:
func (_ANEStorageHelperClass ANEStorageHelperClass) MemoryMapModelAtPathIsPrecompiledModelAttributes(path objectivec.IObject, precompiled bool, attributes []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("memoryMapModelAtPath:isPrecompiled:modelAttributes:"), path, precompiled, objectivec.IObjectSliceToNSArray(attributes))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/memoryMapModelAtPath:modelAttributes:
func (_ANEStorageHelperClass ANEStorageHelperClass) MemoryMapModelAtPathModelAttributes(path objectivec.IObject, attributes []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("memoryMapModelAtPath:modelAttributes:"), path, objectivec.IObjectSliceToNSArray(attributes))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/memoryMapWeightAtPath:
func (_ANEStorageHelperClass ANEStorageHelperClass) MemoryMapWeightAtPath(path objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("memoryMapWeightAtPath:"), path)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/mergeModelCacheStorageInformation:with:
func (_ANEStorageHelperClass ANEStorageHelperClass) MergeModelCacheStorageInformationWith(information objectivec.IObject, with objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("mergeModelCacheStorageInformation:with:"), information, with)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/removeDirectoryAtPath:
func (_ANEStorageHelperClass ANEStorageHelperClass) RemoveDirectoryAtPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEStorageHelperClass.class), objc.Sel("removeDirectoryAtPath:"), path)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/removeFilePath:ifDate:olderThanSecond:
func (_ANEStorageHelperClass ANEStorageHelperClass) RemoveFilePathIfDateOlderThanSecond(path objectivec.IObject, date objectivec.IObject, second objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("removeFilePath:ifDate:olderThanSecond:"), path, date, second)
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/removeShapesDirectoryAtPath:
func (_ANEStorageHelperClass ANEStorageHelperClass) RemoveShapesDirectoryAtPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEStorageHelperClass.class), objc.Sel("removeShapesDirectoryAtPath:"), path)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/setAccessTime:forModelFilePath:
func (_ANEStorageHelperClass ANEStorageHelperClass) SetAccessTimeForModelFilePath(time objectivec.IObject, path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEStorageHelperClass.class), objc.Sel("setAccessTime:forModelFilePath:"), time, path)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/sizeOfDirectoryAtPath:recursionLevel:
func (_ANEStorageHelperClass ANEStorageHelperClass) SizeOfDirectoryAtPathRecursionLevel(path objectivec.IObject, level uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("sizeOfDirectoryAtPath:recursionLevel:"), path, level)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/sizeOfModelCacheAtPath:purgeSubdirectories:
func (_ANEStorageHelperClass ANEStorageHelperClass) SizeOfModelCacheAtPathPurgeSubdirectories(path objectivec.IObject, subdirectories bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("sizeOfModelCacheAtPath:purgeSubdirectories:"), path, subdirectories)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/uniqueFirstLevelSubdirectories:
func (_ANEStorageHelperClass ANEStorageHelperClass) UniqueFirstLevelSubdirectories(subdirectories objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEStorageHelperClass.class), objc.Sel("uniqueFirstLevelSubdirectories:"), subdirectories)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageHelper/updateAccessTimeForFilePath:
func (_ANEStorageHelperClass ANEStorageHelperClass) UpdateAccessTimeForFilePath(path objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEStorageHelperClass.class), objc.Sel("updateAccessTimeForFilePath:"), path)
	return rv
}

