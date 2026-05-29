// Code generated from Apple documentation for Vision. DO NOT EDIT.

// Package vision provides Go bindings for the Vision framework.
//
// Analyze image and video content in your app using computer vision
// algorithms for object detection, text recognition, and image segmentation.
//
// The Vision framework provides pretrained machine learning models for
// computer vision tasks. Use Vision to analyze still images and video for a
// variety of purposes, including:
//
// # Pose analysis
//
//   - [Supporting Pose Types]: Types you use when working with pose analysis. ([Chirality])
//
// # Legacy API
//
//   - [Original Objective-C and Swift API] ([VNRequest], [VNImageBasedRequest], [VNClassifyImageRequest], [VNGenerateImageFeaturePrintRequest], [VNFeaturePrintObservation])//
//
// # Key Types
//
//   - [VNGenerateOpticalFlowRequest] - An object that generates directional change vectors for each pixel in the targeted image.
//   - [VNHomographicImageRegistrationRequest] - An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
//   - [VNTranslationalImageRegistrationRequest] - An image-analysis request that determines the affine transform necessary to align the content of two images.
//   - [VNImageRegistrationRequest] - The abstract superclass for image-analysis requests that align images according to their content.
//   - [VNTargetedImageRequest] - The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
//   - [VNBarcodeObservation] - An object that represents barcode information that an image analysis request detects.
//   - [VNFaceLandmarks2D] - A collection of facial features that a request detects.
//   - [VNVector] - An immutable 2D vector represented by its x-axis and y-axis projections.
//   - [VNImageRequestHandler] - An object that processes one or more image-analysis request pertaining to a single image.
//   - [VNRecognizeTextRequest] - An image-analysis request that finds and recognizes text in an image.
//
// [Original Objective-C and Swift API]: https://developer.apple.com/documentation/vision/original-objective-c-and-swift-api
// [Supporting Pose Types]: https://developer.apple.com/documentation/vision/supporting-pose-types
package vision

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Vision library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Vision.framework/Vision",
	"/usr/lib/libVision.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: Vision: failed to load framework from any known path\n")
	}
}
