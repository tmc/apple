// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMetalDeviceChangeInfo] class.
var (
	_MLMetalDeviceChangeInfoClass     MLMetalDeviceChangeInfoClass
	_MLMetalDeviceChangeInfoClassOnce sync.Once
)

func getMLMetalDeviceChangeInfoClass() MLMetalDeviceChangeInfoClass {
	_MLMetalDeviceChangeInfoClassOnce.Do(func() {
		_MLMetalDeviceChangeInfoClass = MLMetalDeviceChangeInfoClass{class: objc.GetClass("MLMetalDeviceChangeInfo")}
	})
	return _MLMetalDeviceChangeInfoClass
}

// GetMLMetalDeviceChangeInfoClass returns the class object for MLMetalDeviceChangeInfo.
func GetMLMetalDeviceChangeInfoClass() MLMetalDeviceChangeInfoClass {
	return getMLMetalDeviceChangeInfoClass()
}

type MLMetalDeviceChangeInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMetalDeviceChangeInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMetalDeviceChangeInfoClass) Alloc() MLMetalDeviceChangeInfo {
	rv := objc.Send[MLMetalDeviceChangeInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLMetalDeviceChangeInfo.ChangeType]
//   - [MLMetalDeviceChangeInfo.MetalDevice]
//   - [MLMetalDeviceChangeInfo.InitWithMetalDeviceChangeType]
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceChangeInfo
type MLMetalDeviceChangeInfo struct {
	objectivec.Object
}

// MLMetalDeviceChangeInfoFromID constructs a [MLMetalDeviceChangeInfo] from an objc.ID.
func MLMetalDeviceChangeInfoFromID(id objc.ID) MLMetalDeviceChangeInfo {
	return MLMetalDeviceChangeInfo{objectivec.Object{ID: id}}
}
// Ensure MLMetalDeviceChangeInfo implements IMLMetalDeviceChangeInfo.
var _ IMLMetalDeviceChangeInfo = MLMetalDeviceChangeInfo{}

// An interface definition for the [MLMetalDeviceChangeInfo] class.
//
// # Methods
//
//   - [IMLMetalDeviceChangeInfo.ChangeType]
//   - [IMLMetalDeviceChangeInfo.MetalDevice]
//   - [IMLMetalDeviceChangeInfo.InitWithMetalDeviceChangeType]
//
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceChangeInfo
type IMLMetalDeviceChangeInfo interface {
	objectivec.IObject

	// Topic: Methods

	ChangeType() int64
	MetalDevice() objectivec.IObject
	InitWithMetalDeviceChangeType(device objectivec.IObject, type_ int64) MLMetalDeviceChangeInfo
}

// Init initializes the instance.
func (m MLMetalDeviceChangeInfo) Init() MLMetalDeviceChangeInfo {
	rv := objc.Send[MLMetalDeviceChangeInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMetalDeviceChangeInfo) Autorelease() MLMetalDeviceChangeInfo {
	rv := objc.Send[MLMetalDeviceChangeInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMetalDeviceChangeInfo creates a new MLMetalDeviceChangeInfo instance.
func NewMLMetalDeviceChangeInfo() MLMetalDeviceChangeInfo {
	class := getMLMetalDeviceChangeInfoClass()
	rv := objc.Send[MLMetalDeviceChangeInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceChangeInfo/initWithMetalDevice:changeType:
func NewMetalDeviceChangeInfoWithMetalDeviceChangeType(device objectivec.IObject, type_ int64) MLMetalDeviceChangeInfo {
	instance := getMLMetalDeviceChangeInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMetalDevice:changeType:"), device, type_)
	return MLMetalDeviceChangeInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceChangeInfo/initWithMetalDevice:changeType:
func (m MLMetalDeviceChangeInfo) InitWithMetalDeviceChangeType(device objectivec.IObject, type_ int64) MLMetalDeviceChangeInfo {
	rv := objc.Send[MLMetalDeviceChangeInfo](m.ID, objc.Sel("initWithMetalDevice:changeType:"), device, type_)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceChangeInfo/changeType
func (m MLMetalDeviceChangeInfo) ChangeType() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("changeType"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceChangeInfo/metalDevice
func (m MLMetalDeviceChangeInfo) MetalDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metalDevice"))
	return objectivec.Object{ID: rv}
}

