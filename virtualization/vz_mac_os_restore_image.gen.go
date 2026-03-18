// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacOSRestoreImage] class.
var (
	_VZMacOSRestoreImageClass     VZMacOSRestoreImageClass
	_VZMacOSRestoreImageClassOnce sync.Once
)

func getVZMacOSRestoreImageClass() VZMacOSRestoreImageClass {
	_VZMacOSRestoreImageClassOnce.Do(func() {
		_VZMacOSRestoreImageClass = VZMacOSRestoreImageClass{class: objc.GetClass("VZMacOSRestoreImage")}
	})
	return _VZMacOSRestoreImageClass
}

// GetVZMacOSRestoreImageClass returns the class object for VZMacOSRestoreImage.
func GetVZMacOSRestoreImageClass() VZMacOSRestoreImageClass {
	return getVZMacOSRestoreImageClass()
}

type VZMacOSRestoreImageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSRestoreImageClass) Alloc() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a version of macOS to install on to a virtual
// machine.
//
// # Overview
// 
// To set up a new VM compatible with the restore image, use
// [VZMacOSRestoreImage.MostFeaturefulSupportedConfiguration] to obtain the [VZMacOSRestoreImage.HardwareModel] of the
// [VZMacPlatformConfiguration]. Then, create a [VZMacOSRestoreImage] object
// by loading an installation media file. Initialize a [VZMacOSInstaller]
// object with this [VZMacOSRestoreImage] object to install the operating
// system onto a VM.
//
// # Getting Information About the Restore Image
//
//   - [VZMacOSRestoreImage.BuildVersion]: The build version this restore image contains.
//   - [VZMacOSRestoreImage.Supported]: A Boolean value that indicates whether the current host supports this restore image.
//   - [VZMacOSRestoreImage.MostFeaturefulSupportedConfiguration]: This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//   - [VZMacOSRestoreImage.OperatingSystemVersion]: The operating system version this restore image contains.
//   - [VZMacOSRestoreImage.URL]: The URL of this restore image.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage
type VZMacOSRestoreImage struct {
	objectivec.Object
}

// VZMacOSRestoreImageFromID constructs a [VZMacOSRestoreImage] from an objc.ID.
//
// An object that describes a version of macOS to install on to a virtual
// machine.
func VZMacOSRestoreImageFromID(id objc.ID) VZMacOSRestoreImage {
	return VZMacOSRestoreImage{objectivec.Object{ID: id}}
}
// NOTE: VZMacOSRestoreImage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacOSRestoreImage] class.
//
// # Getting Information About the Restore Image
//
//   - [IVZMacOSRestoreImage.BuildVersion]: The build version this restore image contains.
//   - [IVZMacOSRestoreImage.Supported]: A Boolean value that indicates whether the current host supports this restore image.
//   - [IVZMacOSRestoreImage.MostFeaturefulSupportedConfiguration]: This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
//   - [IVZMacOSRestoreImage.OperatingSystemVersion]: The operating system version this restore image contains.
//   - [IVZMacOSRestoreImage.URL]: The URL of this restore image.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage
type IVZMacOSRestoreImage interface {
	objectivec.IObject

	// Topic: Getting Information About the Restore Image

	// The build version this restore image contains.
	BuildVersion() string
	// A Boolean value that indicates whether the current host supports this restore image.
	Supported() bool
	// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
	MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements
	// The operating system version this restore image contains.
	OperatingSystemVersion() foundation.NSOperatingSystemVersion
	// The URL of this restore image.
	URL() foundation.INSURL

	// The Mac hardware model.
	HardwareModel() IVZMacHardwareModel
	SetHardwareModel(value IVZMacHardwareModel)
}

// Init initializes the instance.
func (m VZMacOSRestoreImage) Init() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSRestoreImage) Autorelease() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSRestoreImage creates a new VZMacOSRestoreImage instance.
func NewVZMacOSRestoreImage() VZMacOSRestoreImage {
	class := getVZMacOSRestoreImageClass()
	rv := objc.Send[VZMacOSRestoreImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Fetches the latest restore image supported by this host from the network.
//
// # Discussion
// 
// Construct a [VZMacOSInstaller] object with a [VZMacOSRestoreImage] loaded
// from a file on the local file system. A [VZMacOSRestoreImage] fetched with
// the [FetchLatestSupportedWithCompletionHandler] method has a URL property
// that refers to a restore image on the network.
// 
// To use a network restore image, download the file to disk (using
// [URLSession] or similar API). After downloading the restore image, you can
// initialize a [VZMacOSInstaller] using a URL referring to the local file.
//
// [URLSession]: https://developer.apple.com/documentation/Foundation/URLSession
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/latestSupported
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) FetchLatestSupportedWithCompletionHandler(completionHandler MacOSRestoreImageErrorHandler) {
_block0, _cleanup0 := NewMacOSRestoreImageErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("fetchLatestSupportedWithCompletionHandler:"), _block0)
}

// Load a restore image from a file on the local file system.
//
// fileURL: A file URL that indicates the macOS restore image to load.
//
// # Discussion
// 
// [VZMacOSRestoreImage] can load macOS installation media from a local file.
// If the `fileURL` parameter doesn’t refer to a local file, the system
// raises an exception.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/image(from:)
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) LoadFileURLCompletionHandler(fileURL foundation.INSURL, completionHandler MacOSRestoreImageErrorHandler) {
_block1, _cleanup1 := NewMacOSRestoreImageErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("loadFileURL:completionHandler:"), fileURL, _block1)
}

// The build version this restore image contains.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/buildVersion
func (m VZMacOSRestoreImage) BuildVersion() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("buildVersion"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the current host supports this
// restore image.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/isSupported
func (m VZMacOSRestoreImage) Supported() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSupported"))
	return rv
}

// This object represents the most fully featured configuration that’s
// supported by both the current host and by this restore image.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/mostFeaturefulSupportedConfiguration
func (m VZMacOSRestoreImage) MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return VZMacOSConfigurationRequirementsFromID(objc.ID(rv))
}

// The operating system version this restore image contains.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/operatingSystemVersion
func (m VZMacOSRestoreImage) OperatingSystemVersion() foundation.NSOperatingSystemVersion {
	rv := objc.Send[foundation.NSOperatingSystemVersion](m.ID, objc.Sel("operatingSystemVersion"))
	return foundation.NSOperatingSystemVersion(rv)
}

// The URL of this restore image.
//
// # Discussion
// 
// If the restore image loaded using [LoadFileURLCompletionHandler], the value
// of this property is a file URL.
// 
// If you obtain the restore image by fetching it from a server, use
// [FetchLatestSupportedWithCompletionHandler] and set the value of this
// property to a network URL for the installation media file.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/url
func (m VZMacOSRestoreImage) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The Mac hardware model.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacplatformconfiguration/hardwaremodel
func (m VZMacOSRestoreImage) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("hardwareModel"))
	return VZMacHardwareModelFromID(objc.ID(rv))
}
func (m VZMacOSRestoreImage) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[struct{}](m.ID, objc.Sel("setHardwareModel:"), value)
}

// Fetches the latest restore image supported by this host from the network.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/latestsupported
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) LatestSupported() VZMacOSRestoreImage {
	rv := objc.Send[objc.ID](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("fetchLatestSupportedWithCompletionHandler:"))
	return VZMacOSRestoreImageFromID(objc.ID(rv))
}
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) SetLatestSupported(value VZMacOSRestoreImage) {
	objc.Send[struct{}](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("setFetchLatestSupportedWithCompletionHandler::"), value)
}

// FetchLatestSupported is a synchronous wrapper around [VZMacOSRestoreImage.FetchLatestSupportedWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc VZMacOSRestoreImageClass) FetchLatestSupported(ctx context.Context) (*VZMacOSRestoreImage, error) {
	type result struct {
		val *VZMacOSRestoreImage
		err error
	}
	done := make(chan result, 1)
	mc.FetchLatestSupportedWithCompletionHandler(func(val *VZMacOSRestoreImage, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadFileURL is a synchronous wrapper around [VZMacOSRestoreImage.LoadFileURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc VZMacOSRestoreImageClass) LoadFileURL(ctx context.Context, fileURL foundation.INSURL) (*VZMacOSRestoreImage, error) {
	type result struct {
		val *VZMacOSRestoreImage
		err error
	}
	done := make(chan result, 1)
	mc.LoadFileURLCompletionHandler(fileURL, func(val *VZMacOSRestoreImage, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

