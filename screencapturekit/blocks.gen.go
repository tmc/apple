// Code generated from Apple documentation. DO NOT EDIT.

package screencapturekit

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// CGImageRefErrorHandler handles Closure that processes the screenshot taken from the streaming content.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SCScreenshotManager.CaptureImageInRectCompletionHandler]
//   - [SCScreenshotManager.CaptureImageWithFilterConfigurationCompletionHandler]
type CGImageRefErrorHandler = func(coregraphics.CGImageRef, error)

// NewCGImageRefErrorBlock wraps a Go [CGImageRefErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SCScreenshotManager.CaptureImageInRectCompletionHandler]
//   - [SCScreenshotManager.CaptureImageWithFilterConfigurationCompletionHandler]
func NewCGImageRefErrorBlock(handler CGImageRefErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal coregraphics.CGImageRef, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CMSampleBufferRefErrorHandler handles Closure that processes the capture taken from streaming content.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SCScreenshotManager.CaptureSampleBufferWithFilterConfigurationCompletionHandler]
type CMSampleBufferRefErrorHandler = func(uintptr, error)

// NewCMSampleBufferRefErrorBlock wraps a Go [CMSampleBufferRefErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SCScreenshotManager.CaptureSampleBufferWithFilterConfigurationCompletionHandler]
func NewCMSampleBufferRefErrorBlock(handler CMSampleBufferRefErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal uintptr, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles A completion handler the system calls when this method completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SCStream.StartCaptureWithCompletionHandler]
//   - [SCStream.StopCaptureWithCompletionHandler]
//   - [SCStream.UpdateConfigurationCompletionHandler]
//   - [SCStream.UpdateContentFilterCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SCStream.StartCaptureWithCompletionHandler]
//   - [SCStream.StopCaptureWithCompletionHandler]
//   - [SCStream.UpdateConfigurationCompletionHandler]
//   - [SCStream.UpdateContentFilterCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// SCScreenshotOutputErrorHandler handles Is the handler that will deliver the SCScreenshotOutput object to the client
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SCScreenshotManager.CaptureScreenshotWithFilterConfigurationCompletionHandler]
//   - [SCScreenshotManager.CaptureScreenshotWithRectConfigurationCompletionHandler]
type SCScreenshotOutputErrorHandler = func(*SCScreenshotOutput, error)

// NewSCScreenshotOutputErrorBlock wraps a Go [SCScreenshotOutputErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SCScreenshotManager.CaptureScreenshotWithFilterConfigurationCompletionHandler]
//   - [SCScreenshotManager.CaptureScreenshotWithRectConfigurationCompletionHandler]
func NewSCScreenshotOutputErrorBlock(handler SCScreenshotOutputErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *SCScreenshotOutput
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := SCScreenshotOutputFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// SCShareableContentErrorHandler handles A callback the system invokes with the shareable content, or an error if a failure occurs.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SCShareableContent.GetCurrentProcessShareableContentWithCompletionHandler]
//   - [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindowCompletionHandler]
//   - [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindowCompletionHandler]
//   - [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler]
//   - [SCShareableContent.GetShareableContentWithCompletionHandler]
type SCShareableContentErrorHandler = func(*SCShareableContent, error)

// NewSCShareableContentErrorBlock wraps a Go [SCShareableContentErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SCShareableContent.GetCurrentProcessShareableContentWithCompletionHandler]
//   - [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindowCompletionHandler]
//   - [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindowCompletionHandler]
//   - [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler]
//   - [SCShareableContent.GetShareableContentWithCompletionHandler]
func NewSCShareableContentErrorBlock(handler SCShareableContentErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *SCShareableContent
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := SCShareableContentFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}
