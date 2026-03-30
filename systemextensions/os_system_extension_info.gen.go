// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSSystemExtensionInfo] class.
var (
	_OSSystemExtensionInfoClass     OSSystemExtensionInfoClass
	_OSSystemExtensionInfoClassOnce sync.Once
)

func getOSSystemExtensionInfoClass() OSSystemExtensionInfoClass {
	_OSSystemExtensionInfoClassOnce.Do(func() {
		_OSSystemExtensionInfoClass = OSSystemExtensionInfoClass{class: objc.GetClass("OSSystemExtensionInfo")}
	})
	return _OSSystemExtensionInfoClass
}

// GetOSSystemExtensionInfoClass returns the class object for OSSystemExtensionInfo.
func GetOSSystemExtensionInfoClass() OSSystemExtensionInfoClass {
	return getOSSystemExtensionInfoClass()
}

type OSSystemExtensionInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSSystemExtensionInfoClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSSystemExtensionInfoClass) Alloc() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [OSSystemExtensionInfo.BundleIdentifier]
//   - [OSSystemExtensionInfo.BundleShortVersion]
//   - [OSSystemExtensionInfo.BundleVersion]
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo
type OSSystemExtensionInfo struct {
	objectivec.Object
}

// OSSystemExtensionInfoFromID constructs a [OSSystemExtensionInfo] from an objc.ID.
func OSSystemExtensionInfoFromID(id objc.ID) OSSystemExtensionInfo {
	return OSSystemExtensionInfo{objectivec.Object{ID: id}}
}

// NOTE: OSSystemExtensionInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSSystemExtensionInfo] class.
//
// # Instance Properties
//
//   - [IOSSystemExtensionInfo.BundleIdentifier]
//   - [IOSSystemExtensionInfo.BundleShortVersion]
//   - [IOSSystemExtensionInfo.BundleVersion]
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo
type IOSSystemExtensionInfo interface {
	objectivec.IObject

	// Topic: Instance Properties

	BundleIdentifier() string
	BundleShortVersion() string
	BundleVersion() string
}

// Init initializes the instance.
func (o OSSystemExtensionInfo) Init() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSSystemExtensionInfo) Autorelease() OSSystemExtensionInfo {
	rv := objc.Send[OSSystemExtensionInfo](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionInfo creates a new OSSystemExtensionInfo instance.
func NewOSSystemExtensionInfo() OSSystemExtensionInfo {
	class := getOSSystemExtensionInfoClass()
	rv := objc.Send[OSSystemExtensionInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// The bundle identifier of the extension (CFBundleIdentifier)
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo/bundleIdentifier
func (o OSSystemExtensionInfo) BundleIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// # Discussion
//
// The bundle short version string of the extension
// (CFBundleShortVersionString)
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo/bundleShortVersion
func (o OSSystemExtensionInfo) BundleShortVersion() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bundleShortVersion"))
	return foundation.NSStringFromID(rv).String()
}

// # Discussion
//
// The bundle version of the extension (CFBundleVersion)
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionInfo/bundleVersion
func (o OSSystemExtensionInfo) BundleVersion() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bundleVersion"))
	return foundation.NSStringFromID(rv).String()
}
