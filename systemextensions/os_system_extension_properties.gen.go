// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSSystemExtensionProperties] class.
var (
	_OSSystemExtensionPropertiesClass     OSSystemExtensionPropertiesClass
	_OSSystemExtensionPropertiesClassOnce sync.Once
)

func getOSSystemExtensionPropertiesClass() OSSystemExtensionPropertiesClass {
	_OSSystemExtensionPropertiesClassOnce.Do(func() {
		_OSSystemExtensionPropertiesClass = OSSystemExtensionPropertiesClass{class: objc.GetClass("OSSystemExtensionProperties")}
	})
	return _OSSystemExtensionPropertiesClass
}

// GetOSSystemExtensionPropertiesClass returns the class object for OSSystemExtensionProperties.
func GetOSSystemExtensionPropertiesClass() OSSystemExtensionPropertiesClass {
	return getOSSystemExtensionPropertiesClass()
}

type OSSystemExtensionPropertiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSSystemExtensionPropertiesClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSSystemExtensionPropertiesClass) Alloc() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// Properties that identify a specific version of a system extension.
//
// # Identifying the Extension
//
//   - [OSSystemExtensionProperties.BundleIdentifier]: The bundle identifier of the extension.
//   - [OSSystemExtensionProperties.BundleVersion]: The bundle version of the extension.
//   - [OSSystemExtensionProperties.BundleShortVersion]: The bundle short version string of the extension.
//
// # Locating the Extension’s Installed Location
//
//   - [OSSystemExtensionProperties.URL]: The file URL of the extension bundle.
//
// # Instance Properties
//
//   - [OSSystemExtensionProperties.IsAwaitingUserApproval]
//   - [OSSystemExtensionProperties.IsEnabled]
//   - [OSSystemExtensionProperties.IsUninstalling]
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties
type OSSystemExtensionProperties struct {
	objectivec.Object
}

// OSSystemExtensionPropertiesFromID constructs a [OSSystemExtensionProperties] from an objc.ID.
//
// Properties that identify a specific version of a system extension.
func OSSystemExtensionPropertiesFromID(id objc.ID) OSSystemExtensionProperties {
	return OSSystemExtensionProperties{objectivec.Object{ID: id}}
}

// NOTE: OSSystemExtensionProperties adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSSystemExtensionProperties] class.
//
// # Identifying the Extension
//
//   - [IOSSystemExtensionProperties.BundleIdentifier]: The bundle identifier of the extension.
//   - [IOSSystemExtensionProperties.BundleVersion]: The bundle version of the extension.
//   - [IOSSystemExtensionProperties.BundleShortVersion]: The bundle short version string of the extension.
//
// # Locating the Extension’s Installed Location
//
//   - [IOSSystemExtensionProperties.URL]: The file URL of the extension bundle.
//
// # Instance Properties
//
//   - [IOSSystemExtensionProperties.IsAwaitingUserApproval]
//   - [IOSSystemExtensionProperties.IsEnabled]
//   - [IOSSystemExtensionProperties.IsUninstalling]
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties
type IOSSystemExtensionProperties interface {
	objectivec.IObject

	// Topic: Identifying the Extension

	// The bundle identifier of the extension.
	BundleIdentifier() string
	// The bundle version of the extension.
	BundleVersion() string
	// The bundle short version string of the extension.
	BundleShortVersion() string

	// Topic: Locating the Extension’s Installed Location

	// The file URL of the extension bundle.
	URL() foundation.INSURL

	// Topic: Instance Properties

	IsAwaitingUserApproval() bool
	IsEnabled() bool
	IsUninstalling() bool
}

// Init initializes the instance.
func (o OSSystemExtensionProperties) Init() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSSystemExtensionProperties) Autorelease() OSSystemExtensionProperties {
	rv := objc.Send[OSSystemExtensionProperties](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionProperties creates a new OSSystemExtensionProperties instance.
func NewOSSystemExtensionProperties() OSSystemExtensionProperties {
	class := getOSSystemExtensionPropertiesClass()
	rv := objc.Send[OSSystemExtensionProperties](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The bundle identifier of the extension.
//
// # Discussion
//
// This is the [CFBundleIdentifier] of the extension bundle.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/bundleIdentifier
//
// [CFBundleIdentifier]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CFBundleIdentifier
func (o OSSystemExtensionProperties) BundleIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The bundle version of the extension.
//
// # Discussion
//
// This is the [CFBundleVersion] of the extension bundle.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/bundleVersion
//
// [CFBundleVersion]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CFBundleVersion
func (o OSSystemExtensionProperties) BundleVersion() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bundleVersion"))
	return foundation.NSStringFromID(rv).String()
}

// The bundle short version string of the extension.
//
// # Discussion
//
// This is the [CFBundleShortVersionString] of the extension bundle.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/bundleShortVersion
//
// [CFBundleShortVersionString]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CFBundleShortVersionString
func (o OSSystemExtensionProperties) BundleShortVersion() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bundleShortVersion"))
	return foundation.NSStringFromID(rv).String()
}

// The file URL of the extension bundle.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/url
func (o OSSystemExtensionProperties) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isAwaitingUserApproval
func (o OSSystemExtensionProperties) IsAwaitingUserApproval() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAwaitingUserApproval"))
	return rv
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isEnabled
func (o OSSystemExtensionProperties) IsEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isEnabled"))
	return rv
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionProperties/isUninstalling
func (o OSSystemExtensionProperties) IsUninstalling() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isUninstalling"))
	return rv
}
