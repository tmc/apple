// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCloudDeploymentUtils] class.
var (
	_MLCloudDeploymentUtilsClass     MLCloudDeploymentUtilsClass
	_MLCloudDeploymentUtilsClassOnce sync.Once
)

func getMLCloudDeploymentUtilsClass() MLCloudDeploymentUtilsClass {
	_MLCloudDeploymentUtilsClassOnce.Do(func() {
		_MLCloudDeploymentUtilsClass = MLCloudDeploymentUtilsClass{class: objc.GetClass("MLCloudDeploymentUtils")}
	})
	return _MLCloudDeploymentUtilsClass
}

// GetMLCloudDeploymentUtilsClass returns the class object for MLCloudDeploymentUtils.
func GetMLCloudDeploymentUtilsClass() MLCloudDeploymentUtilsClass {
	return getMLCloudDeploymentUtilsClass()
}

type MLCloudDeploymentUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCloudDeploymentUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCloudDeploymentUtilsClass) Alloc() MLCloudDeploymentUtils {
	rv := objc.Send[MLCloudDeploymentUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCloudDeploymentUtils
type MLCloudDeploymentUtils struct {
	objectivec.Object
}

// MLCloudDeploymentUtilsFromID constructs a [MLCloudDeploymentUtils] from an objc.ID.
func MLCloudDeploymentUtilsFromID(id objc.ID) MLCloudDeploymentUtils {
	return MLCloudDeploymentUtils{objectivec.Object{ID: id}}
}

// Ensure MLCloudDeploymentUtils implements IMLCloudDeploymentUtils.
var _ IMLCloudDeploymentUtils = MLCloudDeploymentUtils{}

// An interface definition for the [MLCloudDeploymentUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLCloudDeploymentUtils
type IMLCloudDeploymentUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c MLCloudDeploymentUtils) Init() MLCloudDeploymentUtils {
	rv := objc.Send[MLCloudDeploymentUtils](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCloudDeploymentUtils) Autorelease() MLCloudDeploymentUtils {
	rv := objc.Send[MLCloudDeploymentUtils](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCloudDeploymentUtils creates a new MLCloudDeploymentUtils instance.
func NewMLCloudDeploymentUtils() MLCloudDeploymentUtils {
	class := getMLCloudDeploymentUtilsClass()
	rv := objc.Send[MLCloudDeploymentUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCloudDeploymentUtils/extractTeamIdentifierAndReturnError:
func (_MLCloudDeploymentUtilsClass MLCloudDeploymentUtilsClass) ExtractTeamIdentifierAndReturnError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCloudDeploymentUtilsClass.class), objc.Sel("extractTeamIdentifierAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
