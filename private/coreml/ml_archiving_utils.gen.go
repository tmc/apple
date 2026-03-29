// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLArchivingUtils] class.
var (
	_MLArchivingUtilsClass     MLArchivingUtilsClass
	_MLArchivingUtilsClassOnce sync.Once
)

func getMLArchivingUtilsClass() MLArchivingUtilsClass {
	_MLArchivingUtilsClassOnce.Do(func() {
		_MLArchivingUtilsClass = MLArchivingUtilsClass{class: objc.GetClass("MLArchivingUtils")}
	})
	return _MLArchivingUtilsClass
}

// GetMLArchivingUtilsClass returns the class object for MLArchivingUtils.
func GetMLArchivingUtilsClass() MLArchivingUtilsClass {
	return getMLArchivingUtilsClass()
}

type MLArchivingUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLArchivingUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLArchivingUtilsClass) Alloc() MLArchivingUtils {
	rv := objc.Send[MLArchivingUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLArchivingUtils
type MLArchivingUtils struct {
	objectivec.Object
}

// MLArchivingUtilsFromID constructs a [MLArchivingUtils] from an objc.ID.
func MLArchivingUtilsFromID(id objc.ID) MLArchivingUtils {
	return MLArchivingUtils{objectivec.Object{ID: id}}
}
// Ensure MLArchivingUtils implements IMLArchivingUtils.
var _ IMLArchivingUtils = MLArchivingUtils{}

// An interface definition for the [MLArchivingUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLArchivingUtils
type IMLArchivingUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a MLArchivingUtils) Init() MLArchivingUtils {
	rv := objc.Send[MLArchivingUtils](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLArchivingUtils) Autorelease() MLArchivingUtils {
	rv := objc.Send[MLArchivingUtils](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLArchivingUtils creates a new MLArchivingUtils instance.
func NewMLArchivingUtils() MLArchivingUtils {
	class := getMLArchivingUtilsClass()
	rv := objc.Send[MLArchivingUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLArchivingUtils/URLOfInputArchive:
func (_MLArchivingUtilsClass MLArchivingUtilsClass) URLOfInputArchive(archive unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLArchivingUtilsClass.class), objc.Sel("URLOfInputArchive:"), archive)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLArchivingUtils/codedObjectURLFromInputArchiver:
func (_MLArchivingUtilsClass MLArchivingUtilsClass) CodedObjectURLFromInputArchiver(archiver unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLArchivingUtilsClass.class), objc.Sel("codedObjectURLFromInputArchiver:"), archiver)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLArchivingUtils/codedObjectURLFromOutputArchiver:
func (_MLArchivingUtilsClass MLArchivingUtilsClass) CodedObjectURLFromOutputArchiver(archiver unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLArchivingUtilsClass.class), objc.Sel("codedObjectURLFromOutputArchiver:"), archiver)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLArchivingUtils/parseModelArchive:modelType:compilerVersion:modelVersion:error:
func (_MLArchivingUtilsClass MLArchivingUtilsClass) ParseModelArchiveModelTypeCompilerVersionModelVersionError(archive unsafe.Pointer, version []objectivec.IObject, version2 []objectivec.IObject) (int, error) {
	var type_ int
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLArchivingUtilsClass.class), objc.Sel("parseModelArchive:modelType:compilerVersion:modelVersion:error:"), archive, unsafe.Pointer(&type_), version, version2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("parseModelArchive:modelType:compilerVersion:modelVersion:error: returned NO with nil NSError")
	}
	return type_, nil
}

