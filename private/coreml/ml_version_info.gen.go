// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLVersionInfo] class.
var (
	_MLVersionInfoClass     MLVersionInfoClass
	_MLVersionInfoClassOnce sync.Once
)

func getMLVersionInfoClass() MLVersionInfoClass {
	_MLVersionInfoClassOnce.Do(func() {
		_MLVersionInfoClass = MLVersionInfoClass{class: objc.GetClass("MLVersionInfo")}
	})
	return _MLVersionInfoClass
}

// GetMLVersionInfoClass returns the class object for MLVersionInfo.
func GetMLVersionInfoClass() MLVersionInfoClass {
	return getMLVersionInfoClass()
}

type MLVersionInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLVersionInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLVersionInfoClass) Alloc() MLVersionInfo {
	rv := objc.Send[MLVersionInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLVersionInfo.MajorVersion]
//   - [MLVersionInfo.MinorVersion]
//   - [MLVersionInfo.OlderThan]
//   - [MLVersionInfo.PatchVersion]
//   - [MLVersionInfo.VariantString]
//   - [MLVersionInfo.VersionNumberString]
//   - [MLVersionInfo.VersionString]
//   - [MLVersionInfo.InitWithMajorMinorPatchVariant]
//
// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo
type MLVersionInfo struct {
	objectivec.Object
}

// MLVersionInfoFromID constructs a [MLVersionInfo] from an objc.ID.
func MLVersionInfoFromID(id objc.ID) MLVersionInfo {
	return MLVersionInfo{objectivec.Object{ID: id}}
}

// Ensure MLVersionInfo implements IMLVersionInfo.
var _ IMLVersionInfo = MLVersionInfo{}

// An interface definition for the [MLVersionInfo] class.
//
// # Methods
//
//   - [IMLVersionInfo.MajorVersion]
//   - [IMLVersionInfo.MinorVersion]
//   - [IMLVersionInfo.OlderThan]
//   - [IMLVersionInfo.PatchVersion]
//   - [IMLVersionInfo.VariantString]
//   - [IMLVersionInfo.VersionNumberString]
//   - [IMLVersionInfo.VersionString]
//   - [IMLVersionInfo.InitWithMajorMinorPatchVariant]
//
// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo
type IMLVersionInfo interface {
	objectivec.IObject

	// Topic: Methods

	MajorVersion() int64
	MinorVersion() int64
	OlderThan(than objectivec.IObject) bool
	PatchVersion() int64
	VariantString() string
	VersionNumberString() string
	VersionString() string
	InitWithMajorMinorPatchVariant(major int64, minor int64, patch int64, variant objectivec.IObject) MLVersionInfo
}

// Init initializes the instance.
func (v MLVersionInfo) Init() MLVersionInfo {
	rv := objc.Send[MLVersionInfo](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MLVersionInfo) Autorelease() MLVersionInfo {
	rv := objc.Send[MLVersionInfo](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLVersionInfo creates a new MLVersionInfo instance.
func NewMLVersionInfo() MLVersionInfo {
	class := getMLVersionInfoClass()
	rv := objc.Send[MLVersionInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/initWithMajor:minor:patch:variant:
func NewVersionInfoWithMajorMinorPatchVariant(major int64, minor int64, patch int64, variant objectivec.IObject) MLVersionInfo {
	instance := getMLVersionInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMajor:minor:patch:variant:"), major, minor, patch, variant)
	return MLVersionInfoFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/olderThan:
func (v MLVersionInfo) OlderThan(than objectivec.IObject) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("olderThan:"), than)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/initWithMajor:minor:patch:variant:
func (v MLVersionInfo) InitWithMajorMinorPatchVariant(major int64, minor int64, patch int64, variant objectivec.IObject) MLVersionInfo {
	rv := objc.Send[MLVersionInfo](v.ID, objc.Sel("initWithMajor:minor:patch:variant:"), major, minor, patch, variant)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/versionInfoWithMajor:minor:patch:variant:
func (_MLVersionInfoClass MLVersionInfoClass) VersionInfoWithMajorMinorPatchVariant(major int64, minor int64, patch int64, variant objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLVersionInfoClass.class), objc.Sel("versionInfoWithMajor:minor:patch:variant:"), major, minor, patch, variant)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/versionInfoWithString:
func (_MLVersionInfoClass MLVersionInfoClass) VersionInfoWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLVersionInfoClass.class), objc.Sel("versionInfoWithString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/versionInfoWithStringProgressive:
func (_MLVersionInfoClass MLVersionInfoClass) VersionInfoWithStringProgressive(progressive objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLVersionInfoClass.class), objc.Sel("versionInfoWithStringProgressive:"), progressive)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/majorVersion
func (v MLVersionInfo) MajorVersion() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("majorVersion"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/minorVersion
func (v MLVersionInfo) MinorVersion() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("minorVersion"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/patchVersion
func (v MLVersionInfo) PatchVersion() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("patchVersion"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/variantString
func (v MLVersionInfo) VariantString() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("variantString"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/versionNumberString
func (v MLVersionInfo) VersionNumberString() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("versionNumberString"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLVersionInfo/versionString
func (v MLVersionInfo) VersionString() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("versionString"))
	return foundation.NSStringFromID(rv).String()
}
