// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLReporterUtils] class.
var (
	_MLReporterUtilsClass     MLReporterUtilsClass
	_MLReporterUtilsClassOnce sync.Once
)

func getMLReporterUtilsClass() MLReporterUtilsClass {
	_MLReporterUtilsClassOnce.Do(func() {
		_MLReporterUtilsClass = MLReporterUtilsClass{class: objc.GetClass("MLReporterUtils")}
	})
	return _MLReporterUtilsClass
}

// GetMLReporterUtilsClass returns the class object for MLReporterUtils.
func GetMLReporterUtilsClass() MLReporterUtilsClass {
	return getMLReporterUtilsClass()
}

type MLReporterUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLReporterUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLReporterUtilsClass) Alloc() MLReporterUtils {
	rv := objc.Send[MLReporterUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLReporterUtils
type MLReporterUtils struct {
	objectivec.Object
}

// MLReporterUtilsFromID constructs a [MLReporterUtils] from an objc.ID.
func MLReporterUtilsFromID(id objc.ID) MLReporterUtils {
	return MLReporterUtils{objectivec.Object{ID: id}}
}

// Ensure MLReporterUtils implements IMLReporterUtils.
var _ IMLReporterUtils = MLReporterUtils{}

// An interface definition for the [MLReporterUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLReporterUtils
type IMLReporterUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r MLReporterUtils) Init() MLReporterUtils {
	rv := objc.Send[MLReporterUtils](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLReporterUtils) Autorelease() MLReporterUtils {
	rv := objc.Send[MLReporterUtils](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLReporterUtils creates a new MLReporterUtils instance.
func NewMLReporterUtils() MLReporterUtils {
	class := getMLReporterUtilsClass()
	rv := objc.Send[MLReporterUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLReporterUtils/archiveModelDetails:withName:toArchive:error:
func (_MLReporterUtilsClass MLReporterUtilsClass) ArchiveModelDetailsWithNameToArchiveError(details unsafe.Pointer, name unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLReporterUtilsClass.class), objc.Sel("archiveModelDetails:withName:toArchive:error:"), details, name, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("archiveModelDetails:withName:toArchive:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLReporterUtils/hashFileAt:error:
func (_MLReporterUtilsClass MLReporterUtilsClass) HashFileAtError(at objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLReporterUtilsClass.class), objc.Sel("hashFileAt:error:"), at, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
