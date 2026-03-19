// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCScreenshotManager] class.
var (
	_SCScreenshotManagerClass     SCScreenshotManagerClass
	_SCScreenshotManagerClassOnce sync.Once
)

func getSCScreenshotManagerClass() SCScreenshotManagerClass {
	_SCScreenshotManagerClassOnce.Do(func() {
		_SCScreenshotManagerClass = SCScreenshotManagerClass{class: objc.GetClass("SCScreenshotManager")}
	})
	return _SCScreenshotManagerClass
}

// GetSCScreenshotManagerClass returns the class object for SCScreenshotManager.
func GetSCScreenshotManagerClass() SCScreenshotManagerClass {
	return getSCScreenshotManagerClass()
}

type SCScreenshotManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCScreenshotManagerClass) Alloc() SCScreenshotManager {
	rv := objc.Send[SCScreenshotManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An instance for the capture of single frames from a stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager
type SCScreenshotManager struct {
	objectivec.Object
}

// SCScreenshotManagerFromID constructs a [SCScreenshotManager] from an objc.ID.
//
// An instance for the capture of single frames from a stream.
func SCScreenshotManagerFromID(id objc.ID) SCScreenshotManager {
	return SCScreenshotManager{objectivec.Object{ID: id}}
}
// NOTE: SCScreenshotManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCScreenshotManager] class.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager
type ISCScreenshotManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SCScreenshotManager) Init() SCScreenshotManager {
	rv := objc.Send[SCScreenshotManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SCScreenshotManager) Autorelease() SCScreenshotManager {
	rv := objc.Send[SCScreenshotManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCScreenshotManager creates a new SCScreenshotManager instance.
func NewSCScreenshotManager() SCScreenshotManager {
	class := getSCScreenshotManagerClass()
	rv := objc.Send[SCScreenshotManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Captures a single frame from a stream as an image, using a filter.
//
// contentFilter: The content filter used to select the stream.
//
// config: Configuration information for how to capture the screenshot.
//
// completionHandler: Closure that processes the screenshot taken from the streaming content.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureImage(contentFilter:configuration:completionHandler:)
func (_SCScreenshotManagerClass SCScreenshotManagerClass) CaptureImageWithFilterConfigurationCompletionHandler(contentFilter ISCContentFilter, config ISCStreamConfiguration, completionHandler CGImageRefErrorHandler) {
_block2, _cleanup2 := NewCGImageRefErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_SCScreenshotManagerClass.class), objc.Sel("captureImageWithFilter:configuration:completionHandler:"), contentFilter, config, _block2)
}
// Captures a single frame directly from a stream’s buffer, using a filter.
//
// contentFilter: The content filter used to select the stream.
//
// config: Configuration information for how to record the stream buffer.
//
// completionHandler: Closure that processes the capture taken from streaming content.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureSampleBuffer(contentFilter:configuration:completionHandler:)
func (_SCScreenshotManagerClass SCScreenshotManagerClass) CaptureSampleBufferWithFilterConfigurationCompletionHandler(contentFilter ISCContentFilter, config ISCStreamConfiguration, completionHandler CMSampleBufferRefErrorHandler) {
_block2, _cleanup2 := NewCMSampleBufferRefErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_SCScreenshotManagerClass.class), objc.Sel("captureSampleBufferWithFilter:configuration:completionHandler:"), contentFilter, config, _block2)
}
//
// rect: The rect for the region in points on the screen space for the screen shot,
// this is display agnostic and supports multiple displays
//
// completionHandler: Is the handler that will deliver the screenshot to the client
//
// # Discussion
// 
// captureImageInRect:completionHandler:
// 
// this method returns an image containing the contents of the rectangle in
// points, specified in display space
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureImage(in:completionHandler:)
func (_SCScreenshotManagerClass SCScreenshotManagerClass) CaptureImageInRectCompletionHandler(rect corefoundation.CGRect, completionHandler CGImageRefErrorHandler) {
_block1, _cleanup1 := NewCGImageRefErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](objc.ID(_SCScreenshotManagerClass.class), objc.Sel("captureImageInRect:completionHandler:"), rect, _block1)
}
//
// contentFilter: Is the filter containing the content to take a screenshot of
//
// config: Is the screenshot configuration containing information on how to format the
// screenshot
//
// completionHandler: Is the handler that will deliver the SCScreenshotOutput object to the
// client
//
// # Discussion
// 
// captureScreenshotWithFilter:configuration:completionHandler:
// 
// this method returns an SCScreenshotOutput object containing CGImages of the
// screenshot requested by the client
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureScreenshot(contentFilter:configuration:completionHandler:)
func (_SCScreenshotManagerClass SCScreenshotManagerClass) CaptureScreenshotWithFilterConfigurationCompletionHandler(contentFilter ISCContentFilter, config ISCScreenshotConfiguration, completionHandler SCScreenshotOutputErrorHandler) {
_block2, _cleanup2 := NewSCScreenshotOutputErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_SCScreenshotManagerClass.class), objc.Sel("captureScreenshotWithFilter:configuration:completionHandler:"), contentFilter, config, _block2)
}
//
// rect: The rect for the region in points on the screen space for the screen shot,
// this is display agnostic and supports multiple displays
//
// config: Is the screenshot configuration containing information on how to format the
// screenshot
//
// completionHandler: Is the handler that will deliver the SCScreenshotOutput object to the
// client
//
// # Discussion
// 
// captureScreenshotWithRect:configuration:completionHandler:
// 
// this method returns an SCScreenshotOutput object containing CGImages of the
// screenshot requested by the client
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureScreenshot(rect:configuration:completionHandler:)
func (_SCScreenshotManagerClass SCScreenshotManagerClass) CaptureScreenshotWithRectConfigurationCompletionHandler(rect corefoundation.CGRect, config ISCScreenshotConfiguration, completionHandler SCScreenshotOutputErrorHandler) {
_block2, _cleanup2 := NewSCScreenshotOutputErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_SCScreenshotManagerClass.class), objc.Sel("captureScreenshotWithRect:configuration:completionHandler:"), rect, config, _block2)
}

// CaptureImageWithFilterConfiguration is a synchronous wrapper around [SCScreenshotManager.CaptureImageWithFilterConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCScreenshotManagerClass) CaptureImageWithFilterConfiguration(ctx context.Context, contentFilter ISCContentFilter, config ISCStreamConfiguration) (coregraphics.CGImageRef, error) {
	type result struct {
		val coregraphics.CGImageRef
		err error
	}
	done := make(chan result, 1)
	sc.CaptureImageWithFilterConfigurationCompletionHandler(contentFilter, config, func(val coregraphics.CGImageRef, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// CaptureSampleBufferWithFilterConfiguration is a synchronous wrapper around [SCScreenshotManager.CaptureSampleBufferWithFilterConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCScreenshotManagerClass) CaptureSampleBufferWithFilterConfiguration(ctx context.Context, contentFilter ISCContentFilter, config ISCStreamConfiguration) (objectivec.IObject, error) {
	type result struct {
		val objectivec.IObject
		err error
	}
	done := make(chan result, 1)
	sc.CaptureSampleBufferWithFilterConfigurationCompletionHandler(contentFilter, config, func(val objectivec.IObject, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// CaptureImageInRect is a synchronous wrapper around [SCScreenshotManager.CaptureImageInRectCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCScreenshotManagerClass) CaptureImageInRect(ctx context.Context, rect corefoundation.CGRect) (coregraphics.CGImageRef, error) {
	type result struct {
		val coregraphics.CGImageRef
		err error
	}
	done := make(chan result, 1)
	sc.CaptureImageInRectCompletionHandler(rect, func(val coregraphics.CGImageRef, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// CaptureScreenshotWithFilterConfiguration is a synchronous wrapper around [SCScreenshotManager.CaptureScreenshotWithFilterConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCScreenshotManagerClass) CaptureScreenshotWithFilterConfiguration(ctx context.Context, contentFilter ISCContentFilter, config ISCScreenshotConfiguration) (*SCScreenshotOutput, error) {
	type result struct {
		val *SCScreenshotOutput
		err error
	}
	done := make(chan result, 1)
	sc.CaptureScreenshotWithFilterConfigurationCompletionHandler(contentFilter, config, func(val *SCScreenshotOutput, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// CaptureScreenshotWithRectConfiguration is a synchronous wrapper around [SCScreenshotManager.CaptureScreenshotWithRectConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCScreenshotManagerClass) CaptureScreenshotWithRectConfiguration(ctx context.Context, rect corefoundation.CGRect, config ISCScreenshotConfiguration) (*SCScreenshotOutput, error) {
	type result struct {
		val *SCScreenshotOutput
		err error
	}
	done := make(chan result, 1)
	sc.CaptureScreenshotWithRectConfigurationCompletionHandler(rect, config, func(val *SCScreenshotOutput, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

