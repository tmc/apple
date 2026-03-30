// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMediaExtensionProperties] class.
var (
	_AVMediaExtensionPropertiesClass     AVMediaExtensionPropertiesClass
	_AVMediaExtensionPropertiesClassOnce sync.Once
)

func getAVMediaExtensionPropertiesClass() AVMediaExtensionPropertiesClass {
	_AVMediaExtensionPropertiesClassOnce.Do(func() {
		_AVMediaExtensionPropertiesClass = AVMediaExtensionPropertiesClass{class: objc.GetClass("AVMediaExtensionProperties")}
	})
	return _AVMediaExtensionPropertiesClass
}

// GetAVMediaExtensionPropertiesClass returns the class object for AVMediaExtensionProperties.
func GetAVMediaExtensionPropertiesClass() AVMediaExtensionPropertiesClass {
	return getAVMediaExtensionPropertiesClass()
}

type AVMediaExtensionPropertiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMediaExtensionPropertiesClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMediaExtensionPropertiesClass) Alloc() AVMediaExtensionProperties {
	rv := objc.Send[AVMediaExtensionProperties](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a Media Extension.
//
// # Inspecting the extension
//
//   - [AVMediaExtensionProperties.ExtensionName]: The name of the Media Extension.
//   - [AVMediaExtensionProperties.ContainingBundleName]: The name of the containing app bundle.
//   - [AVMediaExtensionProperties.ExtensionIdentifier]
//   - [AVMediaExtensionProperties.ExtensionURL]: The file URL of the Media Extension bundle.
//   - [AVMediaExtensionProperties.ContainingBundleURL]: The file URL of the host application for the Media Extension.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties
type AVMediaExtensionProperties struct {
	objectivec.Object
}

// AVMediaExtensionPropertiesFromID constructs a [AVMediaExtensionProperties] from an objc.ID.
//
// An object that describes a Media Extension.
func AVMediaExtensionPropertiesFromID(id objc.ID) AVMediaExtensionProperties {
	return AVMediaExtensionProperties{objectivec.Object{ID: id}}
}

// NOTE: AVMediaExtensionProperties adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMediaExtensionProperties] class.
//
// # Inspecting the extension
//
//   - [IAVMediaExtensionProperties.ExtensionName]: The name of the Media Extension.
//   - [IAVMediaExtensionProperties.ContainingBundleName]: The name of the containing app bundle.
//   - [IAVMediaExtensionProperties.ExtensionIdentifier]
//   - [IAVMediaExtensionProperties.ExtensionURL]: The file URL of the Media Extension bundle.
//   - [IAVMediaExtensionProperties.ContainingBundleURL]: The file URL of the host application for the Media Extension.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties
type IAVMediaExtensionProperties interface {
	objectivec.IObject

	// Topic: Inspecting the extension

	// The name of the Media Extension.
	ExtensionName() string
	// The name of the containing app bundle.
	ContainingBundleName() string
	ExtensionIdentifier() string
	// The file URL of the Media Extension bundle.
	ExtensionURL() foundation.INSURL
	// The file URL of the host application for the Media Extension.
	ContainingBundleURL() foundation.INSURL

	// The properties of the media extension format reader that decodes the asset.
	MediaExtensionProperties() IAVMediaExtensionProperties
	SetMediaExtensionProperties(value IAVMediaExtensionProperties)
}

// Init initializes the instance.
func (m AVMediaExtensionProperties) Init() AVMediaExtensionProperties {
	rv := objc.Send[AVMediaExtensionProperties](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMediaExtensionProperties) Autorelease() AVMediaExtensionProperties {
	rv := objc.Send[AVMediaExtensionProperties](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaExtensionProperties creates a new AVMediaExtensionProperties instance.
func NewAVMediaExtensionProperties() AVMediaExtensionProperties {
	class := getAVMediaExtensionPropertiesClass()
	rv := objc.Send[AVMediaExtensionProperties](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of the Media Extension.
//
// # Discussion
//
// This value corresponds to the extension’s [CFBundleDisplayName].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/extensionName
//
// [CFBundleDisplayName]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CFBundleDisplayName
func (m AVMediaExtensionProperties) ExtensionName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extensionName"))
	return foundation.NSStringFromID(rv).String()
}

// The name of the containing app bundle.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/containingBundleName
func (m AVMediaExtensionProperties) ContainingBundleName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("containingBundleName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/extensionIdentifier
func (m AVMediaExtensionProperties) ExtensionIdentifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extensionIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The file URL of the Media Extension bundle.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/extensionURL
func (m AVMediaExtensionProperties) ExtensionURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extensionURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The file URL of the host application for the Media Extension.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/containingBundleURL
func (m AVMediaExtensionProperties) ContainingBundleURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("containingBundleURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The properties of the media extension format reader that decodes the asset.
//
// See: https://developer.apple.com/documentation/avfoundation/avurlasset/mediaextensionproperties
func (m AVMediaExtensionProperties) MediaExtensionProperties() IAVMediaExtensionProperties {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaExtensionProperties"))
	return AVMediaExtensionPropertiesFromID(objc.ID(rv))
}
func (m AVMediaExtensionProperties) SetMediaExtensionProperties(value IAVMediaExtensionProperties) {
	objc.Send[struct{}](m.ID, objc.Sel("setMediaExtensionProperties:"), value)
}
