// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEDeviceInfo] class.
var (
	_ANEDeviceInfoClass     ANEDeviceInfoClass
	_ANEDeviceInfoClassOnce sync.Once
)

func getANEDeviceInfoClass() ANEDeviceInfoClass {
	_ANEDeviceInfoClassOnce.Do(func() {
		_ANEDeviceInfoClass = ANEDeviceInfoClass{class: objc.GetClass("_ANEDeviceInfo")}
	})
	return _ANEDeviceInfoClass
}

// GetANEDeviceInfoClass returns the class object for _ANEDeviceInfo.
func GetANEDeviceInfoClass() ANEDeviceInfoClass {
	return getANEDeviceInfoClass()
}

type ANEDeviceInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEDeviceInfoClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEDeviceInfoClass) Alloc() ANEDeviceInfo {
	rv := objc.Send[ANEDeviceInfo](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo
type ANEDeviceInfo struct {
	objectivec.Object
}

// ANEDeviceInfoFromID constructs a [ANEDeviceInfo] from an objc.ID.
func ANEDeviceInfoFromID(id objc.ID) ANEDeviceInfo {
	return ANEDeviceInfo{objectivec.Object{ID: id}}
}

// Ensure ANEDeviceInfo implements IANEDeviceInfo.
var _ IANEDeviceInfo = ANEDeviceInfo{}

// An interface definition for the [ANEDeviceInfo] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo
type IANEDeviceInfo interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEDeviceInfo) Init() ANEDeviceInfo {
	rv := objc.Send[ANEDeviceInfo](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEDeviceInfo) Autorelease() ANEDeviceInfo {
	rv := objc.Send[ANEDeviceInfo](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEDeviceInfo creates a new ANEDeviceInfo instance.
func NewANEDeviceInfo() ANEDeviceInfo {
	class := getANEDeviceInfoClass()
	rv := objc.Send[ANEDeviceInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/aneArchitectureType
func (_ANEDeviceInfoClass ANEDeviceInfoClass) AneArchitectureType() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("aneArchitectureType"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/aneBoardType
func (_ANEDeviceInfoClass ANEDeviceInfoClass) AneBoardType() int64 {
	rv := objc.Send[int64](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("aneBoardType"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/aneSubType
func (_ANEDeviceInfoClass ANEDeviceInfoClass) AneSubType() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("aneSubType"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/aneSubTypeProductVariant
func (_ANEDeviceInfoClass ANEDeviceInfoClass) AneSubTypeProductVariant() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("aneSubTypeProductVariant"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/aneSubTypeVariant
func (_ANEDeviceInfoClass ANEDeviceInfoClass) AneSubTypeVariant() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("aneSubTypeVariant"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/bootArgs
func (_ANEDeviceInfoClass ANEDeviceInfoClass) BootArgs() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("bootArgs"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/buildVersion
func (_ANEDeviceInfoClass ANEDeviceInfoClass) BuildVersion() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("buildVersion"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/hasANE
func (_ANEDeviceInfoClass ANEDeviceInfoClass) HasANE() bool {
	rv := objc.Send[bool](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("hasANE"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/isBoolBootArgSetTrue:
func (_ANEDeviceInfoClass ANEDeviceInfoClass) IsBoolBootArgSetTrue(true_ objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("isBoolBootArgSetTrue:"), true_)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/isBootArgPresent:
func (_ANEDeviceInfoClass ANEDeviceInfoClass) IsBootArgPresent(present objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("isBootArgPresent:"), present)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/isExcessivePowerDrainWhenIdle
func (_ANEDeviceInfoClass ANEDeviceInfoClass) IsExcessivePowerDrainWhenIdle() bool {
	rv := objc.Send[bool](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("isExcessivePowerDrainWhenIdle"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/isInternalBuild
func (_ANEDeviceInfoClass ANEDeviceInfoClass) IsInternalBuild() bool {
	rv := objc.Send[bool](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("isInternalBuild"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/isVirtualMachine
func (_ANEDeviceInfoClass ANEDeviceInfoClass) IsVirtualMachine() bool {
	rv := objc.Send[bool](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("isVirtualMachine"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/numANECores
func (_ANEDeviceInfoClass ANEDeviceInfoClass) NumANECores() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("numANECores"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/numANEs
func (_ANEDeviceInfoClass ANEDeviceInfoClass) NumANEs() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("numANEs"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/precompiledModelChecksDisabled
func (_ANEDeviceInfoClass ANEDeviceInfoClass) PrecompiledModelChecksDisabled() bool {
	rv := objc.Send[bool](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("precompiledModelChecksDisabled"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/productName
func (_ANEDeviceInfoClass ANEDeviceInfoClass) ProductName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("productName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDeviceInfo/aneSubTypeAndVariant
func (_ANEDeviceInfoClass ANEDeviceInfoClass) AneSubTypeAndVariant() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDeviceInfoClass.class), objc.Sel("aneSubTypeAndVariant"))
	return objectivec.Object{ID: rv}
}
