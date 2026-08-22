// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

// Package imagecapturecore provides Go bindings for the ImageCaptureCore framework.
//
// Browse for media devices and control them programmatically from your app.
//
// Using ImageCaptureCore, your app can:
//
// # Essentials
//
//   - [ICDeviceBrowser]: An object for finding digital cameras and scanners. ([ICDeviceBrowserDelegate], [ICDevice])
//   - [Photos Library Entitlement]: A Boolean value that indicates whether the app has read-write access to the user’s Photos library.
//
// # Cameras
//
//   - [ICCameraDevice]: An object that represents a camera. ([ICUploadOption], [ICDownloadOption], [ICCameraDeviceDownloadDelegate], [ICDeleteResult], [ICDeleteError])
//   - [ICCameraDeviceDelegate]: Methods for detecting cameras, getting metadata and thumbnails, handling access and capability changes, and performing other actions on connected cameras.
//   - [ICCameraItem]: An abstract class that represents a camera item. ([ICCameraItemMetadataOption], [ICCameraItemThumbnailOption])
//   - [ICCameraFile]: An object that represents a file on a camera. ([ICEXIFOrientationType])
//   - [ICCameraFolder]: An object that represents a folder on a camera.
//
// # Scanners
//
//   - [ICScannerDevice]: An object that represents a scanner. ([ICScannerFunctionalUnitType], [ICScannerFunctionalUnitState])
//   - [ICScannerDeviceDelegate]: Methods for determining availability, selecting a functional unit, and performing scans on connected scanners.
//   - [Scanner Configuration]: Examine a scanner’s functional units and features. ([ICScannerBandData], [ICScannerBitDepth], [ICScannerColorDataFormatType], [ICScannerDocumentType], [ICScannerMeasurementUnit])//
//
// # Key Types
//
//   - [ICCameraFile] - An object that represents a file on a camera.
//   - [ICScannerFunctionalUnit] - An abstract class that represents a scanner functional unit.
//   - [ICCameraDevice] - An object that represents a camera.
//   - [ICDevice] - An abstract object that represents a device.
//   - [ICCameraItem] - An abstract class that represents a camera item.
//   - [ICDeviceBrowser] - An object for finding digital cameras and scanners.
//   - [ICScannerBandData] - The options for each band of data that the scanner reads.
//   - [ICScannerDevice] - An object that represents a scanner.
//   - [ICScannerFunctionalUnitDocumentFeeder] - An object that represents the document feeder unit on a scanner.
//   - [ICScannerFeatureEnumeration] - A feature that can have one of several discrete values, strings or numbers.
//
// [Photos Library Entitlement]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.security.personal-information.photos-library
// [Scanner Configuration]: https://developer.apple.com/documentation/imagecapturecore/scanner-configuration
package imagecapturecore

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ImageCaptureCore library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ImageCaptureCore.framework/ImageCaptureCore",
	"/usr/lib/libImageCaptureCore.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: ImageCaptureCore: failed to load framework from any known path\n")
	}
}
